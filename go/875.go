package main

func main() {
}

func minEatingSpeed(piles []int, h int) int {
	lo, hi := 1, maxOfPiles(piles)

	for lo != hi {
		k := (lo + hi) / 2
		t := totalEatingTime(piles, k)

		if t > h { // too slow
			lo = k + 1
		} else { // too fast
			hi = k
		}
	}

	return lo
}

func totalEatingTime(piles []int, k int) int {
	// assume k > 0
	total := 0
	for _, m := range piles {
		total += (m-1)/k + 1
	}

	return total
}

func maxOfPiles(piles []int) int {
	//assume all piles[i] >= 1
	max := 0
	for _, m := range piles {
		if m > max {
			max = m
		}
	}

	return max
}
