package postgres

import (
	"GoNews/pkg/storage"
	"context"
	"github.com/jackc/pgx/pgxpool"
)

// Store Хранилище данных.
type Store struct {
	db *pgxpool.Pool
}

// New Конструктор объекта хранилища.
func New(pgUri string) (*Store, error) {
	db, err := pgxpool.New(context.Background(), pgUri)
	if err != nil {
		return nil, err
	}
	s := Store{db: db}
	return &s, nil
}

func (s *Store) Posts() ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT
			id,
			author_id,
			title,
			content,
			created_at
		FROM posts;`,
	)
	if err != nil {
		return nil, err
	}

	var posts []storage.Post
	for rows.Next() {
		var p storage.Post
		err = rows.Scan(
			&p.ID,
			&p.AuthorID,
			&p.Title,
			&p.Content,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		posts = append(posts, p)
	}

	return posts, rows.Err()
}

func (s *Store) AddPost(p storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
		INSERT INTO posts (author_id, title, content, created_at)
		VALUES ($1, $2, $3, $4);
		`,
		p.AuthorID,
		p.Title,
		p.Content,
		p.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdatePost(p storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
		UPDATE posts
		SET title=$1, content=$2
		WHERE id=$3;`,
		p.Title,
		p.Content,
		p.ID,
	)

	if err != nil {
		return err
	}

	return nil
}
func (s *Store) DeletePost(p storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
			DELETE FROM posts
			WHERE id = $1;`,
		p.ID,
	)

	if err != nil {
		return err
	}
	return nil
}
