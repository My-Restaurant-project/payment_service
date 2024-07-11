CREATE TABLE IF NOT EXISTS Payments( 
    id CHAR(36) NOT NULL,
    reservation_id CHAR(36) NOT NULL,
    amount DECIMAL(8, 2) NOT NULL,
    payment_method VARCHAR(255) NOT NULL,
    payment_status VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
);