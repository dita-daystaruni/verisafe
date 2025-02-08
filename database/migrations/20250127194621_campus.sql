-- +goose Up
-- +goose StatementBegin
select 'up SQL query'
;

-- enable the uuid
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE campus (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),        -- Unique ID for the campus using UUID
    campus_name VARCHAR(255) NOT NULL,                      -- Name of the campus
    campus_address VARCHAR(255),                            -- Physical address of the campus
    city VARCHAR(100),                                      -- City where the campus is located
    county VARCHAR(100),                                    -- County or province
    zip_code VARCHAR(20),                                   -- Postal code
    country VARCHAR(100),                                   -- Country where the campus is located
    established_year INT,                                   -- Year the campus was established
    picture_url TEXT,
    is_active BOOLEAN DEFAULT TRUE                          -- Whether the campus is currently active
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
select 'down SQL query'
;

DROP TABLE IF EXISTS campus;

DROP EXTENSION IF EXISTS "uuid-ossp";
-- +goose StatementEnd


