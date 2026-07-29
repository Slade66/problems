# 1448. Count Good Nodes in Binary Tree

## Description

- Given a binary tree `root`, a node `X` in the tree is named **good** if in the path from root to `X` there are no nodes with a value greater than `X`.

- Return the number of **good** nodes in the binary tree.

## Examples

### Example 1

- **Input:** `root = [3,1,4,3,null,1,5]`

    ```mermaid
    flowchart TD
        n0["3"]
        n1["1"]
        n2["4"]
        n3["3"]
        null1["null"]
        n5["1"]
        n6["5"]

        n0 -->|L| n1
        n0 -->|R| n2
        n1 -->|L| n3
        n1 -.->|R| null1
        n2 -->|L| n5
        n2 -->|R| n6

        classDef nullNode fill:transparent,stroke:gray,stroke-dasharray:4 4,color:gray;
        class null1 nullNode;
    ```

- **Output:** `4`

- **Explanation:** Root Node (`3`) is always a good node. Node `4` -> (`3`, `4`) is the maximum value in the path starting from the root. Node `5` -> (`3`, `4`, `5`) is the maximum value in the path. Node `3` -> (`3`, `1`, `3`) is the maximum value in the path.

### Example 2

- **Input:** `root = [3,3,null,4,2]`

    ```mermaid
    flowchart TD
        n0["3"]
        n1["3"]
        null1["null"]
        n3["4"]
        n4["2"]

        n0 -->|L| n1
        n0 -.->|R| null1
        n1 -->|L| n3
        n1 -->|R| n4

        classDef nullNode fill:transparent,stroke:gray,stroke-dasharray:4 4,color:gray;
        class null1 nullNode;
    ```

- **Output:** `3`

- **Explanation:** Node `2` -> (`3`, `3`, `2`) is not good, because `3` is higher than it.

### Example 3

- **Input:** `root = [1]`

    ```mermaid
    flowchart TD
        n0["1"]
    ```

- **Output:** `1`

- **Explanation:** Root is considered as **good**.

## Constraints

- The number of nodes in the binary tree is in the range `[1, 10^5]`.

- Each node's value is between `[-10^4, 10^4]`.
