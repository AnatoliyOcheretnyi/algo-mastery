/*
Task: Group Words by Length
Given an array of strings words, return a map where the key is word length
and the value is an array of words with that length.

Goal: implement groupByLength(words) and analyze time/space complexity.
*/

function groupByLength(words) {
  // TODO: implement
  return {}
}

// Test data
const tests = [
  { words: ['go', 'js', 'code', 'a'], expected: { 1: ['a'], 2: ['go', 'js'], 4: ['code'] } },
  { words: ['hi', 'there'], expected: { 2: ['hi'], 5: ['there'] } },
  { words: [], expected: {} },
]

// Uncomment to run quick checks
// for (const t of tests) {
//   console.log(groupByLength(t.words), 'expected', t.expected)
// }
