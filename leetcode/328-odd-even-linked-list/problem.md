# 328. Odd Even Linked List

## Description

- Given the `head` of a singly linked list, group all nodes with odd indices together, followed by all nodes with even indices, and return the reordered list.

- The first node is considered odd, the second node is considered even, and so on.

- The relative order of nodes within both the odd and even groups must remain unchanged from the input.

- Solve the problem using `O(1)` extra space and `O(n)` time.

## Examples

### Example 1

- **Input:** `head = [1,2,3,4,5]`

- **Output:** `[1,3,5,2,4]`

### Example 2

- **Input:** `head = [2,1,3,5,6,4,7]`

- **Output:** `[2,3,6,7,1,5,4]`

## Constraints

- The number of nodes in the linked list is in the range `[0, 10^4]`.

- `-10^6 <= Node.val <= 10^6`
