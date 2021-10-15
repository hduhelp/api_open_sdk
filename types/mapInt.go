package types

type MapInt struct {
	value map[int]bool
}

func (i MapInt) Equal(n MapInt) bool {
	if len(i.value) != len(n.value) {
		return false
	}
	for k := range n.value {
		if i.value[k] != n.value[k] {
			return false
		}
	}
	return true
}
