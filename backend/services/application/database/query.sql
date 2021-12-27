-- name: CreateApplicationRequest :one
INSERT INTO applications (name,
                          social_security,
                          date_of_birth,
                          drivers_license,
                          previous_address,
                          previous_landlord,
                          previous_landlord_number,
                          employer,
                          salary,
                          user_id,
                          apartment_id,
                          attachments)
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12)
RETURNING *;

-- name: GetApplicationRequest :one
SELECT *
FROM applications
WHERE application_request_id = $1;

-- name: ListApplicationRequest :many
SELECT *
FROM applications
ORDER BY application_request_id;

-- name: UpdateApplicationRequest :exec
UPDATE applications
SET name                    = $2,
    previous_address        =$3,
    previous_landlord=$4,
    previous_landlord_number=$5,
    employer=$6,
    salary=$7,
    attachments=$8
WHERE application_request_id = $1;

-- name: UpdateAttachments :exec
UPDATE applications
SET attachments = $2
WHERE application_request_id = $1;

-- name: DeleteApplicationRequest :exec
DELETE
FROM applications
WHERE application_request_id = $1;

-- name: CreateApplicationResponse :one
INSERT INTO application_response (status,
                                  application_id)
VALUES ($1,
        $2)
RETURNING *;

-- name: GetApplicationResponse :one
SELECT *
FROM application_response
WHERE application_response_id = $1;

-- name: ListApplicationResponse :many
SELECT *
FROM application_response
ORDER BY application_response_id;

-- name: UpdateApplicationResponse :exec
UPDATE application_response
SET status = $2
WHERE application_response_id = $1;

-- name: DeleteApplicationResponse :exec
DELETE
FROM application_response
WHERE application_response_id = $1;
