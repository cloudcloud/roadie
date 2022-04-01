package types

// Configer is an expected behaviour to be provided by some form
// of configuration struct.
type Configer interface {
	GetConfigFile() string
	GetHostname() string
	GetListener() string
	GetLogger() Logger
}

// Reference is an individual entry that is located within a Source
// or Destination as a candidate for transferring.
type Reference struct {
	SubPath string `json:"sub_path,omitempty"`
	Entry   string `json:"entry"`
}
