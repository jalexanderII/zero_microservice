CREATE TYPE application_status AS ENUM (
    'STATUS_UNKNOWN',
    'PENDING',
    'APPROVED',
    'DENIED'
    );

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS applications
(
    application_request_id   SERIAL PRIMARY KEY,
    name                     varchar(255) NOT NULL,
    social_security          text         NOT NULL,
    date_of_birth            text         NOT NULL,
    drivers_license          text         NOT NULL,
    previous_address         text,
    previous_landlord        text,
    previous_landlord_number text,
    employer                 text,
    salary                   integer      NOT NULL,
    created_at               timestamp    NOT NULL DEFAULT NOW(),
    user_id                  integer      NOT NULL,
    apartment_id             integer      NOT NULL,
    attachments              text[]
);

CREATE INDEX applications_name_idx ON applications (name);

CREATE TABLE IF NOT EXISTS application_response
(
    application_response_id SERIAL PRIMARY KEY,
    reference_id            uuid DEFAULT uuid_generate_v4(),
    status                  application_status NOT NULL,
    application_id          integer            NOT NULL REFERENCES applications (application_request_id)
);

CREATE INDEX application_response_reference_id_idx ON application_response (reference_id);





