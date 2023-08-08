# Common Characters

Write a function that takes in a non-empty list of non-empty strings and returns a list of characters that are common to 
all strings in the list, ignoring multiplicity.

Note that the strings are not guaranteed to only contain alphanumeric characters. The list you return can be in any order.

## Sample Input
strings = ["abc", "bcd", "cbaccd"]

## Sample Output
["b", "c"] // The characters could be ordered differently.

## Hints

### Hint 1
What data structure could be helpful to remember characters we've seen and how many strings contained those characters?

### Hint 2
We can use a map to keep track of the characters we have seen and how many strings we have seen them in. 
If a character is seen len(strings) times, then it must be in every string.

### Hint 3
Converting a string to a set can quickly get all of the unique characters from that string, which can be helpful since we 
are ignoring multiplicity in this problem.

## Optimal Space & Time Complexity
O(n* m) time | O(m) space - where n is the number of strings, and m is the length of the longest string