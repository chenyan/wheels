package r2

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/chenyan/wheels/flow"
)

const (
	Endpoint           = "https://%s.r2.cloudflarestorage.com"
	EnvVarAccountID    = "CLOUDFLARE_ACCOUNT_ID"
	EnvVarAPIKey       = "CLOUDFLARE_KEY_ID"
	EnvVarAPIKeySecret = "CLOUDFLARE_KEY_SECRET"
	EnvVarBucket       = "CLOUDFLARE_BUCKET"
)

type Client struct {
	client *s3.Client
	bucket string
}

func ClientFromEnv() *Client {
	accountID := flow.GetenvOrPanic(EnvVarAccountID)
	apiKey := flow.GetenvOrPanic(EnvVarAPIKey)
	apiKeySecret := flow.GetenvOrPanic(EnvVarAPIKeySecret)
	bucket := flow.GetenvOrPanic(EnvVarBucket)
	return &Client{
		client: s3.NewFromConfig(aws.Config{
			Region: "auto",
			Credentials: aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
				return aws.Credentials{
					AccessKeyID:     apiKey,
					SecretAccessKey: apiKeySecret,
				}, nil
			}),
			BaseEndpoint: aws.String(fmt.Sprintf(Endpoint, accountID)),
		}),
		bucket: bucket,
	}
}

func (c *Client) Put(ctx context.Context, key string, body io.Reader) error {
	_, err := c.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(key),
		Body:   body,
	})
	return err
}

func (c *Client) Get(ctx context.Context, key string) (io.Reader, error) {
	resp, err := c.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(key),
	})
	return resp.Body, err
}

func (c *Client) Delete(ctx context.Context, key string) error {
	_, err := c.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(key),
	})
	return err
}

func (c *Client) List(ctx context.Context, prefix string) ([]types.Object, error) {
	resp, err := c.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(c.bucket),
		Prefix: aws.String(prefix),
	})
	return resp.Contents, err
}
