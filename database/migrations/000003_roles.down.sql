-- undo the triggers
DROP TRIGGER IF EXISTS after_user_insert ON users;
drop function if exists assign_student_role()
;

-- Remove the view
DROP VIEW IF EXISTS user_role_permissions;

-- To undo the changes
DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS user_roles;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS roles;

