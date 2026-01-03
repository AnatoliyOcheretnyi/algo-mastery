/*
Task: Reverse Words
Given a string s, return a new string with the word order reversed.
Words are separated by one or more spaces. Trim extra spaces in the result.

Goal: implement reverseWords(s) and analyze time/space complexity.
*/

function reverseWords(s) {
  return s.trim().split(' ').filter(Boolean).reverse().join(' ')
}

function reverseWordsWithReg(s) {
  return s.trim().split(/\s+/).reverse().join(' ')
}

// Test data
const tests = [
  { s: 'the sky is blue', expected: 'blue is sky the' },
  { s: '  hello   world  ', expected: 'world hello' },
  { s: 'a', expected: 'a' },
  { s: '  a   good   example ', expected: 'example good a' },
]

// Uncomment to run quick checks
for (const t of tests) {
  console.log(reverseWords(t.s), 'expected', t.expected)
}
