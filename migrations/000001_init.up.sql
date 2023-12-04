CREATE TABLE users
(
    id  SERIAL PRIMARY KEY ,
    login varchar(50) NOT NULL UNIQUE ,
    password varchar NOT NULL,
    current FLOAT,
    withdraw INT
);

CREATE TABLE orders(
    id SERIAL PRIMARY KEY ,
    user_id INT REFERENCES users(id)  ,
    number VARCHAR UNIQUE,
    status VARCHAR,
    uploaded_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    accrual FLOAT DEFAULT 0
);

CREATE TABLE withdraws(
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    number VARCHAR(50) NOT NULL UNIQUE,
    status VARCHAR(30) DEFAULT 'NEW',
    sum FLOAT,
    uploaded_at TIMESTAMP
);