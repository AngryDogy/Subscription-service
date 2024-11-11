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
    'TEXT'
    'FILE'
    'PHOTO'
);

CREATE TABLE IF NOT EXISTS key (
    id SERIAL PRIMARY KEY,
    key_type keyType,
    data bytea,
    subscription_id INTEGER REFERENCES subscription(id),
    proxy_id INTEGER REFERENCES proxy(id) NOT NULL
);

select c.name, k.id, k.key_type, s.id, k.proxy_id, k.data
from key k join subscription s
on k.subscription_id = s.id join country c
on s.country_id = c.id
where s.user_id = 1 and s.expiration_date > now();