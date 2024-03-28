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



from typing import Dict, List

# O(n²)
def twoSum_v1(nums: List[int], target: int) -> List[int]:
        result: List[int] = []
        for i in range(0, len(nums)):
            for j in range(i+1, len(nums)):
                if target == (nums[i] + nums[j]):
                    result.append(i)
                    result.append(j)
        return result

print(twoSum_v1([3,2,3], 6))

# O(n)
def twoSum_v2(nums: List[int], target: int) -> List[int]:
        result: List[int] = []
        for i in range(0, len(nums)):
            if i+1 < len(nums) and target == (nums[i] + nums[i+1]):
                result.append(i)
                result.append(i+1)
        return result

print(twoSum_v2([3,2,3], 6))

# O(n)
def twoSum_v3(nums: List[int], target: int) -> List[int]:
        num_dict: dict = {}
        for i, num in enumerate(nums):
            diff = target - num
            if diff in num_dict:
                return [num_dict[diff], i]
            num_dict[num] = i
        return []

print(twoSum_v3([3,2,3], 6))