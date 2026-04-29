---
title: S3 / MinIO
icon: fa-cloud
primary: "#569A31"
lang: bash
locale: zhs
---

## fa-terminal AWS CLI S3 基础

```bash
aws configure
aws configure set region us-east-1

aws s3 ls
aws s3 ls s3://mybucket/
aws s3 ls s3://mybucket/prefix/ --recursive --human-readable --summarize

aws s3 mb s3://mybucket
aws s3 rb s3://mybucket --force

aws s3 cp file.txt s3://mybucket/
aws s3 cp s3://mybucket/file.txt ./
aws s3 sync ./local/ s3://mybucket/backup/
aws s3 sync s3://mybucket/backup/ ./local/ --delete
```

## fa-box 存储桶操作

```bash
aws s3api create-bucket --bucket mybucket --region us-east-1
aws s3api create-bucket --bucket mybucket --create-bucket-configuration LocationConstraint=eu-west-1

aws s3api list-buckets --query "Buckets[].Name"
aws s3api head-bucket --bucket mybucket
aws s3api delete-bucket --bucket mybucket

aws s3api get-bucket-location --bucket mybucket
aws s3api get-bucket-versioning --bucket mybucket
aws s3api get-bucket-encryption --bucket mybucket
aws s3api get-bucket-policy --bucket mybucket
```

## fa-upload 上传与下载

```bash
aws s3 cp data.csv s3://mybucket/data/
aws s3 cp data.csv s3://mybucket/data/ --storage-class GLACIER
aws s3 cp s3://mybucket/data/report.csv ./reports/
aws s3 sync ./src/ s3://mybucket/src/ --exclude "*.log" --include "*.py"

aws s3api put-object --bucket mybucket --key config.json --body config.json
aws s3api put-object --bucket mybucket --key notes.txt --body "hello world"

aws s3api get-object --bucket mybucket --key data.csv output.csv
aws s3 presign s3://mybucket/file.zip --expires-in 3600
```

## fa-link 预签名 URL

```bash
aws s3 presign s3://mybucket/file.zip --expires-in 3600

aws s3api generate-presigned-url --bucket mybucket --key upload.bin \
  --expires-in 3600 --method PUT

aws s3api generate-presigned-url --bucket mybucket --key file.zip \
  --expires-in 7200 --method GET
```

## fa-code-branch 版本控制

```bash
aws s3api put-bucket-versioning --bucket mybucket \
  --versioning-configuration Status=Enabled

aws s3api list-object-versions --bucket mybucket --prefix doc.pdf
aws s3api get-object --bucket mybucket --key doc.pdf --version-id VERSION_ID output.pdf

aws s3api copy-object --bucket mybucket --key doc.pdf \
  --copy-source mybucket/doc.pdf?versionId=VERSION_ID

aws s3api delete-object --bucket mybucket --key doc.pdf --version-id VERSION_ID
aws s3api delete-objects --bucket mybucket --delete file://delete.json
```

## fa-recycle 生命周期规则

```bash
aws s3api put-bucket-lifecycle-configuration --bucket mybucket \
  --lifecycle-configuration file://lifecycle.json

cat lifecycle.json
{
  "Rules": [{
    "ID": "archive-old",
    "Status": "Enabled",
    "Filter": { "Prefix": "logs/" },
    "Transitions": [{
      "Days": 30, "StorageClass": "STANDARD_IA"
    }, {
      "Days": 90, "StorageClass": "GLACIER"
    }],
    "Expiration": { "Days": 365 }
  }]
}

aws s3api get-bucket-lifecycle-configuration --bucket mybucket
aws s3api delete-bucket-lifecycle --bucket mybucket
```

## fa-layer-group 分片上传

```bash
aws s3api create-multipart-upload --bucket mybucket --key largefile.zip
aws s3api upload-part --bucket mybucket --key largefile.zip \
  --part-number 1 --upload-id UPLOAD_ID --body part1.bin
aws s3api list-parts --bucket mybucket --key largefile.zip --upload-id UPLOAD_ID
aws s3api complete-multipart-upload --bucket mybucket --key largefile.zip \
  --upload-id UPLOAD_ID --multipart-upload file://parts.json

aws s3api abort-multipart-upload --bucket mybucket --key largefile.zip \
  --upload-id UPLOAD_ID

aws s3 cp largefile.zip s3://mybucket/ --expected-size 10737418240
```

## fa-server MinIO 服务部署

```bash
minio server /data
minio server /data1 /data2 /data3 /data4
minio server http://node{1...4}/data{1...4}

export MINIO_ROOT_USER=admin
export MINIO_ROOT_PASSWORD=secretkey123
export MINIO_BROWSER=on
export MINIO_DOMAIN=minio.example.com
export MINIO_SERVER_URL="http://minio.example.com:9000"

minio server /data --console-address ":9001"
```

## fa-computer MinIO 客户端 (mc)

```bash
mc alias set local http://localhost:9000 admin secretkey123
mc alias set s3 https://s3.amazonaws.com $AWS_ACCESS_KEY_ID $AWS_SECRET_ACCESS_KEY

mc ls local/mybucket
mc mb local/newbucket
mc cp file.txt local/mybucket/
mc cp local/mybucket/file.txt ./
mc mirror ./local/ local/mybucket/backup/ --watch --overwrite
mc rm local/mybucket/file.txt
mc rm --recursive --force local/mybucket/prefix/

mc diff local/mybucket1 local/mybucket2
mc find local/mybucket --name "*.log" --older-than 30d
mc du local/mybucket
mc cat local/mybucket/config.json
```

## fa-shield-halved MinIO 存储桶策略

```bash
mc anonymous set download local/mybucket
mc anonymous set upload local/mybucket
mc anonymous set public local/mybucket
mc anonymous set none local/mybucket
mc anonymous get local/mybucket

mc admin policy create local readonly ./readonly-policy.json
mc admin policy attach local readonly --user appuser
mc admin user add local appuser secret123
mc admin user list local
mc admin user disable local appuser
```

## fa-hard-drive MinIO 纠删码

```bash
minio server /data{1...16}

minio server http://node{1...4}/data{1...4}

mc admin info local
mc admin config get local erasure
mc admin heal local --recursive
mc admin heal local/mybucket --scan deep

mc admin prometheus generate local
```

## fa-filter S3 Select

```bash
aws s3api select-object-content --bucket mybucket --key data.csv \
  --expression "SELECT * FROM s3object s WHERE s.status = 'active'" \
  --expression-type SQL --input-serialization '{"CSV": {}}' \
  --output-serialization '{"CSV": {}}' output.csv

aws s3api select-object-content --bucket mybucket --key data.json \
  --expression "SELECT s.name, s.age FROM s3object s WHERE s.age > 25" \
  --expression-type SQL \
  --input-serialization '{"JSON": {"Type": "DOCUMENT"}}' \
  --output-serialization '{"JSON": {}}' output.json
```

## fa-arrows-rotate 跨区域复制

```bash
aws s3api put-bucket-replication --bucket source-bucket \
  --replication-configuration file://replication.json

cat replication.json
{
  "Role": "arn:aws:iam::123456789012:role/s3-replication",
  "Rules": [{
    "ID": "replicate-all",
    "Status": "Enabled",
    "Filter": { "Prefix": "" },
    "Destination": {
      "Bucket": "arn:aws:s3:::dest-bucket",
      "StorageClass": "STANDARD"
    }
  }]
}

aws s3api get-bucket-replication --bucket source-bucket
aws s3api delete-bucket-replication --bucket source-bucket
```
