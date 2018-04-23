package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testing"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type ProductOwn struct {
	ID    uint `gorm:"primary_key"`
	Code  string
	Price uint
}

func (ProductOwn) TableName() string {
	return "new_product"
}

func TestMain(t *testing.T) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	p := &Product{Code: "L1212", Price: 1000}
	db.Create(p)

	// Read
	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)
	t.Log(product.ID)

	// Delete - delete product
	// db.Delete(&product)

	db.Unscoped().Delete(&product)
}

func TestOwnModel(t *testing.T) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&ProductOwn{})

	p1 := &ProductOwn{Code: "L1", Price: 1000}
	db.Create(p1)

	p2 := &ProductOwn{Code: "L2", Price: 2000}
	db.Create(p2)

	db.Delete(p1)
	db.Delete(p2)
}

func (p *Product) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Code", "BeforeCreate")
	return nil
}

func TestPrefix(t *testing.T) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.LogMode(true)

	// TableName() is with higher priority.
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "prefix_" + defaultTableName
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	p := &Product{Code: "L1212", Price: 1000}
	db.Create(p)

	t.Log(p)
}
