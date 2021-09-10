DROP DATABASE IF EXISTS balance_service_db;
CREATE DATABASE balance_service_db
    WITH OWNER postgres
    ENCODING 'utf8';
\connect balance_service_db;


DROP TABLE IF EXISTS balances CASCADE;
CREATE TABLE balances (
    user_id SERIAL NOT NULL PRIMARY KEY,
    balance DECIMAL(12,2) NOT NULL,

    CONSTRAINT balance_value CHECK (balance >= 0)
);

DROP TABLE IF EXISTS transaction_types CASCADE;
CREATE TABLE transaction_types (
    id SERIAL NOT NULL PRIMARY KEY,
    title TEXT NOT NULL
);

DROP TABLE IF EXISTS transactions CASCADE;
CREATE TABLE transactions (
    id SERIAL NOT NULL PRIMARY KEY,
    type_operation INTEGER NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    user_id INTEGER NOT NULL,
    context TEXT NOT NULL,
    value DECIMAL(12,2) NOT NULL,

    FOREIGN KEY (user_id) REFERENCES balances (user_id),
    FOREIGN KEY (type_operation) REFERENCES transaction_types (id),
    CONSTRAINT operation_value CHECK (transactions.value >= 0)
);

INSERT INTO transaction_types (title) VALUES
    ('improve'),
    ('withdraw'),
    ('transfer');
