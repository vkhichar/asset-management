CREATE TYPE status AS ENUM ('active', 'retired','undermaintenance');


CREATE TABLE assets(
        id uuid DEFAULT uuid_generate_v4 (),
        status status,
        category varchar (100) NOT NULL,
        purchase_at timestamp NOT NULL,
        purchase_cost decimal(10,2) NOT NULL,
        name varchar(100) NOT NULL,
        specifications json NOT NULL,
        PRIMARY KEY(id)
);


