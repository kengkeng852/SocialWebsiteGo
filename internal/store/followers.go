package store

import (
	"context"
	"database/sql"
)

type Follower struct {
	UserID     int64  `json:"user_id"`
	FollowerID int64  `json:"follower_id"`
	CreatedAt  string `json:"created_at"`
}

type FollowerStore struct {
	db *sql.DB
}

func (s *FollowerStore) Follow(ctx context.Context, followerID, userID int64) error{
	query := `
		insert into followers (user_id, follower_id) values ($1,$2)
	`

	_, err := s.db.ExecContext(ctx, query, userID, followerID)
	return err
}

func (s *FollowerStore) UnFollow(ctx context.Context, followerID, userID int64) error{
	query := `
		delete from followers 
		where user_id = $1 and follower_id = $2
	`

	_, err := s.db.ExecContext(ctx, query, userID, followerID)
	return err
}