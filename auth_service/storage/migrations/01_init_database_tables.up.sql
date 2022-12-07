CREATE TABLE "user" (
   "id" CHAR(36) PRIMARY KEY,
   "username" VARCHAR(255) UNIQUE NOT NULL,
   "password" VARCHAR(255) NOT NULL,
   "user_type" CHAR(36),
   "created_at" TIMESTAMP DEFAULT now(),
   "updated_at" TIMESTAMP,
   "deleted_at" TIMESTAMP
);

