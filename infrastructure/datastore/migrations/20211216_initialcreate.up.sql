BEGIN;

CREATE TABLE organization (
    id uuid PRIMARY KEY NOT NULL,
    name varchar(100) NOT NULL,
    brewery varchar(100) NOT NULL,
    country varchar(100) NOT NULL,
    price  bigint,
    currency varchar(100) NOT NULL,    
);
COMMIT;
