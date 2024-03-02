-- name: CreateProduct :one
insert into products (name, description, image_url, category_id, created_at, updated_at)
values ($1, $2, $3, $4, $5, $6)
returning *;

-- name: UpdateProduct :one
update
    products
set name        = $1,
    description = $2,
    image_url   = $3,
    category_id = $4,
    updated_at  = $5
where id = $6
returning *;

-- name: DeleteProduct :one
delete
from products
where id = $1
returning *;

-- name: ListProducts :many
select *
from products
order by id
limit $1
offset $2;
