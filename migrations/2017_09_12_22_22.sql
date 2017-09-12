CREATE TABLE "public"."sessions" (
    "id"      SERIAL PRIMARY KEY,
    "token"   VARCHAR(40) NOT NULL,
    "created" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    "enabled" BOOLEAN DEFAULT TRUE NOT NULL
);