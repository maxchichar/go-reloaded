# go-reloaded

## Project Information
**Repo to create for this project:**
https://github.com/maxchichar/go-reloaded.git

**Language:** Go


**Skills:**
- go: 20%
- algo: 15%

**Difficulty:** 1 (Mandatory)


## Introduction

Your project must be written in Go.

The code should respect the good practices.

It is recommended to have test files for unit testing.

The tool you are about to build will receive as arguments the name of a file containing a text that needs some modifications (the input) and the name of the file the modified text should be placed in (the output). Next is a list of possible modifications that your program should execute:

### Text Modifications

1. **Hex Conversion (hex):** Every instance of `(hex)` should replace the word before with the decimal version of the word (in this case the word will always be a hexadecimal number).
   - Example: `"1E (hex) files were added"` → `"30 files were added"`

2. **Binary Conversion (bin):** Every instance of `(bin)` should replace the word before with the decimal version of the word (in this case the word will always be a binary number).
   - Example: `"It has been 10 (bin) years"` → `"It has been 2 years"`

3. **Uppercase (up):** Every instance of `(up)` converts the word before with the Uppercase version of it.
   - Example: `"Ready, set, go (up) !"` → `"Ready, set, GO!"`

4. **Lowercase (low):** Every instance of `(low)` converts the word before with the Lowercase version of it.
   - Example: `"I should stop SHOUTING (low)"` → `"I should stop shouting"`

5. **Capitalize (cap):** Every instance of `(cap)` converts the word before with the capitalized version of it.
   - Example: `"Welcome to the Brooklyn bridge (cap)"` → `"Welcome to the Brooklyn Bridge"`

6. **Multiple Words with (low), (up), (cap):** If a number appears next to it, like so: `(low, <number>)`, `(up, <number>)`, or `(cap, <number>)`, it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly.
   - Example: `"This is so exciting (up, 2)"` → `"This is SO EXCITING"`

7. **Punctuation Spacing:** Every instance of the punctuations `.`, `,`, `!`, `?`, `:` and `;` should be close to the previous word and with space apart from the next one.
   - Example: `"I was sitting over there ,and then BAMM !!"` → `"I was sitting over there, and then BAMM!!"`
   - **Exception:** If there are groups of punctuation like `...` or `!?`, the program should format the text accordingly.
   - Example: `"I was thinking ... You were right"` → `"I was thinking... You were right"`

8. **Quotation Marks ('):** The punctuation mark `'` will always be found with another instance of it and they should be placed to the right and left of the word in the middle of them, without any spaces.
   - Example: `"I am exactly how they describe me: ' awesome '"` → `"I am exactly how they describe me: 'awesome'"`
   - **Multiple Words:** If there are more than one word between the two `'` marks, the program should place the marks next to the corresponding words.
   - Example: `"As Elton John said: ' I am the most well-known homosexual in the world '"` → `"As Elton John said: 'I am the most well-known homosexual in the world'"`

9. **A/An Article Correction:** Every instance of `a` should be turned into `an` if the next word begins with a vowel (a, e, i, o, u) or an h.
   - Example: `"There it was. A amazing rock!"` → `"There it was. An amazing rock!"`

## Allowed Packages

Standard Go packages are allowed.

## Usage

### Example 1: Capitalization and Uppercase
$ cat sample.txt it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.

$ go run . sample.txt result.txt

$ cat result.txt It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.


### Example 2: Number Conversion
$ cat sample.txt Simply add 42 (hex) and 10 (bin) and you will see the result is 68.

$ go run . sample.txt result.txt

$ cat result.txt Simply add 66 and 2 and you will see the result is 68.



### Example 3: Article Correction
$ cat sample.txt There is no greater agony than bearing a untold story inside you.

$ go run . sample.txt result.txt

$ cat result.txt There is no greater agony than bearing an untold story inside you.



### Example 4: Punctuation Formatting
$ cat sample.txt Punctuation tests are ... kinda boring ,what do you think ?

$ go run . sample.txt result.txt

$ cat result.txt Punctuation tests are... kinda boring, what do you think?

## Project Tree
```
go-reloaded/
├── go.mod
├── main.go
├── functions/          # or helpers/  ← put your reused functions here
│   ├── todec.go
│   ├── case.go
│   ├── punctuation.go
│   ├── a_to_an.go
│   ├── split.go
│   └── ...
├── sample.txt          # test inputs
├── result.txt          # your output
└── README.md
```

## Learning Outcomes

This project will help you learn about:

- The Go file system (fs) API
- String and numbers manipulation

