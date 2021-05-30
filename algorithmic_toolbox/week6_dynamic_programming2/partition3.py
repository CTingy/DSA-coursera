def n_partition(n, total, elements):
    if total % n != 0:
        return 0
    count = n
    while count:
        can_partition, elements = partition(total//n, elements)
        # print(can_partition, elements)
        if not can_partition:
            return 0
        count -= 1
    return 1


def partition(value, elements):
    dp = []
    for _ in range(0, value+1):
        dp.append([False for _ in range(len(elements)+1)])

    for i in range(1, value+1):
        for j in range(1, len(elements)+1):
            if dp[i][j-1] or i == elements[j-1]:
                dp[i][j] = True
            else:
                value_diff = i - elements[j-1]
                if value_diff > 0 and dp[value_diff][j-1]:
                    dp[i][j] = True
    
    selected_idx = set()
    temp_value = value
    for j in range(1, len(elements)+1):
        while temp_value and j and dp[temp_value][j]:
            if not dp[temp_value][j-1]:
                selected_idx.add(j-1)
                temp_value -= elements[j-1]
            j -= 1
    
    new_elements = []
    for j in range(len(elements)):
        if j not in selected_idx:
            new_elements.append(elements[j])

    return dp[value][len(elements)], new_elements


def k_partition(nums, k):
	target, m = divmod(sum(nums), k)
	if m: return False
	dp, n = [0]*k, len(nums)
	nums.sort(reverse=True)
	def dfs(i):
		if i == n:
			return len(set(dp)) == 1
		for j in range(k):
			dp[j] += nums[i]
			if dp[j] <= target and dfs(i+1):
				return True
			dp[j] -= nums[i]
			if not dp[j]: break
		return False
	return nums[0] <= target and dfs(0)


# [4,4,6,2,3,8,10,2,10,7]
# 4
if __name__ == '__main__':
    input()
    elements = [int(x) for x in input().split(' ')]
    # elements.sort()
    # print(n_partition(3, sum(elements), elements[::-1]))
    can_partition = k_partition(elements, 3)
    if can_partition:
        print(1)
    else:
        print(0)
