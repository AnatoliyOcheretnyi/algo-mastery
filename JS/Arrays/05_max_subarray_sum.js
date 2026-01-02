/*
Task: Maximum Subarray Sum
Given an array of integers nums, find the maximum sum of any contiguous subarray.
Return the maximum sum.

Goal: implement maxSubarraySum(nums) and analyze time/space complexity.
*/

function maxSubarraySum(nums) {
  let currentSum = nums[0]
  let maxSum = nums[0]

  for (let i = 1; i < nums.length; i++) {
    const num = nums[i]
    currentSum = Math.max(num, currentSum + num)
    maxSum = Math.max(maxSum, currentSum)
  }

  return maxSum
}

// Test data
const tests = [
  { nums: [-2, 1, -3, 4, -1, 2, 1, -5, 4], expected: 6 }, // [4, -1, 2, 1]
  { nums: [1], expected: 1 },
  { nums: [5, 4, -1, 7, 8], expected: 23 },
  { nums: [-1, -2, -3], expected: -1 },
]

// Uncomment to run quick checks
for (const t of tests) {
  console.log(maxSubarraySum(t.nums), 'expected', t.expected)
}
