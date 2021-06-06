package storage

import "context"

type Storage interface {
	Upload(ctx context.Context, file *File) error
}
