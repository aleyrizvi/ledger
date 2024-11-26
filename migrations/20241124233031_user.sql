-- migrate:up
CREATE SEQUENCE IF NOT EXISTS users_id_seq;

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY DEFAULT nextval('users_id_seq'),
    balance INTEGER DEFAULT 0
);

INSERT INTO users (balance) VALUES (0);
INSERT INTO users (balance) VALUES (0);
INSERT INTO users (balance) VALUES (0);

-- migrate:down
DROP TABLE IF EXISTS users;

DROP SEQUENCE IF EXISTS users_id_seq;
