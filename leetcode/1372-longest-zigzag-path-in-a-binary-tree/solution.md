
## 题意理解

- 找出二叉树中连续 “左、右、左、右...” 交替移动的最长路径。

- 例如：

    ```text
        A
         \
          B
         / \
        C   D
       /     \
      E       F
       \
        G
         \
          H
    ```

    - 从 B 出发，可以这样走：B -> 右 -> D -> 左 -> E -> 右 -> G。

    - 方向依次是：右 → 左 → 右。

    - 它经过了 4 个节点：B、D、E、G。

    - 所以长度是：4 - 1 = 3

    - 到达 G 后，下一步必须向左走，因为上一步是向右走。但是 G 没有左子节点，因此路径结束。

- 条件：

    1. 可以从任意节点出发，不一定从根节点出发。

    2. 每走一步必须改变方向。

    3. 路径长度计算的是边数，也就是 “经过的节点数 − 1”。

## 方法一：暴力枚举

- 遍历二叉树中的每一个节点，把当前节点作为 ZigZag 路径的起点，分别计算向左和向右时能够走出的路径长度，取二者中的较大值更新全局最长长度。

- 设计递归函数 `zigZagLengthFromNode(currentNode, requiredNextDirection)`：

    - `currentNode` 表示当前节点，`requiredNextDirection` 表示从当前节点出发时，下一步必须走向左子节点还是右子节点。

    - 返回值表示从当前节点出发，最多还能走过多少条边。

    - 如果当前节点为空，或者指定方向上没有子节点，说明无法继续移动，返回 `0`。

    - 如果当前节点能走到下个节点，就要 +1 了，然后指定下个节点往哪个方向走，加上可能存在的子节点 ZigZag 路径：

        - 当 `requiredNextDirection == MoveLeft` 且左子节点存在时，沿左边走到 `currentNode.Left`，这一步贡献 `1` 条边；为了保持方向交替，下一步必须向右，因此返回 `1 + zigZagLengthFromNode(currentNode.Left, MoveRight)`。

        - 当 `requiredNextDirection == MoveRight` 且右子节点存在时，沿右边走到 `currentNode.Right`，这一步贡献 `1` 条边；为了保持方向交替，下一步必须向左，因此返回 `1 + zigZagLengthFromNode(currentNode.Right, MoveLeft)`。

- 最后一个测试用例超时。
