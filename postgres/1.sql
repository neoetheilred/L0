DROP TABLE IF EXISTS "deliveries";
DROP SEQUENCE IF EXISTS deliveries_id_seq;
CREATE SEQUENCE deliveries_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."deliveries" (
    "id" bigint DEFAULT nextval('deliveries_id_seq') NOT NULL,
    "name" text,
    "phone" text,
    "zip" text,
    "city" text,
    "address" text,
    "region" text,
    "email" text,
    "order_id" bigint,
    CONSTRAINT "deliveries_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "items";
DROP SEQUENCE IF EXISTS items_id_seq;
CREATE SEQUENCE items_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."items" (
    "id" bigint DEFAULT nextval('items_id_seq') NOT NULL,
    "chrt_id" bigint,
    "track_number" text,
    "price" bigint,
    "rid" text,
    "name" text,
    "sale" bigint,
    "size" text,
    "total_price" bigint,
    "nm_id" bigint,
    "brand" text,
    "status" bigint,
    "order_id" bigint,
    CONSTRAINT "items_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "orders";
DROP SEQUENCE IF EXISTS orders_id_seq;
CREATE SEQUENCE orders_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."orders" (
    "id" bigint DEFAULT nextval('orders_id_seq') NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "order_uid" text,
    "track_number" text,
    "entry" text,
    "locale" text,
    "internal_signature" text,
    "customer_id" text,
    "delivery_service" text,
    "shardkey" text,
    "sm_id" bigint,
    "date_created" timestamptz,
    "oof_shard" text,
    CONSTRAINT "orders_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "idx_orders_deleted_at" ON "public"."orders" USING btree ("deleted_at");


DROP TABLE IF EXISTS "payments";
DROP SEQUENCE IF EXISTS payments_id_seq;
CREATE SEQUENCE payments_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."payments" (
    "id" bigint DEFAULT nextval('payments_id_seq') NOT NULL,
    "transaction" text,
    "request_id" text,
    "currency" text,
    "provider" text,
    "amount" bigint,
    "payment_dt" bigint,
    "bank" text,
    "delivery_cost" bigint,
    "goods_total" bigint,
    "custom_fee" bigint,
    "order_id" bigint,
    CONSTRAINT "payments_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


ALTER TABLE ONLY "public"."deliveries" ADD CONSTRAINT "fk_orders_delivery" FOREIGN KEY (order_id) REFERENCES orders(id) NOT DEFERRABLE;

ALTER TABLE ONLY "public"."items" ADD CONSTRAINT "fk_orders_items" FOREIGN KEY (order_id) REFERENCES orders(id) NOT DEFERRABLE;

ALTER TABLE ONLY "public"."payments" ADD CONSTRAINT "fk_orders_payment" FOREIGN KEY (order_id) REFERENCES orders(id) NOT DEFERRABLE;