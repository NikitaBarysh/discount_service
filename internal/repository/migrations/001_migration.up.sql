CREATE TABLE IF NOT EXISTS users(
    id  SERIAL PRIMARY KEY ,
    login varchar(50) NOT NULL UNIQUE ,
    password varchar NOT NULL,
    current INTEGER  DEFAULT 0,
    withdraw INTEGER  DEFAULT 0
);

CREATE TABLE IF NOT EXISTS orders(
    id SERIAL PRIMARY KEY ,
    user_id INT REFERENCES users(id)  ,
    number VARCHAR UNIQUE,
    status VARCHAR,
    uploaded_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    accrual INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS withdraws(
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    number VARCHAR(50) NOT NULL UNIQUE,
    status VARCHAR(30) DEFAULT 'NEW',
    sum INTEGER ,
    uploaded_at TIMESTAMP
);