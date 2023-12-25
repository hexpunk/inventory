package db

import (
	"database/sql"
	"embed"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var empty embed.FS

func TestMigrateNoMigrations(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()

	err := Migrate(db, empty)
	if err != nil {
		t.Errorf(err.Error())

		return
	}

	rows, _ := db.Query("SELECT count(*) FROM schema_migrations;")

	var count int
	for rows.Next() {
		rows.Scan(&count)
	}

	if count != 0 {
		t.Errorf("Expected 0 rows, received %v", count)
	}
}

//go:embed __fixtures__/001-test.sql
//go:embed __fixtures__/002-test.sql
//go:embed __fixtures__/003-test.sql
var firstThreeMigrations embed.FS

func TestMigrate(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()

	err := Migrate(db, firstThreeMigrations)
	if err != nil {
		t.Errorf(err.Error())

		return
	}

	rows, _ := db.Query(`SELECT count(*) = 3 FROM schema_migrations WHERE version IN ("001", "002", "003");`)

	var migrationsTableTest bool
	for rows.Next() {
		rows.Scan(&migrationsTableTest)
	}

	if !migrationsTableTest {
		t.Errorf("Test of schema_migrations failed")

		return
	}

	rows, _ = db.Query(`SELECT count(*) = 3 FROM sqlite_master WHERE name IN ("test_table1", "test_table2", "test_table3");`)

	var schemaTableTest bool
	for rows.Next() {
		rows.Scan(&schemaTableTest)
	}

	if !schemaTableTest {
		t.Errorf("Test of sqlite_master failed")
	}
}

//go:embed __fixtures__/001-test.sql
//go:embed __fixtures__/002-test.sql
//go:embed __fixtures__/003-test.sql
//go:embed __fixtures__/004-test.sql
var allMigrations embed.FS

func TestSecondMigration(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()

	err := Migrate(db, firstThreeMigrations)
	if err != nil {
		t.Errorf(err.Error())

		return
	}

	err = Migrate(db, allMigrations)
	if err != nil {
		t.Errorf(err.Error())

		return
	}

	rows, _ := db.Query(`SELECT count(*) = 4 FROM schema_migrations WHERE version IN ("001", "002", "003", "004");`)

	var migrationsTableTest bool
	for rows.Next() {
		rows.Scan(&migrationsTableTest)
	}

	if !migrationsTableTest {
		t.Errorf("Test of schema_migrations failed")

		return
	}

	rows, _ = db.Query(`SELECT count(*) = 0 FROM sqlite_master WHERE name IN ("test_table1", "test_table2", "test_table3");`)

	var schemaTableTest bool
	for rows.Next() {
		rows.Scan(&schemaTableTest)
	}

	if !schemaTableTest {
		t.Errorf("Test of sqlite_master failed")
	}
}
