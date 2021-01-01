package types

// Destinationer
type Destinationer interface {
	GetLocation() string
	GetRefs() []Reference
	RemoveFile(string) error
	Type() string
}
