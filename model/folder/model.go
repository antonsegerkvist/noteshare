package folder

//
// ModelFolder contains information of a single folder.
//
type ModelFolder struct {
	ID               uint64 `json:"id"`
	Parent           uint64 `json:"parent"`
	Name             string `json:"name"`
	CreatedByUserID  uint64 `json:"createdByUserID"`
	ModifiedByUserID uint64 `json:"modifiedByUserID"`
	CreatedDate      string `json:"createdDate"`
	ModifiedDate     string `json:"modifiedDate"`
}

//
// ModelAddFolder contains information for adding a file.
//
type ModelAddFolder struct {
	Parent uint64 `json:"parent"`
	Name   string `json:"name"`
}
