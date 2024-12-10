-- Make phone number and email address on user to be unique
ALTER TABLE users ADD CONSTRAINT unique_phone UNIQUE (phone);
ALTER TABLE users ADD CONSTRAINT unique_email UNIQUE (email);

-- Add student admission number on user profile
ALTER TABLE userprofile ADD COLUMN
admission_number VARCHAR(7) UNIQUE;

-- Add campus column
ALTER TABLE userprofile ADD COLUMN 
campus VARCHAR(20) NOT NULL DEFAULT 'athi';


-- Add national_id
ALTER TABLE users ADD COLUMN
national_id VARCHAR(20);

-- Create the credentials table
CREATE TABLE IF NOT EXISTS credentials (
  user_id uuid PRIMARY KEY,
  password TEXT NOT NULL,
  last_login TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  modified_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  CONSTRAINT fk_user
  FOREIGN KEY(user_id) REFERENCES users(id)
  ON DELETE CASCADE
);


-- Login Info view
CREATE OR REPLACE VIEW login_info AS
SELECT 
  u.id AS user_id,
  u.username,
  u.email,
  c.password,
  c.last_login
FROM 
  users u
JOIN 
  credentials c ON u.id = c.user_id;

