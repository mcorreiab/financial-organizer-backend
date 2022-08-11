CREATE TABLE users (
    username text UNIQUE NOT NULL PRIMARY KEY,
    password text NOT NULL
);