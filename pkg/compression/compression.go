package compression

import (
	"bytes"
	"compress/gzip"
)

func CompressChunk(data []byte) ([]byte, error) {
	var buffer bytes.Buffer
	writer := gzip.NewWriter(&buffer)
	_, err := writer.Write(data)
	if err != nil {
		return nil, err
	}
	writer.Close()
	return buffer.Bytes(), nil
}
