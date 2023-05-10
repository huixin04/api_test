package purchases

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

// Company struct is a row record of the companies table in the invoice database
// Table struct is database table struct
type Table struct {
	//申請者
	Username string `gorm:"column:username;type:VARCHAR;" json:"username,omitempty"`
	// 請購編號
	PurchaseID string `gorm:"primaryKey;uuid_generate_v4();column:purchase_id;type:UUID;" json:"purchase_id,omitempty"`
	// 公司名稱
	CompanyName string `gorm:"column:company_name;type:VARCHAR;" json:"company_name,omitempty"`
	//請購部門
	Department string `gorm:"column:department;type:VARCHAR;" json:"department,omitempty"`
	//品名
	ProductName string `gorm:"column:product_name;type:VARCHAR;" json:"product_name,omitempty"`
	//數量
	PurchaseQuantity int `gorm:"column:purchase_quantity;type:integer;" json:"purchase_quantity,omitempty"`
	//價格
	Price int `gorm:"column:price;type:integer;" json:"price,omitempty"`
	// 創建時間
	Createdat time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at"`
	// 創建者
	Createdby string `gorm:"column:created_by;type:UUID;" json:"created_by,omitempty"`
	// 更新時間
	Updatedat *time.Time `gorm:"column:updated_at;type:TIMESTAMP;" json:"updated_at,omitempty"`
	// 更新者
	Updatedby *string `gorm:"column:updated_by;type:UUID;" json:"updated_by,omitempty"`
	// 是否刪除
	IsDeleted bool `gorm:"column:is_deleted;type:bool;false" json:"is_deleted,omitempty"`
}

// Base struct is corresponding to table structure file
type Base struct {
	//申請人
	Username string `json:"username,omitempty"`
	//請購編號
	PurchaseID string `json:"purchase_id,omitempty"`
	// 公司名稱
	CompanyName string `json:"company_name,omitempty"`
	//請購部門
	Department string `json:"department,omitempty"`
	//品名
	ProductName string ` json:"product_name,omitempty"`
	//數量
	PurchaseQuantity int ` json:"purchase_quantity,omitempty"`
	//價格
	Price int ` json:"price,omitempty"`
	// 創建時間
	Createdat time.Time `json:"created_at"`
	// 創建者
	Createdby string `json:"created_by,omitempty"`
	// 更新時間
	Updatedat *time.Time `  json:"updated_at,omitempty"`
	// 更新者
	Updatedby *string ` json:"updated_by,omitempty"`
	// 是否刪除
	IsDeleted bool `json:"is_deleted,omitempty"`
}

// Single return structure file
type Single struct {
	//申請人
	Username string `json:"username,omitempty"`
	//請購編號
	PurchaseID string `json:"purchase_id,omitempty"`
	// 公司名稱
	CompanyName string `json:"company_name,omitempty"`
	//請購部門
	Department string `json:"department,omitempty"`
	//品名
	ProductName string ` json:"product_name,omitempty"`
	//數量
	PurchaseQuantity int ` json:"purchase_quantity,omitempty"`
	//價格
	Price int ` json:"price,omitempty"`
	// 創建時間
	Createdat time.Time `json:"created_at"`
	// 創建者
	Createdby string `json:"created_by,omitempty"`
	// 更新時間
	Updatedat *time.Time `json:"updated_at,omitempty"`
	// 更新者
	Updatedby *string `json:"updated_by,omitempty"`
}

// Created struct is used to create
type Created struct {
	// 申請人
	Username string `json:"username,omitempty" binding:"required" validate:"required"`
	// 公司名稱
	CompanyName string `json:"company_name,omitempty" binding:"required" validate:"required"`
	//請購部門
	Department string `json:"department,omitempty" binding:"required" validate:"required"`
	//品名
	ProductName string ` json:"product_name,omitempty" binding:"required" validate:"required"`
	//數量
	PurchaseQuantity int ` json:"purchase_quantity,omitempty" binding:"required" validate:"required"`
	//價格
	Price int ` json:"price,omitempty" binding:"required" validate:"required"`
	// 創建者
	Createdby string `json:"created_by,omitempty" swaggerignore:"true"`
}

// Updated struct is used to update
type Updated struct {
	// 申請人
	Username string `json:"username,omitempty"`
	// 請購編號
	PurchaseID string `json:"purchase_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 公司名稱
	CompanyName string `json:"company_name,omitempty"`
	//請購部門
	Department string `json:"department,omitempty" `
	//品名
	ProductName string ` json:"product_name,omitempty"`
	//數量
	PurchaseQuantity int ` json:"purchase_quantity,omitempty" `
	//價格
	Price int ` json:"price,omitempty" `
	// 更新者
	Updatedby *string ` json:"updated_by,omitempty" swaggerignore:"true"`
	// 是否刪除
	IsDeleted bool `json:"is_deleted,omitempty" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	//申請人
	Username *string `json:"username,omitempty" form:"username"`
	//請購編號
	PurchaseID string `json:"purchase_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 公司名稱
	CompanyName *string `json:"company_name,omitempty"  form:"company_name"`
	//請購部門
	Department *string `json:"department,omitempty" form:"department"`
	//品名
	ProductName *string ` json:"product_name,omitempty" form:"product_name"`
	//數量
	PurchaseQuantity *int ` json:"purchase_quantity,omitempty" form:"purchase_quantity" `
	//價格
	Price *int ` json:"price,omitempty" form:"price" `
	// 是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty" swaggerignore:"true"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files
type List struct {
	Purchase []*struct {

		//申請人
		Username string `json:"username,omitempty" `
		//請購編號
		PurchaseID string `json:"purchase_id,omitempty"`
		// 公司名稱
		CompanyName string `json:"company_name,omitempty" `
		//請購部門
		Department string `json:"department,omitempty" `
		//品名
		ProductName string ` json:"product_name,omitempty"`
		//數量
		PurchaseQuantity int ` json:"purchase_quantity,omitempty"  `
		//價格
		Price int ` json:"price,omitempty"  `
		//創建時間
		CreatedAt time.Time `json:"created_at"`
		// 創建者
		Createdby string `json:"created_by,omitempty"`
	} `json:"purchases"`
	model.OutPage
}

// TableName sets the insert table name for this struct type
func (c *Table) TableName() string {
	return "purchase"
}
