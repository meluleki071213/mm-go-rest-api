package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("Failed to fetch comment by id")
	ErrNotImplemented = errors.New("Function not yet implemented")
)
type Comment struct {
	ID string
	Slug string
	Body string
	Author string
}

// This interface defines all of the methods that our service needs in order to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	UpdateComment(context.Context, string, Comment) (Comment, error)
}
/* Struct on which all our logic will be built on top of. */
type Service struct {
	Store Store
}

//NewService - returns a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service)GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment")
	comment, err := s.Store.GetComment(ctx, id)

	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return comment, nil
}

func (s *Service) UpdateComment(ctx context.Context,id string, updatedComment Comment) (Comment, error) {
	cmt, err := s.Store.UpdateComment(ctx, id, updatedComment)

	if err != nil {
		fmt.Println("error updating comment")
		return Comment{}, err
	}

	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return s.Store.DeleteComment(ctx, id)
}

func (s *Service) PostComment(ctx context.Context, comment Comment) (Comment, error) {
	insertedCmt, err := s.Store.PostComment(ctx, comment)

	if err != nil {
		return Comment{}, err
	}

	return insertedCmt, nil
}