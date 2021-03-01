    CREATE TABLE asset_allocations(
        id SERIAL PRIMARY KEY,
        allocated_by varchar(100) NOT NULL,
        isactive boolean NOT NULL,
        FOREIGN KEY(user_id),
        from_date Date NOT NULL,
        to_date Date 
);