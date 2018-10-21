package file

//
// ModelFile contains information about a single file.
//
type ModelFile struct {
	ID                 uint64 `json:"id"`
	Type               uint64 `json:"type"`
	FileReferenceID    uint64 `json:"fileReferenceID,omitempty"`
	FileReferenceCount uint64 `json:"fileReferenceCount"`
	Parent             uint64 `json:"parent,omitempty"`
	Name               string `json:"name"`
	ModificationDate   string `json:"modificationDate"`
}

//
// ModelAddFile contains information for adding a single file.
//
type ModelAddFile struct {
	Type            uint64 `json:"type"`
	FileReferenceID uint64 `json:"fileReferenceID"`
	Parent          uint64 `json:"parent"`
	Name            string `json:"name"`
}
