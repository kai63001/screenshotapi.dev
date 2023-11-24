package lib

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadToS3(buf []byte, fileName string, bucketName string, accessKey string, secretKey string, endpoint string) error {
	log.Println("Uploading to S3...", fileName, bucketName, accessKey, secretKey, endpoint)
	ctx := context.Background()

	//buf to body
	body := bytes.NewReader(buf)

	endPointer := aws.Endpoint{}
	//check if endpoint is not have region like have only one dot
	if strings.Count(endpoint, ".") == 1 {
		endPointer = aws.Endpoint{
			URL: endpoint,
		}
	} else {
		endPointer = aws.Endpoint{
			URL:           endpoint,
			SigningRegion: getRegion(endpoint),
		}
	}

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return endPointer, nil
	})

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		//region
		config.WithEndpointResolverWithOptions(r2Resolver),
	)
	//auth
	if err != nil {
		return fmt.Errorf("configuration error: %w", err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	// Upload the file to S3
	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   body,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}

	fmt.Printf("Successfully uploaded %q to %q\n", fileName, bucketName)
	return nil
}

// function for get region from endpoint
func getRegion(endpoint string) string {
	//split string
	s := strings.Split(endpoint, ".")
	//get region
	region := s[1]
	return region
}
