CREATE TABLE "follows" (
  "following_user_id" integer NOT NULL ,
  "followed_user_id" integer NOT NULL ,
  "created_at" timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "username" varchar NOT NULL ,
  "role" varchar NOT NULL ,
  "created_at" timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE "posts" (
  "id" integer PRIMARY KEY,
  "title" varchar NOT NULL ,
  "body" text NOT NULL ,
  "user_id" integer NOT NULL ,
  "status" varchar NOT NULL ,
  "created_at" timestamp NOT NULL DEFAULT NOW()
);

COMMENT ON COLUMN "posts"."body" IS 'Content of the post';

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "follows" ADD FOREIGN KEY ("following_user_id") REFERENCES "users" ("id");

ALTER TABLE "follows" ADD FOREIGN KEY ("followed_user_id") REFERENCES "users" ("id");
