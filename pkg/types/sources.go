package types

// Sourcer
type Sourcer interface {
	CopyTo(Reference, Destination) ([]Reference, error)
	GetRefs() []Reference
	Type() string
}
