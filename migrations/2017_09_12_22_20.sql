CREATE TABLE "public"."items" (
    "id"      SERIAL PRIMARY KEY,
    "name"    VARCHAR(100) NOT NULL,
    "number"  VARCHAR(16) NOT NULL,
    "month"   VARCHAR(2) NOT NULL,
    "year"    VARCHAR(4) NOT NULL,
    "created" TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    "updated" TIMESTAMP WITH TIME ZONE,
    "enabled" BOOLEAN DEFAULT TRUE NOT NULL
);