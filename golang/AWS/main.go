package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "fmt"
    "os"
)

// Lists the items in the specified S3 Bucket
func main() {
    if len(os.Args) != 2 {
        exitErrorf("Bucket name required\nUsage: %s bucket_name",
            os.Args[0])
    }

    bucket := os.Args[1]

    // Initialize a session in us-east-1 that the SDK will use to load
    // credentials from the shared credentials file ~/.aws/credentials.
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-east-1")},
    )

    // Create S3 service client
    svc := s3.New(sess)

    // Get the list of items
    resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
    if err != nil {
        exitErrorf("Unable to list items in bucket %q, %v", bucket, err)
    }

    for _, item := range resp.Contents {
        fmt.Println("Object:         ", *item.Key)
        fmt.Println("")
    }

    fmt.Println("Found", len(resp.Contents), "items in bucket", bucket)
    fmt.Println("")
}

func exitErrorf(msg string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, msg+"\n", args...)
    os.Exit(1)
}
