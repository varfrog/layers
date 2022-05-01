package storagemysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"layers/internal/storage"

	"github.com/go-sql-driver/mysql"
	pkgerrors "github.com/pkg/errors"
)

type Item struct {
	Key string `json:"key" validate:"required,max=255"`
	Val string `json:"val" validate:"required"`
}

type storageMysql struct {
	db *sql.DB
}

var _ storage.Facade = (*storageMysql)(nil) // verify interface compliance

func NewStorageMysql(db *sql.DB) storage.Facade {
	return &storageMysql{
		db: db,
	}
}

func (s *storageMysql) GetItemByKey(_ context.Context, key string) (storage.Item, error) {
	row := s.db.QueryRow(`
		SELECT k, v
		FROM app.items
		WHERE k = ?
	`, key)

	var result storage.Item
	err := row.Scan(&result.Key, &result.Val)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return storage.Item{}, storage.ErrNotFound
		}
		return storage.Item{}, fmt.Errorf("scan row: %v", err)
	}

	return result, nil
}

func (s *storageMysql) AddItem(_ context.Context, item storage.Item) error {
	query := `
		INSERT INTO app.items (k, v) VALUES (?, ?)
		ON DUPLICATE KEY UPDATE v=?
	`
	_, err := s.db.Exec(query, item.Key, item.Val, item.Val)
	return err
}

func Connect(dbname, user, pwd, host, port string) (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 user,
		Passwd:               pwd,
		Net:                  "tcp",
		Addr:                 host + ":" + port,
		DBName:               dbname,
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, pkgerrors.Wrapf(err, "db.Ping")
	}

	return db, nil
}
