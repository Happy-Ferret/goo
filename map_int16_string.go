package goo

var _ Map = MapInt16String(nil)

var _ Pointer = &MapInt16String{}

// MapInt16String is a map from int16 to string.
type MapInt16String map[int16]string

// Delete implements Map.
func (m MapInt16String) Delete(k interface{}) {
	delete(m, k.(int16))
}

// Dereference implements Map.
func (m *MapInt16String) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapInt16String) Equals(other Equatable) bool {
	var n = other.(MapInt16String)

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
func (m MapInt16String) Get(k interface{}) interface{} {
	return m[k.(int16)]
}

// GetCheck implements Map.
func (m MapInt16String) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int16)]

	return v, ok
}

// KeyValues implements Map.
func (m MapInt16String) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapInt16String) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapInt16String) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapInt16String) Make(c int) Map {
	return make(MapInt16String, c)
}

// NotEquals implements Map.
func (m MapInt16String) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapInt16String) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapInt16String) Set(k, v interface{}) {
	m[k.(int16)] = v.(string)
}
