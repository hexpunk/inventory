package db

import (
	"database/sql"
	"io/fs"
	"path"
	"regexp"

	_ "github.com/mattn/go-sqlite3"
)

func Migrate(db *sql.DB, fsys fs.FS) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version TEXT NOT NULL
	    );
	`)
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT version FROM schema_migrations ORDER BY version DESC LIMIT 1;")
	if err != nil {
		return err
	}

	var latest string
	for rows.Next() {
		if err := rows.Scan(&latest); err != nil {
			return err
		}
	}

	list, err := fs.Glob(fsys, "**/*.sql")
	if err != nil {
		return err
	}

	version := regexp.MustCompile("^[0-9]+")

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	for _, file := range list {
		if file > latest {
			content, err := fs.ReadFile(fsys, file)
			if err != nil {
				return err
			}

			_, err = tx.Exec(string(content[:]))
			if err != nil {
				tx.Rollback()

				return err
			}

			_, err = tx.Exec("INSERT INTO schema_migrations (version) VALUE (?);", version.FindString(path.Base(file)))
			if err != nil {
				tx.Rollback()

				return err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
