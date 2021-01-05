package types

// Sourcer is the expected behaviour that each source implementation will follow.
type Sourcer interface {
	CopyTo(Reference, Destination) ([]Reference, error)
	GetRefs() []Reference
	GetSubRefs(string) []Reference
	Type() string
}
