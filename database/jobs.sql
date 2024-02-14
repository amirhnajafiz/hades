CREATE TABLE jobs (
    id         SERIAL PRIMARY KEY,
    manifest   NVARCHAR(16384),
    namespace  VARCHAR(1024),
    name       VARCHAR(1024),
    cronjob    VARCHAR(1024),
    retry      INTEGER,
    bound      INTEGER,
    created_at TIME,
    updated_at TIME,
    deleted_at TIME
);