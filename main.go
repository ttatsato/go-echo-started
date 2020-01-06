package main

import (
	"app/Controllers"
	"fmt"
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
 	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"time"
)

// gorm.Modelの定義
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Product struct {
	gorm.Model
	Name string `gorm:"default:'product_name'"`
	Code string
	Price uint
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", getUser)
	e.POST("/order", Controllers.SaveProduct)
	ConnectMySql()
	e.Logger.Fatal(e.Start(":1323"))

}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "id is " + id)
}

/**
 * @see http://gorm.io/ja_JP/docs
 */
func ConnectMySql () {
	db, err := gorm.Open("mysql", "root:@(localhost)/order_sass?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("データベースへの接続に失敗しました")
		panic("データベースへの接続に失敗しました")
	} else {
		fmt.Println("データベースへの接続に成功しました")
	}
	defer db.Close()

	// スキーマのマイグレーション
	db.AutoMigrate(&Product{})
	insertRecord := &Product{Code: "L1212", Price: 1000}
	// Create
	createNewProduct(db, insertRecord)

	// Read
	var product Product
	db.First(&product, 0) // idが1の製品を探します
	createNewProduct(db, insertRecord)
	// Update - 製品価格を2,000に更新します
	db.Model(&product).Update("Price", 2000)

	// Delete - 製品を削除します
	db.Delete(&product)
}

func createNewProduct (db *gorm.DB, insertRecord *Product) error {
	tx := db.Begin()
	// データベース操作をトランザクション内で行います（ここからは'db'ではなく'tx'を使います）
	var count = 0
	db.Model(&Product{}).Where("code = ?", insertRecord.Code).Count(&count)
	if count > 0 {
		// エラーを返した場合はロールバックされます
		tx.Rollback()
		return nil
	}

	if err := tx.Create(insertRecord).Error; err != nil {
		tx.Rollback()
		return err
	}

	// nilを返すとコミットされます
	return tx.Commit().Error
}

func (p *Product) BeforeSave() (err error) {
	fmt.Println("Model Product BeforeSave")
	return
}

func (p *Product) AfterCreate(scope *gorm.Scope) (err error) {
	fmt.Println("Model Product AfterCreate")
	fmt.Println("レコードを追加しました")
	return
}