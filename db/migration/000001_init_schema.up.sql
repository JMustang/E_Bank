CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "accoubt_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_accoubt_id" bigint NOT NULL,
  "to_accoubt_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "entries" ("accoubt_id");

CREATE INDEX ON "transfers" ("from_accoubt_id");

CREATE INDEX ON "transfers" ("to_accoubt_id");

CREATE INDEX ON "transfers" ("from_accoubt_id", "to_accoubt_id");

COMMENT ON COLUMN "entries"."amount" IS 'Pode ser positivo ou negativo';

COMMENT ON COLUMN "transfers"."amount" IS 'Deve ser positivo';

ALTER TABLE "entries" ADD FOREIGN KEY ("accoubt_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_accoubt_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_accoubt_id") REFERENCES "accounts" ("id");