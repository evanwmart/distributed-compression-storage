package storage

import (
	"fmt"
	"os"
)

func StoreChunk(chunkID string, data []byte) error {
	filePath := fmt.Sprintf("chunk_storage/%s.chunk", chunkID)
	return os.WriteFile(filePath, data, 0644)
}
