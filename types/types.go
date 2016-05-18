package types

const (
	Unsupported = iota
	Null
	Int
	Float32
	Float64
	String
)

// ItemType defines the type of btree Item
type T uint8

// getItemType returns the Item's type value
func Of(v interface{}) T {
	switch v.(type) {
	case nil:
		return Null
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
