CREATE DATABASE cars OWNER postgres;
GRANT ALL PRIVILEGES ON DATABASE cars TO postgres;
\c cars
CREATE TABLE IF NOT EXISTS users
(
    id bigserial NOT NULL ,
    first_name text NOT NULL,
    last_name text NOT NULL,
    birth_date timestamp without time zone NOT NULL,
    driver_license bigint NOT NULL,
    luck_number bigint NOT NULL,
    created_at timestamp without time zone NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
)