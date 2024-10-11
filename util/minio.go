package util

import (
	"context"
	"log"

	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/lifecycle"
)

func saveExcellToMinio() {
	ctx := context.Background()
	endPoint := "localhost:9000"
	accessKey := "iK6guhVM24ZEtstT81vr"
	secretKey := "bsURvwA65wsfX5PaY3VQMbErNWqA2AlwhEJjapZ3"
	userSSL := false

	minioClient, err := minio.New(endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: userSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	bucketName := "testbucket"
	location := "us-east-1"
	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	config := lifecycle.NewConfiguration()
	config.Rules = []lifecycle.Rule{
		{
			ID:     "expire-bucket",
			Status: "Enabled",
			Expiration: lifecycle.Expiration{
				Days: 1,
			},
		},
	}

	err = minioClient.SetBucketLifecycle(ctx, bucketName, config)
	if err != nil {
		log.Fatalln(err)
	}

	objectName := "Book1.xlsx"
	filePath := "/home/khanhdt/Desktop/khanhdt/go-playground/Book1.xlsx"
	contentType := "application/octet-stream"
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}
