package aws

import (
	"bufio"
	"fmt"
	"testing"
)

func TestGetS3Object(t *testing.T) {
	t.Skip()
	reader, err := GetS3Data("some-bucket", "some-path/file.txt", true)

	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
