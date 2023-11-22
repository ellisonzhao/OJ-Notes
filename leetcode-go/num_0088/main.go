package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	pos := m + n - 1
	i, j := m-1, n-1
	for pos >= 0 {
		for i >= 0 && j >= 0 {
			if nums1[i] >= nums2[j] {
				nums1[pos] = nums1[i]
				i--
			} else {
				nums1[pos] = nums2[j]
				j--
			}
			pos--
		}
		for i >= 0 {
			nums1[pos] = nums1[i]
			i--
			pos--
		}
		for j >= 0 {
			nums1[pos] = nums2[j]
			j--
			pos--
		}
	}
}

func main() {

}
