package aws

import (
	"bufio"
	"compress/gzip"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/s3"
)

// GetS3Data for getting s3 data for given bucket and key
// Supported Compression is gzip
func GetS3Data(bucket, key string, compressed bool) (io.Reader, error) {
	out, err := s3svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	// defer out.Body.Close()
	// Read All bytes
	var r io.Reader
	if compressed {
		r, err = gzip.NewReader(out.Body)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
	return bufio.NewReader(out.Body), nil
}

// PutKinesisRecords for putting records
func PutKinesisRecords(records [][]byte, stream, partitionKey string) error {
	// request structure
	input := &kinesis.PutRecordsInput{
		StreamName: aws.String(stream),
	}

	krecs := make([]*kinesis.PutRecordsRequestEntry, len(records))

	for i, record := range records {
		krecs[i] = &kinesis.PutRecordsRequestEntry{
			Data:         record,
			PartitionKey: aws.String(partitionKey),
		}
	}

	input.SetRecords(krecs)
	_, err := kinesissvc.PutRecords(input)
	if err != nil {
		return err
	}
	// TODO: return some stats like failed records if any
	return nil
}
