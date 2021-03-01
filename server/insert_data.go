package server

import (
	"github.com/vkhichar/asset-management/config"
	"github.com/vkhichar/asset-management/repository"
	"database/sql"
	"github.com/jmoiron/sqlx"
)



func Insert_data(){
	err := config.Init()
	if err != nil {
		fmt.Printf("main: error while initialising config: %s", err.Error())
		return
	}
	repository.InitDB()

	var d *sqlx.DB

	d = repository.GetDB()

	tx:=d.MustBegin()
	tx.MustExec("Insert into users (name,email,password,is_admin) values ($1,$2,$3,$4)","Jan Doe","jandoe@gmail.com","12345",true)
	tx.MustExec("Insert into users (name,email,password,is_admin) values ($1,$2,$3,$4)","Alisa Ray","alisaray@gmail.com","hello",false)
	tx.MustExec("Insert into users (name,email,password,is_admin) values ($1,$2,$3,$4)","Tom Walters","tomwalters@gmail.com","tom123",false)
	tx.MustExec("Insert into users (name,email,password,is_admin) values ($1,$2,$3,$4)","Alice Stephen","alicestep@gmail.com","alice776",true)

	tx.MustExec("Insert into assets (asset_id,category,status_asset,purchase_date,initial_cost,asset_name,specification) values ($1,$2,$3,$4,$5,$6,$7)","LAP123","Laptop","Active","01/07/2017","50000","Dell Latitude E5550",'{"RAM":"4GB","HDD":"500GB","Generation":"i5"}')
	tx.MustExec("Insert into assets (asset_id,category,status_asset,purchase_date,initial_cost,asset_name,specification) values ($1,$2,$3,$4,$5,$6,$7)","LAP988","Laptop","Maintenance","01/10/2017","40000","IBM Thinkpad",'{"RAM":"4GB","HDD":"500GB","Generation":"i5"}')
	tx.MustExec("Insert into assets (asset_id,category,status_asset,purchase_date,initial_cost,asset_name,specification) values ($1,$2,$3,$4,$5,$6,$7)","LAP456","Laptop","Active","10/05/2016","45000","Lenovo Gaming Pad",'{"RAM":"8GB","HDD":"500GB","Generation":"i8"}')
	tx.MustExec("Insert into assets (asset_id,category,status_asset,purchase_date,initial_cost,asset_name,specification) values ($1,$2,$3,$4,$5,$6,$7)","LAP786","Laptop","Active","10/10/2015","60000","MacBook",'{"RAM":"4GB","HDD":"500GB","Colour":"Space Grey"}')
	tx.MustExec("Insert into assets (asset_id,category,status_asset,purchase_date,initial_cost,asset_name,specification) values ($1,$2,$3,$4,$5,$6,$7)","MOU110","Mouse","Active","11/08/2019","500","Intel",'{"Wireless":"Yes","Sensor":"Laser"}')
	tx.MustExec("Insert into assets (asset_id,category,status_asset,purchase_date,initial_cost,asset_name,specification) values ($1,$2,$3,$4,$5,$6,$7)","MOU1001","Mouse","Not Active","29/12/2016","450","Logitech",'{"Wireless":"No","Sensor":"Laser"}')
	tx.MustExec("Insert into assets (asset_id,category,status_asset,purchase_date,initial_cost,asset_name,specification) values ($1,$2,$3,$4,$5,$6,$7)","KEY457","Keyboard","Active","26/07/2019","750","HP",'{"Height":"34 mm","Width":"13 cms"}')
	tx.MustExec("Insert into assets (asset_id,category,status_asset,purchase_date,initial_cost,asset_name,specification) values ($1,$2,$3,$4,$5,$6,$7)","EARP987","Earphones","Active","22/05/2020","400","Boat",'{"Color":"Black"}')

	tx.MustExec("Insert into asset_allocations (asset_id,user_id,allocated_by,isactive,from_date) values ($1,$2,$3,$4,$5)","LAP123","2","1",true,"02/07/2017")
	tx.MustExec("Insert into asset_allocations (asset_id,user_id,allocated_by,isactive,from_date,to_date) values ($1,$2,$3,$4,$5,$6)","MOU1001","2","4",false,"02/07/2017","07/10/2019")
	tx.MustExec("Insert into asset_allocations (asset_id,user_id,allocated_by,isactive,from_date) values ($1,$2,$3,$4,$5)","LAP456","3","1",true,"10/07/2016")
	tx.MustExec("Insert into asset_allocations (asset_id,user_id,allocated_by,isactive,from_date) values ($1,$2,$3,$4,$5)","KEY457","3","4",true,"27/07/2019")
	tx.MustExec("Insert into asset_allocations (asset_id,user_id,allocated_by,isactive,from_date) values ($1,$2,$3,$4,$5)","EARP987","1","4",true,"01/06/2020")
	tx.MustExec("Insert into asset_allocations (asset_id,user_id,allocated_by,isactive,from_date) values ($1,$2,$3,$4,$5)","LAP786","4","1",true,"25/10/2015")
	tx.MustExec("Insert into asset_allocations (asset_id,user_id,allocated_by,isactive,from_date,to_date) values ($1,$2,$3,$4,$5,$6)","LAP988","2","4",true,"01/06/2020")

	tx.MustExec("Insert into maintenance_activities (asset_id,maintenance_description,cost,admission_date,discharge_date) values ($1,$2,$3,$4,$5)","KEY457","Enter Key not working","100","22/12/2019","29/12/2019")
	tx.MustExec("Insert into maintenance_activities (asset_id,maintenance_description,cost,admission_date,discharge_date) values ($1,$2,$3,$4,$5)","LAP786","Display damaged","1500","10/02/2017","20/02/2017")
	tx.MustExec("Insert into maintenance_activities (asset_id,maintenance_description,cost,admission_date) values ($1,$2,$3,$4)","LAP988","USB Port not working","1000","22/01/2021")
}