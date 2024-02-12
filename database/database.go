package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go-tenancy/storage/models"
	"log"
	"os"
)

var source *sql.DB

var names = [3]string{"The", "Brush", "James"}

func Initialize() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(
		"DROP TABLE IF EXISTS tenancies"); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS tenancies (id VARCHAR(8) PRIMARY KEY, key VARCHAR(64), name VARCHAR(64))"); err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare("INSERT INTO tenancies(id, key, name) VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(statement)

	for i := 0; i < len(names); i++ {
		instance := models.Tenancy{}.Create(names[i])

		_, err := statement.Exec(instance.Id, instance.Key, instance.Data.Name)
		if err != nil {
			log.Fatal(err)
		}
	}

	source = db
}

func Tenancies(tenancies *[]*models.Tenancy) {
	rows, err := source.Query("SELECT id, key, name FROM tenancies")

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
		instance := models.Tenancy{}.Create("")

		if err := rows.Scan(&instance.Id, &instance.Key, &instance.Data.Name); err != nil {
			log.Fatal(err)
		}

		*tenancies = append(*tenancies, instance)
	}
}

func TenancyByKey(tenancies *[]*models.Tenancy, key string, value string) {
	statement, err := source.Prepare("SELECT id, key, name FROM tenancies WHERE " + key + " = $1 ORDER BY name")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := statement.Query(value)
	if err != nil {
		log.Fatal(err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	instance := models.Tenancy{}.Create("")

	for rows.Next() {
		if err := rows.Scan(&instance.Id, &instance.Key, &instance.Data.Name); err != nil {
			log.Fatal(err)
		}

		*tenancies = append(*tenancies, instance)
	}
}
