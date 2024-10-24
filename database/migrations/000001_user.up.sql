-- Enable uuid extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- User's table
CREATE TABLE users(
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  username VARCHAR(128) UNIQUE NOT NULL,
  firstname VARCHAR(255) NOT NULL,
  othernames VARCHAR(255) NOT NULL,
  phone VARCHAR(20),
  email VARCHAR(128),
  gender VARCHAR(10) DEFAULT 'unknown',
  active BOOLEAN DEFAULT true,
  created_at DATE NOT NULL DEFAULT CURRENT_DATE,
  modified_at DATE NOT NULL DEFAULT CURRENT_DATE
);

-- Creating user index
CREATE INDEX idx_users_username ON users (username);


-- User Profile
CREATE TABLE userprofile(
  user_id uuid PRIMARY KEY,
  bio TEXT,
  vibe_points INT NOT NULL DEFAULT 0,
  profile_picture_url TEXT NOT NULL,
  last_seen DATE NOT NULL DEFAULT CURRENT_DATE,
  created_at DATE NOT NULL DEFAULT CURRENT_DATE,
  modified_at DATE NOT NULL DEFAULT CURRENT_DATE,

  CONSTRAINT fk_user
  FOREIGN KEY(user_id) REFERENCES users(id)
  ON DELETE CASCADE
);
