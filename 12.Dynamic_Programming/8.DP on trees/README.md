

## DP on Trees 

1. General Syntax 
    - The general syntax for a recursive function is as follows:
    ```python
    def recursive_function(root):
        if root is None:
            return 0
        left = recursive_function(root.left)
        right = recursive_function(root.right)
        return left + right
    ```

2. how DP can be applied to trees (identification)
    - The problem can be solved using a recursive function that returns the answer for the current node and the answer for the subtree rooted at the current node.
    - The recursive function can be called for the left and right children of the current node.
    - The answer for the current node can be calculated using the answers for the left and right children of the current node.
    - The answer for the subtree rooted at the current node can be calculated using the answers for the left and right children of the current node.

## Problems 

1. Diameter of a Binary Tree
    - Given a binary tree, find the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.
    - Example
        - Input: root = [1,2,3,4,5]
        - Output: 3
        - Explanation: The diameter of the tree is the path [4,2,1,3] or [5,2,1,3].
    - [Leetcode](https://leetcode.com/problems/diameter-of-binary-tree/)

2. Maximum Path Sum from any node to any
    - Given a binary tree, find the maximum path sum. The path may start and end at any node in the tree.
    - Example
        - Input: root = [-10,9,20,null,null,15,7]
        - Output: 42
        - Explanation: The maximum path sum is the path [15,20,7].
    - [Leetcode](https://leetcode.com/problems/binary-tree-maximum-path-sum/)

3. Maximum path sum from any leaf node to any leaf node
    - Given a binary tree, find the maximum path sum. The path may start and end at any node in the tree.
    - Example
        - Input: root = [-10,9,20,null,null,15,7]
        - Output: 42
        - Explanation: The maximum path sum is the path [15,20,7].
    - [Leetcode](https://leetcode.com/problems/binary-tree-maximum-path-sum/)

4. Diameter of N-ary Tree
    - Given a binary tree, find the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.
    - Example
        - Input: root = [1,2,3,4,5]
        - Output: 3
        - Explanation: The diameter of the tree is the path [4,2,1,3] or [5,2,1,3].
    - [Leetcode](https://leetcode.com/problems/diameter-of-n-ary-tree/)