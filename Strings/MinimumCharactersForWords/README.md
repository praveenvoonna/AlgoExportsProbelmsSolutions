# Minimum Characters For Words

Write a function that takes in an array of words and returns the smallest array of characters needed to form all of the words. 
The characters don't need to be in any particular order.

For example, the characters ["y", "r", "o", "u"] are needed to form the words ["your", "you", "or", "yo"].

Note: the input words won't contain any spaces; however, they might contain punctuation and/or special characters.

## Sample Input
words = ["this", "that", "did", "deed", "them!", "a"]

## Sample Output
["t", "t", "h", "i", "s", "a", "d", "d", "e", "e", "m", "!"]
// The characters could be ordered differently.

## Hints
### Hint 1
There are a few different ways to solve this problem, but all of them use the same general approach. 
You'll need to determine not only all of the unique characters required to form the input words, but also their required frequencies. 
What determines the required frequencies of characters to form multiple words?

### Hint 2
The word that contains the highest frequency of any character dictates how many of those characters are required. 
For example, given words = ["A", "AAAA"] you need 4 A s, because the word that contains the most of amount of As has 4.

### Hint 3
Use a hash table to keep track of the maximum frequencies of all unique characters that occur across all words. 
Count the frequency of each character in each word, and use those per-word frequencies to update your maximum-character-frequency hash table. 
Once you've determined the maximum frequency of each character across all words, you can use the built-up hash table to generate your output array.

## Optimal Space & Time Complexity
O(n * I) time | O(c) space - where n is the number of words, I is the length of the longest word, and c is the number of unique characters 
across all words See notes under video explanation for details about the space complexity.

