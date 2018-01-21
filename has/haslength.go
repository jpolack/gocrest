package has

import (
	"fmt"
	"github.com/corbym/gocrest"
	"reflect"
)

//Length can be called with arrays, maps, *gocrest.Matcher and strings but not numeric types.
//has.Length(is.GreaterThan(x)) is a valid call.
//Returns a matcher that matches if the length matches the given criteria
func Length(expected interface{}) *gocrest.Matcher {
	const description = "value with length %v"
	matcher := new(gocrest.Matcher)
	matcher.Describe = fmt.Sprintf(description, expected)
	matcher.Matches = func(actual interface{}) bool {
		if actual == nil {
			return false
		}
		lenOfActual := reflect.ValueOf(actual).Len()
		matcher.Actual = fmt.Sprintf("length was %d", lenOfActual)
		typeMatcher, ok := expected.(*gocrest.Matcher)
		if ok {
			matcher.Describe = fmt.Sprintf(description, typeMatcher.Describe)
			return typeMatcher.Matches(lenOfActual)
		}
		return lenOfActual == expected
	}
	return matcher
}
