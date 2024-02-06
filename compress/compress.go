package compress

import (
	"bytes"
	"compress/gzip"
)

func Compress(marshal []byte) ([]byte, error) {
	var compressedData bytes.Buffer
	gzipWriter := gzip.NewWriter(&compressedData)

	_, err := gzipWriter.Write(marshal)
	if err != nil {
		gzipWriter.Close()
		return nil, err
	}

	if err := gzipWriter.Close(); err != nil {
		return nil, err
	}

	return compressedData.Bytes(), nil
}
