package model

import (
	"database/sql"
	"fmt"
	"os"
	"sample/models"
)

var db *models.DB

// EstablishDB DBへの接続
func EstablishDB() (*sql.DB, error) {
	_db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")) + "?parseTime=true&loc=Asia%2FTokyo&charset=utf8mb4")
	if err != nil {
		return nil, fmt.Errorf("Failed In Connecting To Databases:%w", err)
	}
	db = models.NewDB(_db)

	err = db.Migrate()
	if err != nil {
		return nil, fmt.Errorf("failed to migrate: %w", err)
	}

	return _db, nil
}