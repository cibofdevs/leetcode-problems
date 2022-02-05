package Problem0003

func lengthOfLongestSubstring(s string) int {
	// location[s[i]] == j means:
	// The i-th string in s last appeared at the j position of s, so there is no s[i] in s[j+1:i]
	// location[s[i]] == -1 means: the first occurrence of s[i] in s
	location := [256]int{} // only 256 long because, assuming the input string is only ASCII characters
	for i := range location {
		location[i] = -1 // first set all characters not seen
	}

	maxLen, left := 0, 0

	for i := 0; i < len(s); i++ {
		// indicates that s[i] has been repeated in s[left:i+1]
		// and the last occurrence of s[i] is at location[s[i]]
		if location[s[i]] >= left {
			left = location[s[i]] + 1 // remove the s[i] character and the part before it in s[left:i+1]
		} else if i+1-left > maxLen {
			// fmt.Println(s[left:i+1])
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}

	return maxLen
}