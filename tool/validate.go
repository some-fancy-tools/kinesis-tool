package tool

import (
	"errors"
	"fmt"
	"os"
)

func (s3 *S3) validate() error {
	if s3.Bucket == "" {
		return fmt.Errorf(ErrInvalidField, "bucket")
	}
	if s3.Key == "" {
		return fmt.Errorf(ErrInvalidField, "objectkey")
	}
	return nil
}

func (f *File) validate() error {
	s, err := os.Stat(f.Path)
	if os.IsNotExist(err) {
		return errors.New(ErrPathDoesNotExists)
	}
	if err != nil {
		return err
	}
	if s.IsDir() {
		return errors.New(ErrPathIsDir)
	}
	return nil
}

func (k *Kinesis) validate() error {
	if k.Stream == "" {
		return fmt.Errorf(ErrInvalidField, "stream")
	}
	if k.PartitionKey == "" {
		return fmt.Errorf(ErrInvalidField, "partitionkey")
	}
	return nil
}

// Validate function for validating tool configuration.
func (t *Tool) Validate() error {
	// File and S3 validation
	if t.File.Path == "" && t.S3.Bucket == "" && t.S3.Key == "" {
		return errors.New(ErrAtleaseOneSource)
	}
	if t.File.Path != "" {
		if err := t.File.validate(); err != nil {
			return err
		}
	} else {
		if err := t.S3.validate(); err != nil {
			return err
		}
	}
	if err := t.Kinesis.validate(); err != nil {
		return err
	}
	return nil
}
