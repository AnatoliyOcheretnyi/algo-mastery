/*
Task: Valid Anagram
Given two strings s and t, return true if t is an anagram of s.

Goal: implement isAnagram(s, t) and analyze time/space complexity.
*/

function isAnagram(s, t) {
  // TODO: implement
  return false
}

// Test data
const tests = [
  { s: 'anagram', t: 'nagaram', expected: true },
  { s: 'rat', t: 'car', expected: false },
  { s: 'a', t: 'a', expected: true },
  { s: 'ab', t: 'a', expected: false },
]

// Uncomment to run quick checks
// for (const t of tests) {
//   console.log(isAnagram(t.s, t.t), 'expected', t.expected)
// }
