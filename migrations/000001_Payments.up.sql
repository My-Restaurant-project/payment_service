CREATE TABLE IF NOT EXISTS Payments( 
    id UUID NOT NULL,
    reservation_id UUID NOT NULL,
    amount DECIMAL(8, 2) NOT NULL,
    payment_method VARCHAR(255) NOT NULL,
    payment_status VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY(id)
);