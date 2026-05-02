func characterReplacement(s string, k int) int {
	characters := []rune(s)
	left := 0
	longest := 0
	repetition := make(map[rune]int)
	maxFreq := 0

	for right := 0; right < len(characters); right++ {
		repetition[characters[right]]++
		if repetition[characters[right]] > maxFreq {
			maxFreq = repetition[characters[right]]
		}

		for (right-left+1)-maxFreq > k {
			repetition[characters[left]]--
			left++
		}

		if longest < right-left+1 {
			longest = right - left + 1
		}

	}

 return longest

}
