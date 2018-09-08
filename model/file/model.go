package file

//
// ModelFile contains information about a single file.
//
// ID
// This is a unique file identifier.
//
// Type:
// Refers what type of file this is.
//
// FileReferenceID:
// Is used for weak and hard links and refers to the original file id.
//
// FileReferenceCount:
// Is used for hard links and is the number of hard link references that are
// pointing to this file.
//
// Parent:
// Refers to the parent file id (of type folder) of this file.
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
