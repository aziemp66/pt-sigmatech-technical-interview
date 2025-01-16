CREATE TYPE "PaymentStatus" AS ENUM ('unpaid', 'paid');

CREATE TABLE "Customer" (
	"customer_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
	"NIK" varchar(16) UNIQUE NOT NULL,
	"full_name" varchar(100) NOT NULL,
	"legal_name" varchar(100),
	"birth_place" varchar(50),
	"birth_date" date,
	"salary" decimal(15, 2),
	"id_card_photo" text,
	"selfie_photo" text
);

CREATE TABLE "CreditLimit" (
	"limit_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
	"customer_id" uuid,
	"tenor" int NOT NULL,
	"limit_amount" decimal(15, 2) NOT NULL
);

CREATE TABLE "Transaction" (
	"transaction_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
	"customer_id" uuid,
	"contract_number" varchar(50) UNIQUE NOT NULL,
	"on_the_road_price" decimal(15, 2) NOT NULL,
	"admin_fee" decimal(15, 2),
	"installment_amount" decimal(15, 2),
	"interest_amount" decimal(15, 2),
	"asset_name" varchar(100),
	"transaction_date" timestamp NOT NULL
);

CREATE TABLE "Payment" (
	"payment_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
	"transaction_id" uuid,
	"installment_number" int NOT NULL,
	"payment_date" timestamp,
	"payment_due" timestamp NOT NULL,
	"payment_amount" decimal(15, 2) NOT NULL,
	"payment_status" PaymentStatus NOT NULL DEFAULT 'unpaid'
);

CREATE TABLE "ActivityLog" (
	"log_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
	"customer_id" uuid,
	"action" varchar(255) NOT NULL,
	"timestamp" timestamp DEFAULT (current_timestamp)
);

ALTER TABLE
	"CreditLimit"
ADD
	FOREIGN KEY ("customer_id") REFERENCES "Customer" ("customer_id");

ALTER TABLE
	"Transaction"
ADD
	FOREIGN KEY ("customer_id") REFERENCES "Customer" ("customer_id");

ALTER TABLE
	"Payment"
ADD
	FOREIGN KEY ("transaction_id") REFERENCES "Transaction" ("transaction_id");

ALTER TABLE
	"ActivityLog"
ADD
	FOREIGN KEY ("customer_id") REFERENCES "Customer" ("customer_id");