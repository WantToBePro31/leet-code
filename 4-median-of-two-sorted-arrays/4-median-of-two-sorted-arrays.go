func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}

	if len(nums1) + len(nums2) == 1 {
		if len(nums1) == 1 {
			return float64(nums1[0])
		}
		return float64(nums2[0])
	}

	if len(nums1) + len(nums2) == 2 {
		if len(nums1) == 0 {
			return float64(nums2[0]+nums2[1])/float64(2)
		}
		if len(nums2) == 0 {
			return float64(nums1[0]+nums1[1])/float64(2)
		}
		return float64(nums1[0]+nums2[0])/float64(2)
	}

	if len(nums1) != 0 && len(nums2) == 0 {
		if len(nums1) % 2 == 1 {
			return float64(nums1[len(nums1)/2])
		}
		return float64(nums1[len(nums1)/2]+nums1[len(nums1)/2 - 1])/float64(2)
	}

	if len(nums1) == 0 && len(nums2) != 0 {
		if len(nums2) % 2 == 1 {
			return float64(nums2[len(nums2)/2])
		}
		return float64(nums2[len(nums2)/2]+nums2[len(nums2)/2 - 1])/float64(2)
	}

	var mixedArray []int
	var ind int
	limit := 0
	for ind = 0; ind < len(nums1); ind++ {
		if limit == len(nums2) {
			break
		}
		if nums1[ind] > nums2[limit] {
			for nums2[limit] < nums1[ind] {
				mixedArray = append(mixedArray, nums2[limit])
				limit++
				if limit == len(nums2) {
					break
				}
			}
		}
		mixedArray = append(mixedArray, nums1[ind])
	}
	mixedArray = append(mixedArray, nums1[ind:]...)
	mixedArray = append(mixedArray, nums2[limit:]...)
	fmt.Println(mixedArray)
	if len(mixedArray) % 2 == 1 {
		return float64(mixedArray[len(mixedArray)/2])
	}
	return float64(mixedArray[len(mixedArray)/2]+mixedArray[len(mixedArray)/2 - 1])/float64(2)
}