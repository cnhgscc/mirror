package gomysql

import (
	"database/sql"
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db sync.Map
)

func Init(c *Config) error {
	dtl := "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(dtl, c.User, c.Pass, c.Addr, c.DB)

	if c.Max == 0 {
		c.Max = 10
	}

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	sqlDB.SetMaxOpenConns(c.Max)
	sqlDB.SetMaxIdleConns(c.Max)

	ormDB, err := gorm.Open(
		mysql.New(mysql.Config{
			Conn: sqlDB,
		}),
		&gorm.Config{})

	if err != nil {
		return err
	}
	if c.Scope != "" {
		c.Scope = c.DB
	}
	db.Store(c.Scope, ormDB)
	return nil
}

func Use(name string) *gorm.DB {
	c, _ := db.Load(name)
	return c.(*gorm.DB)
}
