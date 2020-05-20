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
  -bucket string
        Bucket Name
  -compressed
        Should be true if Key is compressed
  -debug
        Enable debug logs
  -key string
        Key Name
  -partitionKey string
        Kinesis Partition Key
  -profile string
        Profile to be used for AWS (default "default")
  -region string
        Region to be used for AWS (default "us-east-1")
  -stream string
        Kinesis Stream Name
```

## Common Options

* debug
* profile
* region

## S3 Options

* bucket
* key
* `compressed`: It can accept GZIP compressed files as well, it decompresses and then puts the records.

## Kinesis Options

* stream
* partitionKey

