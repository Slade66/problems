# 700. Search in a Binary Search Tree

## Description

- You are given the `root` of a binary search tree (BST) and an integer `val`.

- Find the node in the BST that the node's value equals `val` and return the subtree rooted with that node. If such a node does not exist, return `null`.

## Examples

### Example 1

- **Input:** `root = [4,2,7,1,3]`, `val = 2`

    ```mermaid
    flowchart TD
        n0["4"]
        n1["2"]
        n2["7"]
        n3["1"]
        n4["3"]

        n0 -->|L| n1
        n0 -->|R| n2
        n1 -->|L| n3
        n1 -->|R| n4
    ```

- **Output:** `[2,1,3]`

    ```mermaid
    flowchart TD
        n1["2"]
        n3["1"]
        n4["3"]

        n1 -->|L| n3
        n1 -->|R| n4
    ```

### Example 2

- **Input:** `root = [4,2,7,1,3]`, `val = 5`

    ```mermaid
    flowchart TD
        n0["4"]
        n1["2"]
        n2["7"]
        n3["1"]
        n4["3"]

        n0 -->|L| n1
        n0 -->|R| n2
        n1 -->|L| n3
        n1 -->|R| n4
    ```

- **Output:** `[]`

## Constraints

- The number of nodes in the tree is in the range `[1, 5000]`.

- `1 <= Node.val <= 10^7`

- `root` is a binary search tree.

- `1 <= val <= 10^7`
