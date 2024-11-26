-- migrate:up
CREATE TABLE IF NOT EXISTS transactions (
    transaction_id character varying NOT NULL UNIQUE,
    user_id integer NOT NULL,
    amount integer NOT NULL,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);


-- migrate:down
DROP TABLE IF EXISTS transactions;
