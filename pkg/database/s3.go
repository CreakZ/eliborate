package database

import (
	"yurii-lib/pkg/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
)

func InitS3Client() *s3.S3 {
	// Сессионный конфиг
	cfg := aws.Config{
		Credentials: credentials.NewStaticCredentials(
			viper.GetString(config.S3AccessKey),
			viper.GetString(config.S3SecretKey),
			"",
		),
		Endpoint: aws.String(viper.GetString(config.S3Endpoint)),
		Region:   aws.String("ru-1"),

		// Использование path-style адресации необходимо для доступа к стороннему Amazon-like S3 API
		S3ForcePathStyle: aws.Bool(true),
	}

	sess := session.Must(session.NewSession(&cfg))

	svc := s3.New(sess)

	_, err := svc.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(viper.GetString(config.S3ImageBucketName)),
	})
	if err != nil {
		switch err.Error() {
		case s3.ErrCodeNoSuchBucket:
			_, err = svc.CreateBucket(&s3.CreateBucketInput{
				Bucket: aws.String(viper.GetString(config.S3ImageBucketName)),
			})
			if err != nil {
				panic(err)
			}

			return svc
		default:
			panic(err)
		}
	}

	return svc
}
