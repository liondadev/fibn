package fibn

var store map[int64]int64

func init() {
	store = make(map[int64]int64)
}

func Calc(n int64) int64 {
	if n := store[n]; n != 0 {
		return n
	}
	if n <= 1 {
		return 1
	}
	r := Calc(n-1) + Calc(n-2)
	if r <= 1 {
		return r
	}
	store[n] = r
	return r
}
