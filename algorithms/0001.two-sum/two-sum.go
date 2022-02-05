package Problem0001

func twoSum(nums []int, target int) []int {
	// m is responsible for saving the sequence number of the map[int]int
	m := make(map[int]int, len(nums))

	// Get the serial number of b through a for loop
	for i, b := range nums {
		// Get the serial number of a = target - b by querying the map
		if j, ok := m[target-b]; ok {
			// ok is true
			// Explain that before i, there is nums[j] == a
			return []int{j, i}
			// Note that the order is j, i because j<i
		}

		// store the values of i and i in the map
		m[nums[i]] = i
	}

	return nil
}
