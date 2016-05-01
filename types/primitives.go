package types

const (
	Unsupported = iota
	Int
	Float32
	Float64
	String
)

// ItemType defines the type of btree Item
type T uint8

// getItemType returns the Item's type value
func Of(Item interface{}) T {
	switch Item.(type) {
	case int:
		return Int
	case float32:
		return Float32
	case float64:
		return Float64
	case string:
		return String
	default:
		return Unsupported
	}
}
