package tool

import (
	"bufio"
	"io"
	"os"

	"git.dcpri.me/some-fancy-tools/kinesis-tool/aws"
)

const (
	batchLimit = 1024 * 1024
)

// Tool for main tool configurations
type Tool struct {
	Debug      bool
	Compressed bool
	File       File
	S3         S3
	Kinesis    Kinesis
}

// S3 for Bucket and Key
type S3 struct {
	Bucket string
	Key    string
}

// Kinesis for kinesis details
type Kinesis struct {
	Stream       string
	PartitionKey string
}

type File struct {
	Path string
}

// Run function for running the actual tool.
func (t *Tool) Run() error {
	var reader io.Reader
	var err error
	if t.File.Path != "" {
		reader, err = os.OpenFile(t.File.Path, os.O_RDONLY, 0644)
		if err != nil {
			return err
		}
	} else {
		reader, err = aws.GetS3Data(t.S3.Bucket, t.S3.Key, t.Compressed)
		if err != nil {
			return err
		}
	}

	scanner := bufio.NewScanner(reader)

	// Initialize
	currentBufferSize := 0
	currentBatch := [][]byte{}

	for scanner.Scan() {
		// fmt.Printf("Scanning... %d\n", len(currentBatch))
		scannedBytes := scanner.Bytes()
		if currentBufferSize = currentBufferSize + len(scannedBytes); currentBufferSize > batchLimit || len(currentBatch) == 500 {
			if err := push(t.Debug, currentBatch, t.Kinesis.Stream, t.Kinesis.PartitionKey); err != nil {
				return err
			}

			// Cleanup for next batch
			currentBatch = [][]byte{}
			currentBufferSize = len(scannedBytes)
		}

		currentBatch = append(currentBatch, scannedBytes)
	}
	if err := scanner.Err(); err != nil {
		return err
		// log.Println("Error occurred while scanning: ", err)
	}

	return push(t.Debug, currentBatch, t.Kinesis.Stream, t.Kinesis.PartitionKey)
}
