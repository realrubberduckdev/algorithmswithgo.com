package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var reversed string
	for _, r := range word {
		reversed = string(r) + reversed
	}
	return reversed
}

func Reverse2(word string) string {
	start := len(word) - 1
	var reversed string
	for i := start; i >= 0; i-- {
		reversed += string(word[i])
	}
	return reversed
}
