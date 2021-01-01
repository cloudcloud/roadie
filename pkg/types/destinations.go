package types

// Destinationer is the expected behaviour for a specific type of
// destination to follow.
type Destinationer interface {
	GetLocation() string
	GetRefs() []Reference
	RemoveFile(string) error
	Type() string
}
