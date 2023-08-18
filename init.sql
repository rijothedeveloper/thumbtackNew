CREATE TABLE "services" (
  "id" bigSerial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "popular_services" (
  "service_id" bigInt NOT NULL
);

ALTER TABLE "popular_services" ADD FOREIGN KEY ("service_id") REFERENCES "services" ("id");
