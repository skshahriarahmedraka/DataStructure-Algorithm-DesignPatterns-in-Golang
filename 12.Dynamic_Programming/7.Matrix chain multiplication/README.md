

## Matrix Chain Multiplication

1. Matrix Chain Multiplication
    - Given a sequence of matrices, find the most efficient way to multiply these matrices together.
    - Example
        - Input: [40, 20, 30, 10, 30]
        - Output: 26000
        - Explanation: There are 4 matrices of dimensions 40x20, 20x30, 30x10 and 10x30. Let the input 4 matrices be A, B, C and D. The most efficient way to multiply these matrices is ((AB)C)D. The number of operations are 20*30*10 + 40*20*10 + 40*10*30 = 26000.
    - [Leetcode](https://leetcode.com/problems/minimum-cost-to-merge-stones/)
2. Printing Matrix Chain Multiplication
    - Given a sequence of matrices, find the most efficient way to multiply these matrices together and print the order of multiplication.
    - Example
        - Input: [40, 20, 30, 10, 30]
        - Output: ((AB)C)D
    - [Leetcode](https://leetcode.com/problems/minimum-cost-to-merge-stones/)
3. Evaluate Expression to True / Boolean Parenthesization
    - Given a boolean expression with following symbols.
        - Symbols
            - 'T' -> true
            - 'F' -> false
        - And following operators filled between symbols
            - & -> boolean AND
            - | -> boolean OR
            - ^ -> boolean XOR
    - Count the number of ways we can parenthesize the expression so that the value of expression evaluates to true.
    - Example
        - Input: "T|T&F^T"
        - Output: 4
        - Explanation: The expression can be parenthesized in 4 ways ((T|T)&(F^T)), (T|(T&(F^T))), (((T|T)&F)^T) and (T|((T&F)^T)).
    - [Leetcode](https://leetcode.com/problems/boolean-parenthesization/)
4. Minimum and Maximum Values of an Expression
    - Given a string `s` which represents an expression, we need to find the minimum and maximum value which can be obtained by evaluating this expression by different parenthesization.
    - Example
        - Input: "1+2*3+4*5"
        - Output: Minimum: 27, Maximum: 105
        - Explanation: Minimum evaluation: (1+((2*3)+(4*5))) = 27, Maximum evaluation: (((1+2)*3)+(4*5)) = 105
    - [Leetcode](https://leetcode.com/problems/different-ways-to-add-parentheses/)

5. Palindrome Partitioning
    - Given a string `s`, partition `s` such that every substring of the partition is a palindrome. Return all possible palindrome partitioning of `s`.
    - Example
        - Input: "aab"
        - Output: [["a","a","b"],["aa","b"]]
        - Explanation: The palindrome partitions are ["a","a","b"] and ["aa","b"].
    - [Leetcode](https://leetcode.com/problems/palindrome-partitioning/)
6. Scramble String
    - Given a string `s1`, we may represent it as a binary tree by partitioning it to two non-empty substrings recursively. To scramble the string, we may choose any non-leaf node and swap its two children. Given two strings `s1` and `s2` of the same length, determine if `s2` is a scrambled string of `s1`.
    - Example
        - Input: s1 = "great", s2 = "rgeat"
        - Output: true
        - Explanation: One possible scramble is "rgeat" which is a scrambled form of "great".
    - [Leetcode](https://leetcode.com/problems/scramble-string/)
7. Egg Dropping Puzzle
    - You are given `k` eggs, and you have access to a building with `n` floors from `1` to `n`. Each egg is identical in function, and if an egg breaks, you cannot drop it again. You know that there exists a floor `f` with `0 <= f <= n` such that any egg dropped at a floor higher than `f` will break, and any egg dropped at or below floor `f` will not break. Each move, you may take an egg (if you have an unbroken one) and drop it from any floor `x` (with `1 <= x <= n`). Your goal is to know with certainty what the value of `f` is. What is the minimum number of moves that you need to know with certainty what `f` is, regardless of the initial value of `f`?
    - Example
        - Input: k = 2, n = 10
        - Output: 4
        - Explanation: The minimum number of moves is 4. We can drop the first egg from floor 4, 7, 9 and 10.
    - [Leetcode](https://leetcode.com/problems/super-egg-drop/)
