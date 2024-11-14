-- Active: 1731435335596@@localhost@3362
-- go_http
USE go_http;

CREATE TABLE social_accounts (
    uuid CHAR(36) NOT NULL PRIMARY KEY,
    user_uuid CHAR(36) NOT NULL,
    provider VARCHAR(255) NOT NULL,
    provider_user_id VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL
);

CREATE INDEX idx_user_uuid ON social_accounts (user_uuid);
CREATE INDEX idx_provider ON social_accounts (provider);
CREATE INDEX idx_provider_user_id ON social_accounts (provider_user_id);
CREATE INDEX idx_created_at ON social_accounts (created_at);
CREATE INDEX idx_updated_at ON social_accounts (updated_at);
