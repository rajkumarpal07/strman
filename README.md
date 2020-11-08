# Strman package
A Go string manipulation library. By [Rajkumar Palani](https://www.rajkumarpalani.com)

[![License](https://img.shields.io/:license-mit-blue.svg)](./LICENSE.md)


A Go library for working with Strings. You can learn about all the String utility functions implemented in strman library by reading the documentation.

## Installation

~~~~
go get github.com/rajkumarpal07/strman
~~~~

## Package import

```
import (
  "github.com/rajkumarpal07/strman"
)
```

## Documentation
### Supported Methods
1. **Appends** - Appends strings to another string.
2. **AppendsArray** - Appends an array of String to a string.
3. **At** - Get the character at index.
4. **Base64Decode** - Decodes data encoded with MIME base64.
5. **Base64Encode** - Encodes data with MIME base64.
6. **Between** - Returns an array with strings between start and end.
7. **Capitalize** - Converts the first character of string to upper case and the remaining to lower case.
8. **Chars** - Returns a String array consisting of the characters in the String.
9. **CharsCount** - Counts the number of occurrences of each character in the string.
10. **CollapseWhitespace** - Replace consecutive whitespace characters with a single space.
11. **Contains** - Verifies that the needle is contained in the value. Case sensitive.
12. **ContainsAll** - Verifies that all needles are contained in value. Case insensitive.
13. **ContainsAllCaseSensitive** - Verifies that all needles are contained in value. Case sensitive.
14. **ContainsAny** - Verifies that one or more of needles are contained in value. Case insensitive.
15. **ContainsAnyCaseSensitive** - Verifies that one or more of needles are contained in value.
16. **CountSubstr** - Count the number of times substr appears in value. Case Sensitive.
17. **CountSubstring** - Count the number of times substr appears in value.
18. **CountsSubstr** - Counts the number of times substr appears in value.
19. **Dasherize** - Converts a underscored or camelized string into an dasherized one.
20. **DecEncode** - Convert string chars to decimal unicode (16 digits).
21. **EndsWith** - Tests if value ends with search. 
22. **EndsWithCase** - Tests if value ends with search. Case sensitive.
23. **EnsureLeft** - Ensures that the value begins with prefix. If it doesn't exist, it's prepended. Case sensitive.
24. **EnsureLeftWithCase** - Ensures that the value begins with prefix. If it doesn't exist, it's prepended.
25. **EnsureRight** - Ensures that the value ends with suffix. If it doesn't, it's appended. Case sensitive.
26. **EnsureRightWithCase** - Ensures that the value ends with suffix. If it doesn't, it's appended. Case sensitive.
27. **First** - Returns the first 3 chars of String(ASCII).
28. **FirstNChars** - Returns the first n chars of String(ASCII).
29. **Format** - Formats a string using parameters. Will be added in a future release.
30. **FormatNumber** - Returns a string representation of the number where groups of three digits are delimited by comma
31. **HTMLDecode** - Converts all HTML entities to applicable characters.
32. **HTMLEncode** - Convert all applicable characters to HTML entities.
33. **Head** - Return the first char of String(ASCII).
34. **HexDecode** - Convert hexadecimal unicode (4 digits) string to string chars.
35. **HexEncode** - Convert string chars to hexadecimal unicode (4 digits).
36. **Humanize** - Converts an underscored, camelized, or dasherized string into a humanized one.
37. **IndexOf** - Returns the index within the calling String of the first occurrence of the specified value.
38. **Insert** - Inserts 'substring' into the 'value' at the 'index' provided.
39. **IsBlank** - Checks if string is empty.
40. **IsEnclosedBetween** - Verifies whether String is enclosed by same encloser.
41. **IsEnclosedBetweenTwo** - Verifies whether String is enclosed by 2 different enclosers.
42. **IsLowerCase** - Verifies if String is lowercase.
43. **IsString** - Checks whether parameter is String.
44. **IsUpperCase** - Verifies if String is uppercase.
45. **Join** - concatenates all the elements of the strings array into a single String.
46. **Last** - Returns the last n chars of String.
47. **LastIndexOf** - returns the index within the calling String object of the last occurrence of the specified value.
48. **LastIndexOfWithCase** - returns the index of the last occurrence of the specified value. Case sensitive.
49. **LeftPad** - Returns a new string such that the beginning of the string is padded.
50. **LeftPad2Len** - Returns a new string of a given length such that the beginning of the string is padded.
51. **LeftTrim** - Removes all spaces on left.
52. **Length** - Returns length of String. Delegates to len method.
53. **Lines** - Split lines to an array.
54. **LowerFirst** - Converts the first character of string to lower case.
55. **Prepend** - Returns a new String starting with 'prepends' string.
56. **PrependArray** - Returns a new String starting with 'prepends' string array.
57. **RemoveEmptyStrings** - Remove empty Strings from string array.
58. **RemoveLeft** - Returns a new String with the prefix removed, if present. Case sensitive.
59. **RemoveLeftWithCase** - Returns a new String with the prefix removed, if present.
60. **RemoveNonWords** - Remove all non word characters.
61. **RemoveRight** - Returns a new string with the 'suffix' removed, if present. Case sensitive.
62. **RemoveRightWithCase** - Returns a new string with the 'suffix' removed, if present.
63. **RemoveSpaces** - Removes all spaces.
64. **Repeat** - Returns a repeated string given a multiplier.
65. **Replace** - Replace all occurrences of 'search' value to 'newvalue'.
66. **Reverse** - Reverse the input String.
67. **RightPad** - Returns a new string such that the ending of the string is padded.
68. **RightPad2Len** - Returns a new string of a given length such that the ending of the string is padded.
69. **RightTrim** - Removes all spaces on right.
70. **SafeTruncate** - Truncate the string securely, not cutting a word in half.
71. **Shuffle** - Returns a string with its characters in random order.
72. **Slice** - A substring method.
73. **Slugify** - Convert a String to a slug.
74. **Split** - Alias to String split function.
75. **Surround** - Surrounds a 'value' with the given 'prefix' and 'suffix'.
76. **SwapCase** - Returns a copy of the string in which all the case-based characters have had their case swapped.
77. **Tail** - Returns tail of the String.
78. **ToCamelCase** - Transform to camelCase.
79. **ToDecamelize** - Decamelize String.
80. **ToKebabCase** - Transform to kebab-case.
81. **ToSnakeCase** - Transform to snake_case.
82. **ToStudlyCase** - Transform to StudlyCaps.
83. **Transliterate** - Remove all non valid characters.
84. **TrimEnd** - Removes trailing chars from string.
85. **TrimEndSpaces** - Removes trailing whitespace from string.
86. **TrimStart** - Removes leading chars from string.
87. **TrimStartSpaces** - Removes leading whitespace from string.
88. **Truncate** - Truncates the unsecured form string, cutting the independent string of required position.
89. **UnEqual** - Tests if two Strings are InEqual.
90. **Underscored** - Changes passed in string to all lower case and adds underscore between words.
91. **UpperFirst** - Converts the first character of string to upper case.
92. **Words** - Splits a String to words.
93. **WordsDelimitedBy** - Splits a String to words with a delimiter.
94. **Zip** - Aggregates the contents of 2 strings arrays into a single list of tuples.



## Inspiration

This library is inspired by [dleitee/strman](https://github.com/dleitee/strman) and 
[shekhargulati/strman](https://github.com/shekhargulati/strman-java).

License
-------
strman is licensed under the MIT License - see the `LICENSE` file for details.
