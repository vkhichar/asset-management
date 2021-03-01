CREATE TABLE maintenance_activities(
    SERIAL PRIMARY KEY(maintenance_id),
    FOREIGN KEY(asset_id),
    description varchar(100) NOT NULL,
    cost varchar(100) NOT NULL,
    admission_date Date NOT NULL,
    discharge_date DATE NOT NULL
); 