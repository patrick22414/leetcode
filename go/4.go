package main

import "fmt"

func main() {
	var nums1, nums2 []int

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	fmt.Println(nums1, nums2, findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 3}
	nums2 = []int{2}
	fmt.Println(nums1, nums2, findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 2}
	nums2 = []int{4, 5}
	fmt.Println(nums1, nums2, findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n { // make sure nums1 is shorter
		nums1, nums2, m, n = nums2, nums1, n, m
	}

	s := 0     // start
	e := 2 * m // end

	var cut1, cut2, l1, r1, l2, r2 int

	for {
		cut1 = (s + e) / 2  // cut in nums1
		cut2 = m + n - cut1 // cut in nums2

		if cut1 != 0 { // left side of cut1
			l1 = nums1[(cut1-1)/2]
		} else {
			l1 = -1e6
		}
		if cut1 != 2*m { // right side of cut1
			r1 = nums1[cut1/2]
		} else {
			r1 = 1e6
		}

		if cut2 != 0 { // left side of cut2
			l2 = nums2[(cut2-1)/2]
		} else {
			l2 = -1e6
		}
		if cut2 != 2*n { // right side of cut2
			r2 = nums2[cut2/2]
		} else {
			r2 = 1e6
		}

		if l1 > r2 {
			e = cut1
		} else if l2 > r1 {
			s = cut1 + 1
		} else {
			// fmt.Println(l1, r1, l2, r2)
			return float64(max(l1, l2)+min(r1, r2)) / 2
		}
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
