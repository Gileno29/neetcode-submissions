func twoSum(nums []int, target int) []int {

seen := make(map[int]int)
	var result []int
	for i := 0; i < len(nums); i++ {
		seen[nums[i]] = i

	}

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if _, ok := seen[complement]; ok && seen[complement] != i {
			result = append(result, i)
			result = append(result, seen[complement])
			return result

		}
	}

	return result


return result


	
}
