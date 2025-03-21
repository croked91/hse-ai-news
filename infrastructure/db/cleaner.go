package db

import (
	"database/sql"
	"testing"
)

// MustClean очищает указанные таблицы в базе данных: news, n_last_ctx, compressed_ctx
func MustClean(t *testing.T, db *sql.DB) {
	tables := []string{"news", "n_last_ctx", "compressed_ctx", "session"}

	t.Cleanup(func() {
		for _, table := range tables {
			query := "DELETE FROM " + table
			_, err := db.Exec(query)
			if err != nil {
				panic(err)
			}
		}

		db.Close()
	})

}
