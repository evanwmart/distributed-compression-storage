package main

import (
	"fmt"
	"log"
	"os"
	"github.com/evanwmart/distributed-compression-storage/pkg/chunk"
	"github.com/evanwmart/distributed-compression-storage/pkg/compression"
	"github.com/evanwmart/distributed-compression-storage/pkg/storage"
	"github.com/evanwmart/distributed-compression-storage/pkg/metadata"
)

func main() {
	// Ensure storage directory exists
	os.Mkdir("chunk_storage", 0755)

	// Initialize metadata store
	metaStore := metadata.NewMetadataStore()

	// Chunk the file
	filePath := "test_file.txt"
	chunks, err := chunk.ChunkFile(filePath)
	if err != nil {
		log.Fatalf("Failed to chunk file: %v", err)
	}

	// Process each chunk
	var chunkIDs []string
	for i, chunkData := range chunks {
		// Compress the chunk
		compressedChunk, err := compression.CompressChunk(chunkData)
		if err != nil {
			log.Fatalf("Failed to compress chunk: %v", err)
		}

		// Store the compressed chunk
		chunkID := fmt.Sprintf("chunk-%d", i)
		err = storage.StoreChunk(chunkID, compressedChunk)
		if err != nil {
			log.Fatalf("Failed to store chunk: %v", err)
		}
		chunkIDs = append(chunkIDs, chunkID)
	}

	// Save metadata
	fileID := "file-12345"
	metaStore.AddFile(fileID, chunkIDs)

	// Output success message
	fmt.Printf("File '%s' processed and stored successfully!\n", filePath)
}
