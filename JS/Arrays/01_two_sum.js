/*
Task: Two Sum
Given an array of integers nums and an integer target, return the indices of
the two numbers such that they add up to target.
Assume exactly one solution exists. You may not use the same element twice.

Goal: implement twoSum(nums, target) and analyze time/space complexity.
*/

function twoSum(nums, target) {
  const result = []
  for (let i = 0; i < nums.length; i++) {
    for (let j = i + 1; j < nums.length; j++) {
      if (nums[i] + nums[j] === target) {
        return [i, j]
      }
    }
  }
  return result
}

// Alternative: O(n) time using a Map
function twoSumWithMap(nums, target) {
  const seen = new Map()
  for (let i = 0; i < nums.length; i++) {
    const need = target - nums[i]
    if (seen.has(need)) {
      return [seen.get(need), i]
    }
    seen.set(nums[i], i)
  }
  return []
}

// Test data
const tests = [
  { nums: [2, 7, 11, 15], target: 9, expected: [0, 1] },
  { nums: [3, 2, 4], target: 6, expected: [1, 2] },
  { nums: [-1, -2, -3, -4, -5], target: -8, expected: [2, 4] },
]

// Uncomment to run quick checks
for (const t of tests) {
  console.log(twoSum(t.nums, t.target), 'expected', t.expected)
  console.log(twoSumWithMap(t.nums, t.target), 'expected', t.expected)
}
