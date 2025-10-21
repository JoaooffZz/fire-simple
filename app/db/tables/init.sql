CREATE TABLE IF NOT EXISTS clients_ip (
    client_ip TEXT PRIMARY KEY NOT NULL,
    is_secure BOOLEAN NOT NULL DEFAULT TRUE,
    unix_first_access_date BIGINT NOT NULL
);