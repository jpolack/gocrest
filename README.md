# gocrest

A hamcrest-like assertion library for Go. 

Inspired by [Hamcrest](https://github.com/hamcrest). 

## Package import

```
import (
  gocrest "github.com/corbym/gocrest"
)
```

## Example:
```
gocrest.AssertThat(testing, "hi", gocrest.EqualTo("bye"))
```

output:

```
expected: value equal to bye but was: hi
```

# Matchers so far..

- EqualTo(x)
- IsNil()
- Contains(x) -- acts like containsAll
- Not(m *Matcher) -- logical not of matcher's result
- MatchesPattern(regex string) -- a string regex expression
- HasFunction(string) - checks if a Type has a function (method)
- AllOf(... *Matcher) - returns true if all matchers match
- AnyOf(... *Matcher) - return true if any matcher matches