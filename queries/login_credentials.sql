-- name: GetUserCredentials :one
select *
from login_info
where username = $1 or email = $2 or user_id = $3 or admission_number = $4
limit 1
;

