CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid NOT NULL DEFAULT uuid_generate_v4 (),
    username text UNIQUE NOT NULL PRIMARY KEY,
    password text NOT NULL
);