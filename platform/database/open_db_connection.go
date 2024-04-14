package database

import (
	"log"

	"github.com/golden-ocean/fiber-ocean/ent"
)

// OpenDBConnection func for opening database connection.
func OpenDBConnection() *ent.Client {
	client, err := MySQLConnection()
	if err != nil {
		log.Fatal("Mysql数据库连接出错!", err)
	}
	return client
}
