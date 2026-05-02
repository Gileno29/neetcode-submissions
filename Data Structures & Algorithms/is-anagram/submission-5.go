func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}

	
	str1:= make(map[rune]int)
	str2:= make(map[rune]int)
	for i:=0; i< len(s); i++{
		if _, ok:=str1[rune(s[i])]; !ok{
			str1[rune(s[i])]++
		}
		str1[rune(s[i])]++

	}

	for i:=0; i< len(t); i++{
		if _, ok:=str2[rune(t[i])]; !ok{
			str2[rune(t[i])]++
		}
		str2[rune(t[i])]++
	}

	fmt.Println(str1, str2)
	for i:=0; i< len(t); i++{
		if _, ok:=str1[rune(t[i])]; !ok{
			return false
		}
		if str1[rune(t[i])]!=str2[rune(t[i])]{
			return false
		}
	}

	 return true



}
