package account

//
// ModelAccount contains information of a single account.
//
type ModelAccount struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
