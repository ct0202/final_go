CREATE TABLE bicycles (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    available BOOLEAN NOT NULL
);

CREATE TABLE orders (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    bicycle_id VARCHAR(36) NOT NULL,
    total_price DOUBLE PRECISION NOT NULL,
    status VARCHAR(50) NOT NULL,
    FOREIGN KEY (bicycle_id) REFERENCES bicycles(id)
);

CREATE TABLE rentals (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    bicycle_id VARCHAR(36) NOT NULL,
    start_time BIGINT NOT NULL,
    end_time BIGINT,
    status VARCHAR(50) NOT NULL,
    FOREIGN KEY (bicycle_id) REFERENCES bicycles(id)
);