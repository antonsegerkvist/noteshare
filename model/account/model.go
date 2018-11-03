package account

//
// ModelAccount contains information of a single account.
//
type ModelAccount struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

//
// ModelAccountLayout contains information of the layout on the frontpage.
//
type ModelAccountLayout struct {
	Layout *string `json:"layout"`
}
