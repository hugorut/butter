package filesystem

import (
	"butter/sys"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

// Return a new filesystem from the OS configuration
func NewFileSystem() FileSystem {
	switch os.Getenv("FILE_PROVIDER") {
	case "s3":
		return NewS3FileSystem(
			sys.EnvOrDefault("S3_REGION", "eu-west-1"),
			sys.EnvOrDefault("S3_BUCKET", "butter"),
			&credentials.EnvProvider{},
		)
	case "os":
		return NewOSFileSystem()
	}

	return NewOSFileSystem()
}

// NewS3FileSystem is a construct function takes both the region, bucket, and credential provider of your s3 filesystem.
// the region and bucket parameters are self explanatory and represent configuration that you can find in you aws dashboard
// the final argument, the aws provider, this is a struct which is in charge of getting your aws credentials
// it is recommended to use the aws.EnvProvider with the filesystem
func NewS3FileSystem(region, bucket string, provider credentials.Provider) *S3FileSystem {
	return &S3FileSystem{
		bucket,
		&aws.Config{
			Region:      aws.String(region),
			Credentials: credentials.NewCredentials(provider),
		},
		new(S3Call),
		new(sys.OSTime),
	}
}

// NewOSFileSystem is a construct function that returns a pointer to a OSFileSystem.
func NewOSFileSystem() *OSFileSystem {
	return &OSFileSystem{
		&osFS{},
	}
}
