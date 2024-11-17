CREATE TABLE IF NOT EXISTS "user"(
    id BIGINT PRIMARY KEY,
    had_trial BOOLEAN
);

CREATE TABLE IF NOT EXISTS country (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS proxy (
    id SERIAL PRIMARY KEY,
    address VARCHAR(255) UNIQUE NOT NULL,
    country_id INTEGER REFERENCES country(id) NOT NULL
);

CREATE TABLE IF NOT EXISTS subscription (
    id SERIAL PRIMARY KEY,
    expiration_date TIMESTAMP,
    user_id BIGINT REFERENCES "user"(id) NOT NULL,
    country_id INTEGER REFERENCES country(id)
);

CREATE TYPE keyType AS ENUM (
    'TEXT',
    'FILE',
    'PHOTO'
);

CREATE TABLE IF NOT EXISTS key (
    id SERIAL PRIMARY KEY,
    key_type keyType,
    data bytea,
    subscription_id INTEGER REFERENCES subscription(id),
    proxy_id INTEGER REFERENCES proxy(id) NOT NULL,
    id_in_proxy VARCHAR(255) NOT NULL
);