package types

// ExecutePayload
type ExecutePayload struct {
	DestinationName string `json:"destination_name"`
	EntryName       string `json:"entry_name"`
	SourceName      string `json:"source_name"`
}

// ExecuteResult
type ExecuteResult struct {
	Error      error       `json:"error"`
	References []Reference `json:"references"`
}

// RemovePayload
type RemovePayload struct {
	DestinationName string `json:"destination_name"`
	EntryName       string `json:"entry_name"`
}
