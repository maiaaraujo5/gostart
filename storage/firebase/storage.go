package firebase

import (
	"bytes"
	bucket "cloud.google.com/go/storage"
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/google/uuid"
	"github.com/maiaaraujo5/gostart/storage"
	"google.golang.org/api/option"
	"io"
)

type Storage struct {
	Storage *bucket.BucketHandle
	Options *Config
}

func NewStorage(ctx context.Context) (storage.Storage, error) {
	options, err := NewConfig()
	if err != nil {
		return nil, err
	}

	config := &firebase.Config{
		StorageBucket: options.StorageBucket,
	}

	firebaseOptions := option.WithCredentialsFile(options.PathToCredentialsFile)

	app, err := firebase.NewApp(ctx, config, firebaseOptions)
	if err != nil {
		return nil, err
	}

	client, err := app.Storage(ctx)
	if err != nil {
		return nil, err
	}

	b, err := client.DefaultBucket()
	if err != nil {
		return nil, err
	}

	return &Storage{
		Storage: b,
		Options: options,
	}, nil
}

func (s *Storage) Upload(ctx context.Context, file *storage.File) error {

	metadata := make(map[string]string)
	metadata["firebaseStorageDownloadTokens"] = uuid.NewString()

	ctx, cancel := context.WithTimeout(ctx, s.Options.UploadTimeout)
	defer cancel()

	object := s.Storage.Object(file.Name)
	writer := object.NewWriter(ctx)

	for key, value := range file.Metadata {
		metadata[key] = value
	}

	writer.ObjectAttrs.Metadata = metadata

	if _, err := io.Copy(writer, bytes.NewReader(file.Content)); err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	if err := object.ACL().Set(ctx, bucket.AllUsers, bucket.RoleReader); err != nil {
		return err
	}

	return nil
}
