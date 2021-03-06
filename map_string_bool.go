package goo

var _ Map = MapStringBool(nil)

var _ Pointer = &MapStringBool{}

// MapStringBool is a map from string to bool.
type MapStringBool map[string]bool

// Delete implements Map.
func (m MapStringBool) Delete(k interface{}) {
	delete(m, k.(string))
}

// Dereference implements Map.
func (m *MapStringBool) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapStringBool) Equals(other Equatable) bool {
	var n = other.(MapStringBool)

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
func (m MapStringBool) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringBool) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringBool) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringBool) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringBool) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringBool) Make(c int) Map {
	return make(MapStringBool, c)
}

// NotEquals implements Map.
func (m MapStringBool) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapStringBool) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapStringBool) Set(k, v interface{}) {
	m[k.(string)] = v.(bool)
}
