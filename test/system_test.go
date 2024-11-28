package test

import (
	"os"
	"testing"
	"github.com/evanwmart/distributed-compression-storage/pkg/chunk"
)

func TestChunking(t *testing.T) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "test_file.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up

	// Write test data to the file
	testData := make([]byte, 1024) // 1 KB of test data
	if _, err := tmpFile.Write(testData); err != nil {
		t.Fatalf("Failed to write test data: %v", err)
	}

	// Call the ChunkFile function
	chunks, err := chunk.ChunkFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Chunking failed: %v", err)
	}

	// Check the output for expected behavior
	if len(chunks) == 0 {
		t.Errorf("Expected at least one chunk, got %d", len(chunks))
	} else {
		t.Logf("Chunking successful: %d chunks created", len(chunks))
	}
}
