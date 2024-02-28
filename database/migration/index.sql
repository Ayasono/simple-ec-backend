CREATE TABLE "products" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar,
  "description" text,
  "category_id" int,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "skus" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "product_id" int,
  "sku_name" varchar,
  "price" decimal(10,2),
  "stock_quantity" int,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "purchase_types" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "type_name" varchar
);

CREATE TABLE "product_skus_purchase_types" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "sku_id" int,
  "purchase_type_id" int,
  "override_price" decimal(10,2),
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "users" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "username" varchar,
  "email" varchar,
  "password_hash" varchar,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "addresses" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "user_id" int,
  "address_line1" varchar,
  "address_line2" varchar,
  "city" varchar,
  "state" varchar,
  "country" varchar,
  "zip_code" varchar,
  "address_type" varchar,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "payments" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "user_id" int,
  "payment_type" varchar,
  "provider" varchar,
  "account_number" varchar,
  "expiry_date" date,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "orders" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "user_id" int,
  "address_id" int,
  "payment_id" int,
  "order_status" varchar,
  "total_price" decimal(10,2),
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "order_items" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "order_id" int,
  "sku_id" int,
  "purchase_type_id" int,
  "quantity" int,
  "price_per_item" decimal(10,2),
  "total_price" decimal(10,2),
  "created_at" timestamp,
  "updated_at" timestamp
);

ALTER TABLE "skus" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "product_skus_purchase_types" ADD FOREIGN KEY ("sku_id") REFERENCES "skus" ("id");

ALTER TABLE "product_skus_purchase_types" ADD FOREIGN KEY ("purchase_type_id") REFERENCES "purchase_types" ("id");

ALTER TABLE "addresses" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "payments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("address_id") REFERENCES "addresses" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("payment_id") REFERENCES "payments" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("sku_id") REFERENCES "skus" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("purchase_type_id") REFERENCES "purchase_types" ("id");