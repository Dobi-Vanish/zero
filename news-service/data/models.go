package data

import (
	"context"
	"database/sql"
	"github.com/lib/pq"
	"gopkg.in/reform.v1"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

type PostgresRepository struct {
	DB *reform.DB
}

func NewPostgresRepository(db *reform.DB) *PostgresRepository {
	return &PostgresRepository{
		DB: db,
	}
}

type NewsWithCategories struct {
	Id         int           `reform:"id"`
	Title      string        `reform:"title"`
	Content    string        `reform:"content"`
	Categories pq.Int64Array `reform:"categories"`
}

// GetNewsList get list of all news
func (u *PostgresRepository) GetNewsList(limit, offset int) ([]*NewsWithCategories, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        SELECT
            n.Id,
            n.Title,
            n.Content,
            array_agg(nc.CategoryId) AS Categories
        FROM
            News n
        LEFT JOIN
            NewsCategories nc ON n.Id = nc.NewsId
        GROUP BY
            n.Id, n.Title, n.Content
        ORDER BY
            n.Id DESC 
        LIMIT $1 OFFSET $2
    `

	rows, err := u.DB.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var news []*NewsWithCategories

	for rows.Next() {
		var newsDB NewsWithCategories
		err := rows.Scan(
			&newsDB.Id,
			&newsDB.Title,
			&newsDB.Content,
			&newsDB.Categories,
		)
		if err != nil {
			return nil, 0, err
		}

		news = append(news, &newsDB)
	}

	var total int
	countQuery := `SELECT COUNT(*) FROM News`
	err = u.DB.QueryRowContext(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return news, total, nil
}

// EditNews edits the news with provided info
func (u *PostgresRepository) EditNews(id int, title, content string, categories []int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx, err := u.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if title != "" || content != "" {
		query := `
            UPDATE News
            SET
                Title = COALESCE(NULLIF($1, ''), Title),
                Content = COALESCE(NULLIF($2, ''), Content)
            WHERE Id = $3
        `
		_, err = tx.ExecContext(ctx, query, title, content, id)
		if err != nil {
			return err
		}
	}

	if categories != nil {
		_, err = tx.ExecContext(ctx, "DELETE FROM NewsCategories WHERE NewsId = $1", id)
		if err != nil {
			return err
		}

		if len(categories) > 0 {
			insertQuery := `
                INSERT INTO NewsCategories (NewsId, CategoryId)
                VALUES ($1, $2)
            `
			for _, categoryId := range categories {
				_, err = tx.ExecContext(ctx, insertQuery, id, categoryId)
				if err != nil {
					return err
				}
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
