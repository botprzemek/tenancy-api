package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go-tenancy/identifier"
	"go-tenancy/tenancy"
	"log"
	"os"
	"strconv"
)

var source *sql.DB

func Initialize() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(
		"DROP TABLE IF EXISTS tenancies"); err != nil {
		log.Fatal(err)
	}

	var query = fmt.Sprintf("CREATE TABLE IF NOT EXISTS tenancies (id VARCHAR(%v) PRIMARY KEY, name VARCHAR(64))", strconv.Itoa(int(identifier.Size())))
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(
		"INSERT INTO tenancies (id, name) VALUES ('d0f161f6', 'Cwen'), ('f43treg1', 'Brah') ON CONFLICT DO NOTHING"); err != nil {
		log.Fatal(err)
	}

	source = db
}

func Tenancies(tenancies *[]*tenancy.Tenancy) {
	rows, err := source.Query("SELECT id, name FROM tenancies")

	if err != nil {
		log.Fatal(err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	for rows.Next() {
		instance := tenancy.Create()

		if err := rows.Scan(&instance.Id, &instance.Name); err != nil {
			log.Fatal(err)
		}

		*tenancies = append(*tenancies, instance)
	}
}
