CREATE TABLE "users"
(
    "id"         bigint            NOT NULL GENERATED BY DEFAULT AS IDENTITY,
    "first_name" character varying NOT NULL,
    "last_name"  character varying NOT NULL,
    "birthday"   date              NOT NULL,
    "updated_at" timestamptz       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_at" timestamptz       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);

CREATE TYPE "task_status" AS ENUM (
    'todo',
    'in_progress',
    'done'
);

CREATE TABLE "tasks"
(
    "id"         bigint            NOT NULL GENERATED BY DEFAULT AS IDENTITY,
    "user_id"    bigint NULL,
    "child_id"   bigint NULL,
    "title"      character varying NOT NULL,
    "note"       text              NOT NULL,
    "status"     task_status       NOT NULL,
    "updated_at" timestamptz       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_at" timestamptz       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id"),
    CONSTRAINT "task_users" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE SET NULL,
    CONSTRAINT "task_tasks" FOREIGN KEY ("child_id") REFERENCES "tasks" ("id") ON DELETE SET NULL
);

CREATE TYPE "project_status" AS ENUM (
    'todo',
    'in_progress',
    'done'
);

CREATE TABLE "projects"
(
    "id"          bigint            NOT NULL GENERATED BY DEFAULT AS IDENTITY,
    "title"       character varying NOT NULL,
    "description" text              NOT NULL,
    "status"      project_status    NOT NULL,
    "updated_at"  timestamptz       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_at"  timestamptz       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);

CREATE TABLE "project_tasks"
(
    "project_id" bigint      NOT NULL,
    "task_id"    bigint      NOT NULL,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("project_id", "task_id"),
    CONSTRAINT "project_task_projects" FOREIGN KEY ("project_id") REFERENCES "projects" ("id") ON DELETE SET NULL,
    CONSTRAINT "project_task_tasks" FOREIGN KEY ("task_id") REFERENCES "tasks" ("id") ON DELETE SET NULL
);
