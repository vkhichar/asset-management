CREATE TABLE assets(
        PRIMARY KEY(asset_id),
        FOREIGN KEY (category),
        status_asset varchar(100) NOT NULL,
        purchase_date Date NOT NULL,
        inital_cost varchar(100) NOT NULL,
        asset_name varchar(100) NOT NULL,
        specification NOT NULL
);