package cmd

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var putCmd = &cobra.Command{
	Use:   "put",
	Short: "put a file to S3",
	Long:  "put a file upload to S3",
	Run:   runPut,
}

func init() {
	rootCmd.AddCommand(putCmd)
	putCmd.PersistentFlags().String("file", "", "file upload to S3")

}

func runPut(cmd *cobra.Command, args []string) {
	log.Println("run walarch put cmd.")

	accessKey, _ := cmd.Flags().GetString("access_key")
	secretKey, _ := cmd.Flags().GetString("secret_key")
	endpoint, _ := cmd.Flags().GetString("endpoint")
	region, _ := cmd.Flags().GetString("region")
	bucket, _ := cmd.Flags().GetString("bucket")
	file, _ := cmd.Flags().GetString("file")

	if accessKey == "" || secretKey == "" || endpoint == "" || region == "" || bucket == "" || file == "" {
		log.Fatalln("walarch args failed, please check:")
	}

	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(region),
		Endpoint:         aws.String(endpoint),
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		log.Fatalln("connect to S3 failed:", err)
	}

	s3client := s3.New(sess)

	f, err := os.Stat(file)
	if err != nil {
		log.Fatalln(err)
	}

	key := f.Name()

	value, err := os.ReadFile(file)
	if err != nil {
		log.Fatalln("read file failed:", err)
	}
	log.Println("prepare upload ", key, " value size ", len(value)/(1024*1024), " M")

	out, err := s3client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(value),
	})
	if err != nil {
		log.Fatalln("put object to S3 failed: ", err)
	}

	log.Println("put object to S3 success. ", out)
}
