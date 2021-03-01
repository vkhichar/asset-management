CREATE TABLE assets(
        asset_id varchar(100) PRIMARY KEY,
        status_asset varchar(100) NOT NULL,
        category varchar (100) NOT NULL,
        purchase_date Date NOT NULL,
        inital_cost varchar(100) NOT NULL,
        asset_name varchar(100) NOT NULL,
        specification json NOT NULL
);