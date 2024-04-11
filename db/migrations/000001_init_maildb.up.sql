START TRANSACTION;
CREATE SCHEMA IF NOT EXISTS "mail";

CREATE TABLE
    mail.clients (
        id BIGSERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        server_id BIGINT NOT NULL,
        template_id BIGINT NOT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
        updated_at timestamp with time zone NOT NULL DEFAULT (now()),
        CONSTRAINT pk_mail_clients PRIMARY KEY (id)
    );

CREATE TABLE 
    mail.templates (
        id BIGSERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        html TEXT DEFAULT NULL,
        status VARCHAR(255) DEFAULT "active",
        created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
        updated_at timestamp with time zone NOT NULL DEFAULT (now())
        CONSTRAINT pk_mail_templates PRIMARY KEY (id)
);
CREATE TABLE 
    mail.servers (
        id BIGSERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        host VARCHAR(255) NOT NULL,
        port VARCHAR(255) NOT NULL,
        username VARCHAR(255) NOT NULL,
        password VARCHAR(255) NOT NULL,
        tls VARCHAR(255) DEFAULT "ssl/tls",
        skip_tls BOOLEAN DEFAULT "FALSE",
        max_connections BIGINT DEFAULT 10,
        idle_timeout BIGINT DEFAULT 15,
        retries BIGINT DEFAULT 2,
        wait_timeout BIGINT DEFAULT 5,
        created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
        updated_at timestamp with time zone NOT NULL DEFAULT (now()),
        CONSTRAINT pk_mail_servers PRIMARY KEY (id)
);
