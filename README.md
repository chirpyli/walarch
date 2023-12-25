# walarch
PostgreSQL WAL archive to S3 tools

## 用法
tools for archive postgresql wal to s3

Usage:
walarch [command]

Available Commands:
get         get a file from S3
put         put a file to S3

Flags:
--access_key string   S3 Access Key
--bucket string       S3 bucket
--endpoint string     S3 endpoint
--region string       S3 region
--secret_key string   S3 Secret Key

### 上传文件到S3
put a file upload to S3

Usage:
walarch put [flags]

Flags:
--file string   file upload to S3

Global Flags:
--access_key string   S3 Access Key
--bucket string       S3 bucket
--endpoint string     S3 endpoint
--region string       S3 region
--secret_key string   S3 Secret Key

### 从S3下载文件
get a file download from S3

Usage:
walarch get [flags]

Flags:
--file string   file download from S3
--path string   path of file to restore

Global Flags:
--access_key string   S3 Access Key
--bucket string       S3 bucket
--endpoint string     S3 endpoint
--region string       S3 region
--secret_key string   S3 Secret Key

