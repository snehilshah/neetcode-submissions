func getConcatenation(nums []int) []int {
	cat := make([]int, 0, len(nums))
	for _ = range 2{
    	for _, v := range nums{
			cat = append(cat, v)
		}
	}
	return cat
}
