-- name: CreateCampus :one
INSERT INTO campus (
  campus_name, campus_address, city, county, zip_code, country, established_year, picture_url
) VALUES ( 
$1, $2, $3, $4, $5, $6, $7, $8

)
RETURNING *;


-- name: GetAllCampuses :many
SELECT * FROM campus
LIMIT $1
OFFSET $2;


-- name: UpdateCampusByID :one
UPDATE campus
  SET  
    campus_name = COALESCE($1, campus_name),
    campus_address = COALESCE($2, campus_address),
    city = COALESCE($3, city),
    county = COALESCE($4, county),
    zip_code = COALESCE($5, zip_code),
    country = COALESCE($6, country),
    established_year = COALESCE($7, established_year),
    picture_url = COALESCE($8, picture_url)
  WHERE id = $9
RETURNING *;



-- name: GetCampusByID :one
SELECT * FROM campus WHERE id = $1;

-- name: DeleteCampus :exec
DELETE FROM campus
  WHERE id = $1;
