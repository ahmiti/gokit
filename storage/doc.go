// Package storage provides a uniform interface for object storage.
//
// Supported drivers: s3, gcs, azureblob, local.
//
// Example:
//
//   import _ "github.com/ahmiti/gokit/storage/s3"
//
//   s, err := storage.Open(ctx, "s3", "bucket=my-bucket,region=us-east-1")
//   err = s.Put(ctx, "key.txt", bytes.NewReader(data), "text/plain")
//   reader, err := s.Get(ctx, "key.txt")
package storage
