/*
 Navicat Premium Data Transfer

 Source Server         : Postgres-db
 Source Server Type    : PostgreSQL
 Source Server Version : 90623
 Source Host           : localhost:5432
 Source Catalog        : aveonline
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 90623
 File Encoding         : 65001

 Date: 23/09/2021 04:37:10
*/


-- ----------------------------
-- Sequence structure for invoices_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."invoices_id_seq";
CREATE SEQUENCE "public"."invoices_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."invoices_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for medicines_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."medicines_id_seq";
CREATE SEQUENCE "public"."medicines_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."medicines_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for promotions_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."promotions_id_seq";
CREATE SEQUENCE "public"."promotions_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."promotions_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Table structure for invoice_medicines
-- ----------------------------
DROP TABLE IF EXISTS "public"."invoice_medicines";
CREATE TABLE "public"."invoice_medicines" (
  "invoice_id" int8 NOT NULL,
  "medicine_id" int8 NOT NULL
)
;
ALTER TABLE "public"."invoice_medicines" OWNER TO "postgres";

-- ----------------------------
-- Records of invoice_medicines
-- ----------------------------
BEGIN;
INSERT INTO "public"."invoice_medicines" VALUES (1, 9);
INSERT INTO "public"."invoice_medicines" VALUES (1, 10);
INSERT INTO "public"."invoice_medicines" VALUES (2, 1);
INSERT INTO "public"."invoice_medicines" VALUES (2, 5);
INSERT INTO "public"."invoice_medicines" VALUES (3, 1);
INSERT INTO "public"."invoice_medicines" VALUES (3, 5);
INSERT INTO "public"."invoice_medicines" VALUES (4, 9);
INSERT INTO "public"."invoice_medicines" VALUES (4, 10);
INSERT INTO "public"."invoice_medicines" VALUES (5, 9);
INSERT INTO "public"."invoice_medicines" VALUES (5, 10);
INSERT INTO "public"."invoice_medicines" VALUES (6, 9);
INSERT INTO "public"."invoice_medicines" VALUES (6, 10);
INSERT INTO "public"."invoice_medicines" VALUES (7, 9);
INSERT INTO "public"."invoice_medicines" VALUES (7, 10);
COMMIT;

-- ----------------------------
-- Table structure for invoices
-- ----------------------------
DROP TABLE IF EXISTS "public"."invoices";
CREATE TABLE "public"."invoices" (
  "id" int8 NOT NULL DEFAULT nextval('invoices_id_seq'::regclass),
  "creation_date" text COLLATE "pg_catalog"."default",
  "total_pay" numeric,
  "promotion_id" int8
)
;
ALTER TABLE "public"."invoices" OWNER TO "postgres";

-- ----------------------------
-- Records of invoices
-- ----------------------------
BEGIN;
INSERT INTO "public"."invoices" VALUES (2, '2021-09-22', 123, 9);
INSERT INTO "public"."invoices" VALUES (3, '2021-09-23', 78.42308, 7);
INSERT INTO "public"."invoices" VALUES (7, '2021-09-23', 564.3, 11);
INSERT INTO "public"."invoices" VALUES (6, '2021-09-23', 564.3, 11);
INSERT INTO "public"."invoices" VALUES (5, '2021-09-23', 564.3, 11);
INSERT INTO "public"."invoices" VALUES (4, '2021-09-22', 564.3, 11);
INSERT INTO "public"."invoices" VALUES (1, '2021-09-21', 564.3, 11);
COMMIT;

