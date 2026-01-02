/*
Task: Valid Parentheses (Stack)
Given a string s containing only '()[]{}', return true if the input string is valid.
An input is valid if open brackets are closed by the same type and in the correct order.

Goal: implement isValid(s) and analyze time/space complexity.
*/

function isValid(s) {
  // TODO: implement
  return false
}

// Test data
const tests = [
  { s: '()', expected: true },
  { s: '()[]{}', expected: true },
  { s: '(]', expected: false },
  { s: '([)]', expected: false },
  { s: '{[]}', expected: true },
]

// Uncomment to run quick checks
// for (const t of tests) {
//   console.log(isValid(t.s), 'expected', t.expected)
// }
