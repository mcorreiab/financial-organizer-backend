CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid NOT NULL DEFAULT uuid_generate_v4 (),
    username text UNIQUE NOT NULL PRIMARY KEY,
    password text NOT NULL
);

CREATE TABLE expenses (
    id uuid NOT NULL DEFAULT uuid_generate_v4 () PRIMARY KEY,
    expense_name text NOT NULL,
    expense_value bigint NOT NULL,
    expense_user_id text NOT NULL references users(username)
);