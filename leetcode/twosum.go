// https://leetcode.com/problems/two-sum/solution/

// my attempt

func twoSum(nums []int, target int) []int {
    numsTable := make(map[int]int)
    
    for i := 0; i < len(nums); i++ {
        number = target - nums[i]
        
        if number in numsTable {
					return numsTable[final], i
        } else
        	numsTable[nums[i]] = i
		}
}

// solution

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for index, num := range nums {
			if value, ok := hashMap[target - num]; ok {
					return []int{index, value}
			} else {
					hashMap[num] = index
			}
	}
	return nums
}

