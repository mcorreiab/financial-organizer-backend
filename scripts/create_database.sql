CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid NOT NULL DEFAULT uuid_generate_v4 () PRIMARY KEY,
    username text UNIQUE NOT NULL UNIQUE,
    password text NOT NULL
);

CREATE TABLE expenses (
    id uuid NOT NULL DEFAULT uuid_generate_v4 () PRIMARY KEY,
    expense_name text NOT NULL,
    expense_value bigint NOT NULL,
    expense_user_id uuid NOT NULL references users(id)
);