//go:build testmode

package infrastructure

import (
	"fmt"

	"gorm.io/gorm"
	"log"
)

func ClearDB(db *gorm.DB) {
	var tables []string
	query := "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'"
	db.Raw(query).Scan(&tables)

	for _, table := range tables {
		db.Exec(fmt.Sprintf("DELETE FROM %s", table))
		log.Printf("Deleted data from table: %s\n", table)
	}
}
