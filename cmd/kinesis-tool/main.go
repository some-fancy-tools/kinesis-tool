package main

import (
	"bufio"
	"flag"
	"log"

	"git.dcpri.me/some-fancy-tools/kinesis-tool/aws"
)

var (
	bucket, key  string
	compressed   bool
	stream       string
	partitionKey string
)

const (
	batchLimit = 1 * 1024 * 1024
)

func init() {
	// S3 Flags
	flag.StringVar(&bucket, "bucket", "", "Bucket Name")
	flag.StringVar(&key, "key", "", "Key Name")
	flag.BoolVar(&compressed, "compressed", false, "Should be true if Key is compressed")

	// Kinesis Flags
	flag.StringVar(&stream, "stream", "", "Kinesis Stream Name")
	flag.StringVar(&partitionKey, "partitionKey", "", "Kinesis Partition Key")

	flag.Parse()
}

func main() {

	reader, err := aws.GetS3Data(bucket, key, compressed)
	if err != nil {
		log.Println(err)
	}

	scanner := bufio.NewScanner(reader)

	// Initialize
	currentBufferSize := 0
	currentBatch := [][]byte{}

	for scanner.Scan() {
		// fmt.Printf("Scanning... %d\n", len(currentBatch))
		ibts := scanner.Bytes()
		if currentBufferSize = currentBufferSize + len(ibts); currentBufferSize > batchLimit {
			err := aws.PutKinesisRecords(currentBatch, stream, partitionKey)
			if err != nil {
				log.Println("Error occurred while doing PutKinesisRecords: ", err)
				break
			}
			// Cleanup for next batch
			currentBatch = [][]byte{}
			currentBufferSize = len(ibts)
		}

		currentBatch = append(currentBatch, ibts)
	}

	err = aws.PutKinesisRecords(currentBatch, stream, partitionKey)
	if err != nil {
		log.Println("Error occurred while doing PutKinesisRecords: ", err)
	}

	if scanner.Err() != nil {
		log.Println("Error occurred while scanning: ", err)
	}
}
