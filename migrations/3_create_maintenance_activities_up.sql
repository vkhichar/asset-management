CREATE TABLE maintenance_activities(
    maintenance_id SERIAL PRIMARY KEY,
    asset_id varchar(100) NOT NULL,
   CONSTRAINT fk_asset
      FOREIGN KEY(asset_id) 
	  REFERENCES assets(asset_id),
    maintenance_description varchar(100) NOT NULL,
    cost varchar(100) NOT NULL,
    admission_date Date NOT NULL,
    discharge_date DATE NOT NULL
); 