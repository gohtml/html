package html

// A HTMLNode set implemented by a slice.
// A slice is more time and memofy efficient than a map when
// the number of elements is small.
type htmlNodeSet []HTMLNode

// Puts a new element into the set.
func (set *htmlNodeSet) Put(s HTMLNode) {
	for _, el := range *set {
		if el == s {
			return
		}
	}

	*set = append(*set, s)
}

// Deletes a new element from the set.
func (set *htmlNodeSet) Del(s HTMLNode) {
	for i, el := range *set {
		if el == s {
			*set = append((*set)[:i], (*set)[i+1:]...)
			return
		}
	}
}
