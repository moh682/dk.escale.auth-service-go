CREATE EXTENSION IF NOT EXISTS "pgcrypto";


/* **** TABLES **** */
CREATE TABLE "permissions"(
  "id" INTEGER NOT NULL PRIMARY KEY,
  "name" INTEGER NOT NULL PRIMARY KEY,
  "createdAt" DATE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "roles"(
    "uid" UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    "company_id" UUID NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "createdBy" INTEGER NOT NULL,
    "permission_id" INTEGER NOT NULL,
    "createdAt" DATE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "employees"(
    "uid" UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    "company_id" UUID NOT NULL,
    "role_id" INTEGER NOT NULL,
    "department_id" INTEGER NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "mobilePhone" VARCHAR(255) NULL,
    "createdAt" DATE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "companies"(
    "uid" UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    "cvr" INTEGER NOT NULL UNIQUE,
    "name" VARCHAR(255) NOT NULL,
    "createdAt" DATE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "roles_permissions_mapping"(
    "permission_id" INTEGER NOT NULL PRIMARY KEY,
    "role_id" INTEGER NOT NULL PRIMARY KEY,
    "createdBy" INTEGER NOT NULL,
    "createdAt" DATE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "departments"(
    "uid" UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    "location_id" INTEGER NOT NULL,
    "company_id" INTEGER NOT NULL,
    "createdAt" INTEGER NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "locations"(
    "uid" INTEGER NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    "city" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "country" VARCHAR(255) NOT NULL,
    "zipCode" INTEGER NOT NULL,
    "createdAt" INTEGER NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "companies_departments_mapping"(
    "company_id" UUID NOT NULL PRIMARY KEY,
    "deparment_id" UUID NOT NULL PRIMARY KEY
);

/* **** FOREIGN KEYS **** */
ALTER TABLE "roles" ADD CONSTRAINT "roles_company_id_foreign" FOREIGN KEY("company_id") REFERENCES "companies"("uid");
ALTER TABLE "employees" ADD CONSTRAINT "employees_company_id_foreign" FOREIGN KEY("company_id") REFERENCES "companies"("uid");
ALTER TABLE "employees" ADD CONSTRAINT "employees_role_id_foreign" FOREIGN KEY("role_id") REFERENCES "roles"("uid");
ALTER TABLE "employees" ADD CONSTRAINT "employees_department_id_foreign" FOREIGN KEY("department_id") REFERENCES "departments"("id");
ALTER TABLE "roles" ADD CONSTRAINT "roles_permission_id_foreign" FOREIGN KEY("permission_id") REFERENCES "roles_permissions_mapping"("role_id");
ALTER TABLE "roles" ADD CONSTRAINT "roles_createdby_foreign" FOREIGN KEY("createdBy") REFERENCES "employees"("uid");
ALTER TABLE "departments" ADD CONSTRAINT "departments_location_id_foreign" FOREIGN KEY("location_id") REFERENCES "locations"("uid");
ALTER TABLE "departments" ADD CONSTRAINT "departments_company_id_foreign" FOREIGN KEY("company_id") REFERENCES "companies"("uid");