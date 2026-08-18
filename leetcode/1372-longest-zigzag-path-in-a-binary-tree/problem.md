# 1372. Longest ZigZag Path in a Binary Tree

## Description

- You are given the `root` of a binary tree.

- A **ZigZag path** for a binary tree is defined as follows:

    - Choose **any** node in the binary tree and a direction (right or left).

    - If the current direction is right, move to the right child of the current node; otherwise, move to the left child.

    - Change the direction from right to left or from left to right.

    - Repeat the second and third steps until you cannot move in the tree.

- ZigZag length is defined as the number of nodes visited minus `1`. A single node has a length of `0`.

- Return the longest **ZigZag** path contained in that tree.

## Examples

### Example 1

- **Input:** `root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1]`

    ```mermaid
    flowchart TD
        n0["1"]
        null1["null"]
        n2["1"]
        n3["1"]
        n4["1"]
        n7["1"]
        n8["1"]
        null2["null"]
        n10["1"]
        null3["null"]
        n14["1"]

        n0 -.->|L| null1
        n0 -->|R| n2
        n2 -->|L| n3
        n2 -->|R| n4
        n4 -->|L| n7
        n4 -->|R| n8
        n7 -.->|L| null2
        n7 -->|R| n10
        n10 -.->|L| null3
        n10 -->|R| n14

        classDef nullNode fill:transparent,stroke:gray,stroke-dasharray:4 4,color:gray;
        class null1,null2,null3 nullNode;
    ```

- **Output:** `3`

- **Explanation:** One longest ZigZag path starts at the right child of the root and follows right -> left -> right.

### Example 2

- **Input:** `root = [1,1,1,null,1,null,null,1,1,null,1]`

    ```mermaid
    flowchart TD
        n0["1"]
        n1["1"]
        n2["1"]
        null1["null"]
        n4["1"]
        n7["1"]
        n8["1"]
        null2["null"]
        n10["1"]

        n0 -->|L| n1
        n0 -->|R| n2
        n1 -.->|L| null1
        n1 -->|R| n4
        n4 -->|L| n7
        n4 -->|R| n8
        n7 -.->|L| null2
        n7 -->|R| n10

        classDef nullNode fill:transparent,stroke:gray,stroke-dasharray:4 4,color:gray;
        class null1,null2 nullNode;
    ```

- **Output:** `4`

- **Explanation:** The longest ZigZag path starts at the root and follows left -> right -> left -> right.

### Example 3

- **Input:** `root = [1]`

    ```mermaid
    flowchart TD
        n0["1"]
    ```

- **Output:** `0`

## Constraints

- The number of nodes in the tree is in the range `[1, 5 * 10^4]`.

- `1 <= Node.val <= 100`
