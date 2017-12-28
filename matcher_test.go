package gocrest_test

import (
	"testing"
	"gocrest"
	"strings"
)

var stubTestingT *gocrest.StubTestingT

func init() {
	stubTestingT = new(gocrest.StubTestingT)
}

func TestAssertThatTwoValuesAreEqualOrNot(testing *testing.T) {
	var equalsItems = []struct {
		actual     interface{}
		expected   interface{}
		shouldFail bool
	}{
		{actual: 1, expected: 1, shouldFail: false},
		{actual: 1.12, expected: 1.12, shouldFail: false},
		{actual: 1, expected: 2, shouldFail: true},
		{actual: "hi", expected: "bees", shouldFail: true},
	}
	for _, test := range equalsItems {
		stubTestingT = new(gocrest.StubTestingT)
		gocrest.AssertThat(stubTestingT, test.actual, gocrest.EqualTo(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, EqualTo(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}

func TestEqualToFailsWithDescriptionTest(testing *testing.T) {
	gocrest.AssertThat(stubTestingT, 1, gocrest.EqualTo(2))
	if stubTestingT.MockTestOutput != "expected: value equal to 2 but was: 1" {
		testing.Errorf("did not get expected description, got %s", stubTestingT.MockTestOutput)
	}
}

func TestAssertThatFailsTest(testing *testing.T) {
	gocrest.AssertThat(stubTestingT, 1, gocrest.EqualTo(2))
	if !stubTestingT.HasFailed() {
		testing.Error("1 EqualTo 2 did not fail test")
	}
}

func TestAssertThatTwoValuesAreGreaterThanOrNot(testing *testing.T) {
	var equalsItems = []struct {
		actual     interface{}
		expected   interface{}
		shouldFail bool
	}{
		{actual: 1, expected: 1, shouldFail: true},
		{actual: 2, expected: 1, shouldFail: false},
		{actual: 1.12, expected: 1.12, shouldFail: true},
		{actual: float32(1.12), expected: float32(1.0), shouldFail: false},
		{actual: float64(1.24), expected: float64(1.0), shouldFail: false},
		{actual: uint(3), expected: uint(1), shouldFail: false},
		{actual: uint16(4), expected: uint16(1), shouldFail: false},
		{actual: uint32(6), expected: uint32(1), shouldFail: false},
		{actual: uint64(7), expected: uint64(1), shouldFail: false},
		{actual: uint64(8), expected: uint64(1), shouldFail: false},
		{actual: int16(9), expected: int16(1), shouldFail: false},
		{actual: int32(10), expected: int32(1), shouldFail: false},
		{actual: int64(11), expected: int64(1), shouldFail: false},
		{actual: int64(12), expected: int64(1), shouldFail: false},
		{actual: "zzz", expected: "aaa", shouldFail: false},
		{actual: "aaa", expected: "zzz", shouldFail: true},
		{actual: complex64(1.0), expected: complex64(1.0), shouldFail: true}, // cannot compare complex types, so fails
	}
	for _, test := range equalsItems {
		stubTestingT = new(gocrest.StubTestingT)
		gocrest.AssertThat(stubTestingT, test.actual, gocrest.GreaterThan(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, GreaterThan(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}

func TestGreaterThanFailsWithDescriptionTest(testing *testing.T) {
	gocrest.AssertThat(stubTestingT, 1, gocrest.GreaterThan(2))
	if stubTestingT.MockTestOutput != "expected: value greater than 2 but was: 1" {
		testing.Errorf("did not get expected description, got %s", stubTestingT.MockTestOutput)
	}
}
func TestAssertThatTwoValuesAreLessThanOrNot(testing *testing.T) {
	var equalsItems = []struct {
		actual     interface{}
		expected   interface{}
		shouldFail bool
	}{
		{actual: 1, expected: 1, shouldFail: true},
		{actual: 1, expected: 2, shouldFail: false},
		{actual: 1.12, expected: 1.12, shouldFail: true},
		{actual: float32(1.0), expected: float32(1.12), shouldFail: false},
		{actual: float64(1.0), expected: float64(1.24), shouldFail: false},
		{actual: uint(0), expected: uint(3), shouldFail: false},
		{actual: uint16(1), expected: uint16(4), shouldFail: false},
		{actual: uint32(1), expected: uint32(6), shouldFail: false},
		{actual: uint64(1), expected: uint64(7), shouldFail: false},
		{actual: uint64(1), expected: uint64(8), shouldFail: false},
		{actual: int16(1), expected: int16(9), shouldFail: false},
		{actual: int32(1), expected: int32(10), shouldFail: false},
		{actual: int64(1), expected: int64(11), shouldFail: false},
		{actual: "aaa", expected: "zzz", shouldFail: false},
		{actual: "zzz", expected: "aaa", shouldFail: true},
		{actual: complex64(1.0), expected: complex64(1.0), shouldFail: true}, // cannot compare complex types, so fails
	}
	for _, test := range equalsItems {
		stubTestingT = new(gocrest.StubTestingT)
		gocrest.AssertThat(stubTestingT, test.actual, gocrest.LessThan(test.expected))
		if stubTestingT.HasFailed() != test.shouldFail {
			testing.Errorf("assertThat(%v, LessThan(%v)) gave unexpected test result (wanted failed %v, got failed %v)", test.actual, test.expected, test.shouldFail, stubTestingT.HasFailed())
		}
	}
}

func TestLessThanFailsWithDescriptionTest(testing *testing.T) {
	gocrest.AssertThat(stubTestingT, 2, gocrest.LessThan(1))
	if stubTestingT.MockTestOutput != "expected: value less than 1 but was: 2" {
		testing.Errorf("did not get expected description, got %s", stubTestingT.MockTestOutput)
	}
}

func TestNotReturnsTheOppositeOfGivenMatcher(testing *testing.T) {
	gocrest.AssertThat(stubTestingT, 1, gocrest.Not(gocrest.EqualTo(2)))
	if !stubTestingT.HasFailed() {
		testing.Error("Not(EqualTo) did not fail the test")
	}
}

func TestNotReturnsTheNotDescriptionOfGivenMatcher(testing *testing.T) {
	gocrest.AssertThat(stubTestingT, 2, gocrest.Not(gocrest.EqualTo(2)))
	if stubTestingT.MockTestOutput != "expected: not(value equal to 2) but was: 2" {
		testing.Errorf("did not get expected description, got %s", stubTestingT.MockTestOutput)
	}
}

func TestIsNilMatches(testing *testing.T) {
	gocrest.AssertThat(testing, nil, gocrest.IsNil())
}

func TestIsNilFails(testing *testing.T) {
	gocrest.AssertThat(stubTestingT, 2, gocrest.IsNil())
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestIsNilHasDescriptionTest(testing *testing.T) {
	gocrest.AssertThat(stubTestingT, 1, gocrest.IsNil())
	if stubTestingT.MockTestOutput != "expected: value equal to <nil> but was: 1" {
		testing.Errorf("did not get expected description, got %s", stubTestingT.MockTestOutput)
	}
}

func TestContainsDescriptionTest(testing *testing.T) {
	list := []string{"Foo", "Bar"}
	expectedList := []string{"Baz", "Bing"}
	gocrest.AssertThat(stubTestingT, list, gocrest.Contains(expectedList))
	if stubTestingT.MockTestOutput != "expected: something that contains [Baz Bing] but was: [Foo Bar]" {
		testing.Errorf("did not get expected description, got %s", stubTestingT.MockTestOutput)
	}
}

func TestContainsFailsForTwoStringArraysTest(testing *testing.T) {
	actualList := []string{"Foo", "Bar"}
	expectedList := []string{"Baz", "Bing"}
	gocrest.AssertThat(stubTestingT, actualList, gocrest.Contains(expectedList))
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestContainsFailsForTwoIntArraysTest(testing *testing.T) {
	actualList := []int{12, 13}
	expectedList := []int{14, 15}
	gocrest.AssertThat(stubTestingT, actualList, gocrest.Contains(expectedList))
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestContainsForString(testing *testing.T) {
	actualList := []string{"Foo", "Bar"}
	expected := "Foo"
	gocrest.AssertThat(testing, actualList, gocrest.Contains(expected))
}

func TestContainsFailsForString(testing *testing.T) {
	actualList := []string{"Foo", "Bar"}
	expected := "Moo"
	gocrest.AssertThat(stubTestingT, actualList, gocrest.Contains(expected))
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestContainsForSlice(testing *testing.T) {
	actualList := []string{"Foo", "Bar", "Bong", "Boom"}
	expected := []string{"Baz", "Bing", "Bong"}
	gocrest.AssertThat(testing, actualList[2:2], gocrest.Contains(expected[2:2]))
}

func TestContainsForList(testing *testing.T) {
	actualList := []string{"Foo", "Bar", "Bong", "Boom"}
	expected := []string{"Boom", "Bong", "Bar"}
	gocrest.AssertThat(testing, actualList, gocrest.Contains(expected))
}

func TestMapContainsMap(testing *testing.T) {
	actualList := map[string]string{
		"bing":  "boop",
		"bling": "bling",
	}
	expected := map[string]string{
		"bing": "boop",
	}

	gocrest.AssertThat(testing, actualList, gocrest.Contains(expected))
}

func TestStringContainsString(testing *testing.T) {
	actualList := "abcd"
	expected := "bc"
	gocrest.AssertThat(testing, actualList, gocrest.Contains(expected))
}

func TestMatchesPatternMatchesString(testing *testing.T) {
	actual := "blarney stone"
	expected := "^blarney.*"
	gocrest.AssertThat(testing, actual, gocrest.MatchesPattern(expected))
}

func TestMatchesPatternDoesNotMatchString(testing *testing.T) {
	actual := "blarney stone"
	expected := "^123.?.*"
	gocrest.AssertThat(stubTestingT, actual, gocrest.MatchesPattern(expected))
	if !stubTestingT.HasFailed() {
		testing.Error("did not fail test")
	}
}

func TestMatchesPatternDescription(testing *testing.T) {
	actual := "blarney stone"
	expected := "~123.?.*"
	gocrest.AssertThat(stubTestingT, actual, gocrest.MatchesPattern(expected))
	if stubTestingT.MockTestOutput != "expected: a value that matches pattern ~123.?.* but was: blarney stone" {
		testing.Errorf("incorrect description: %s", stubTestingT.MockTestOutput)
	}
}

func TestMatchesPatternWithErrorDescription(testing *testing.T) {
	actual := "blarney stone"
	expected := "+++"
	gocrest.AssertThat(stubTestingT, actual, gocrest.MatchesPattern(expected))
	if !strings.Contains(stubTestingT.MockTestOutput, "error parsing regexp: missing argument to repetition operator: `+`") {
		testing.Errorf("incorrect description: %s", stubTestingT.MockTestOutput)
	}
}

func TestHasFunctionPasses(testing *testing.T) {
	type MyType interface {
		N() int
		F() string
	}
	actual := new(MyType)
	expected := "F"
	gocrest.AssertThat(testing, actual, gocrest.HasFunctionNamed(expected))
}

func TestHasFunctionDoesNotPass(testing *testing.T) {
	type MyType interface {
		F() string
	}
	actual := new(MyType)
	expected := "E"
	gocrest.AssertThat(stubTestingT, actual, gocrest.HasFunctionNamed(expected))
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestHasFunctionDescribesMismatch(testing *testing.T) {
	type MyType interface {
		F() string
		B() string
	}
	actual := new(MyType)
	expected := "X"
	gocrest.AssertThat(stubTestingT, actual, gocrest.HasFunctionNamed(expected))
	if stubTestingT.MockTestOutput != "expected: interface with function X but was: MyType{B()F()}" {
		testing.Errorf("incorrect description:%s", stubTestingT.MockTestOutput)
	}
}

func TestAllOfMatches(testing *testing.T) {
	actual := "abcdef"
	gocrest.AssertThat(testing, actual, gocrest.AllOf(gocrest.EqualTo("abcdef"), gocrest.Contains("e")))
}

func TestAllOfFailsToMatch(testing *testing.T) {
	actual := "abc"
	gocrest.AssertThat(stubTestingT, actual, gocrest.AllOf(gocrest.EqualTo("abc"), gocrest.Contains("e")))
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestAllOfHasCorrectDescription(testing *testing.T) {
	actual := "abc"
	gocrest.AssertThat(stubTestingT, actual, gocrest.AllOf(gocrest.EqualTo("abc"), gocrest.Contains("e")))
	if stubTestingT.MockTestOutput != "expected: all of (value equal to abc and something that contains e) but was: abc" {
		testing.Errorf("incorrect description:%s", stubTestingT.MockTestOutput)
	}
}

func TestAnyOfMatches(testing *testing.T) {
	actual := "abcdef"
	gocrest.AssertThat(testing, actual, gocrest.AnyOf(gocrest.EqualTo("abcdef"), gocrest.Contains("g")))
}

func TestAnyOfFailsToMatch(testing *testing.T) {
	actual := "abc"
	gocrest.AssertThat(stubTestingT, actual, gocrest.AnyOf(gocrest.EqualTo("efg"), gocrest.Contains("e")))
	if !stubTestingT.HasFailed() {
		testing.Fail()
	}
}

func TestAnyOfHasCorrectDescription(testing *testing.T) {
	actual := "abc"
	gocrest.AssertThat(stubTestingT, actual, gocrest.AnyOf(gocrest.EqualTo("efg"), gocrest.Contains("e")))
	if stubTestingT.MockTestOutput != "expected: any of (value equal to efg or something that contains e) but was: abc" {
		testing.Errorf("incorrect description:%s", stubTestingT.MockTestOutput)
	}
}
