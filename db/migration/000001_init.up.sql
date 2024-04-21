CREATE TABLE "bk_role" (
    "role_id" SERIAL PRIMA RY KEY,
    "role_name" VARCHAR NOT NULL
);

CREATE TABLE "bk_users" (
    "user_id" BIGSERIAL PRIMARY KEY,
    "user_email" VARCHAR NOT NULL UNIQUE,
    "user_phone_number" VARCHAR NOT NULL UNIQUE,
    "user_password" VARCHAR NOT NULL,
    "users_role_id" INTEGER NOT NULL REFERENCES "bk_role" ("role_id")
);

CREATE TABLE "bk_tokens" (
    "token_id" BIGSERIAL PRIMARY KEY,
    "token_value" VARCHAR NOT NULL,
    "token_user_id" BIGINT NOT NULL REFERENCES "bk_users" ("user_id"),
    "token_expiry_date" DATE
);

CREATE UNIQUE INDEX "user_email_index" ON "bk_users"("user_email");

CREATE UNIQUE INDEX "user_phone_number_index" ON "bk_users"("user_phone_number");

INSERT INTO
    bk_role(role_name)
VALUES
    ('business'),
    ('investor'),
    ('admin');