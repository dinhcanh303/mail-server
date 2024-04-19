START TRANSACTION;
CREATE SCHEMA IF NOT EXISTS "mail";

CREATE TABLE
    mail.clients (
        id BIGSERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        server_id BIGINT NOT NULL,
        template_id BIGINT NOT NULL,
        is_default BOOLEAN NOT NULL DEFAULT FALSE,
        created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
        updated_at timestamp with time zone NOT NULL DEFAULT (now())
    );

CREATE TABLE 
    mail.templates (
        id BIGSERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        html TEXT DEFAULT NULL,
        status VARCHAR(255) DEFAULT 'active',
        is_default BOOLEAN NOT NULL DEFAULT FALSE,
        created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
        updated_at timestamp with time zone NOT NULL DEFAULT (now())
);
CREATE TABLE 
    mail.servers (
        id BIGSERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        host VARCHAR(255) NOT NULL,
        port BIGINT NOT NULL,
        auth_protocol VARCHAR(255) DEFAULT 'plain',
        username VARCHAR(255) NOT NULL,
        password VARCHAR(255) NOT NULL,
        from_name VARCHAR(255) DEFAULT 'mail',
        from_address VARCHAR(255) DEFAULT 'noreply@server.yoursite.com',
        tls_type VARCHAR(255) DEFAULT 'TLS',
        tls_skip_verify BOOLEAN DEFAULT false,
        max_connections BIGINT DEFAULT 10,
        idle_timeout BIGINT DEFAULT 15,
        retries BIGINT DEFAULT 2,
        wait_timeout BIGINT DEFAULT 5,
        is_default BOOLEAN NOT NULL DEFAULT FALSE,
        created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
        updated_at timestamp with time zone NOT NULL DEFAULT (now())
);
CREATE TABLE 
    mail.histories (
        id BIGSERIAL PRIMARY KEY,
        from_ TEXT NOT NULL,
        to_ TEXT NOT NULL,
        subject TEXT DEFAULT NULL,
        cc TEXT DEFAULT NULL,
        bcc TEXT DEFAULT NULL,
        content JSONB NOT NULL DEFAULT '{}',
        status VARCHAR(255) DEFAULT NULL,
        created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,
        updated_at timestamp with time zone NOT NULL DEFAULT (now())
);
--insert default 
INSERT INTO mail.templates 
(
    id,
    name,
    html,
    status,
    is_default
) VALUES (1,'default','<!doctype html>
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, minimum-scale=1">
        <base target="_blank">

        <style>
            body {
                background-color: #F0F1F3;
                font-family: Helvetica, sans-serif;
                font-size: 15px;
                line-height: 26px;
                margin: 0;
                color: #444;
            }

            pre {
                background: #f4f4f4f4;
                padding: 2px;
            }

            table {
                width: 100%;
                border: 1px solid #ddd;
            }
            table td {
                border-color: #ddd;
                padding: 5px;
            }

            .wrap {
                background-color: #fff;
                padding: 30px;
                max-width: 525px;
                margin: 0 auto;
                border-radius: 5px;
            }

            .button {
                background: #0055d4;
                border-radius: 3px;
                text-decoration: none !important;
                color: #fff !important;
                font-weight: bold;
                padding: 10px 30px;
                display: inline-block;
            }
            .button:hover {
                background: #111;
            }

            .footer {
                text-align: center;
                font-size: 12px;
                color: #888;
            }
                .footer a {
                    color: #888;
                }

            .gutter {
                padding: 30px;
            }

            img {
                max-width: 100%;
            }

            a {
                color: #0055d4;
            }
                a:hover {
                    color: #111;
                }
            @media screen and (max-width: 600px) {
                .wrap {
                    max-width: auto;
                }
                .gutter {
                    padding: 10px;
                }
            }
        </style>
    </head>
<body style="background-color: #F0F1F3;font-family:Helvetica, sans-serif;font-size: 15px;line-height: 26px;margin: 0;color: #444;">
    <div class="gutter" style="padding: 30px;">&nbsp;</div>
    <div class="wrap" style="background-color: #fff;padding: 30px;max-width: 525px;margin: 0 auto;border-radius: 5px;">
        {{ template "content" . }}
    </div>
    
    <div class="footer" style="text-align: center;font-size: 12px;color: #888;">
        <p>
            {{ L.T "email.unsubHelp" }}
            <a href="{{ UnsubscribeURL }}" style="color: #888;">{{ L.T "email.unsub" }}</a>
        </p>
        <p>Powered by <a href="https://test.app" target="_blank" style="color: #888;">test</a></p>
    </div>
    <div class="gutter" style="padding: 30px;">&nbsp;{{ .TrackView }}</div>
</body>
</html>
','active', TRUE);
INSERT INTO mail.servers 
(
    id,
    name,
    host,
    port,
    username,
    password,
    tls_type,
    tls_skip_verify,
    max_connections,
    idle_timeout,
    retries,
    wait_timeout,
    is_default
) VALUES (1,'default','smtp.yoursite.com','465','username','password','TLS',false,10,15,5,10,true);
INSERT INTO mail.clients 
(
    id,
    name,
    server_id,
    template_id,
    is_default
) VALUES (1,'default',1,1,true);
COMMIT;