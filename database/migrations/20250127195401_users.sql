-- +goose Up
-- +goose StatementBegin
select 'up SQL query'
;
-- +goose StatementEnd
-- enable the uuid
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users(
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  username VARCHAR(128) UNIQUE NOT NULL,
  firstname VARCHAR(255) NOT NULL,
  othernames VARCHAR(255) NOT NULL,
  phone VARCHAR(20) UNIQUE,
  email VARCHAR(128) UNIQUE,
  gender VARCHAR(10) DEFAULT 'unknown',
  active BOOLEAN DEFAULT true,
  national_id VARCHAR(20),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  modified_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


CREATE TABLE userprofile(
  user_id uuid PRIMARY KEY,
  admission_number VARCHAR(7) UNIQUE,
  bio TEXT,
  vibe_points INT NOT NULL DEFAULT 0,
  date_of_birth DATE NOT NULL DEFAULT NOW(),
  profile_picture_url TEXT NOT NULL,
  last_seen TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  modified_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  CONSTRAINT fk_user
  FOREIGN KEY(user_id) REFERENCES users(id)
  ON DELETE CASCADE
);


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
  c.last_login,
  p.admission_number
FROM 
  users u
JOIN 
  credentials c ON u.id = c.user_id JOIN userprofile p ON u.id = p.user_id;


-- +goose Down
-- +goose StatementBegin
select 'down SQL query'
;

DROP VIEW IF EXISTS login_info;
DROP TABLE IF EXISTS credentials;
DROP TABLE IF EXISTS userprofile;
DROP TABLE IF EXISTS users;

DROP EXTENSION IF EXISTS "uuid-ossp";
-- +goose StatementEnd


