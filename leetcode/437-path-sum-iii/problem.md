# 437. Path Sum III

## Description

- Given the `root` of a binary tree and an integer `targetSum`, return the number of paths where the sum of the values along the path equals `targetSum`.

- The path does not need to start or end at the root or a leaf, but it must go downwards (i.e., traveling only from parent nodes to child nodes).

## Examples

### Example 1

- **Input:** `root = [10,5,-3,3,2,null,11,3,-2,null,1]`, `targetSum = 8`

    ```mermaid
    flowchart TD
        n0["10"]
        n1["5"]
        n2["-3"]
        n3["3"]
        n4["2"]
        null1["null"]
        n6["11"]
        n7["3"]
        n8["-2"]
        null2["null"]
        n10["1"]

        n0 -->|L| n1
        n0 -->|R| n2
        n1 -->|L| n3
        n1 -->|R| n4
        n2 -.->|L| null1
        n2 -->|R| n6
        n3 -->|L| n7
        n3 -->|R| n8
        n4 -.->|L| null2
        n4 -->|R| n10

        classDef nullNode fill:transparent,stroke:gray,stroke-dasharray:4 4,color:gray;
        class null1,null2 nullNode;
    ```

- **Output:** `3`

- **Explanation:** The paths that sum to `8` are `5 -> 3`, `5 -> 2 -> 1`, and `-3 -> 11`.

### Example 2

- **Input:** `root = [5,4,8,11,null,13,4,7,2,null,null,5,1]`, `targetSum = 22`

    ```mermaid
    flowchart TD
        n0["5"]
        n1["4"]
        n2["8"]
        n3["11"]
        null1["null"]
        n5["13"]
        n6["4"]
        n7["7"]
        n8["2"]
        null2["null"]
        null3["null"]
        n11["5"]
        n12["1"]

        n0 -->|L| n1
        n0 -->|R| n2
        n1 -->|L| n3
        n1 -.->|R| null1
        n2 -->|L| n5
        n2 -->|R| n6
        n3 -->|L| n7
        n3 -->|R| n8
        n5 -.->|L| null2
        n5 -.->|R| null3
        n6 -->|L| n11
        n6 -->|R| n12

        classDef nullNode fill:transparent,stroke:gray,stroke-dasharray:4 4,color:gray;
        class null1,null2,null3 nullNode;
    ```

- **Output:** `3`

## Constraints

- The number of nodes in the tree is in the range `[0, 1000]`.

- `-10^9 <= Node.val <= 10^9`

- `-1000 <= targetSum <= 1000`
