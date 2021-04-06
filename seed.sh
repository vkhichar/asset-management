echo "Username: $1"
echo "Password: $2"
echo "DBName: $3" 

psql -U $1 $3 << EOF

insert into users (name,email,password,is_admin) values ('Jan Doe','jandoe@gmail.com','12345',true);
insert into users (name,email,password,is_admin) values ('Alisa Ray','alisaray@gmail.com','hello',false);
insert into users (name,email,password,is_admin) values ('Tom Walters','tomwalters@gmail.com','tom123',false);
insert into users (name,email,password,is_admin) values ('Alice Stephen','alicestep@gmail.com','alice776',true);

insert into assets (id,status,category,purchase_at,purchase_cost,name,specifications) values ('ffb4b1a4-7bf5-11eb-9439-0242ac130002','active','Laptop','01/07/2017',50000,'Dell Latitude E5550','{"RAM":"4GB","HDD":"500GB","Generation":"i5"}');
insert into assets (id,status,category,purchase_at,purchase_cost,name,specifications) values ('ffb4b488-7bf5-11eb-9439-0242ac130002','undermaintenance','Laptop','01/10/2017',40000,'IBM Thinkpad','{"RAM":"4GB","HDD":"500GB","Generation":"i5"}');
insert into assets (id,status,category,purchase_at,purchase_cost,name,specifications) values ('ffb4b6cc-7bf5-11eb-9439-0242ac130002','active','Laptop','10/05/2016',45000,'Lenovo Gaming Pad','{"RAM":"8GB","HDD":"500GB","Generation":"i8"}');
insert into assets (id,status,category,purchase_at,purchase_cost,name,specifications) values ('ffb4b898-7bf5-11eb-9439-0242ac130002','active','Laptop','10/10/2015',60000,'MacBook','{"RAM":"4GB","HDD":"500GB","Colour":"Space Grey"}');
insert into assets (id,status,category,purchase_at,purchase_cost,name,specifications) values ('ffb4b988-7bf5-11eb-9439-0242ac130002','active','Mouse','11/08/2019',500,'Intel','{"Wireless":"Yes","Sensor":"Laser"}');
insert into assets (id,status,category,purchase_at,purchase_cost,name,specifications) values ('ffb4ba50-7bf5-11eb-9439-0242ac130002','retired','Laptop','29/12/2016',450,'Logitech','{"Wireless":"No","Sensor":"Laser"}');
insert into assets (id,status,category,purchase_at,purchase_cost,name,specifications) values ('ffb4bb0e-7bf5-11eb-9439-0242ac130002','active','Keyboard','26/07/2019',750,'HP','{"Height":"34 mm","Width":"13 cms"}');
insert into assets (id,status,category,purchase_at,purchase_cost,name,specifications) values ('ffb4be1a-7bf5-11eb-9439-0242ac130002','active','Earphones','22/05/2020',400,'Boat','{"Color":"Black"}');

insert into asset_allocations (user_id,asset_id,allocated_by,allocated_from) values (2,'ffb4b488-7bf5-11eb-9439-0242ac130002',1,'02/07/2017');
insert into asset_allocations (user_id,asset_id,allocated_by,allocated_from,allocated_till) values (2,'ffb4b988-7bf5-11eb-9439-0242ac130002',4,'02/07/2017','07/10/2019');
insert into asset_allocations (user_id,asset_id,allocated_by,allocated_from) values (3,'ffb4bb0e-7bf5-11eb-9439-0242ac130002',1,'10/07/2016');
insert into asset_allocations (user_id,asset_id,allocated_by,allocated_from) values (3,'ffb4b1a4-7bf5-11eb-9439-0242ac130002',4,'27/07/2019');
insert into asset_allocations (user_id,asset_id,allocated_by,allocated_from) values (1,'ffb4ba50-7bf5-11eb-9439-0242ac130002',4,'01/06/2020');
insert into asset_allocations (user_id,asset_id,allocated_by,allocated_from) values (4,'ffb4be1a-7bf5-11eb-9439-0242ac130002',1,'25/10/2015');
insert into asset_allocations (user_id,asset_id,allocated_by,allocated_from) values (2,'ffb4bb0e-7bf5-11eb-9439-0242ac130002',4,'01/06/2020');

insert into maintenance_activities (asset_id,description,cost,started_at,ended_at) values ('ffb4bb0e-7bf5-11eb-9439-0242ac130002','Enter Key not working',100,'22/12/2019','29/12/2019');
insert into maintenance_activities (asset_id,description,cost,started_at,ended_at) values ('ffb4b6cc-7bf5-11eb-9439-0242ac130002','Display damaged',1500,'10/02/2017','20/02/2017');
insert into maintenance_activities (asset_id,description,cost,started_at) values ('ffb4b488-7bf5-11eb-9439-0242ac130002','USB Port not working',1000,'22/01/2021');

EOF