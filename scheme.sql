CREATE TABLE IF NOT EXISTS "user"(
    id BIGINT PRIMARY KEY,
    had_trial BOOLEAN
);

CREATE TABLE IF NOT EXISTS country (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE
);

CREATE TABLE IF NOT EXISTS proxy (
    id SERIAL PRIMARY KEY,
    address VARCHAR(255) UNIQUE,
    country_id SERIAL REFERENCES country(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS subscription (
    id SERIAL PRIMARY KEY,
    user_id SERIAL REFERENCES "user"(id) ON DELETE CASCADE,
    country_id SERIAL REFERENCES country(id) ON DELETE CASCADE,
    expiration_date DATE
);

CREATE TYPE keyType as enum (
    'TEXT'
    'FILE'
    'PHOTO'
);

CREATE TABLE IF NOT EXISTS key (
    id SERIAL PRIMARY KEY,
    data bytea,
    key_type keyType,
    subscription_id SERIAL REFERENCES subscription(id) ON DELETE CASCADE
);