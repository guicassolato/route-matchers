# Grammar

```bnf
               <Set> ::= <Value> | <Expression>
        <Expression> ::= <Set> "∪" <Set> | <Set> "∩" <Set> | <Set> "'" | "(" <Expression> ")"
             <Value> ::= <PrefixOrSuffixAttr> | <KeyValueAttr> | <Method>

<PrefixOrSuffixAttr> ::= <Hostname> | "*" "." <Hostname> | <Path> | <Path> "*"
          <Hostname> ::= <Word> | <Word> "." <Hostname>
              <Path> ::= "/" <PathSection>
       <PathSection> ::= "" | <Word> | <Word> "/" <PathSection>
            <Method> ::= "GET" | "HEAD" | "POST" | "PUT" | "DELETE" | "CONNECT" | "OPTIONS" | "TRACE" | "PATCH"

      <KeyValueAttr> ::= <Param> | <Header>
             <Param> ::= <Key> "=" <Word> | <Key> "=" "*"
            <Header> ::= <Key> ":" <Word> | <Key> ":" "*"
               <Key> ::= <Word>

              <Word> ::= <Letter> | <Letter> <Word> | <Word> <Digit> | <Word> <Symbol> <Word>
            <Letter> ::= "a" | "b" | "c" | "d" | "e" | "f" | "g" | "h" | "i" | "j" | "k" | "l" | "m" | "n" | "o" | "p" | "q" | "r" | "s" | "t" | "u" | "v" | "w" | "x" | "y" | "z"
             <Digit> ::= "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
            <Symbol> ::= "_" | "-"
```

## Examples

- **`(GET ∪ POST) ∩ toys.acme.com ∩ /dolls/*`**<br/>
  <sup>`GET` or `POST` requests to `toys.acme.com/dolls/*`.</sup>

  ```
  1.  <Set>
  2.  <Expression>
  3.  <Set> "∩" <Set>
  4.  <Expression> "∩" <Set>
  5.  "(" <Expression> ")" "∩" <Set>
  6.  "(" <Set> "∪" <Set> ")" "∩" <Set>
  7.  "(" <Value> "∪" <Set> ")" "∩" <Set>
  8.  "(" <Method> "∪" <Set> ")" "∩" <Set>
  9.  "(" "GET" "∪" <Set> ")" "∩" <Set>
  10. "(" "GET" "∪" <Value> ")" "∩" <Set>
  11. "(" "GET" "∪" <Method> ")" "∩" <Set>
  12. "(" "GET" "∪" "POST" ")" "∩" <Set>
  13. "(" "GET" "∪" "POST" ")" "∩" <Expression>
  14. "(" "GET" "∪" "POST" ")" "∩" <Set> "∩" <Set>
  15. "(" "GET" "∪" "POST" ")" "∩" <Value> "∩" <Set>
  16. "(" "GET" "∪" "POST" ")" "∩" <PrefixOrSuffixAttr> "∩" <Set>
  17. "(" "GET" "∪" "POST" ")" "∩" <Hostname> "∩" <Set>
  18. "(" "GET" "∪" "POST" ")" "∩" <Word> "." <Hostname> "∩" <Set>
  19. "(" "GET" "∪" "POST" ")" "∩" <Letter> <Word> "." <Hostname> "∩" <Set>
  20. "(" "GET" "∪" "POST" ")" "∩" "t" <Word> "." <Hostname> "∩" <Set>
  21. "(" "GET" "∪" "POST" ")" "∩" "t" <Letter> <Word> "." <Hostname> "∩" <Set>
  22. "(" "GET" "∪" "POST" ")" "∩" "t" "o" <Word> "." <Hostname> "∩" <Set>
  23. "(" "GET" "∪" "POST" ")" "∩" "t" "o" <Letter> <Word> "." <Hostname> "∩" <Set>
  24. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" <Letter> "." <Hostname> "∩" <Set>
  25. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." <Hostname> "∩" <Set>
  26. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." <Word> "." <Hostname> "∩" <Set>
  27. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." <Letter> <Word> "." <Hostname> "∩" <Set>
  28. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" <Word> "." <Hostname> "∩" <Set>
  29. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" <Letter> <Word> "." <Hostname> "∩" <Set>
  30. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" <Word> "." <Hostname> "∩" <Set>
  31. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" <Letter> <Word> "." <Hostname> "∩" <Set>
  32. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" <Word> "." <Hostname> "∩" <Set>
  33. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" <Letter> "." <Hostname> "∩" <Set>
  34. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." <Hostname> "∩" <Set>
  35. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." <Word> "∩" <Set>
  36. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." <Letter> <Word> "∩" <Set>
  37. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." "c" <Word> "∩" <Set>
  38. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." "c" <Letter> "∩" <Set>
  39. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." "c" "o" <Word> "∩" <Set>
  40. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." "c" "o" <Letter> "∩" <Set>
  41. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" <Set>
  42. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" <Value>
  43. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" <Path> "*"
  44. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "/" <PathSection> "*"
  45. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "/" <Word> "/" <PathSection> "*"
  46. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "/" <Letter> <Word> "/" <PathSection> "*"
  47. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "/" "d" <Word> "/" <PathSection> "*"
  48. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "/" "d" <Letter> <Word> "/" <PathSection> "*"
  49. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "/" "d" "o" <Word> "/" <PathSection> "*"
  50. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "/" "d" "o" <Letter> <Word> "/" <PathSection> "*"
  51. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "/" "d" "o" "l" <Word> "/" <PathSection> "*"
  52. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "/" "d" "o" "l" <Letter> <Word> "/" <PathSection> "*"
  53. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "/" "d" "o" "l" "l" <Word> "/" <PathSection> "*"
  54. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "/" "d" "o" "l" "l" <Letter> "/" <PathSection> "*"
  55. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "/" "d" "o" "l" "l" "s" "/" <PathSection> "*"
  56. "(" "GET" "∪" "POST" ")" "∩" "t" "o" "y" "s" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "/" "d" "o" "l" "l" "s" "/" "" "*"
  ```

