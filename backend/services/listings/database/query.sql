-- name: CreateApartment :one
INSERT INTO apartments (name,
                        full_address,
                        street,
                        city,
                        state,
                        zip_code,
                        neighborhood,
                        unit,
                        lat,
                        lng,
                        rent,
                        sqft,
                        beds,
                        baths,
                        available_on,
                        days_on_market,
                        description,
                        amenities,
                        upload_ids,
                        is_archived,
                        building_id,
                        realtor_id)
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
        $12,
        $13,
        $14,
        $15,
        $16,
        $17,
        $18,
        $19,
        $20,
        $21,
        $22)
RETURNING *;

-- name: GetApartment :one
SELECT *
FROM apartments
WHERE apartment_id = $1;

-- name: ListApartments :many
SELECT *
FROM apartments
ORDER BY apartment_id;

-- name: UpdateApartment :exec
UPDATE apartments
SET name          = $2,
    full_address  = $3,
    street= $4,
    city= $5,
    state= $6,
    zip_code= $7,
    neighborhood= $8,
    unit= $9,
    lat= $10,
    lng= $11,
    rent= $12,
    sqft= $13,
    beds= $14,
    baths= $15,
    available_on= $16,
    days_on_market= $17,
    description= $18,
    amenities= $19,
    upload_ids= $20,
    is_archived= $21,
    building_id= $22,
    realtor_id= $23
WHERE apartment_id = $1;

-- name: DeleteApartment :exec
DELETE
FROM apartments
WHERE apartment_id = $1;

-- name: CreateBuilding :one
INSERT INTO buildings (name,
                       full_address,
                       street,
                       city,
                       state,
                       zip_code,
                       neighborhood,
                       lat,
                       lng,
                       description,
                       amenities,
                       upload_ids,
                       owner_id)
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
        $12,
        $13)
RETURNING *;

-- name: GetBuilding :one
SELECT *
FROM buildings
WHERE building_id = $1;

-- name: ListBuildings :many
SELECT *
FROM buildings
ORDER BY building_id;

-- name: UpdateBuilding :exec
UPDATE buildings
SET name         = $2,
    full_address = $3,
    street= $4,
    city= $5,
    state= $6,
    zip_code= $7,
    neighborhood= $8,
    lat= $9,
    lng= $10,
    description= $11,
    amenities= $12,
    upload_ids= $13,
    owner_id= $14
WHERE building_id = $1;

-- name: DeleteBuilding :exec
DELETE
FROM buildings
WHERE building_id = $1;

-- name: CreateRealtor :one
INSERT INTO realtors (name,
                      user_id,
                      email,
                      phone_number,
                      company)
VALUES ($1,
        $2,
        $3,
        $4,
        $5)
RETURNING *;

-- name: GetRealtor :one
SELECT *
FROM realtors
WHERE realtor_id = $1;

-- name: ListRealtors :many
SELECT *
FROM realtors
ORDER BY realtor_id;

-- name: UpdateRealtor :exec
UPDATE realtors
SET name        = $2,
    email       = $3,
    phone_number= $4,
    company= $5
WHERE realtor_id = $1;

-- name: DeleteRealtor :exec
DELETE
FROM realtors
WHERE realtor_id = $1;

-- name: CreateOwner :one
INSERT INTO owners (name,
                    user_id,
                    email,
                    phone_number,
                    company)
VALUES ($1,
        $2,
        $3,
        $4,
        $5)
RETURNING *;

-- name: GetOwner :one
SELECT *
FROM owners
WHERE owner_id = $1;

-- name: ListOwners :many
SELECT *
FROM owners
ORDER BY owner_id;

-- name: UpdateOwner :exec
UPDATE owners
SET name        = $2,
    email       = $3,
    phone_number= $4,
    company= $5
WHERE owner_id = $1;

-- name: DeleteOwner :exec
DELETE
FROM owners
WHERE owner_id = $1;

-- name: AppendContentApartment :exec
UPDATE apartments
SET upload_ids = array_append(upload_ids, $2)
WHERE apartment_id = $1;

-- name: AppendContentBuilding :exec
UPDATE buildings
SET upload_ids = array_append(upload_ids, $2)
WHERE building_id = $1;