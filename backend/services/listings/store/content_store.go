package store

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
	listingsPB "github.com/jalexanderII/zero_microservice/gen/listings"
)

// ContentStore is an interface to store content
type ContentStore interface {
	// Save saves new content to the store
	Save(SourceId int32, contentType listingsPB.ContentInfo_ContentType, contentData bytes.Buffer) (int32, *ContentInfo, error)
}

// DiskImageStore stores image on disk, and its info on memory
type DiskImageStore struct {
	mutex         sync.RWMutex
	contentFolder string
	content       map[string]*ContentInfo
}

// ContentInfo contains information of the content
type ContentInfo struct {
	SourceId    int32
	ContentType listingsPB.ContentInfo_ContentType
	Path        string
}

// NewDiskImageStore returns a new DiskImageStore
func NewDiskImageStore(contentFolder string) *DiskImageStore {
	return &DiskImageStore{
		contentFolder: contentFolder,
		content:       make(map[string]*ContentInfo),
	}
}

// Save adds a new content to a store
func (store *DiskImageStore) Save(
	sourceID int32,
	contentType listingsPB.ContentInfo_ContentType,
	contentData bytes.Buffer,
) (int32, *ContentInfo, error) {
	contentID, err := uuid.NewRandom()
	if err != nil {
		return 0, nil, fmt.Errorf("cannot generate content id: %w", err)
	}

	contentPath := fmt.Sprintf("%s/%s%s", store.contentFolder, contentID, contentType)

	file, err := os.Create(contentPath)
	if err != nil {
		return 0, nil, fmt.Errorf("cannot create content file: %w", err)
	}

	_, err = contentData.WriteTo(file)
	if err != nil {
		return 0, nil, fmt.Errorf("cannot write content to file: %w", err)
	}

	store.mutex.Lock()
	defer store.mutex.Unlock()

	contentInfo := &ContentInfo{
		SourceId:    sourceID,
		ContentType: contentType,
		Path:        contentPath,
	}

	store.content[contentID.String()] = contentInfo

	return int32(contentID.ID()), contentInfo, nil
}
