/*
Task: Find Duplicates
Given an array of integers nums (may contain duplicates), return an array of
values that appear more than once. Each duplicate value should appear only once
in the result, in the order of its first occurrence.

Goal: implement findDuplicates(nums) and analyze time/space complexity.
*/

function findDuplicates(nums) {
  const dublicates = []
  const seen = new Map()
  const added = new Map()

  for (const num of nums) {
    if (seen.has(num)) {
      if (!added.has(num)) {
        dublicates.push(num)
        added.set(num)
      }
    } else {
      seen.set(num)
    }
  }
  return dublicates
}

// Test data
const tests = [
  { nums: [1, 2, 3, 2, 4, 1], expected: [1, 2] },
  { nums: [5, 5, 5, 5], expected: [5] },
  { nums: [1, 2, 3, 4], expected: [] },
  { nums: [-1, -1, 0, 2, 2, 2], expected: [-1, 2] },
]

// Uncomment to run quick checks
for (const t of tests) {
  console.log(findDuplicates(t.nums), 'expected', t.expected)
}
