package storages

import (
	"context"
	"database/sql"
	"github.com/hoanbentley/URL-shortener/internal/entities"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type liteDB struct {
	db *sql.DB
}

func NewLiteDB() *liteDB {
	db, err := sql.Open("sqlite3", "./rabbit.db")
	if err != nil {
		log.Println("liteDB , connect db err:", err)
	} else {
		log.Println("liteDB , connect db success")
		return &liteDB{
			db: db,
		}
	}
	return nil
}

func (l *liteDB) GenerateUrl(ctx context.Context, t *entities.Urls) error {
	stmt := `INSERT INTO urls (short_code, full_url, expiry, number_of_hits) VALUES (?, ?, ?, ?)`
	_, err := l.db.ExecContext(ctx, stmt, &t.ShortCode, &t.FullUrl, &t.Expiry, &t.NumberOfHits)
	if err != nil {
		return err
	}

	return nil
}

func (l *liteDB) ListUrl(ctx context.Context) ([]*entities.Urls, error) {
	stmt := `SELECT short_code, full_url, expiry,number_of_hits FROM urls`
	rows, err := l.db.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []*entities.Urls
	for rows.Next() {
		url := &entities.Urls{}
		err := rows.Scan(&url.ShortCode, &url.FullUrl, &url.Expiry, &url.NumberOfHits)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

func (l *liteDB) GetUrl(ctx context.Context, shortCode string) (*entities.Urls, error) {
	stmt := `SELECT short_code, full_url, expiry,number_of_hits FROM urls WHERE short_code = ?`
	row := l.db.QueryRowContext(ctx, stmt, shortCode)

	url := &entities.Urls{}
	err := row.Scan(&url.ShortCode, &url.FullUrl, &url.Expiry, &url.NumberOfHits)
	if err != nil {
		return nil, err
	}

	return url, nil
}
