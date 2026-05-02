func lengthOfLongestSubstring(s string) int {
	characters:= []rune(s)
	left:=0
	maxLength:=0
	repetition:= make(map[rune]int)

	for rigth:=0; rigth<len(characters); rigth++{
		if _, found:=repetition[characters[rigth]]; !found{
			repetition[characters[rigth]]=1
		}else{
			repetition[characters[rigth]]++
		}

		for repetition[characters[rigth]]>1{
			repetition[characters[left]]--
			left++
		}

		if maxLength < rigth-left+1{
			maxLength= rigth-left+1
		}
	}

	return maxLength

}
