-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
ALTER TABLE user_roles
  DROP CONSTRAINT user_roles_user_id_fkey,  -- Drop the existing foreign key constraint
  ADD CONSTRAINT user_roles_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;  -- Add the new foreign key with ON DELETE CASCADE

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

-- Revert the foreign key constraint change: Drop the constraint with ON DELETE CASCADE and restore the previous one
ALTER TABLE user_roles
  DROP CONSTRAINT user_roles_user_id_fkey,  -- Drop the new foreign key constraint
  ADD CONSTRAINT user_roles_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id);  -- Restore the previous foreign key without ON DELETE CASCADE
