package storage

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/luispfcanales/online-shop/config"
)

type MysqlRepository struct {
	db *sql.DB
}

func New() *MysqlRepository {
	db, err := sql.Open("mysql", config.GetURLConnection())
	if err != nil {
		log.Fatalf("cannot connection db : %v", err)
	}
	return &MysqlRepository{
		db: db,
	}
}

func (r *MysqlRepository) Ping() error {
	defer r.db.Close()
	return r.db.Ping()
}
