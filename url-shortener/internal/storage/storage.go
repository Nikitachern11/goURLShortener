package storage

import (
	"errors"
)

var ErrURLNotFound = errors.New("url not found")

type Storage interface {
	Save(ctx context.Context, url string) (string, error)
	Get(ctx context.Context, shortURL string) (string, error)
}
