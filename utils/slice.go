package utils

// Difference returns the elements in slice1 that are not in slice2.
func Difference(slice1, slice2 []int) []int {
	// Create a map to track elements in slice2
	slice2Map := make(map[int]bool)
	for _, v := range slice2 {
		slice2Map[v] = true
	}

	// Find elements in slice1 that are not in slice2
	var diff []int
	for _, v := range slice1 {
		if !slice2Map[v] {
			diff = append(diff, v)
		}
	}

	return diff
}
