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
select p.id,
       p.name,
       p.description,
       p.image_url,
       p.category_id,
       c.name as category_name,
       COALESCE(s.single_price, 0)::float4 as single_price,
       COALESCE(s.subs_price, 0)::float4 as subs_price
from products p
         join public.categories c on p.category_id = c.id
         join public.skus s on p.id = s.product_id
order by p.id
limit $1 offset $2;

-- name: GetProductCategories :many
select id, name
from categories
order by id;
