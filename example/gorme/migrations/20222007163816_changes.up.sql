CREATE TABLE "users" ("id" bigserial,"first_name" text,"last_name" text,"birthday" timestamptz,"updated_at" timestamptz,"created_at" timestamptz,PRIMARY KEY ("id"));
CREATE TABLE "tasks" ("id" bigserial,"title" text,"note" text,"status" text,"updated_at" timestamptz,"created_at" timestamptz,PRIMARY KEY ("id"));
