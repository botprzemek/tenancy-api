package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go-tenancy/tenancy"
	"log"
	"os"
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

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS tenancies (id VARCHAR(8) PRIMARY KEY, name VARCHAR(64), key VARCHAR(64))"); err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare("INSERT INTO tenancies(id, name, key) VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(statement)

	var names = [3]string{"The", "Brush", "James"}

	for i := 0; i < len(names); i++ {
		instance := tenancy.Create()

		tenancy.SetName(instance, names[i])

		_, err := statement.Exec(instance.Id, instance.Data.Name, instance.Key)
		if err != nil {
			log.Fatal(err)
		}
	}

	source = db
}

func Tenancies(tenancies *[]*tenancy.Tenancy) {
	rows, err := source.Query("SELECT id, name, key FROM tenancies")

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

		if err := rows.Scan(&instance.Id, &instance.Data.Name, &instance.Key); err != nil {
			log.Fatal(err)
		}

		*tenancies = append(*tenancies, instance)
	}
}

func TenancyByKey(tenancies *[]*tenancy.Tenancy, key string, value string) {
	statement, err := source.Prepare("SELECT id, name, key FROM tenancies WHERE " + key + " = $1")

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

	instance := tenancy.Create()

	for rows.Next() {
		if err := rows.Scan(&instance.Id, &instance.Data.Name, &instance.Key); err != nil {
			log.Fatal(err)
		}

		*tenancies = append(*tenancies, instance)
	}
}