- **`*.acme.com ∩ x-canary:true`**<br/>
  <sup>Requests to `*.acme.com` containing `x-canary:true` header.</sup>

  ```
  1.  <Set>
  2.  <Expression>
  3.  <Set> "∩" <Set>
  4.  <Value> "∩" <Set>
  5.  <PrefixOrSuffixAttr> "∩" <Set>
  6.  "*" "." <Hostname> "∩" <Set>
  7.  "*" "." <Word> "." <Hostname> "∩" <Set>
  8.  "*" "." <Letter> <Word> "." <Hostname> "∩" <Set>
  9.  "*" "." "a" <Word> "." <Hostname> "∩" <Set>
  10. "*" "." "a" <Letter> <Word> "." <Hostname> "∩" <Set>
  11. "*" "." "a" "c" <Word> "." <Hostname> "∩" <Set>
  12. "*" "." "a" "c" <Letter> <Word> "." <Hostname> "∩" <Set>
  13. "*" "." "a" "c" "m" <Word> "." <Hostname> "∩" <Set>
  14. "*" "." "a" "c" "m" <Letter> "." <Hostname> "∩" <Set>
  15. "*" "." "a" "c" "m" "e" "." <Hostname> "∩" <Set>
  16. "*" "." "a" "c" "m" "e" "." <Word> "∩" <Set>
  17. "*" "." "a" "c" "m" "e" "." <Letter> <Word> "∩" <Set>
  18. "*" "." "a" "c" "m" "e" "." "c" <Word> "∩" <Set>
  19. "*" "." "a" "c" "m" "e" "." "c" <Letter> <Word> "∩" <Set>
  20. "*" "." "a" "c" "m" "e" "." "c" "o" <Word> "∩" <Set>
  21. "*" "." "a" "c" "m" "e" "." "c" "o" <Letter> "∩" <Set>
  22. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" <Set>
  23. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" <Value>
  24. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" <KeyValueAttr>
  25. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" <Header>
  26. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" <Key> ":" <Word>
  27. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" <Word> ":" <Word>
  28. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" <Word> <Symbol> <Word> ":" <Word>
  29. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" <Letter> <Symbol> <Word> ":" <Word>
  30. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" <Symbol> <Word> ":" <Word>
  31. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" "-" <Word> ":" <Word>
  32. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" "-" <Letter> <Word> ":" <Word>
  33. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" "-" "c" <Word> ":" <Word>
  34. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" "-" "c" <Letter> <Word> ":" <Word>
  35. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" "-" "c" "a" <Word> ":" <Word>
  36. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" "-" "c" "a" <Letter> <Word> ":" <Word>
  37. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" "-" "c" "a" "n" <Word> ":" <Word>
  38. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" "-" "c" "a" "n" <Letter> <Word> ":" <Word>
  39. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" "-" "c" "a" "n" "a" <Word> ":" <Word>
  40. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" "-" "c" "a" "n" "a" <Letter> <Word> ":" <Word>
  41. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" "-" "c" "a" "n" "a" "r" <Word> ":" <Word>
  42. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" "-" "c" "a" "n" "a" "r" <Letter> ":" <Word>
  43. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" "-" "c" "a" "n" "a" "r" "y" ":" <Word>
  44. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" "-" "c" "a" "n" "a" "r" "y" ":" <Letter> <Word>
  45. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" "-" "c" "a" "n" "a" "r" "y" ":" "t" <Word>
  46. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" "-" "c" "a" "n" "a" "r" "y" ":" "t" <Letter> <Word>
  47. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" "-" "c" "a" "n" "a" "r" "y" ":" "t" "r" <Word>
  48. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" "-" "c" "a" "n" "a" "r" "y" ":" "t" "r" <Letter> <Word>
  49. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" "-" "c" "a" "n" "a" "r" "y" ":" "t" "r" "u" <Word>
  50. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" "-" "c" "a" "n" "a" "r" "y" ":" "t" "r" "u" <Letter>
  51. "*" "." "a" "c" "m" "e" "." "c" "o" "m" "∩" "x" "-" "c" "a" "n" "a" "r" "y" ":" "t" "r" "u" "e"
  ```
