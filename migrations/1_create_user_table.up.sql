CREATE TABLE "dockertable" (
    "id" UUid DEFAULT gen_random_uuid() NOT NULL,
    "user_id" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("id")
);