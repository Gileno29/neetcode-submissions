func hasDuplicate(nums []int) bool {

	if len(nums)==0{
		return false
	}
	
	saw:= make(map[int]int)
	saw[nums[0]]++

	
	for i:=1; i<len(nums); i++{
		if _, ok:= saw[nums[i]]; ok{
			return true
		}
		saw[nums[i]]++
	}
	return false
}