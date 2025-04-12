---
title: Regular Expressions for Python
date: 2025-04-11
tags: [regex, regular expressions, pattern matching, text processing, python, find and replace]
subject: Python

---

## Introduction

Regular expressions (regex) are a powerful tool for pattern matching and text manipulation. They allow you to search, extract, and replace specific patterns within strings, making them invaluable for tasks like data validation, parsing, and search and replace operations. This guide will introduce you to the fundamental concepts and syntax of regular expressions.

## Basic Syntax

*   **Literal characters:** Match exact characters. For example, the regex `hello` will match the string "hello".
*   **Character classes:** Match sets of characters.

    *   `[a-z]`: Matches any lowercase letter.
    *   `[A-Z]`: Matches any uppercase letter.
    *   `[0-9]`: Matches any digit.
    *   `[a-zA-Z0-9]`: Matches any alphanumeric character.
    *   `[^a-z]`: Matches any character *except* lowercase letters.
*   **Anchors:** Match the beginning or end of a string.

    *   `^`: Matches the beginning of the string.  For example, `^hello` matches strings that start with "hello".
    *   `$`: Matches the end of the string. For example, `hello$` matches strings that end with "hello".
*   **Quantifiers:** Specify how many times a character or group should be repeated.

    *   `*`: Matches zero or more occurrences.  For example, `a*` matches "", "a", "aa", "aaa", etc.
    *   `+`: Matches one or more occurrences. For example, `a+` matches "a", "aa", "aaa", etc., but not "".
    *   `?`: Matches zero or one occurrence. For example, `a?` matches "" or "a".
    *   `{n}`: Matches exactly n occurrences. For example, `a{3}` matches "aaa".
    *   `{n,}`: Matches n or more occurrences. For example, `a{2,}` matches "aa", "aaa", "aaaa", etc.
    *   `{n,m}`: Matches between n and m occurrences. For example, `a{2,4}` matches "aa", "aaa", or "aaaa".

## Special Characters

*   `.`: Matches any character (except newline).
*   `\d`: Matches a digit (equivalent to `[0-9]`).
*   `\w`: Matches a word character (alphanumeric and underscore) (equivalent to `[a-zA-Z0-9_]`).
*   `\s`: Matches whitespace (space, tab, newline).
*   `\D`: Matches a non-digit character.
*   `\W`: Matches a non-word character.
*   `\S`: Matches a non-whitespace character.

## Grouping and Capturing

Parentheses `()` are used to group parts of a regular expression. They also create capturing groups, which allow you to extract the matched portions of the text.

*   `(pattern)`: Groups the `pattern`.
*   `\1`, `\2`, etc.:  Refer to the captured groups (backreferences).

**Example:**

To extract the area code and phone number from a string like "(123) 456-7890":

```regex
\((\d{3})\) (\d{3}-\d{4})
