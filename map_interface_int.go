package goo

var _ Map = MapInterfaceInt(nil)

var _ Pointer = &MapInterfaceInt{}

// MapInterfaceInt is a map from interface{} to int.
type MapInterfaceInt map[interface{}]int

// Delete implements Map.
func (m MapInterfaceInt) Delete(k interface{}) {
	delete(m, k.(interface{}))
}

// Dereference implements Map.
func (m *MapInterfaceInt) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapInterfaceInt) Equals(other Equatable) bool {
	var n = other.(MapInterfaceInt)

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
func (m MapInterfaceInt) Get(k interface{}) interface{} {
	return m[k.(interface{})]
}

// GetCheck implements Map.
func (m MapInterfaceInt) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(interface{})]

	return v, ok
}

// KeyValues implements Map.
func (m MapInterfaceInt) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapInterfaceInt) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapInterfaceInt) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapInterfaceInt) Make(c int) Map {
	return make(MapInterfaceInt, c)
}

// NotEquals implements Map.
func (m MapInterfaceInt) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapInterfaceInt) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapInterfaceInt) Set(k, v interface{}) {
	m[k.(interface{})] = v.(int)
}
