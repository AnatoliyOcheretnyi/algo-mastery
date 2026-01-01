/*
Task: Move Zeroes
Given an array nums, move all 0's to the end while maintaining the relative
order of the non-zero elements. Do it in-place.

Goal: implement moveZeroes(nums) and analyze time/space complexity.
*/

function moveZeroes(nums) {
  let write = 0;

  for (let i = 0; i < nums.length; i++) {
    if (nums[i] !== 0) {
      nums[write] = nums[i];
      write++;
    }
  }

  while (write < nums.length) {
    nums[write] = 0;
    write++;
  }
}

// Test data
const tests = [
  { nums: [0, 1, 0, 3, 12], expected: [1, 3, 12, 0, 0] },
  { nums: [0, 0, 0], expected: [0, 0, 0] },
  { nums: [4, 2, 4, 0, 0, 3, 0, 5, 1], expected: [4, 2, 4, 3, 5, 1, 0, 0, 0] },
];

// Uncomment to run quick checks
for (const t of tests) {
  const arr = t.nums.slice();
  moveZeroes(arr);
  console.log(arr, "expected", t.expected);
}