-- ----------------------------
-- Table structure for medicines
-- ----------------------------
DROP TABLE IF EXISTS "public"."medicines";
CREATE TABLE "public"."medicines" (
  "id" int8 NOT NULL DEFAULT nextval('medicines_id_seq'::regclass),
  "name" text COLLATE "pg_catalog"."default",
  "price" numeric,
  "location" text COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."medicines" OWNER TO "postgres";

-- ----------------------------
-- Records of medicines
-- ----------------------------
BEGIN;
INSERT INTO "public"."medicines" VALUES (1, 'MD 01', 56.78, 'Mexico');
INSERT INTO "public"."medicines" VALUES (5, 'MD 02', 45.6, 'México');
INSERT INTO "public"."medicines" VALUES (6, 'MD 03', 3.4, 'ssss');
INSERT INTO "public"."medicines" VALUES (7, 'MD 04', 45.6, 'asdads');
INSERT INTO "public"."medicines" VALUES (8, 'MD 05', 56.7, 'asdad');
INSERT INTO "public"."medicines" VALUES (9, 'MD 06', 33, 'dddd');
INSERT INTO "public"."medicines" VALUES (10, 'MD 06', 34, 'adasd');
INSERT INTO "public"."medicines" VALUES (11, 'MD 08', 23, '123');
COMMIT;

-- ----------------------------
-- Table structure for promotions
-- ----------------------------
DROP TABLE IF EXISTS "public"."promotions";
CREATE TABLE "public"."promotions" (
  "id" int8 NOT NULL DEFAULT nextval('promotions_id_seq'::regclass),
  "description" text COLLATE "pg_catalog"."default",
  "percentage" numeric,
  "start_date" date,
  "finish_date" date,
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6)
)
;
ALTER TABLE "public"."promotions" OWNER TO "postgres";

-- ----------------------------
-- Records of promotions
-- ----------------------------
BEGIN;
INSERT INTO "public"."promotions" VALUES (7, 'prueba de promocion 01', 15, '2021-09-27', '2021-09-27', '2021-09-21 19:34:32.386994-05', '2021-09-21 19:34:32.386994-05');
INSERT INTO "public"."promotions" VALUES (8, 'prueba de promocion 02', 10, '2021-09-26', '2021-09-26', '2021-09-21 19:40:08.907003-05', '2021-09-21 19:40:08.907003-05');
INSERT INTO "public"."promotions" VALUES (9, 'prueba de promocion 03', 69, '2021-09-09', '2021-09-10', '2021-09-21 19:41:12.554086-05', '2021-09-21 19:41:12.554086-05');
INSERT INTO "public"."promotions" VALUES (11, 'prueba de promocion 04', 5, '2021-09-22', '2021-09-22', '2021-09-22 03:29:37.32507-05', '2021-09-22 03:29:37.32507-05');
INSERT INTO "public"."promotions" VALUES (12, 'prueba de promocion 05', 10, '2021-09-25', '2021-09-25', '2021-09-22 07:03:12.381584-05', '2021-09-22 07:03:12.381584-05');
COMMIT;

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."invoices_id_seq"
OWNED BY "public"."invoices"."id";
SELECT setval('"public"."invoices_id_seq"', 8, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."medicines_id_seq"
OWNED BY "public"."medicines"."id";
SELECT setval('"public"."medicines_id_seq"', 12, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."promotions_id_seq"
OWNED BY "public"."promotions"."id";
SELECT setval('"public"."promotions_id_seq"', 13, true);

-- ----------------------------
-- Primary Key structure for table invoice_medicines
-- ----------------------------
ALTER TABLE "public"."invoice_medicines" ADD CONSTRAINT "invoice_medicines_pkey" PRIMARY KEY ("invoice_id", "medicine_id");

-- ----------------------------
-- Primary Key structure for table invoices
-- ----------------------------
ALTER TABLE "public"."invoices" ADD CONSTRAINT "invoices_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table medicines
-- ----------------------------
ALTER TABLE "public"."medicines" ADD CONSTRAINT "medicines_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table promotions
-- ----------------------------
ALTER TABLE "public"."promotions" ADD CONSTRAINT "promotions_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table invoice_medicines
-- ----------------------------
ALTER TABLE "public"."invoice_medicines" ADD CONSTRAINT "fk_invoice_medicines_invoice" FOREIGN KEY ("invoice_id") REFERENCES "public"."invoices" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."invoice_medicines" ADD CONSTRAINT "fk_invoice_medicines_medicine" FOREIGN KEY ("medicine_id") REFERENCES "public"."medicines" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table invoices
-- ----------------------------
ALTER TABLE "public"."invoices" ADD CONSTRAINT "fk_invoices_promotion" FOREIGN KEY ("promotion_id") REFERENCES "public"."promotions" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
