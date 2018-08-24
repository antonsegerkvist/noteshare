package file

//
// ModelFile contains information about a single file.
//
type ModelFile struct {
	ID   uint64 `json:"id"`
	Type uint64 `json:"type"`
	Name string `json:"name"`
}
