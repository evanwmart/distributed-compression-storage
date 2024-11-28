package metadata

import (
	"sync"
)

type MetadataStore struct {
	mu       sync.Mutex
	metadata map[string][]string // FileID -> ChunkIDs
}

func NewMetadataStore() *MetadataStore {
	return &MetadataStore{metadata: make(map[string][]string)}
}

func (m *MetadataStore) AddFile(fileID string, chunkIDs []string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.metadata[fileID] = chunkIDs
}

func (m *MetadataStore) GetFile(fileID string) ([]string, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	chunks, exists := m.metadata[fileID]
	return chunks, exists
}
