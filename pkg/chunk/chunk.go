package chunk

import (
	"io"
	"os"
)

const ChunkSize = 10 * 1024 * 1024 // 10 MB

func ChunkFile(filePath string) ([][]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var chunks [][]byte
	buffer := make([]byte, ChunkSize)

	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			break
		}

		chunk := make([]byte, n)
		copy(chunk, buffer[:n])
		chunks = append(chunks, chunk)
	}
	return chunks, nil
}
