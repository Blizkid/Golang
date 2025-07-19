	var siteMap map[int]int /*创建集合 */

	func singleNumber(nums []int) int {
	
		siteMap = make(map[int]int)
		for _,v := range nums {
			var count = siteMap[v]
			count += 1
			siteMap[v] = count
		}

		for v := range siteMap {
			if siteMap[v] == 1 {
				return v
			}
		}
        return -1
	}