package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		s := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			s[x] = uint8((i + x) / 2)
		}
		res[i] = s
	}
	return res
}

func main() {
	pic.Show(Pic)
}
