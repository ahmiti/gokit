package s3

import (
	"context"
	"io"
	"strings"
	"time"

	"github.com/ahmiti/gokit/storage"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func init() {
	storage.Register("s3", New)
}

type s3Storage struct {
	client *s3.Client
	bucket string
}

func New(ctx context.Context, dsn string) (storage.Storage, error) {
	// dsn format: bucket=my-bucket,region=us-east-1
	parts := strings.Split(dsn, ",")
	bucket := ""
	region := "us-east-1"
	for _, part := range parts {
		kv := strings.SplitN(part, "=", 2)
		if len(kv) == 2 {
			if kv[0] == "bucket" {
				bucket = kv[1]
			}
			if kv[0] == "region" {
				region = kv[1]
			}
		}
	}
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(cfg)
	return &s3Storage{client: client, bucket: bucket}, nil
}

func (s *s3Storage) Put(ctx context.Context, key string, reader io.Reader, contentType string) error {
	uploader := manager.NewUploader(s.client)
	_, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        reader,
		ContentType: aws.String(contentType),
	})
	return err
}

func (s *s3Storage) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	downloader := manager.NewDownloader(s.client)
	buf := &manager.WriteAtBuffer{}
	_, err := downloader.Download(ctx, buf, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return io.NopCloser(strings.NewReader(string(buf.Bytes()))), nil
}

func (s *s3Storage) Stat(ctx context.Context, key string) (*storage.ObjectInfo, error) {
	resp, err := s.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return &storage.ObjectInfo{
		Key:          key,
		Size:         resp.ContentLength,
		ContentType:  aws.ToString(resp.ContentType),
		ETag:         aws.ToString(resp.ETag),
		LastModified: aws.ToTime(resp.LastModified),
	}, nil
}

func (s *s3Storage) Delete(ctx context.Context, key string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	return err
}

func (s *s3Storage) SignedURL(ctx context.Context, key string, ttl time.Duration) (string, error) {
	presigner := s3.NewPresignClient(s.client)
	req, err := presigner.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(ttl))
	if err != nil {
		return "", err
	}
	return req.URL, nil
}

func (s *s3Storage) List(ctx context.Context, prefix string) ([]*storage.ObjectInfo, error) {
	resp, err := s.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucket),
		Prefix: aws.String(prefix),
	})
	if err != nil {
		return nil, err
	}
	objects := make([]*storage.ObjectInfo, 0, len(resp.Contents))
	for _, obj := range resp.Contents {
		objects = append(objects, &storage.ObjectInfo{
			Key:          aws.ToString(obj.Key),
			Size:         obj.Size,
			LastModified: aws.ToTime(obj.LastModified),
			ETag:         aws.ToString(obj.ETag),
		})
	}
	return objects, nil
}

func (s *s3Storage) Close() error {
	return nil
}
