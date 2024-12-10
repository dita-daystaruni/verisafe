-- Login Info view
CREATE OR REPLACE VIEW login_info AS
SELECT 
  u.id AS user_id,
  u.username,
  u.email,
  c.password,
  c.last_login,
  p.admission_number
FROM 
  users u
JOIN 
  credentials c ON u.id = c.user_id JOIN userprofile p ON u.id = p.user_id;

-- Creation of roles table
CREATE TABLE roles (
  id SERIAL PRIMARY KEY,
  name VARCHAR(128) NOT NULL,
  description TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  modified_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Permissions table
CREATE TABLE permissions (
  id SERIAL PRIMARY KEY,
  name VARCHAR(128) NOT NULL,
  description TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  modified_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- user roles table
CREATE TABLE user_roles (
  user_id UUID REFERENCES users(id),
  role_id INT REFERENCES roles(id),
  assigned_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  modified_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY(user_id, role_id)
);

-- Role permissions
CREATE TABLE role_permissions (
  role_id INT REFERENCES roles(id),
  permission_id INT REFERENCES permissions(id),
  PRIMARY KEY(role_id, permission_id)
);

CREATE OR REPLACE VIEW user_role_permissions AS
SELECT 
    u.id AS user_id,
    u.username,
    u.firstname,
    u.othernames,
    r.name AS role_name,
    p.name AS permission_name
FROM 
    users u
JOIN 
    user_roles ur ON u.id = ur.user_id
JOIN 
    roles r ON ur.role_id = r.id
LEFT JOIN 
    role_permissions rp ON r.id = rp.role_id
LEFT JOIN 
    permissions p ON rp.permission_id = p.id;


-- To create two default roles admin and student
INSERT INTO roles 
(id,name,description) VALUES 
( 1, 'administrator', 'The administrator role'),
( 2, 'student','The default student role' );

-- Create various default permissions for the various user types
INSERT INTO permissions 
(id,name, description)
VALUES 
(1,'read:user', 'Read all users'),
(2,'modify:user', 'Modify user data'),
(3,'delete:user', 'Delete user data'),
(4,'create:user', 'Create a user'),
(5,'read:roles', 'Read all roles'),
(6,'create:roles', 'Create a role'),
(7,'modify:roles', 'Update roles'),
(8,'delete:roles', 'Delete a role'),
(9,'assign:roles', 'Assign roles to users'),
(10, 'read:userprofile', 'Retrieve user profiles'),
(11, 'modify:userprofile', 'Update user profiles'),
(12, 'create:userprofile', 'Create user profiles'),
(13, 'delete:userprofile', 'Delete user profiles');

-- Assign the permissions to the roles
INSERT INTO role_permissions
(role_id, permission_id)
VALUES
(1, 1),
(1, 2),
(1, 3),
(1, 4),
(1, 5),
(1, 6),
(1, 7),
(1, 8),
(1, 9),
(1, 10),
(1, 11),
(1, 12),
(1, 13),
(2, 10),
(2, 11),
(2, 12),
(2, 13);

-- Trigger to assign the default student role to a user
create or replace function assign_student_role()
returns trigger
as $$
BEGIN
    -- Insert the 'student' role for the new user
    INSERT INTO user_roles (user_id, role_id)
    SELECT NEW.id, r.id 
    FROM roles r 
    WHERE r.name = 'student';
    
    RETURN NEW;
END;
$$
language plpgsql
;


-- Bind the trigger
CREATE TRIGGER after_user_insert
AFTER INSERT ON users
FOR EACH ROW
EXECUTE FUNCTION assign_student_role();

