-- Creation of roles table
CREATE TABLE roles (
  id SERIAL PRIMARY KEY,
  name VARCHAR(128) NOT NULL,
  description TEXT,
  created_at DATE NOT NULL DEFAULT CURRENT_DATE,
  modified_at DATE NOT NULL DEFAULT CURRENT_DATE
);

-- Permissions table
CREATE TABLE permissions (
  id SERIAL PRIMARY KEY,
  name VARCHAR(128) NOT NULL,
  description TEXT,
  created_at DATE NOT NULL DEFAULT CURRENT_DATE,
  modified_at DATE NOT NULL DEFAULT CURRENT_DATE
);

-- user roles table
CREATE TABLE user_roles (
  user_id UUID REFERENCES users(id),
  role_id INT REFERENCES roles(id),
  assigned_at DATE NOT NULL DEFAULT CURRENT_DATE,
  modified_at DATE NOT NULL DEFAULT CURRENT_DATE,
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
(1,'user management', 'Manage all users'),
(2,'role management', 'Manage all user roles'),
(3,'personal profile', 'Manage only personal profile');

-- Assign the permissions to the roles
INSERT INTO role_permissions
(role_id, permission_id)
VALUES
(1, 1),
(1, 2),
(2, 3);
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

