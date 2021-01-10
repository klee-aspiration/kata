-- migrate:up

CREATE EXTENSION "uuid-ossp";

CREATE TABLE "user" (
    uuid uuid DEFAULT uuid_generate_v4(),
    email TEXT NOT NULL
);

-- migrate:down

DROP TABLE "user";
DROP EXTENSION "uuid-ossp";
