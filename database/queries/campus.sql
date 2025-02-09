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
    campus_name = COALESCE(campus_name, $1),
    campus_address = COALESCE(campus_address, $2),
    city = COALESCE(city, $3),
    county = COALESCE(county, $4),
    zip_code = COALESCE(zip_code, $5),
    country = COALESCE(country, $6),
    established_year = COALESCE(established_year, $7),
    picture_url = COALESCE(picture_url, $8)
  WHERE id = $9
RETURNING *;



-- name: GetCampusByID :one
SELECT * FROM campus WHERE id = $1;

-- name: DeleteCampus :exec
DELETE FROM campus
  WHERE id = $1;
