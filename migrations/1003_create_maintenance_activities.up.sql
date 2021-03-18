CREATE TABLE maintenance_activities(
    id SERIAL PRIMARY KEY,
    asset_id uuid NOT NULL,
   CONSTRAINT FK_ASSET
      FOREIGN KEY(asset_id) 
	  REFERENCES assets(id) 
      ON DELETE SET NULL,
    description varchar(100) NOT NULL,
    cost decimal(10,2) NOT NULL,
    started_at TIMESTAMP DEFAULT NOW(),
    ended_at timestamp
);