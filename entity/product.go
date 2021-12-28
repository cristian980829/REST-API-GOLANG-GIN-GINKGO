package entity

type Product struct {
	ID            string `json:"id,omitempty" gorm:"primary_key;type:varchar(36)"`
	Name          string `json:"name" binding:"required" gorm:"type:varchar(45)"`
	Description   string `json:"description" gorm:"type:text(300)"`
	Status        string `json:"status" gorm:"type:varchar(45)"`
	Creation_date string `json:"creation_date" gorm:"type:varchar(50)"`
	Update_date   string `json:"update_date" gorm:"type:varchar(50)"`
	Account_id       string  `json:"account_id" gorm:"type:varchar(25)"`
	Format_product   string  `json:"format_product"`
	Value_unit       float32 `json:"value_unit" gorm:"type:real"`
	Unit_name        string  `json:"unit_name" gorm:"type:varchar(45)"`
	Unit_description string  `json:"unit_description" gorm:"type:text(300)"`
	Stock            int32   `json:"stock"`
}
