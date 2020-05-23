# Kinesis Tool

This tool helps you to put records in specified Kinesis Stream getting it from S3 Object.

# Introduction

It is using AWS SDK for Go for Kinesis and S3. This has some parallel downloads and PutRecords call to Kinesis.

It does have some limits built-in to avoid errors, e.g. Max of 500 records can be sent.

# Usage

Usage can be found by option `-h`.
```
> kinesis-tool -h
Usage of kinesis-tool:
  -compressed
        Should be true if file is compressed
  -debug
        Enable debug logs
  -file string
        Local File Path
  -kinesis-key string
        Kinesis Partition Key
  -kinesis-stream string
        Kinesis Stream Name
  -profile string
        Profile to be used for AWS (default "default")
  -region string
        Region to be used for AWS (default "us-east-1")
  -s3-bucket string
        S3 Bucket Name
  -s3-key string
        S3 Key Name
```

## Common Options

* `debug`: Enables the debug logs.
* `profile`: AWS Profile to use.
* `region`: AWS Region to use.
* `compressed`: It can accept GZIP compressed files as well, it decompresses and then puts the records.

## File Option

> In case we want to send something from local we can use this

* `file`: Path to file which can be read line by line and then send to Kinesis.

## S3 Options

* `s3-bucket`: As we can understand from the name it is the bucket name from where we want to use the file.
* `s3-key`: The object or key path with prefix.

## Kinesis Options

* `kinesis-stream`: Kinesis stream to put the records to.
* `kinesis-key`: Partition key by which we send the records.
