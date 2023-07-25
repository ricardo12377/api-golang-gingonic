package config

import (
	"api/oportunities/schemas"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	dbLogger = GetLogger("sqlite")
)

func InitializarSqlite () (*gorm.DB, error) {

	dbPath := "./db/main.db"
	
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err){
		dbLogger.Info("database file is not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		dbLogger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(schemas.Opening{})
	if err != nil {
		dbLogger.Errorf("sqlite AutoMigrate error: %v", err)
		return nil, err
	}

	return db, nil
}