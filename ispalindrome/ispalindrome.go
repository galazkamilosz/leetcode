package ispalindrome

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xCopy := x
	y := 0

	for xCopy >= 1 {
		piece := xCopy % 10
		y = y*10 + piece
		xCopy /= 10
	}
	return x == y
}
