Simple permutation generator written in Golang

It allows for custom word length, word separator and charset

possible arguments:

```
    --sep, --separator
        sets a custom separator string, by default it's space

    --len, --length
        sets a custom minimal word length, by default it's 3

    --chrs, ––charset
        a charset to be used, by default the 26-letter lowercase Latin alphabet is used

```

example uses:

```
./main --len=3 --sep=... --chrs=abc
./main
./main --len=10
```
