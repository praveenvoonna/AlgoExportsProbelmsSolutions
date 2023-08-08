# One Edit

You're given two strings stringOne and stringTwo. Write a function that determines if these two strings can be made equal using only one edit.

There are 3 possible edits:
• Replace: One character in one string is swapped for a different character.
• Add: One character is added at any index in one string.
• Remove: One character is removed at any index in one string.

Note that both strings will contain at least one character. If the strings are the same, your function should return true.

## Sample Input
stringOne = "hello"
stringTwo = "hollo"

## Sample Output
True // A single replace at index 1 of either string can make the strings equal

## Hints

### Hint 1
If the difference in lengths of the strings is greater than 1, then there is no way to make them equal with a single edit.

### Hint 2
If the lengths of the strings are the same, then the only possible edit is a replace, because adding or removing a character would 
make the strings different lengths.

### Hint 3
If the strings are different lengths, the only possible moves are adding and removing a character. 
These are essentially the same operation, because they represent the case where one string has a character that another does not.

## Optimal Space & Time Complexity
O(n) time | O(1) space - where n is the length of the shorter string