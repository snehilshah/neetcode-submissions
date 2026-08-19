func hasDuplicate(nums []int) bool {
    seen := make(map[int]struct{})
    for _, v := range nums{
        _, prs := seen[v]
        if prs{
            return true
        }
        seen[v] = struct{}{}
    }
    return false
}