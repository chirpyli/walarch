package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "walarch",
	Short: "postgresql wal archive to s3 tools",
	Long:  "tools for archive postgresql wal to s3",
}

func init() {
	rootCmd.PersistentFlags().String("endpoint", "", "S3 endpoint")
	rootCmd.PersistentFlags().String("access_key", "", "S3 Access Key")
	rootCmd.PersistentFlags().String("secret_key", "", "S3 Secret Key")
	rootCmd.PersistentFlags().String("region", "", "S3 region")
	rootCmd.PersistentFlags().String("bucket", "", "S3 bucket")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("execute failed: ", err)
	}
}
