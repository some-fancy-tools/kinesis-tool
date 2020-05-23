package tool

import (
	"log"

	"git.dcpri.me/some-fancy-tools/kinesis-tool/aws"
)

func push(debug bool, currentBatch [][]byte, stream, partitionKey string) error {
	if debug {
		log.Printf("Pushing Records: %d\n", len(currentBatch))
	}

	if err := aws.PutKinesisRecords(currentBatch, stream, partitionKey); err != nil {
		log.Println("Error occurred while doing PutKinesisRecords: ", err)
		return err
	}
	return nil
}
