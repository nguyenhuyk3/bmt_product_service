CREATE TYPE "genres" AS ENUM (
  'action',
  'adventure',
  'animation',
  'comedy',
  'crime',
  'drama',
  'fantasy',
  'historical',
  'horror',
  'mystery',
  'romance',
  'sci_fi',
  'thriller',
  'war',
  'western',
  'dark_comedy',
  'documentary',
  'musical',
  'sports',
  'superhero',
  'psychological_thriller',
  'slasher',
  'biopic',
  'noir',
  'family'
);

CREATE TYPE "languages" AS ENUM (
  'vietnamese',
  'english'
);

CREATE TYPE "seat_types" AS ENUM (
  'vip',
  'standard'
);

CREATE TYPE "statuses" AS ENUM (
  'failed',
  'pending',
  'success'
);

CREATE TABLE "films" (
  "id" serial PRIMARY KEY NOT NULL,
  "title" text NOT NULL,
  "description" text NOT NULL,
  "release_date" date NOT NULL,
  "duration" interval NOT NULL
);

CREATE TABLE "fillm_changes" (
  "film_id" int NOT NULL,
  "changed_by" varchar(32) NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "film_genres" (
  "film_id" int,
  "genre" genres
);

CREATE TABLE "other_film_informations" (
  "film_id" int PRIMARY KEY,
  "status" statuses,
  "poster_url" text,
  "trailer_url" text
);

CREATE TABLE "cinemas" (
  "id" serial PRIMARY KEY NOT NULL,
  "name" text UNIQUE NOT NULL,
  "location" text NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "auditoriums" (
  "id" serial PRIMARY KEY,
  "cinema_id" int NOT NULL,
  "name" text NOT NULL,
  "seat_capacity" int NOT NULL DEFAULT 0,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "seats" (
  "id" serial PRIMARY KEY,
  "auditorium_id" int NOT NULL,
  "seat_type" seat_types NOT NULL,
  "seat_number" varchar(2) NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "showtimes" (
  "id" serial PRIMARY KEY,
  "film_id" int NOT NULL,
  "auditorium_id" int NOT NULL,
  "start_time" timestamp NOT NULL,
  "end_time" timestamp NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE INDEX ON "films" ("id");

CREATE UNIQUE INDEX ON "film_genres" ("film_id", "genre");

CREATE UNIQUE INDEX ON "other_film_informations" ("film_id");

CREATE INDEX ON "cinemas" ("id");

CREATE INDEX ON "auditoriums" ("id");

CREATE UNIQUE INDEX ON "seats" ("seat_type", "seat_number");

CREATE INDEX ON "showtimes" ("film_id");

CREATE INDEX ON "showtimes" ("auditorium_id");

ALTER TABLE "fillm_changes" ADD FOREIGN KEY ("film_id") REFERENCES "films" ("id");

ALTER TABLE "film_genres" ADD FOREIGN KEY ("film_id") REFERENCES "films" ("id");

ALTER TABLE "other_film_informations" ADD FOREIGN KEY ("film_id") REFERENCES "films" ("id");

ALTER TABLE "auditoriums" ADD FOREIGN KEY ("cinema_id") REFERENCES "cinemas" ("id");

ALTER TABLE "seats" ADD FOREIGN KEY ("auditorium_id") REFERENCES "auditoriums" ("id");

ALTER TABLE "showtimes" ADD FOREIGN KEY ("film_id") REFERENCES "films" ("id");

ALTER TABLE "showtimes" ADD FOREIGN KEY ("auditorium_id") REFERENCES "auditoriums" ("id");
