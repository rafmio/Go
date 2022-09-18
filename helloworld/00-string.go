String
In Go strings are different from other languages like Java, C++, Python, etc. 
It is a sequence of variable-width character where each and evry character is represented by one or more bytes using UTF-8 Encoding.

Strings are immutable chain of arbitrary bytes (including bytes with zero value) or strings are a read-only slice of bytes and bytes of the strings can be represented in the Unicode text using UTF-8 encoding.

In Go, strings are immutable once a string is created the value of the string cannot be changed. String are read-only. If you try to change it, then the compiler will throw an error. 

You can iterate over string using for rang loop. This loop can iterate over the Unicode code point for a string. 

You can trim a string in different ways using under strings package, so you have to import strings in your program to access these functions:
- Trim
- TrimLeft
- TrimRight
- TrimSpace 
- TrimSuffix
- TrimPrefix



// https://www.geeksforgeeks.org/strings-in-golang/
// https://www.geeksforgeeks.org/how-to-trim-a-string-in-golang/
