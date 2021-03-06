package goo

var _ Map = MapUint64Struct(nil)

var _ Pointer = &MapUint64Struct{}

// MapUint64Struct is a map from uint64 to struct{}.
type MapUint64Struct map[uint64]struct{}

// Delete implements Map.
func (m MapUint64Struct) Delete(k interface{}) {
	delete(m, k.(uint64))
}

// Dereference implements Map.
func (m *MapUint64Struct) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapUint64Struct) Equals(other Equatable) bool {
	var n = other.(MapUint64Struct)

	if len(n) != len(m) {
		return false
	}

	for k, v := range m {
		if nv, ok := n[k]; !ok {
			return false
		} else if nv != v {
			return false
		}
	}

	return true
}

// Get implements Map.
func (m MapUint64Struct) Get(k interface{}) interface{} {
	return m[k.(uint64)]
}

// GetCheck implements Map.
func (m MapUint64Struct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(uint64)]

	return v, ok
}

// KeyValues implements Map.
func (m MapUint64Struct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapUint64Struct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapUint64Struct) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapUint64Struct) Make(c int) Map {
	return make(MapUint64Struct, c)
}

// NotEquals implements Map.
func (m MapUint64Struct) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapUint64Struct) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapUint64Struct) Set(k, v interface{}) {
	m[k.(uint64)] = v.(struct{})
}
