-- Active: 1731435335596@@localhost@3362
-- go_http
USE go_http;

CREATE TABLE users (
    uuid CHAR(36) NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL
);

CREATE INDEX idx_created_at ON users (created_at);
CREATE INDEX idx_updated_at ON users (updated_at);
