-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- create the default roles student and admin
INSERT INTO roles (id, role_name, description, created_at, modified_at)
VALUES 
    (1, 'admin', 'Academia''s administrator role.', NOW(), NOW()),
    (2, 'student', 'Academia''s default student role', NOW(), NOW())
ON CONFLICT (id) DO NOTHING;


-- +goose StatementBegin
-- Function to assign new permissions to the admin role
CREATE OR REPLACE FUNCTION assign_permission_to_admin() RETURNS TRIGGER AS 'BEGIN
    INSERT INTO role_permissions (role_id, permission_id, created_at, modified_at)
    VALUES (1, NEW.id, NOW(), NOW());
    RETURN NEW;
END;
' LANGUAGE plpgsql;

-- +goose StatementEnd


-- Trigger to call the function after a new permission is inserted
CREATE TRIGGER assign_permission_to_admin_trigger
AFTER INSERT ON permissions
FOR EACH ROW
EXECUTE FUNCTION assign_permission_to_admin();

-- Create default permissions for various users

-- Permissions for user-related actions
INSERT INTO permissions (permission_name, created_at, modified_at) VALUES
('create:profile', NOW(), NOW()),
('read:profile', NOW(), NOW()),
('update:profile', NOW(), NOW()),
('delete:profile', NOW(), NOW()),
('read:user', NOW(), NOW()),
('read:users', NOW(), NOW()),
('delete:user', NOW(), NOW())
ON CONFLICT (permission_name) DO NOTHING;

-- Permissions for role-related actions
INSERT INTO permissions (permission_name, created_at, modified_at) VALUES
('create:role', NOW(), NOW()),
('read:role', NOW(), NOW()),
('update:role', NOW(), NOW()),
('delete:role', NOW(), NOW()),
('create:role-permission-assignment', NOW(), NOW()),
('delete:role-permission-assignment', NOW(), NOW()),
('create:user-role-assignment', NOW(), NOW()),
('delete:user-role-assignment', NOW(), NOW())
ON CONFLICT (permission_name) DO NOTHING;

-- Permissions for permission-related actions
INSERT INTO permissions (permission_name, created_at, modified_at) VALUES
('create:permission', NOW(), NOW()),
('read:permission', NOW(), NOW()),
('update:permission', NOW(), NOW()),
('delete:permission', NOW(), NOW())
ON CONFLICT (permission_name) DO NOTHING;


-- For the student role, assign user related permissions except deleting
-- Insert user-related permissions (excluding delete actions) for the student role
INSERT INTO role_permissions (role_id, permission_id, created_at, modified_at)
SELECT 2 AS role_id, id AS permission_id, NOW() AS created_at, NOW() AS modified_at
FROM permissions
WHERE permission_name IN (
    'create:profile',
    'read:profile',
    'update:profile',
    'read:user'
);

-- For all users assign the default student role
INSERT INTO user_roles (user_id, role_id, created_at, modified_at)
SELECT id, 2, NOW(), NOW()
FROM users
WHERE NOT EXISTS (
    SELECT 1 
    FROM user_roles 
    WHERE user_roles.user_id = users.id 
    AND user_roles.role_id = 2
);


-- +goose StatementBegin
-- Assign the default student role to all users that will be created now henceforth
CREATE OR REPLACE FUNCTION assign_default_student_role()
RETURNS TRIGGER AS 'BEGIN
    INSERT INTO user_roles (user_id, role_id, created_at, modified_at)
    VALUES (NEW.id, 2, NOW(), NOW());
    RETURN NEW;
END
' LANGUAGE plpgsql;
-- +goose StatementEnd

CREATE TRIGGER assign_default_student_role_trigger
AFTER INSERT ON users
FOR EACH ROW
EXECUTE FUNCTION assign_default_student_role();


-- +goose Down
-- +goose StatementBegin

-- Message to indicate the rollback process
SELECT 'Executing down SQL queries for rollback...';

-- Remove all permissions
DELETE FROM role_permissions WHERE role_id IN (1, 2); -- Ensure dependent role permissions are removed first
DELETE FROM permissions;

-- Revoke roles from users
DELETE FROM user_roles WHERE role_id IN (1, 2);

-- Remove all roles
DELETE FROM roles WHERE id IN (1, 2);

-- Drop trigger and function for assigning permissions to the admin role
DROP TRIGGER IF EXISTS assign_permission_to_admin_trigger ON permissions;
DROP FUNCTION IF EXISTS assign_permission_to_admin();

-- Drop trigger and function for assigning the default student role
DROP TRIGGER IF EXISTS assign_default_student_role_trigger ON users;
DROP FUNCTION IF EXISTS assign_default_student_role();

-- +goose StatementEnd
