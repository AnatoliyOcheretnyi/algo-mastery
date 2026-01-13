/*
Task: Merge Sorted Arrays
Given two sorted arrays nums1 and nums2, return a new sorted array containing all elements.

Goal: implement mergeSorted(nums1, nums2) and analyze time/space complexity.
*/

function mergeSorted(nums1, nums2) {
  // TODO: implement
  return []
}

// Test data
const tests = [
  { nums1: [1, 3, 5], nums2: [2, 4, 6], expected: [1, 2, 3, 4, 5, 6] },
  { nums1: [], nums2: [1, 2], expected: [1, 2] },
  { nums1: [0, 2, 2], nums2: [1, 3], expected: [0, 1, 2, 2, 3] },
]

// Uncomment to run quick checks
// for (const t of tests) {
//   console.log(mergeSorted(t.nums1, t.nums2), 'expected', t.expected)
// }
