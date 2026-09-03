func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	seen := make(map[byte]int, len(s))

	for i, _ := range s {
		sval := s[i]
		tval := t[i]	
		seen[sval]++
		seen[tval]--
	}

	for _, v := range seen{
		if v != 0{
			return false
		}
	}
	return true
}
