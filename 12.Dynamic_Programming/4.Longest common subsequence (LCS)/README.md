
## Longest common subsequence (LCS)

1. largest common sub string
    - Given two strings `s1` and `s2`, find the length of the longest subsequence present in both of them.
    - A subsequence is a sequence that can be derived from another sequence by deleting some or no elements without changing the order of the remaining elements.
    - Example
        - Input: s1 = "abdca", s2 = "cbda"
        - Output: 3
        - Explanation: The longest common subsequence is "bda".
    - [Leetcode](https://leetcode.com/problems/longest-common-subsequence/)
2. Print LCS 
    - Given two strings `s1` and `s2`, find the longest subsequence present in both of them.
    - A subsequence is a sequence that can be derived from another sequence by deleting some or no elements without changing the order of the remaining elements.
    - Example
        - Input: s1 = "abdca", s2 = "cbda"
        - Output: "bda"
    - [Leetcode](https://leetcode.com/problems/longest-common-subsequence/)
3. Shortest common supersequence
    - Given two strings `s1` and `s2`, find the length of the shortest string that has both `s1` and `s2` as subsequences.
    - Example
        - Input: s1 = "abac", s2 = "cab"
        - Output: 5
        - Explanation: The shortest common supersequence is "cabac".
    - [Leetcode](https://leetcode.com/problems/shortest-common-supersequence/)
4. Print SCS
    - Given two strings `s1` and `s2`, find the shortest string that has both `s1` and `s2` as subsequences.
    - Example
        - Input: s1 = "abac", s2 = "cab"
        - Output: "cabac"
    - [Leetcode](https://leetcode.com/problems/shortest-common-supersequence/)
5. Minimum number of deletions and insertions to convert a string into another
    - Given two strings `s1` and `s2`, find the minimum number of deletions and insertions needed to convert `s1` into `s2`.
    - Example
        - Input: s1 = "abc", s2 = "fbc"
        - Output: 1 deletion and 1 insertion
        - Explanation: We need to delete the character 'a' from `s1` and insert the character 'f' in `s1` to convert it into `s2`.
    - [Leetcode](https://leetcode.com/problems/delete-operation-for-two-strings/)
6. Largest repeating sequence
    - Given a string `s`, find the length of the largest repeating subsequence.
    - Example
        - Input: s = "tomorrow"
        - Output: 2
        - Explanation: The largest repeating subsequence is "or".
    - [Leetcode](https://leetcode.com/problems/longest-repeating-subsequence/)
7. Length of largest subsequence of "A" which is a subsequence in "B"
    - Given two strings `s1` and `s2`, find the length of the longest subsequence of `s1` which is a subsequence in `s2`.
    - Example
        - Input: s1 = "abcde", s2 = "ace"
        - Output: 3
        - Explanation: The longest subsequence of `s1` which is a subsequence in `s2` is "ace".
    - [Leetcode](https://leetcode.com/problems/longest-common-subsequence/)
8. subsequence pattern matching
    - Given a string `s` and a pattern `p`, find if `p` is a subsequence of `s`.
    - Example
        - Input: s = "abppplee", p = "apple"
        - Output: true
        - Explanation: The pattern "apple" is a subsequence of the string "abppplee".
    - [Leetcode](https://leetcode.com/problems/is-subsequence/)
9. Count how many times "A" appears as a subsequence in "B"
    - Given two strings `s1` and `s2`, find the number of times `s1` appears as a subsequence in `s2`.
    - Example
        - Input: s1 = "abc", s2 = "abcabc"
        - Output: 3
        - Explanation: The string "abc" appears as a subsequence in the string "abcabc" three times.
    - [Leetcode](https://leetcode.com/problems/distinct-subsequences/)
10. Longest palindromic subsequence
    - Given a string `s`, find the length of the longest palindromic subsequence.
    - Example
        - Input: s = "bbbab"
        - Output: 4
        - Explanation: The longest palindromic subsequence is "bbbb".
    - [Leetcode](https://leetcode.com/problems/longest-palindromic-subsequence/)
11. Longest palindromic sub string
    - Given a string `s`, find the length of the longest palindromic sub string.
    - Example
        - Input: s = "babad"
        - Output: 3
        - Explanation: The longest palindromic sub string is "aba".
    - [Leetcode](https://leetcode.com/problems/longest-palindromic-substring/)
12. count of palindromic sub strings
    - Given a string `s`, find the number of palindromic sub strings.
    - Example
        - Input: s = "abc"
        - Output: 3
        - Explanation: The palindromic sub strings are "a", "b", and "c".
    - [Leetcode](https://leetcode.com/problems/palindromic-substrings/)
13. minimum number of deletations to make a string palindrome
    - Given a string `s`, find the minimum number of deletions needed to make `s` a palindrome.
    - Example
        - Input: s = "abc"
        - Output: 2
        - Explanation: We need to delete the characters 'b' and 'c' to make the string "abc" a palindrome.
    - [Leetcode](https://leetcode.com/problems/minimum-number-of-deletions-to-make-palindrome/)
14. minimum number of insertions to make a string palindrome
    - Given a string `s`, find the minimum number of insertions needed to make `s` a palindrome.
    - Example
        - Input: s = "abc"
        - Output: 2
        - Explanation: We need to insert the characters 'b' and 'c' to make the string "abc" a palindrome.
    - [Leetcode](https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/)