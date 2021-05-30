def n_partition(k, total, elements):
    if total % k != 0:
        return False
    number = k
    while number:
        can_partition, elements = partition(total//k, elements)
        # print(can_partition, elements)
        if not can_partition:
            return False
        number -= 1
    return True


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


class Solution:
    def canPartitionKSubsets(self, nums: List[int], k: int) -> bool:
        nums.sort()
        return n_partition(k, sum(nums), nums[::-1])
