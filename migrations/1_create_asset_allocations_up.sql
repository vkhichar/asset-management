    CREATE TABLE asset_allocations(
        asset_id varchar (100) ,
        user_id varchar (100),
        allocated_by varchar(100) NOT NULL,
        isactive boolean NOT NULL,
        from_date Date NOT NULL,
        to_date Date,
        PRIMARY KEY(asset_id,user_id) 
);