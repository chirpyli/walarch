package cmd

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get a file from S3",
	Long:  "get a file download from S3",
	Run:   runGet,
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().String("file", "", "file download from S3")
	getCmd.PersistentFlags().String("path", "", "path of file to restore")
}

func runGet(cmd *cobra.Command, args []string) {
	log.Println("run walarch get cmd")

	accessKey, _ := cmd.Flags().GetString("access_key")
	secretKey, _ := cmd.Flags().GetString("secret_key")
	endpoint, _ := cmd.Flags().GetString("endpoint")
	region, _ := cmd.Flags().GetString("region")
	bucket, _ := cmd.Flags().GetString("bucket")
	file, _ := cmd.Flags().GetString("file")
	path, _ := cmd.Flags().GetString("path")

	if accessKey == "" || secretKey == "" || endpoint == "" || region == "" || bucket == "" || file == "" || path == "" {
		log.Fatalln("walarch args failed, please check:")
	}

	sess, err1 := session.NewSession(&aws.Config{
		Region:           aws.String(region),
		Endpoint:         aws.String(endpoint),
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err1 != nil {
		log.Fatalln("connect to S3 failed:", err1)
	}

	s3client := s3.New(sess)

	if err := DownloadFileFromS3(s3client, bucket, file, path); err != nil {
		log.Fatalln("download file from S3 failed: ", err)
	}

	log.Println("get object %v success. put at %v", file, path)
}

func DownloadFileFromS3(s3client *s3.S3, bucketName string, objectKey string, filePathName string) error {
	result, err := s3client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Println("couldn't get object %v:%v. %v\n", bucketName, objectKey, err)
		return err
	}
	defer result.Body.Close()

	file, err := os.Create(filePathName)
	if err != nil {
		log.Println("couldn't create file %v: ", filePathName, err)
		return err
	}
	defer file.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		log.Println("couldn't read object body from %v:", objectKey, err)
	}

	_, err = file.Write(body)
	return err
}
