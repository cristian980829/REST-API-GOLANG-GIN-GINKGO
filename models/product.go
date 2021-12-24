package models

type Product struct {
	ID     				uint   	`json:"id" gorm:"primary_key"`
	Name  				string 	`json:"name"`
	Description 		string 	`json:"description"`
	Status				string 	`json:"status"`
	Creation_date		string 	`json:"creation_date"`
	Update_date			string 	`json:"update_date"`
	Account_id			string 	`json:"account_id"`
	Format_product		string 	`json:"format_product"`
	Value_unit			float32 `json:"value_unit"`
	Unit_name			string 	`json:"unit_name"`
	Unit_description 	string 	`json:"unit_description"`
	Stock				int32 	`json:"stock"`
}

