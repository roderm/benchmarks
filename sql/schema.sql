CREATE TABLE IF NOT EXISTS "company" (
    "id" UUID DEFAULT gen_random_uuid(),
    "name" VARCHAR(64) NOT NULL,
    "branch" VARCHAR(64) NOT NULL,
    "url" VARCHAR(128),
    "founded" DATE,
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "employee" (
    "id" UUID DEFAULT gen_random_uuid(),
    "company_id" UUID NOT NULL,
    "firstname" VARCHAR(64),
    "lastname" VARCHAR(64), 
    "email" VARCHAR(64), 
    "birthdate" DATE,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("company_id") REFERENCES "company"("id")
);

CREATE TABLE IF NOT EXISTS "product" (
    "id" UUID DEFAULT gen_random_uuid(),
    "company_id" UUID NOT NULL,
    "name" VARCHAR(128),
    "prod_type" VARCHAR(64),
    "manufactered" INT,
    "sold" INT,
    "price" FLOAT,
    "released" DATE,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("company_id") REFERENCES "company"("id")
);