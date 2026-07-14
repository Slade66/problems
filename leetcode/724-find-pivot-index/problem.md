# 724. Find Pivot Index

## Description

- Given an array of integers `nums`, calculate the pivot index of this array.

- The pivot index is the index where the sum of all the numbers strictly to the left of the index is equal to the sum of all the numbers strictly to the right.

- If the index is on the left edge of the array, then the left sum is `0` because there are no elements to the left. This also applies to the right edge of the array.

- Return the leftmost pivot index. If no such index exists, return `-1`.

- This question is the same as 1991: [Find the Middle Index in Array](https://leetcode.com/problems/find-the-middle-index-in-array/).

## Examples

### Example 1

- Input: `nums = [1, 7, 3, 6, 5, 6]`

- Output: `3`

- Explanation: The pivot index is `3`. The left sum is `1 + 7 + 3 = 11`, and the right sum is `5 + 6 = 11`.

### Example 2

- Input: `nums = [1, 2, 3]`

- Output: `-1`

- Explanation: There is no index that satisfies the conditions in the problem statement.

### Example 3

- Input: `nums = [2, 1, -1]`

- Output: `0`

- Explanation: The pivot index is `0`. The left sum is `0`, and the right sum is `1 + -1 = 0`.

## Constraints

- `1 <= nums.length <= 10^4`

- `-1000 <= nums[i] <= 1000`
