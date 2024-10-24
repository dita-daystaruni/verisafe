-- Drop the unique constraint on user phone and email
ALTER TABLE users DROP CONSTRAINT IF EXISTS unique_phone;
ALTER TABLE users DROP CONSTRAINT IF EXISTS unique_email;
-- Add the date created table
ALTER TABLE users DROP COLUMN IF EXISTS
date_of_birth;

-- Remove the admission number field
ALTER TABLE userprofile DROP COLUMN IF EXISTS
admission_number;

-- Remove Campus
ALTER TABLE userprofile DROP COLUMN IF EXISTS
campus;

-- Add national_id
ALTER TABLE users DROP COLUMN IF EXISTS
national_id;

-- Drop the login info
DROP VIEW IF EXISTS login_info;

-- Drop the credentials table
DROP TABLE IF EXISTS credentials;

