"""sumary_line

Challenge: Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution,
and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]

Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.

GOAL: Can you come up with an algorithm that is less than O(n²) time complexity?
"""

from typing import List


# O(n²)
def twoSum_v1(nums: List[int], target: int) -> List[int]:
    for i in range(0, len(nums)):
        for j in range(i + 1, len(nums)):
            if target == (nums[i] + nums[j]):
                return [i, j]
    return []


print(twoSum_v1([3, 2, 3], 6))


# O(n)
def twoSum_v3(nums: List[int], target: int) -> List[int]:
    num_dict: dict = {}
    for i, num in enumerate(nums):
        diff = target - num
        if diff in num_dict:
            return [num_dict[diff], i]
        num_dict[num] = i
    return []


print(twoSum_v3([11, 3, 7, 2, 5], 9))


def twoSum_v4(nums: List[int], target: int) -> List[int]:
    hashmap: dict = dict()
    position: int = 0
    while position < len(nums):
        if (target - nums[position]) not in hashmap:
            hashmap[nums[position]] = position
            position += 1
        else:
            return [hashmap[target - nums[position]], position]
    return []

print(twoSum_v4([11, 3, 7, 2, 5], 9))