// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"math/rand"
	"time"

	s "github.com/inancgumus/prettyslice"
	"github.com/inancgumus/screen"
)

func main() {
	s.PrintBacking = true
	s.MaxPerLine = 30
	s.Width = 150

	var nums []int

	screen.Clear()
	for cap(nums) <= 128 {
		screen.MoveTopLeft()

		s.Show("nums", nums)
		nums = append(nums, rand.Intn(9)+1)

		time.Sleep(time.Second / 4)
	}
}
