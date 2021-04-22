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

func (l *liteDB) GenerateUrl(ctx context.Context, url *entities.Urls) error {
	stmt := `INSERT INTO urls (short_code, full_url, expiry, number_of_hits) VALUES (?, ?, ?, ?)`
	_, err := l.db.ExecContext(ctx, stmt, &url.ShortCode, &url.FullUrl, &url.Expiry, &url.NumberOfHits)
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

func (l *liteDB) UpdateUrl(ctx context.Context, url *entities.Urls) error {
	stmt := `UPDATE urls SET expiry = ?, number_of_hits = ? WHERE short_code = ?`
	_, err := l.db.ExecContext(ctx, stmt, &url.Expiry, &url.NumberOfHits, &url.ShortCode)
	if err != nil {
		return err
	}
	return nil
}

func (l *liteDB) DeleteUrl(ctx context.Context, shortCode string) error {
	stmt := `DELETE FROM urls WHERE short_code = ?`
	_, err := l.db.ExecContext(ctx, stmt, shortCode)
	if err != nil {
		return err
	}
	return nil
}

// ValidateUser returns true/false if match user_id AND password
func (l *liteDB) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT user_id FROM users WHERE user_id = ? AND password = ?`
	row := l.db.QueryRowContext(ctx, stmt, userID, pwd)
	u := &entities.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}
	return true
}
