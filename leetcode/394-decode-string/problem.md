# 394. Decode String

## Description

- Given an encoded string, return its decoded string.

- The encoding rule is `k[encoded_string]`, where the `encoded_string` inside the square brackets is repeated exactly `k` times. `k` is guaranteed to be a positive integer.

- The input string is always valid: it contains no extra whitespace and its square brackets are well-formed. The original data contains no digits; digits appear only in repeat counts. For example, inputs such as `3a` and `2[4]` do not occur.

- The test cases guarantee that the decoded output length does not exceed `10^5`.

## Examples

### Example 1

- Input: `s = "3[a]2[bc]"`

- Output: `"aaabcbc"`

### Example 2

- Input: `s = "3[a2[c]]"`

- Output: `"accaccacc"`

### Example 3

- Input: `s = "2[abc]3[cd]ef"`

- Output: `"abcabccdcdcdef"`

## Constraints

- `1 <= s.length <= 30`

- `s` consists of lowercase English letters, digits, and square brackets `[]`.

- `s` is guaranteed to be a valid input.

- All integers in `s` are in the range `[1, 300]`.
