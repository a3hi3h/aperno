CREATE TABLE "questions" (
  "q_uuid" uuid PRIMARY KEY,
  "q_type" int NOT NULL,
  "q_level" int NOT NULL,
  "question" json NOT NULL,
  "created_at" timestamp
);

CREATE TABLE "exam" (
  "q_uuid" uuid PRIMARY KEY,
  "e_type" int,
  "e_level" int,
  "e_name" varchar,
  "q_list" json,
  "created_at" timestamp
);

CREATE TABLE "skill" (
  "s_uuid" uuid PRIMARY KEY,
  "skill_name" varchar,
  "skill_type" int,
  "exam_list" json,
  "created_at" timestamp
);

CREATE TABLE "job" (
  "j_uuid" uuid PRIMARY KEY,
  "job_name" varchar,
  "job_type" int,
  "job_status" int,
  "job_desc" json,
  "job_skill" json
);

CREATE TABLE "users" (
  "u_uuid" uuid PRIMARY KEY,
  "u_first_name" varchar not NULL,
  "u_last_name" varchar not NULL,
  "u_type" int NOT NULL,
  "u_org" uuid,
  "u_detail" json
);

CREATE TABLE "usersgroup" (
  "ug_uuid" uuid PRIMARY KEY,
  "ug_name" varchar,
  "ug_type" int,
  "ug_org" uuid,
  "ug_list" json
);


CREATE TABLE "organization" (
  "o_uuid" uuid PRIMARY KEY,
  "country_code" int,
  "org_name" varchar NOT NULL,
  "org_type" varchar,
  "org_cat" int,
  "job_list" json,
  "admin_id" uuid,
  "org_billing_details" json,
  "last_billing" uuid
);

CREATE TABLE "hiree" (
  "h_uuid" uuid PRIMARY KEY,
  "h_u_uuid" uuid,
  "h_details" json,
  "h_job_details" json,
  "h_skills" json
);

CREATE TABLE "billing" (
  "b_uuid" uuid PRIMARY KEY,
  "b_prev" uuid,
  "b_type" int,
  "b_created" timestamp,
  "b_paid" timestamp,
  "b_last_reminder" timestamp,
  "b_org" uuid,
  "b_detail" json,
  "b_paid_detail" json
);

CREATE INDEX ON "organization" ("org_name");

ALTER TABLE "exam" ADD FOREIGN KEY ("q_list") REFERENCES "questions" ("q_uuid");

ALTER TABLE "skill" ADD FOREIGN KEY ("exam_list") REFERENCES "exam" ("q_list");

ALTER TABLE "job" ADD FOREIGN KEY ("job_skill") REFERENCES "skill" ("s_uuid");

