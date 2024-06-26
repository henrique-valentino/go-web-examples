package postgres

import (
	"fmt"

	"example.com/goreddit"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PostStore struct {
	*sqlx.DB
}

func (s *PostStore) Post(id uuid.UUID) (goreddit.Post, error) {
	var p goreddit.Post
	if err := s.Get(&p, `SELECT * FROM posts WHERE id = $1`, id); err != nil {
		return goreddit.Post{}, fmt.Errorf(`error getting post: %w`, err)
	}
	return p, nil
}

func (s *PostStore) PostsByThread(threadId uuid.UUID) ([]goreddit.Post, error) {
	var pp []goreddit.Post
	if err := s.Get(&pp, `SELECT * FROM posts WHEHRE thread_id = $1`, threadId); err != nil {
		return []goreddit.Post{}, fmt.Errorf(`error getting posts: %w`, err)
	}
	return pp, nil
}

func (s *PostStore) CreatePost(p *goreddit.Post) error {
	if err := s.Get(p, `INSERT INTO posts VALUES ($1, $2, $3, $4, $5) RETURNING *`,
		p.ID,
		p.ThreadId,
		p.Title,
		p.Content,
		p.Votes,
	); err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}
	return nil
}

func (s *PostStore) UpdatePost(p *goreddit.Post) error {
	if err := s.Get(p, `UPDATE posts SET thread_id = $1,  title = $2, content = $3, votes = $4 WHERE id = $5 RETURNING *`,
		p.ThreadId,
		p.Title,
		p.Content,
		p.Votes,
		p.ID,
	); err != nil {
		return fmt.Errorf("error updating post: %w", err)
	}
	return nil
}

func (s *PostStore) DeletePost(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE posts WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error deleting post: %w", err)
	}
	return nil
}
