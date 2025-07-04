-- Users table
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    balance BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Cars table
CREATE TABLE IF NOT EXISTS cars (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    code VARCHAR(50) NOT NULL UNIQUE,
    category VARCHAR(50) NOT NULL,
    rental_cost BIGINT NOT NULL,
    is_available BOOLEAN NOT NULL DEFAULT TRUE,
    owner_id BIGINT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_car_owner FOREIGN KEY (owner_id)
        REFERENCES users (id)
        ON DELETE SET NULL
);

-- Rental Histories table
CREATE TABLE IF NOT EXISTS rental_histories (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    car_id BIGINT NOT NULL,
    cost BIGINT NOT NULL,
    rented_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    return_at TIMESTAMPTZ,

    CONSTRAINT fk_rental_user FOREIGN KEY (user_id)
        REFERENCES users (id)
        ON DELETE CASCADE,

    CONSTRAINT fk_rental_car FOREIGN KEY (car_id)
        REFERENCES cars (id)
        ON DELETE CASCADE
);

-- Transaction Histories table
CREATE TABLE IF NOT EXISTS transaction_histories (
    id BIGSERIAL PRIMARY KEY,
    sender_id BIGINT,
    receiver_id BIGINT,
    rental_id BIGINT UNIQUE,
    amount BIGINT NOT NULL,
    description VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_tx_sender FOREIGN KEY (sender_id)
        REFERENCES users (id)
        ON DELETE SET NULL,

    CONSTRAINT fk_tx_receiver FOREIGN KEY (receiver_id)
        REFERENCES users (id)
        ON DELETE SET NULL,

    CONSTRAINT fk_tx_rental FOREIGN KEY (rental_id)
        REFERENCES rental_histories (id)
        ON DELETE SET NULL
);
