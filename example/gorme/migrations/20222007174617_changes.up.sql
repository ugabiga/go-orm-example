CREATE TABLE
  "users" (
    "id" bigserial,
    "first_name" text,
    "last_name" text,
    "birthday" timestamptz,
    "updated_at" timestamptz,
    "created_at" timestamptz,
    PRIMARY KEY ("id")
  );

CREATE TABLE
  "tasks" (
    "id" bigserial,
    "user_id" bigint,
    "title" text,
    "note" text,
    "status" text DEFAULT 'todo',
    "updated_at" timestamptz,
    "created_at" timestamptz,
    PRIMARY KEY ("id"),
    CONSTRAINT "fk_users_tasks" FOREIGN KEY ("user_id") REFERENCES "users"("id")
  );

CREATE TABLE
  "projects" (
    "id" bigserial,
    "title" text,
    "description" text,
    "status" text DEFAULT 'todo',
    "updated_at" timestamptz,
    "created_at" timestamptz,
    PRIMARY KEY ("id")
  );

CREATE TABLE
  "project_tasks" (
    "project_id" bigint,
    "task_id" bigint,
    PRIMARY KEY ("project_id", "task_id"),
    CONSTRAINT "fk_project_tasks_project" FOREIGN KEY ("project_id") REFERENCES "projects"("id"),
    CONSTRAINT "fk_project_tasks_task" FOREIGN KEY ("task_id") REFERENCES "tasks"("id")
  );