/*
Task: Remove Duplicates from Sorted Array
Given a sorted array nums, remove the duplicates in-place so that each element
appears only once and return the new length.
The first part of nums should contain the unique values in the same order.

Goal: implement removeDuplicates(nums) and analyze time/space complexity.
*/

function removeDuplicates(nums) {
  if (nums.length === 0) return 0

  let write = 1

  for (let i = 1; i < nums.length; i++) {
    if (nums[i] !== nums[write - 1]) {
      nums[write] = nums[i]
      write++
    }
  }
  return write
}

// Test data
const tests = [
  { nums: [1, 1, 2], expectedLength: 2, expectedPrefix: [1, 2] },
  { nums: [0, 0, 1, 1, 1, 2, 2, 3, 3, 4], expectedLength: 5, expectedPrefix: [0, 1, 2, 3, 4] },
  { nums: [1], expectedLength: 1, expectedPrefix: [1] },
]

// Uncomment to run quick checks
for (const t of tests) {
  const arr = t.nums.slice()
  const len = removeDuplicates(arr)
  console.log(len, arr.slice(0, len), 'expected', t.expectedLength, t.expectedPrefix)
}
