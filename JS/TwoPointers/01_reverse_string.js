/*
Task: Reverse String (Two Pointers)
Given an array of characters chars, reverse it in-place.
Do not return anything; modify chars directly.

Goal: implement reverseString(chars) and analyze time/space complexity.
*/

function reverseString(chars) {
  // TODO: implement
}

// Test data
const tests = [
  { chars: ['h', 'e', 'l', 'l', 'o'], expected: ['o', 'l', 'l', 'e', 'h'] },
  { chars: ['H', 'a', 'n', 'n', 'a', 'h'], expected: ['h', 'a', 'n', 'n', 'a', 'H'] },
  { chars: ['a'], expected: ['a'] },
]

// Uncomment to run quick checks
// for (const t of tests) {
//   const arr = t.chars.slice()
//   reverseString(arr)
//   console.log(arr, 'expected', t.expected)
// }
