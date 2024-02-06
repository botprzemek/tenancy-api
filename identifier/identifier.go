package identifier

import (
	"crypto/rand"
	"fmt"
	"io"
)

const size int8 = 8

func Size() int8 {
	return size
}

func Create() (*[]byte, error) {
	buf := make([]byte, size)

	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return nil, fmt.Errorf("failed to read bytes: %v", err)
	}

	return &buf, nil
}

func Get() string {
	buf, _ := Create()
	return fmt.Sprintf("%x", *buf)[0:size]
}
