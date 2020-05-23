package main

import (
	"flag"
	"log"

	"git.dcpri.me/some-fancy-tools/kinesis-tool/tool"
)

var (
	bucket, key  string
	compressed   bool
	stream       string
	partitionKey string
	debug        bool
	filepath     string
)

func init() {
	// Generic Flags
	flag.BoolVar(&debug, "debug", false, "Enable debug logs")
	flag.BoolVar(&compressed, "compressed", false, "Should be true if file is compressed")

	// S3 Flags
	flag.StringVar(&bucket, "s3-bucket", "", "S3 Bucket Name")
	flag.StringVar(&key, "s3-key", "", "S3 Key Name")

	// File Flags
	flag.StringVar(&filepath, "file", "", "Local File Path")

	// Kinesis Flags
	flag.StringVar(&stream, "kinesis-stream", "", "Kinesis Stream Name")
	flag.StringVar(&partitionKey, "kinesis-key", "", "Kinesis Partition Key")

	flag.Parse()
}

func main() {

	s3 := tool.S3{
		Bucket: bucket,
		Key:    key,
	}

	kinesis := tool.Kinesis{
		PartitionKey: partitionKey,
		Stream:       stream,
	}

	f := tool.File{
		Path: filepath,
	}

	t := tool.Tool{
		Compressed: compressed,
		Debug:      debug,
		File:       f,
		S3:         s3,
		Kinesis:    kinesis,
	}

	err := t.Validate()
	if err != nil {
		log.Fatal("Error while validating configuration: ", err)
	}
	err = t.Run()
	if err != nil {
		log.Println(err)
	}
}
