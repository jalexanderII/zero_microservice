CREATE TYPE content_type AS ENUM (
    'TYPE_UNKNOWN',
    'IMAGE',
    'VIDEO'
    );

CREATE TYPE content_source AS ENUM (
    'SOURCE_UNKNOWN',
    'APARTMENT',
    'BUILDING'
    );

CREATE TABLE IF NOT EXISTS content
(
    content_id     SERIAL PRIMARY KEY,
    filename       text,
    content_type   content_type,
    content_source content_source,
    source_id      integer NOT NULL
);

CREATE TABLE IF NOT EXISTS realtors
(
    realtor_id   SERIAL PRIMARY KEY,
    name         varchar(255) NOT NULL,
    email        text,
    phone_number text,
    company      text
);

CREATE INDEX realtors_name_idx ON realtors (name);

CREATE TABLE IF NOT EXISTS buildings
(
    building_id  SERIAL PRIMARY KEY,
    name         varchar(255) NOT NULL,
    full_address text         NOT NULL,
    street       text         NOT NULL,
    city         text         NOT NULL,
    state        text         NOT NULL,
    zip_code     integer      NOT NULL,
    neighborhood text         NOT NULL,
    lat          integer      NOT NULL,
    lng          integer      NOT NULL,
    description  text,
    amenities    text[],
    upload_ids   text[],
    realtor_id   integer      NOT NULL REFERENCES realtors (realtor_id)
);

CREATE INDEX buildings_name_idx ON buildings (name);

CREATE TABLE IF NOT EXISTS apartments
(
    apartment_id   SERIAL PRIMARY KEY,
    name           varchar(255) NOT NULL,
    full_address   text         NOT NULL,
    street         text         NOT NULL,
    city           text         NOT NULL,
    state          text         NOT NULL,
    zip_code       integer      NOT NULL,
    neighborhood   text         NOT NULL,
    unit           text,
    lat            integer      NOT NULL,
    lng            integer      NOT NULL,
    rent           integer      NOT NULL,
    sqft           integer,
    beds           integer      NOT NULL,
    baths          integer      NOT NULL,
    available_on   timestamp    NOT NULL DEFAULT NOW(),
    created_at     timestamp    NOT NULL DEFAULT NOW(),
    days_on_market integer,
    description    text,
    amenities      text[],
    upload_ids     text[],
    is_archived    boolean      NOT NULL DEFAULT FALSE,
    building_id    integer      NOT NULL REFERENCES buildings (building_id),
    realtor_id     integer      NOT NULL REFERENCES realtors (realtor_id)
);

CREATE INDEX apartments_name_idx ON apartments (name);





