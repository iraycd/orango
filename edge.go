package orango

type Edge struct {
	Id    string `json:"_id,omitempty"  `
	From  string `json:"_from"`
	To    string `json:"_to"  `
	Error bool   `json:"error,omitempty"`
}
