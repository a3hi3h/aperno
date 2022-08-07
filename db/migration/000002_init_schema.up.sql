CREATE TABLE "users" (
  "id" uuid PRIMARY KEY NOT NULL,
  "status" int NOT NULL,
  "type" int NOT NULL,
  "orgid" uuid,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "hashedpwd" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "last_modified" timestamp
);

CREATE TABLE "org" (
  "id" uuid PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "userid" uuid NOT NULL,
  "accountid" uuid
);

CREATE TABLE "questions" (
  "id" uuid PRIMARY KEY NOT NULL,
  "type" int NOT NULL,
  "level" int NOT NULL,
  "format" int NOT NULL,
  "userid" uuid NOT NULL,
  "orgid" uuid,
  "question" varchar NOT NULL,
  "question_time" int,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "last_modified" timestamp
);

CREATE TABLE "examround" (
  "id" uuid PRIMARY KEY NOT NULL,
  "type" int NOT NULL,
  "level" int NOT NULL,
  "time" int NOT NULL,
  "userid" uuid NOT NULL,
  "orgid" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "last_modified" timestamp
);

CREATE TABLE "accounts" (
  "id" uuid PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "status" varchar NOT NULL,
  "type" int not null,
  "userid" uuid NOT NULL,
  "orgid" uuid NOT NULL,
  "city" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "last_modified" timestamp
);

CREATE TABLE "usersessions" (
  "id" uuid PRIMARY KEY NOT NULL,
  "userid" uuid NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "expires_at" timestamp not null,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "users" ("email");

ALTER TABLE "users" ADD FOREIGN KEY ("orgid") REFERENCES "org" ("id");

ALTER TABLE "org" ADD FOREIGN KEY ("userid") REFERENCES "users" ("id");

ALTER TABLE "org" ADD FOREIGN KEY ("accountid") REFERENCES "accounts" ("id");

ALTER TABLE "questions" ADD FOREIGN KEY ("userid") REFERENCES "users" ("id");

ALTER TABLE "questions" ADD FOREIGN KEY ("orgid") REFERENCES "org" ("id");

ALTER TABLE "examRound" ADD FOREIGN KEY ("userid") REFERENCES "users" ("id");

ALTER TABLE "examRound" ADD FOREIGN KEY ("orgid") REFERENCES "org" ("id");

ALTER TABLE "accounts" ADD FOREIGN KEY ("userid") REFERENCES "users" ("id");

ALTER TABLE "accounts" ADD FOREIGN KEY ("orgid") REFERENCES "org" ("id");

ALTER TABLE "userSessions" ADD FOREIGN KEY ("userid") REFERENCES "users" ("id");
