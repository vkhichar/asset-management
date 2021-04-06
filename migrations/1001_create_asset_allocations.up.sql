    CREATE TABLE asset_allocations(
        id SERIAL PRIMARY KEY,
        user_id integer NOT NULL,
        asset_id uuid NOT NULL,
        allocated_by integer NOT NULL,
        allocated_from timestamp NOT NULL,
        allocated_till timestamp DEFAULT NULL,
        CONSTRAINT FK_ASSET
            FOREIGN KEY(asset_id) 
	        REFERENCES assets(id) 
            ON DELETE SET NULL,
        CONSTRAINT FK_USER
            FOREIGN KEY(user_id) 
	        REFERENCES users(id) 
            ON DELETE SET NULL,
        CONSTRAINT FK_ADMIN
            FOREIGN KEY(allocated_by) 
	        REFERENCES users(id) 
            ON DELETE SET NULL
    );