CREATE TABLE "services" (
  "id" bigSerial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "popular_services" (
  "service_id" bigInt NOT NULL
);

ALTER TABLE "popular_services" ADD FOREIGN KEY ("service_id") REFERENCES "services" ("id");


INSERT INTO services (name) VALUES ('House Cleaning');
INSERT INTO services (name) VALUES ('Handyman');
INSERT INTO services (name) VALUES ('Lawn Moving and Trimming');
INSERT INTO services (name) VALUES ('Local Moving (under 50 miles)');
INSERT INTO services (name) VALUES ('Furniture Moving and Heavy Lifting');

INSERT INTO popular_services (service_id) VALUES (1);
INSERT INTO popular_services (service_id) VALUES (2);
INSERT INTO popular_services (service_id) VALUES (3);
INSERT INTO popular_services (service_id) VALUES (4);
INSERT INTO popular_services (service_id) VALUES (5);