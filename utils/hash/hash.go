package utils

import (
	"crypto/rand"
	"fmt"
	"io"
)

func Create(size int8) (*[]byte, error) {
	buf := make([]byte, size)

	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return nil, fmt.Errorf("failed to read bytes: %v", err)
	}

	return &buf, nil
}

func Get(size int8) string {
	buf, _ := Create(size)
	return fmt.Sprintf("%x", *buf)[0:size]
}

func Id() string {
	const size int8 = 8
	return Get(size)
}

func Key() string {
	const size int8 = 64
	return Get(size)
}
