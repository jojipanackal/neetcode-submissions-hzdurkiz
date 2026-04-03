func longestConsecutive(nums []int) int {
    var result int
    in_nums := make(map[int]bool)

    for _, num := range nums {
        in_nums[num] = true
    }

    for _, num := range nums {
        count := 0
        if !in_nums[num-1] {
            count = 1;
            temp := num
            for in_nums[temp+1] {
                temp++
                count++
            }
            if count > result {
                result = count
            }
        }
    }
    return result
}
