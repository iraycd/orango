package orango

type Document struct {
	Id  string `json:"_id,omitempty"              `
	Rev string `json:"_rev,omitempty"             `
	Key string `json:"_key,omitempty"             `

	Error   bool   `json:"error,omitempty"`
	Message string `json:"errorMessage,omitempty"`
}
