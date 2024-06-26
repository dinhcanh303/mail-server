-- name: CreateServer :one
INSERT INTO mail.servers (
    name,
    host,
    port,
    auth_protocol,
    username,
    password,
    from_name,
    from_address,
    tls_type,
    max_connections,
    idle_timeout,
    retries,
    wait_timeout
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) RETURNING *;

-- name: GetServer :one
SELECT * FROM mail.servers WHERE id = $1;

-- name: GetServers :many
SELECT * FROM mail.servers LIMIT $1 OFFSET $2;

-- name: UpdateServer :one
UPDATE mail.servers SET
    name = COALESCE(sqlc.narg(name),name),
    host = COALESCE(sqlc.narg(host),host),
    port = COALESCE(sqlc.narg(port),port),
    auth_protocol = COALESCE(sqlc.narg(auth_protocol),auth_protocol),
    username = COALESCE(sqlc.narg(username),username),
    password = COALESCE(sqlc.narg(password),password),
    from_name = COALESCE(sqlc.narg(from_name),from_name),
    from_address = COALESCE(sqlc.narg(from_address),from_address),
    tls_type = COALESCE(sqlc.narg(tls_type),tls_type),
    max_connections = COALESCE(sqlc.narg(max_connections),max_connections),
    idle_timeout = COALESCE(sqlc.narg(idle_timeout),idle_timeout),
    retries = COALESCE(sqlc.narg(retries),retries),
    wait_timeout = COALESCE(sqlc.narg(wait_timeout),wait_timeout)
WHERE id = sqlc.arg(id) RETURNING *;

-- name: DeleteServer :exec
DELETE FROM mail.servers WHERE id = $1;


-- name: CreateTemplate :one
INSERT INTO mail.templates (
    name,
    html,
    status
) VALUES ($1,$2,$3) RETURNING *;

-- name: GetTemplate :one
SELECT * FROM mail.templates WHERE id = $1;

-- name: GetTemplates :many
SELECT * FROM mail.templates LIMIT $1 OFFSET $2;

-- name: GetTemplatesActive :many
SELECT * FROM mail.templates WHERE status = 'active';

-- name: UpdateTemplate :one
UPDATE mail.templates SET 
    name = COALESCE(sqlc.narg(name),name),
    html = COALESCE(sqlc.narg(html),html),
    status = COALESCE(sqlc.narg(status),status)
WHERE id = sqlc.arg(id) RETURNING *;

-- name: DeleteTemplate :exec
DELETE FROM mail.templates WHERE id = $1;


-- name: CreateClient :one
INSERT INTO mail.clients (
    name,
    server_id,
    template_id,
    api_key
) VALUES ($1,$2,$3,$4) RETURNING *;

-- name: GetClient :one
SELECT * FROM mail.clients WHERE id = $1;

-- name: GetClientByApiKey :one
SELECT * FROM mail.clients WHERE api_key = $1;

-- name: GetClients :many
SELECT * FROM mail.clients LIMIT $1 OFFSET $2;

-- name: UpdateClient :one
UPDATE mail.clients SET 
    name = COALESCE(sqlc.narg(name),name),
    server_id = COALESCE(sqlc.narg(server_id),server_id),
    template_id = COALESCE(sqlc.narg(template_id),template_id)
WHERE id = sqlc.arg(id) RETURNING *;

-- name: DeleteClient :exec
DELETE FROM mail.clients WHERE id = $1;


-- name: CreateHistory :one
INSERT INTO mail.histories (
    api_key,
    to_,
    subject,
    cc,
    bcc,
    content,
    status
) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING *;

-- name: GetHistory :one
SELECT * FROM mail.histories WHERE id = $1;

-- name: GetHistories :many
SELECT * FROM mail.histories LIMIT $1 OFFSET $2;

-- name: UpdateHistory :one
UPDATE mail.histories SET 
    to_ = COALESCE(sqlc.narg(to_),to_),
    subject = COALESCE(sqlc.narg(subject),subject),
    cc = COALESCE(sqlc.narg(cc),cc),
    bcc = COALESCE(sqlc.narg(bcc),bcc),
    content = COALESCE(sqlc.narg(content),content),
    status = COALESCE(sqlc.narg(status),status)
WHERE id = sqlc.arg(id) RETURNING *;

-- name: DeleteHistory :exec
DELETE FROM mail.histories WHERE id = $1;
