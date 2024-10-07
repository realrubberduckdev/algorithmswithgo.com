package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {

	for _, i := range list {
		if i == num {
			return true
		}
	}
	return false
}

func NumInList2(list []int, num int) bool {

	count := len(list)
	for i := 0; i < count; i++ {
		if list[i] == num {
			return true
		}
	}
	return false
}
