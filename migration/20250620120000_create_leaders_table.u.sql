CREATE TABLE leaders (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    position VARCHAR(255) NOT NULL,
    jamaah VARCHAR(255) NOT NULL,
    periode VARCHAR(255) NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    INDEX idx_leaders_deleted_at (deleted_at)
);
