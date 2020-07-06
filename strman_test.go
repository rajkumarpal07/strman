package strman

/*
*
*  * The MIT License
*  *
*  * Copyright 2020 Rajkumar Palani <rajkumarpalani07@gmail.com>.
*  *
*  * Permission is hereby granted, free of charge, to any person obtaining a copy
*  * of this software and associated documentation files (the "Software"), to deal
*  * in the Software without restriction, including without limitation the rights
*  * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
*  * copies of the Software, and to permit persons to whom the Software is
*  * furnished to do so, subject to the following conditions:
*  *
*  * The above copyright notice and this permission notice shall be included in
*  * all copies or substantial portions of the Software.
*  *
*  * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
*  * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
*  * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
*  * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
*  * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
*  * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
*  * THE SOFTWARE.
*
 */
//strman_test.go ... Created on :: 03-July-2020

import (
	"reflect"
	"testing"
)

// Helper functions
type testList []struct {
	name, got, expected string
}

func (tl testList) validate(t *testing.T) {
	for _, test := range tl {
		if test.got != test.expected {
			t.Errorf("On %v, expected '%v', but got '%v'",
				test.name, test.expected, test.got)
		}
	}
}

//For checking boolean return statements
type testListBool []struct {
	name          string
	got, expected bool
}

func (t2 testListBool) validateBool(t *testing.T) {
	for _, test := range t2 {
		if test.got != test.expected {
			t.Errorf("On %v, expected '%v', but got '%v'",
				test.name, test.expected, test.got)
		}
	}
}

//For checking int return statements
type testListInt []struct {
	name          string
	got, expected int
}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func EqualIntArr(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Equal tells whether a and b contain the same elements.
func EqualStrArr(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

//Test Functions Begin...
func (t2 testListInt) validateInt(t *testing.T) {
	for _, test := range t2 {
		if test.got != test.expected {
			t.Errorf("On %v, expected '%v', but got '%v'",
				test.name, test.expected, test.got)
		}
	}
}

func TestAppendsShouldAppendStringsToEndOfValue(t *testing.T) {
	testList{
		{"TestAppends-0", Appends("f", "o", "o", "b", "a", "r"), "foobar"},
		{"TestAppends-1", Appends("foobar"), "foobar"},
		{"TestAppends-2", Appends("", "foobar"), "foobar"},
	}.validate(t)
}
func TestAppendArrayShouldAppendStringArrayToEndOfValue(t *testing.T) {
	testList{
		{"TestAppendArray-0", AppendsArray("f", []string{"o", "o", "b", "a", "r"}), "foobar"},
		{"TestAppendArray-1", AppendsArray("foobar", []string{}), "foobar"},
		{"TestAppendArray-2", AppendsArray("", []string{"foobar"}), "foobar"},
	}.validate(t)
}

func TestAtShouldFindCharacterAtIndex(t *testing.T) {
	testList{
		{"TestAt-0", At("foobar", 0), "f"},
		{"TestAt-1", At("foobar", 1), "o"},
		{"TestAt-2", At("foobar", -1), "r"},
		{"TestAt-3", At("foobar", -2), "a"},
		{"TestAt-4", At("foobar", 10), ""},
		{"TestAt-5", At("foobar", -10), ""},
	}.validate(t)
}

func TestBetweenShouldReturnArrayWithStringsBetweenStartAndEnd(t *testing.T) {
	st1, _ := Between("[abc][def]", "[", "]")
	st2 := []string{"abc", "def"}

	result := "nil"
	if reflect.DeepEqual(st1, st2) {
		result = "true"
	}
	testList{
		{"TestBetween-0", result, "true"},
	}.validate(t)
}
func TestBetweenshouldReturnEmptyArrayWhenStartAndEndDoesNotExist(t *testing.T) {
	st1, _ := Between("[abc][def]", "{", "}")
	st3 := []string{}
	result := "nil"
	if reflect.DeepEqual(st1, st3) {
		result = "true"
	}
	testList{
		{"TestBetweenEmpty-0", result, "true"},
	}.validate(t)
}

func TestCharsShouldReturnAllCharactersInString(t *testing.T) {
	fixture1, _ := Chars("title")
	fixture2 := []string{"t", "i", "t", "l", "e"}
	result := "nil"
	if reflect.DeepEqual(fixture1, fixture2) {
		result = "true"
	}
	testList{
		{"TestChars-0", result, "true"},
	}.validate(t)
}

func TestCollapseWhitespaceShouldReplaceConsecutiveWhitespaceWithSingleSpace(t *testing.T) {
	fixture := "foo      bar"
	fixture1 := "     foo     bar    "
	fixture2 := " foo      bar      bazz     hello    world    "

	testList{
		{"TestCollapseWhitespace-0", CollapseWhitespace(fixture), "foo bar"},
		{"TestCollapseWhitespace-1", CollapseWhitespace(fixture1), "foo bar"},
		{"TestCollapseWhitespace-2", CollapseWhitespace(fixture2), "foo bar bazz hello world"},
	}.validate(t)
}

func TestContains_shouldReturnTrueWhenStringContainsNeedle(t *testing.T) {
	fixture := "foo bar"
	fixture1 := "bar foo"
	fixture2 := "foobar"
	fixture3 := "foo"
	testListBool{
		{"TestContains-0", Contains(fixture, "foo", false), true},
		{"TestContains-1", Contains(fixture1, "foo", false), true},
		{"TestContains-2", Contains(fixture2, "foo", false), true},
		{"TestContains-3", Contains(fixture3, "foo", false), true},
		{"TestContains-4", Contains(fixture, "FOO", true), false},
		{"TestContains-5", Contains(fixture1, "FOO", true), false},
		{"TestContains-6", Contains(fixture2, "FOO", true), false},
		{"TestContains-7", Contains(fixture3, "FOO", true), false},
	}.validateBool(t)
}

func TestContainsAllShouldReturnTrueOnlyWhenAllNeedlesAreContainedInValue(t *testing.T) {
	fixture := "foo bar"
	fixture1 := "bar foo"
	fixture2 := "foobar"

	res1 := []string{"foo", "bar"}
	res2 := []string{"foo", "raj"}

	testListBool{
		{"TestContainsAll-0", ContainsAll(fixture, res1), true},
		{"TestContainsAll-1", ContainsAll(fixture1, res1), true},
		{"TestContainsAll-2", ContainsAll(fixture2, res1), true},
		{"TestContainsAll-4", ContainsAll(fixture, res2), false},
		{"TestContainsAll-5", ContainsAll(fixture1, res2), false},
		{"TestContainsAll-6", ContainsAll(fixture2, res2), false},
	}.validateBool(t)
}

func TestContainsAllCaseSensitiveShouldReturnTrueOnlyWhenAllNeedlesAreContainedInValue(t *testing.T) {
	fixture := "foo bar"
	fixture1 := "bar foo"
	fixture2 := "foobar"

	res1 := []string{"foo", "bar"}
	res2 := []string{"FOO", "bar"}

	testListBool{
		{"TestContainsAllCaseSensitive-0", ContainsAllCaseSensitive(fixture, res1, true), true},
		{"TestContainsAllCaseSensitive-1", ContainsAllCaseSensitive(fixture1, res1, true), true},
		{"TestContainsAllCaseSensitive-2", ContainsAllCaseSensitive(fixture2, res1, true), true},
		{"TestContainsAllCaseSensitive-4", ContainsAllCaseSensitive(fixture, res2, true), false},
		{"TestContainsAllCaseSensitive-5", ContainsAllCaseSensitive(fixture1, res2, true), false},
		{"TestContainsAllCaseSensitive-6", ContainsAllCaseSensitive(fixture2, res2, true), false},
	}.validateBool(t)
}

func TestContainsAnyShouldReturnTrueWhenAnyOfSearchNeedleExistInInputValue(t *testing.T) {
	fixture := "foo bar"
	fixture1 := "bar foo"
	fixture2 := "foobar"
	res1 := []string{"foo", "bar", "test"}
	testListBool{
		{"TestContainsAny-0", ContainsAny(fixture, res1), true},
		{"TestContainsAny-1", ContainsAny(fixture1, res1), true},
		{"TestContainsAny-2", ContainsAny(fixture2, res1), true},
	}.validateBool(t)
}

func TestContainsAnyCaseSensitiveShouldReturnFalseWhenNoneOfSearchNeedleExistInInputValue(t *testing.T) {
	fixture := "foo bar"
	fixture1 := "bar foo"
	fixture2 := "foobar"
	res1 := []string{"FOO", "BAR", "test"}
	testListBool{
		{"TestContainsAnyCaseSensitive-0", ContainsAnyCaseSensitive(fixture, res1, true), false},
		{"TestContainsAnyCaseSensitive-1", ContainsAnyCaseSensitive(fixture1, res1, true), false},
		{"TestContainsAnyCaseSensitive-2", ContainsAnyCaseSensitive(fixture2, res1, true), false},
	}.validateBool(t)
}

func TestCountSubstrShouldCountSubStrCountCaseSensitiveWithoutOverlapInValue(t *testing.T) {
	fixture := "aaaAAAaaa"
	testListInt{
		{"TestCountSubstr-0", CountSubstr(fixture, "aaa"), 2},
	}.validateInt(t)
}

func TestCountSubstringShouldCountSubStrCountCaseInsensitiveWithoutOverlapInValue(t *testing.T) {
	fixture := "aaaAAAaaa"
	testListInt{
		{"TestCountSubstring-0", CountSubstring(fixture, "aaa", false, false), 3},
		{"TestCountSubstring-1", CountSubstring(fixture, "aaa", false, true), 7},
	}.validateInt(t)
}

func TestEndsWith_ShouldBeTrueWhenStringEndsWithSearchStr(t *testing.T) {
	fixture := "foo bar"
	fixture1 := "bar"
	testListBool{
		{"TestEndsWith-0", EndsWith(fixture, "bar"), true},
		{"TestEndsWith-1", EndsWith(fixture1, "bar"), true},
	}.validateBool(t)
}

func TestEndsWithCase_ShouldBeTrueWhenStringEndsWithSearchStr(t *testing.T) {
	fixture := "foo bar"
	fixture1 := "BAR"
	testListBool{
		{"TestEndsWithCase-0", EndsWithCase(fixture, "bar", true), true},
		{"TestEndsWithCase-1", EndsWithCase(fixture1, "BAR", true), true},
	}.validateBool(t)
}

func TestEnsureLeftShouldEnsureValueStartsWithFoo(t *testing.T) {
	fixture := "foobar"
	fixture1 := "bar"
	testList{
		{"TestEnsureLeft-0", EnsureLeft(fixture, "foo"), "foobar"},
		{"TestEnsureLeft-1", EnsureLeft(fixture1, "foo"), "foobar"},
	}.validate(t)
}

func TestEnsureLeftWithCaseShouldEnsureValueStartsWithFoo(t *testing.T) {
	fixture := "foobar"
	fixture1 := "bar"
	testList{
		{"TestEnsureLeftWithCase-0", EnsureLeftWithCase(fixture, "foo", true), "foobar"},
		{"TestEnsureLeftWithCase-1", EnsureLeftWithCase(fixture1, "FOO", true), "FOObar"},
	}.validate(t)
}

func TestBase64EncodeShouldEncodeAString(t *testing.T) {
	testList{
		{"TestBase64Encode-0", Base64Encode("strman"), "c3RybWFu"},
		{"TestBase64Encode-1", Base64Encode("foo"), "Zm9v"},
		{"TestBase64Encode-2", Base64Encode("bar"), "YmFy"},
		{"TestBase64Encode-3", Base64Encode("bár!"), "YsOhciE="},
		{"TestBase64Encode-4", Base64Encode("ராஜ்"), "4K6w4K6+4K6c4K+N"},
		{"TestBase64Encode-5", Base64Encode("குமார்"), "4K6V4K+B4K6u4K6+4K6w4K+N"},
	}.validate(t)
}

func TestBase64DecodeShouldDecodeABase64DecodedValueToString(t *testing.T) {
	testList{
		{"TestBase64Decode-0", Base64Decode("c3RybWFu"), "strman"},
		{"TestBase64Decode-1", Base64Decode("Zm9v"), "foo"},
		{"TestBase64Decode-2", Base64Decode("YmFy"), "bar"},
		{"TestBase64Decode-3", Base64Decode("YsOhciE="), "bár!"},
		{"TestBase64Decode-4", Base64Decode("4K6w4K6+4K6c4K+N"), "ராஜ்"},
		{"TestBase64Decode-5", Base64Decode("4K6V4K+B4K6u4K6+4K6w4K+N"), "குமார்"},
	}.validate(t)
}

func TestDecEncodeShouldEncodeStringToDecimal(t *testing.T) {
	testList{
		{"TestDecEncode-0", DecEncode("A"), "65"},
		{"TestDecEncode-1", DecEncode("AA"), "6565"},
		{"TestDecEncode-2", DecEncode("Á"), "193"},
		{"TestDecEncode-3", DecEncode("漢"), "28450"},
	}.validate(t)
}

func TestEnsureRightShouldEnsureStringEndsWithBar(t *testing.T) {
	fixture := "foo"
	fixture1 := "foobar"
	fixture2 := "fooBAR"
	testList{
		{"TestEnsureRight-0", EnsureRight(fixture, "bar"), "foobar"},
		{"TestEnsureRight-1", EnsureRight(fixture1, "bar"), "foobar"},
		{"TestEnsureRight-2", EnsureRight(fixture2, "bar"), "fooBARbar"},
	}.validate(t)
}

func TestEnsureRightWithCaseShouldEnsureStringEndsWithBar(t *testing.T) {
	fixture := "foo"
	fixture1 := "foobar"
	fixture2 := "fooBAR"
	testList{
		{"TestEnsureRight-0", EnsureRightWithCase(fixture, "bar", true), "foobar"},
		{"TestEnsureRight-1", EnsureRightWithCase(fixture1, "bar", true), "foobar"},
		{"TestEnsureRight-2", EnsureRightWithCase(fixture2, "bar", true), "fooBARbar"},
		{"TestEnsureRight-3", EnsureRightWithCase(fixture2, "BAR", true), "fooBAR"},
	}.validate(t)
}

func TestFirstShouldReturnFirstThreeCharsOfString(t *testing.T) {
	fixture := "foo"
	fixture1 := "foobar"
	testList{
		{"TestFirst-0", First(fixture), "foo"},
		{"TestFirst-1", First(fixture1), "foo"},
	}.validate(t)
}

func TestFirstNCharsShouldReturnFirstThreeCharsOfString(t *testing.T) {
	fixture := "foo"
	fixture1 := "foobar"
	testList{
		{"TestFirstNChars-0", FirstNChars(fixture, 3), "foo"},
		{"TestFirstNChars-1", FirstNChars(fixture1, 3), "foo"},
	}.validate(t)
}

func TestHeadShouldReturnFirstCharOfString(t *testing.T) {
	fixture := "foo"
	fixture1 := "foobar"
	testList{
		{"TestHead-0", Head(fixture), "f"},
		{"TestHead-1", Head(fixture1), "f"},
	}.validate(t)
}

func TestHexDecodeShouldDecodeHexCodeToString(t *testing.T) {
	testList{
		{"TestHexDecode-0", HexDecode("e6bca2"), "漢"},
		{"TestHexDecode-1", HexDecode("41"), "A"},
		{"TestHexDecode-2", HexDecode("c381"), "Á"},
		{"TestHexDecode-3", HexDecode("4141"), "AA"},
	}.validate(t)
}

func TestHexEncodeShouldEncodeStringToHexadecimalFormat(t *testing.T) {
	testList{
		{"TestHexEncode-0", HexEncode("漢"), "e6bca2"},
		{"TestHexEncode-1", HexEncode("A"), "41"},
		{"TestHexEncode-2", HexEncode("Á"), "c381"},
		{"TestHexEncode-3", HexEncode("AA"), "4141"},
	}.validate(t)
}

func TestIndexOfShouldBeTrueWhenNeedleExists(t *testing.T) {
	fixture := "foobar"
	testListInt{
		{"TestIndexOf-0", IndexOf(fixture, "f", true), 0},
		{"TestIndexOf-1", IndexOf(fixture, "o", true), 1},
		{"TestIndexOf-2", IndexOf(fixture, "b", true), 3},
		{"TestIndexOf-3", IndexOf(fixture, "a", true), 4},
		{"TestIndexOf-4", IndexOf(fixture, "r", true), 5},
		{"TestIndexOf-5", IndexOf(fixture, "t", true), -1},
	}.validateInt(t)
}

func TestUnEqualShouldTestInequalityOfStrings(t *testing.T) {
	testListBool{
		{"TestUnEqual-0", UnEqual("a", "b"), true},
		{"TestUnEqual-1", UnEqual("a", "a"), false},
		{"TestUnEqual-2", UnEqual("0", "1"), true},
	}.validateBool(t)
}

func TestInsertShouldInsertStringAtIndex(t *testing.T) {
	testList{
		{"TestInsert-0", Insert("fbar", "oo", 1), "foobar"},
		{"TestInsert-1", Insert("foo", "bar", 3), "foobar"},
		{"TestInsert-2", Insert("foobar", "x", 5), "foobaxr"},
		{"TestInsert-4", Insert("foobar", "x", 6), "foobarx"},
		{"TestInsert-5", Insert("foo bar", "asadasd", 100), "foo bar"},
	}.validate(t)
}

func TestIsLowerCaseShouldBeTrueWhenStringIsLowerCase(t *testing.T) {
	testListBool{
		{"TestIsLowerCase-0", IsLowerCase(""), true},
		{"TestIsLowerCase-1", IsLowerCase("foo"), true},
		{"TestIsLowerCase-2", IsLowerCase("foobarfoo"), true},
		{"TestIsLowerCase-3", IsLowerCase("Foo"), false},
		{"TestIsLowerCase-4", IsLowerCase("foobarfooA"), false},
	}.validateBool(t)
}

func TestIsUpperCaseShouldBeFalseWhenStringIsNotLowerCase(t *testing.T) {
	testListBool{
		{"TestIsUpperCase-0", IsUpperCase(""), true},
		{"TestIsUpperCase-1", IsUpperCase("FOO"), true},
		{"TestIsUpperCase-2", IsUpperCase("FOOBARFOO"), true},
		{"TestIsUpperCase-3", IsUpperCase("Foo"), false},
		{"TestIsUpperCase-4", IsUpperCase("foobarfooA"), false},
	}.validateBool(t)
}

func TestLastShouldReturnLastNChars(t *testing.T) {
	testList{
		{"TestLast-0", Last("foo", 3), "foo"},
		{"TestLast-1", Last("foobarfoo", 3), "foo"},
		{"TestLast-2", Last("", 3), ""},
		{"TestLast-4", Last("f", 3), "f"},
		{"TestLast-5", Last("foo bar", 100), "foo bar"},
	}.validate(t)
}

func TestLeftPadShouldAddPaddingOnTheLeft(t *testing.T) {
	testList{
		{"TestLeftPad-0", LeftPad("1", "0", 5), "000001"},
		{"TestLeftPad-1", LeftPad("01", "0", 5), "0000001"},
		{"TestLeftPad-2", LeftPad("001", "0", 5), "00000001"},
		{"TestLeftPad-4", LeftPad("0001", "0", 5), "000000001"},
		{"TestLeftPad-5", LeftPad("00001", "0", 5), "0000000001"},
	}.validate(t)
}

func TestLeftPad2Len(t *testing.T) {
	testList{
		{"TestLeftPad2Len-0", LeftPad2Len("1", "0", 5), "00001"},
		{"TestLeftPad2Len-1", LeftPad2Len("01", "0", 5), "00001"},
		{"TestLeftPad2Len-2", LeftPad2Len("001", "0", 5), "00001"},
		{"TestLeftPad2Len-4", LeftPad2Len("0001", "0", 5), "00001"},
		{"TestLeftPad2Len-5", LeftPad2Len("00001", "0", 5), "00001"},
	}.validate(t)
}

func TestRightPad2Len(t *testing.T) {
	testList{
		{"TestRightPad2Len-0", RightPad2Len("1", "0", 5), "10000"},
		{"TestRightPad2Len-1", RightPad2Len("10", "0", 5), "10000"},
		{"TestRightPad2Len-2", RightPad2Len("100", "0", 5), "10000"},
		{"TestRightPad2Len-4", RightPad2Len("1000", "0", 5), "10000"},
		{"TestRightPad2Len-5", RightPad2Len("10000", "0", 5), "10000"},
		{"TestRightPad2Len-5", RightPad2Len("10000000", "0", 5), "10000"},
	}.validate(t)
}

func TestIsString(t *testing.T) {
	testListBool{
		{"TestIsString-0", IsString(1), false},
		{"TestIsString-1", IsString(false), false},
		{"TestIsString-2", IsString(1.2), false},
		{"TestIsString-3", IsString([]string{}), false},
		{"TestIsString-4", IsString("foobarfooA"), true},
		{"TestIsString-4", IsString(""), true},
	}.validateBool(t)
}

func TestLastIndexOfShouldFindIndexOfNeedle(t *testing.T) {
	fixture := "foobarfoobar"
	testListInt{
		{"TestLastIndexOf-0", LastIndexOf(fixture, "f"), 6},
		{"TestLastIndexOf-1", LastIndexOf(fixture, "o"), 8},
		{"TestLastIndexOf-2", LastIndexOf(fixture, "b"), 9},
		{"TestLastIndexOf-3", LastIndexOf(fixture, "a"), 10},
		{"TestLastIndexOf-4", LastIndexOf(fixture, "r"), 11},
		{"TestLastIndexOf-5", LastIndexOf(fixture, "t"), -1},
	}.validateInt(t)
}

func TestLastIndexOfWithCaseShouldFindIndexOfNeedleCaseInsensitive(t *testing.T) {
	fixture := "foobarfoobar"
	testListInt{
		{"TestLastIndexOfWithCase-0", LastIndexOfWithCase(fixture, "F", false), 6},
		{"TestLastIndexOfWithCase-1", LastIndexOfWithCase(fixture, "O", false), 8},
		{"TestLastIndexOfWithCase-2", LastIndexOfWithCase(fixture, "B", false), 9},
		{"TestLastIndexOfWithCase-3", LastIndexOfWithCase(fixture, "A", false), 10},
		{"TestLastIndexOfWithCase-4", LastIndexOfWithCase(fixture, "R", false), 11},
		{"TestLastIndexOfWithCase-5", LastIndexOfWithCase(fixture, "T", false), -1},
	}.validateInt(t)
}

func TestLeftTrimShouldRemoveSpacesOnLeft(t *testing.T) {
	testList{
		{"TestLeftTrim-0", LeftTrim("     strman"), "strman"},
		{"TestLeftTrim-1", LeftTrim("     strman  "), "strman  "},
	}.validate(t)
}

func TestPrependShouldPrependStrings(t *testing.T) {
	testList{
		{"TestPrepend-0", Prepend("r", "f", "o", "o", "b", "a"), "foobar"},
		{"TestPrepend-1", Prepend("foobar"), "foobar"},
		{"TestPrepend-2", Prepend("", "foobar"), "foobar"},
		{"TestPrepend-3", Prepend("bar", "foo"), "foobar"},
	}.validate(t)
}

func TestPrependArrayShouldPrependStrings(t *testing.T) {
	testList{
		{"TestPrependArray-0", PrependArray("r", []string{"f", "o", "o", "b", "a"}), "foobar"},
		{"TestPrependArray-1", PrependArray("foobar", []string{}), "foobar"},
		{"TestPrependArray-2", PrependArray("", []string{"foobar"}), "foobar"},
		{"TestPrependArray-3", PrependArray("bar", []string{"foo"}), "foobar"},
	}.validate(t)
}

func TestRemoveEmptyStringsShouldRemoveEmptyStrings(t *testing.T) {
	res1, _ := RemoveEmptyStrings([]string{"aa", "", "   ", "bb", "cc"})
	target := []string{"aa", "bb", "cc"}

	if !(EqualStrArr(res1, target)) {
		panic("strman:TestRemoveEmptyStringsShouldRemoveEmptyStrings:: Failed!")
	}
}

func TestRemoveLeftShouldRemoveStringFromLeft(t *testing.T) {
	fixture := "foobar"
	fixture1 := "bar"
	fixture2 := "barfoo"
	fixture3 := "foofoo"
	testList{
		{"TestRemoveLeft-0", RemoveLeft(fixture, "foo"), "bar"},
		{"TestRemoveLeft-1", RemoveLeft(fixture1, "foo"), "bar"},
		{"TestRemoveLeft-2", RemoveLeft(fixture2, "foo"), "barfoo"},
		{"TestRemoveLeft-3", RemoveLeft(fixture3, "foo"), "foo"},
	}.validate(t)
}

func TestRemoveLeftWithCaseShouldRemoveStringFromLeftCaseInSensitive(t *testing.T) {
	fixture := "foobar"
	fixture1 := "bar"
	testList{
		{"TestRemoveLeftWithCase-0", RemoveLeftWithCase(fixture, "FOO", false), "bar"},
		{"TestRemoveLeftWithCase-1", RemoveLeftWithCase(fixture1, "FOO", false), "bar"},
	}.validate(t)
}

func TestRemoveNonWordsShouldRemoveAllNonWords(t *testing.T) {
	fixture := "fooÁbar"
	fixture1 := "foo&bar-"
	fixture2 := "foo bar"
	testList{
		{"TestRemoveNonWords-0", RemoveNonWords(fixture), "foobar"},
		{"TestRemoveNonWords-1", RemoveNonWords(fixture1), "foobar"},
		{"TestRemoveNonWords-2", RemoveNonWords(fixture2), "foobar"},
	}.validate(t)
}

func TestRemoveRightShouldRemoveStringFromRight(t *testing.T) {
	fixture := "foobar"
	fixture1 := "foo"
	fixture2 := "barfoo"
	fixture3 := "barbar"
	testList{
		{"TestRemoveRight-0", RemoveRight(fixture, "bar"), "foo"},
		{"TestRemoveRight-1", RemoveRight(fixture1, "bar"), "foo"},
		{"TestRemoveRight-2", RemoveRight(fixture2, "bar"), "barfoo"},
		{"TestRemoveRight-3", RemoveRight(fixture3, "bar"), "bar"},
	}.validate(t)
}

func TestRemoveRightWithCaseShouldRemoveStringFromRight(t *testing.T) {
	fixture := "foobar"
	fixture1 := "foo"
	fixture2 := "barfoo"
	fixture3 := "barbar"
	testList{
		{"TestRemoveRightWithCase-0", RemoveRightWithCase(fixture, "BAR", false), "foo"},
		{"TestRemoveRightWithCase-1", RemoveRightWithCase(fixture1, "BAR", false), "foo"},
		{"TestRemoveRightWithCase-2", RemoveRightWithCase(fixture2, "BAR", false), "barfoo"},
		{"TestRemoveRightWithCase-3", RemoveRightWithCase(fixture3, "BAR", false), "bar"},
	}.validate(t)
}

func TestRemoveSpacesShouldRemoveSpacesInTheString(t *testing.T) {
	fixture := "foobar"
	fixture1 := "foo bar "
	fixture2 := " foo bar"
	fixture3 := " foo bar "
	testList{
		{"TestRemoveSpaces-0", RemoveSpaces(fixture), "foobar"},
		{"TestRemoveSpaces-1", RemoveSpaces(fixture1), "foobar"},
		{"TestRemoveSpaces-2", RemoveSpaces(fixture2), "foobar"},
		{"TestRemoveSpaces-3", RemoveSpaces(fixture3), "foobar"},
	}.validate(t)
}

func TestRepeatShouldRepeatAStringNTimes(t *testing.T) {
	testList{
		{"TestRepeat-0", Repeat("1", 1), "1"},
		{"TestRepeat-1", Repeat("1", 2), "11"},
		{"TestRepeat-2", Repeat("1", 3), "111"},
		{"TestRepeat-3", Repeat("1", 4), "1111"},
		{"TestRepeat-4", Repeat("1", 5), "11111"},
	}.validate(t)
}

func TestReplaceShouldReplaceAllOccurrencesOfString(t *testing.T) {
	testList{
		{"TestReplace-0", Replace("foo bar", "foo", "bar", true), "bar bar"},
		{"TestReplace-1", Replace("foo bar foo", "foo", "bar", true), "bar bar bar"},
		{"TestReplace-2", Replace("FOO bar", "foo", "bar", false), "bar bar"},
		{"TestReplace-3", Replace("FOO bar foo", "foo", "bar", false), "bar bar bar"},
	}.validate(t)
}

func TestReverseShouldReverseInputString(t *testing.T) {
	testList{
		{"TestReverse-0", Reverse(""), ""},
		{"TestReverse-1", Reverse("foo"), "oof"},
		{"TestReverse-2", Reverse("rajkumar"), "ramukjar"},
		{"TestReverse-3", Reverse("bar"), "rab"},
		{"TestReverse-4", Reverse("foo_"), "_oof"},
		{"TestReverse-5", Reverse("f"), "f"},
	}.validate(t)
}

func TestRightPadShouldRightPadAString(t *testing.T) {
	testList{
		{"TestRightPad-0", RightPad("1", "0", 5), "10000"},
		{"TestRightPad-1", RightPad("10", "0", 5), "10000"},
		{"TestRightPad-2", RightPad("100", "0", 5), "10000"},
		{"TestRightPad-3", RightPad("1000", "0", 5), "10000"},
		{"TestRightPad-4", RightPad("10000", "0", 5), "10000"},
		{"TestRightPad-5", RightPad("10000000", "0", 5), "10000000"},
	}.validate(t)
}

func TestRightTrimShouldRemoveSpacesFromTheRight(t *testing.T) {
	testList{
		{"TestRightTrim-0", RightTrim("strman   "), "strman"},
		{"TestRightTrim-1", RightTrim("   strman"), "   strman"},
		{"TestRightTrim-2", RightTrim("strman"), "strman"},
	}.validate(t)
}

func TestSafeTruncateShouldSafelyTruncateStrings(t *testing.T) {
	testList{
		{"TestSafeTruncate-0", SafeTruncate("foo bar", 0, "."), ""},
		{"TestSafeTruncate-1", SafeTruncate("foo bar", 4, "."), "foo."},
		{"TestSafeTruncate-2", SafeTruncate("foo bar", 3, "."), "."},
		{"TestSafeTruncate-3", SafeTruncate("foo bar", 2, "."), "."},
		{"TestSafeTruncate-4", SafeTruncate("foo bar", 7, "."), "foo bar"},
		{"TestSafeTruncate-5", SafeTruncate("foo bar", 8, "."), "foo bar"},
		{"TestSafeTruncate-6", SafeTruncate("A Golang string manipulation library.", 12, "..."), "A Golang..."},
		{"TestSafeTruncate-7", SafeTruncate("A Golang string manipulation library.", 11, "..."), "A Golang..."},
		{"TestSafeTruncate-8", SafeTruncate("A Golang string manipulation library.", 10, "..."), "A..."},
		{"TestSafeTruncate-9", SafeTruncate("A Golang string manipulation library.", 9, "..."), "A..."},
	}.validate(t)
}

func TestTruncateShouldTruncateString(t *testing.T) {
	testList{
		{"TestTruncate-0", Truncate("foo bar", 0, "."), ""},
		{"TestTruncate-1", Truncate("foo bar", 3, "."), "fo."},
		{"TestTruncate-2", Truncate("foo bar", 2, "."), "f."},
		{"TestTruncate-3", Truncate("foo bar", 4, "."), "foo."},
		{"TestTruncate-4", Truncate("foo bar", 7, "."), "foo bar"},
		{"TestTruncate-5", Truncate("foo bar", 8, "."), "foo bar"},
		{"TestTruncate-6", Truncate("A Golang string manipulation library.", 12, "..."), "A Golang ..."},
		{"TestTruncate-7", Truncate("A Golang string manipulation library.", 11, "..."), "A Golang..."},
		{"TestTruncate-8", Truncate("A Golang string manipulation library.", 10, "..."), "A Golan..."},
		{"TestTruncate-9", Truncate("A Golang string manipulation library.", 9, "..."), "A Gola..."},
	}.validate(t)
}

func TestHTMLDecode(t *testing.T) {
	testList{
		{"TestHTMLDecode-0", HTMLDecode("&aacute;"), "\u00E1"},
		{"TestHTMLDecode-1", HTMLDecode("&SHcy;"), "Ш"},
		{"TestHTMLDecode-2", HTMLDecode("&ZHcy;"), "Ж"},
		{"TestHTMLDecode-3", HTMLDecode("&boxdl;"), "┐"},
	}.validate(t)
}

func TestHTMLEncode(t *testing.T) {
	testList{
		{"TestHTMLEncode-0", HTMLEncode("fred, barney, & pebbles"), "fred, barney, &amp; pebbles"},
		{"TestHTMLEncode-1", HTMLEncode("á"), "á"},
		{"TestHTMLEncode-2", HTMLEncode("áéíóú"), "áéíóú"},
		{"TestHTMLEncode-3", HTMLEncode("Ш"), "Ш"},
		{"TestHTMLEncode-4", HTMLEncode("Ж"), "Ж"},
		{"TestHTMLEncode-5", HTMLEncode("┐"), "┐"},
	}.validate(t)
}

func TestShuffleShouldShuffleAString(t *testing.T) {
	res1 := Shuffle("rajkumar")
	if reflect.DeepEqual(res1, "rajkumar") {
		panic("strman:TestShuffleShouldShuffleAString-0:: Failed!")
	}
	res2 := Shuffle("")
	if !reflect.DeepEqual(res2, "") {
		panic("strman:TestShuffleShouldShuffleAString-1:: Failed!")
	}
	res3 := Shuffle("s")
	if !reflect.DeepEqual(res3, "s") {
		panic("strman:TestShuffleShouldShuffleAString-2:: Failed!")
	}
}

func TestSliceShouldSliceupGivenString(t *testing.T) {
	fixture := "foobarfoo"
	testList{
		{"TestSlice-0", Slice(fixture, 0, 2), "fo"},
		{"TestSlice-1", Slice(fixture, 2, 5), "oba"},
		{"TestSlice-2", Slice(fixture, 7, 9), "oo"},
	}.validate(t)
}

func TestSlugifyShouldBeFooAndBar(t *testing.T) {
	testList{
		{"TestSlugify-0", Slugify("foo&bar"), "foo-and-bar"},
		{"TestSlugify-1", Slugify("foo&bar."), "foo-and-bar"},
		{"TestSlugify-2", Slugify("foo&bar "), "foo-and-bar"},
		{"TestSlugify-3", Slugify(" foo&bar"), "foo-and-bar"},
		{"TestSlugify-4", Slugify(" foo&bar "), "foo-and-bar"},
		{"TestSlugify-5", Slugify("foo&bar"), "foo-and-bar"},
		{"TestSlugify-6", Slugify("fóõ-and---bár"), "foo-and-bar"},
		{"TestSlugify-7", Slugify("foo  &    bar"), "foo-and-bar"},
		{"TestSlugify-8", Slugify("FOO  &   bar"), "foo-and-bar"},
	}.validate(t)
}

func TestTransliterateShouldTransliterateTheText(t *testing.T) {
	fixture := "fóõ bár"
	testList{
		{"TestTransliterate-0", Transliterate(fixture), "foo bar"},
	}.validate(t)
}

func TestSurroundShouldSurroundStringWithPrefixAndSuffix(t *testing.T) {
	testList{
		{"TestSurround-0", Surround("foo", "bar", "bar"), "barfoobar"},
		{"TestSurround-1", Surround("rajkumar", "***", "***"), "***rajkumar***"},
		{"TestSurround-2", Surround("", ">", ">"), ">>"},
		{"TestSurround-3", Surround("bar", "", ""), "bar"},
		{"TestSurround-4", Surround("f", "", ""), "f"},
		{"TestSurround-5", Surround("div", "<", ">"), "<div>"},
	}.validate(t)
}

func TestToCamelCaseShouldConvertStringToCamelCase(t *testing.T) {
	testList{
		{"TestToCamelCase-0", ToCamelCase("CamelCase"), "camelCase"},
		{"TestToCamelCase-1", ToCamelCase("camelCase"), "camelCase"},
		{"TestToCamelCase-2", ToCamelCase("Camel case"), "camelCase"},
		{"TestToCamelCase-3", ToCamelCase("Camel  case"), "camelCase"},
		{"TestToCamelCase-4", ToCamelCase("camel Case"), "camelCase"},
		{"TestToCamelCase-5", ToCamelCase("camel-case"), "camelCase"},
		{"TestToCamelCase-6", ToCamelCase("-camel--case"), "camelCase"},
		{"TestToCamelCase-7", ToCamelCase("camel_case"), "camelCase"},
		{"TestToCamelCase-8", ToCamelCase("     camel_case"), "camelCase"},
		{"TestToCamelCase-9", ToCamelCase("c"), "c"},
	}.validate(t)

}

func TestToDeCamelCaseShouldDeCamelCaseAString(t *testing.T) {
	testList{
		{"TestToDecamelize-0", ToDecamelize("deCamelize", ""), "de camelize"},
		{"TestToDecamelize-1", ToDecamelize("de-Camelize", ""), "de camelize"},
		{"TestToDecamelize-2", ToDecamelize("de camelize", ""), "de camelize"},
		{"TestToDecamelize-3", ToDecamelize("de  camelize", ""), "de camelize"},
		{"TestToDecamelize-4", ToDecamelize("de Camelize", ""), "de camelize"},
		{"TestToDecamelize-5", ToDecamelize("de-camelize", ""), "de camelize"},
		{"TestToDecamelize-6", ToDecamelize("-de--camelize", ""), "de camelize"},
		{"TestToDecamelize-7", ToDecamelize("de_camelize", ""), "de camelize"},
		{"TestToDecamelize-8", ToDecamelize("     de_camelize", ""), "de camelize"},
		{"TestToDecamelize-9", ToDecamelize("camelCoRoTYase", "_"), "camel_co_ro_t_yase"},
	}.validate(t)
}

func TestToKebabCaseShouldKebabCaseAString(t *testing.T) {
	testList{
		{"TestToKebabCase-0", ToKebabCase("deCamelize"), "de-camelize"},
		{"TestToKebabCase-1", ToKebabCase("de-Camelize"), "de-camelize"},
		{"TestToKebabCase-2", ToKebabCase("de camelize"), "de-camelize"},
		{"TestToKebabCase-3", ToKebabCase("de  camelize"), "de-camelize"},
		{"TestToKebabCase-4", ToKebabCase("de Camelize"), "de-camelize"},
		{"TestToKebabCase-5", ToKebabCase("de-camelize"), "de-camelize"},
		{"TestToKebabCase-6", ToKebabCase("-de--camelize"), "de-camelize"},
		{"TestToKebabCase-7", ToKebabCase("de_camelize"), "de-camelize"},
		{"TestToKebabCase-8", ToKebabCase("     de_camelize"), "de-camelize"},
		{"TestToKebabCase-9", ToKebabCase("camelCoRooTBase"), "camel-co-roo-t-base"},
	}.validate(t)
}

func TestToSnakeCaseShouldSnakeCaseAString(t *testing.T) {
	testList{
		{"TestToSnakeCase-0", ToSnakeCase("deCamelize"), "de_camelize"},
		{"TestToSnakeCase-1", ToSnakeCase("de-Camelize"), "de_camelize"},
		{"TestToSnakeCase-2", ToSnakeCase("de camelize"), "de_camelize"},
		{"TestToSnakeCase-3", ToSnakeCase("de  camelize"), "de_camelize"},
		{"TestToSnakeCase-4", ToSnakeCase("de Camelize"), "de_camelize"},
		{"TestToSnakeCase-5", ToSnakeCase("de-camelize"), "de_camelize"},
		{"TestToSnakeCase-6", ToSnakeCase("-de--camelize"), "de_camelize"},
		{"TestToSnakeCase-7", ToSnakeCase("de_camelize"), "de_camelize"},
		{"TestToSnakeCase-8", ToSnakeCase("     de_camelize"), "de_camelize"},
		{"TestToSnakeCase-9", ToSnakeCase("camelCoRooTBase"), "camel_co_roo_t_base"},
	}.validate(t)
}

func TestJoinShouldJoinArrayOfStringIntoASingleString(t *testing.T) {
	fixture := []string{"hello", "world", "123"}
	testList{
		{"TestJoin-0", Join(fixture, ":"), "hello:world:123"},
		{"TestJoin-1", Join([]string{}, ":"), ""},
	}.validate(t)
}

func TestCapitalizeShouldCapitalizeFirstCharacterOfString(t *testing.T) {
	testList{
		{"TestCapitalize-0", Capitalize("FRED"), "Fred"},
		{"TestCapitalize-1", Capitalize("fRED"), "Fred"},
		{"TestCapitalize-2", Capitalize("fred"), "Fred"},
	}.validate(t)
}

func TestLowerFirstShouldLowercasedFirstCharacterOfString(t *testing.T) {
	testList{
		{"TestLowerFirst-0", LowerFirst("FRED"), "fRED"},
		{"TestLowerFirst-1", LowerFirst("fred"), "fred"},
		{"TestLowerFirst-2", LowerFirst("Fred"), "fred"},
	}.validate(t)
}

func TestUpperFirstShouldLowercasedFirstCharacterOfString(t *testing.T) {
	testList{
		{"TestUpperFirst-0", UpperFirst("FRED"), "FRED"},
		{"TestUpperFirst-1", UpperFirst("fred"), "Fred"},
		{"TestUpperFirst-2", UpperFirst("Fred"), "Fred"},
	}.validate(t)
}

func TestIsEnclosedBetweenShouldChekcWhetherStringIsEnclosed(t *testing.T) {
	testListBool{
		{"TestIsEnclosedBetween-0", IsEnclosedBetween("%rajkumar%", "%"), true},
		{"TestIsEnclosedBetween-1", IsEnclosedBetween("rajkumar", "%"), false},
		{"TestIsEnclosedBetween-2", IsEnclosedBetween("*rajkumar*", "*"), true},
		{"TestIsEnclosedBetween-3", IsEnclosedBetween("rajkumar", "*"), false},
	}.validateBool(t)
}

func TestIsEnclosedBetweenTwoShouldChekcWhetherStringIsEnclosed(t *testing.T) {
	testListBool{
		{"TestIsEnclosedBetweenTwo-0", IsEnclosedBetweenTwo("{{rajkumar}}", "{{", "}}"), true},
		{"TestIsEnclosedBetweenTwo-1", IsEnclosedBetweenTwo("rajkumar", "{{", "}}"), false},
		{"TestIsEnclosedBetweenTwo-2", IsEnclosedBetweenTwo("<rajkumar>", "<", ">"), true},
	}.validateBool(t)
}

func TestTrimStartSpacesShouldRemoveAllWhitespaceAtStart(t *testing.T) {
	testList{
		{"TestTrimStartSpaces-0", TrimStartSpaces("   abc   "), "abc   "},
		{"TestTrimStartSpaces-1", TrimStartSpaces("abc   "), "abc   "},
		{"TestTrimStartSpaces-2", TrimStartSpaces("abc"), "abc"},
		{"TestTrimStartSpaces-3", TrimStartSpaces(""), ""},
	}.validate(t)
}

func TestTrimStartShouldRemoveAllSpecialCharsAtStart(t *testing.T) {
	testList{
		{"TestTrimStart-0", TrimStart("-_-abc-_-", "_-"), "abc-_-"},
		{"TestTrimStart-1", TrimStart("-_-!abc-_-", "_-!"), "abc-_-"},
		{"TestTrimStart-2", TrimStart("-_-#abc-_-", "_-#"), "abc-_-"},
	}.validate(t)
}

func TestTrimEndSpacesShouldRemoveAllTrailingWhitespace(t *testing.T) {
	testList{
		{"TestTrimEndSpaces-0", TrimEndSpaces("   abc   "), "   abc"},
		{"TestTrimEndSpaces-1", TrimEndSpaces("abc   "), "abc"},
		{"TestTrimEndSpaces-2", TrimEndSpaces("abc"), "abc"},
		{"TestTrimEndSpaces-3", TrimEndSpaces(""), ""},
	}.validate(t)
}

func TestTrimEndShouldRemoveAllSpecialCharsAtStart(t *testing.T) {
	testList{
		{"TestTrimEnd-0", TrimEnd("-_-abc-_-", "_-"), "-_-abc"},
		{"TestTrimEnd-1", TrimEnd("-_-abc!-_-", "_-!"), "-_-abc"},
		{"TestTrimEnd-2", TrimEnd("-_-abc#-_-", "_-#"), "-_-abc"},
	}.validate(t)
}

func TestCharsCountShouldReturnCharsCount(t *testing.T) {
	resmap := make(map[string]int)
	resmap["-"] = 10
	resmap["A"] = 1
	resmap["B"] = 2
	resmap["C"] = 3
	resmap["a"] = 1
	resmap["b"] = 2
	resmap["c"] = 3

	map1 := CharsCount("-----abbcccCCCBBA-----")
	if !reflect.DeepEqual(map1, resmap) {
		panic("strman:TestCharsCountShouldReturnCharsCount-0:: Failed!")
	}
}

func TestIsBlankShouldReturnTrueIfEmpty(t *testing.T) {
	testListBool{
		{"TestIsBlank-0", IsBlank(""), true},
		{"TestIsBlank-1", IsBlank("abc"), false},
	}.validateBool(t)
}

func TestUnderscoredShouldReturnUnderscoredString(t *testing.T) {
	testList{
		{"TestUnderscored-0", Underscored("MozTransform"), "moz_transform"},
		{"TestUnderscored-1", Underscored(""), ""},
	}.validate(t)
}

func TestZipShouldReturnExpectedListOfTuples_whenMoreThanThreeInputs(t *testing.T) {
	a := []string{"111111", "222222", "33", "444444", "555555"}
	b := []string{"0", "9", "8", "7", "6"}
	testvar, _ := Zip(a, b)
	targetvar := []StrTuple{{"111111", "0"}, {"222222", "9"}, {"33", "8"}, {"444444", "7"}, {"555555", "6"}}

	if !reflect.DeepEqual(testvar, targetvar) {
		panic("strman:TestCharsCountShouldReturnCharsCount-0:: Failed!")
	}

}

func TestLinesShouldSplitToLines(t *testing.T) {
	testListInt{
		{"TestLines-0", len(Lines("Hello\r\nWorld")), 2},
		{"TestLines-1", len(Lines("Hello\rWorld")), 2},
		{"TestLines-2", len(Lines("Hello World")), 1},
		{"TestLines-3", len(Lines("\r\n\n\r ")), 4},
		{"TestLines-4", len(Lines("Hello\r\r\nWorld")), 3},
		{"TestLines-5", len(Lines("Hello\r\rWorld")), 3},
	}.validateInt(t)

}

func TestDasherizeShouldDasherizeInputString(t *testing.T) {
	testList{
		{"TestDasherize-0", Dasherize("the_dasherize_string_method"), "the-dasherize-string-method"},
		{"TestDasherize-1", Dasherize("TheDasherizeStringMethod"), "the-dasherize-string-method"},
		{"TestDasherize-2", Dasherize("thisIsATest"), "this-is-a-test"},
		{"TestDasherize-3", Dasherize("this Is A Test"), "this-is-a-test"},
		{"TestDasherize-4", Dasherize("input with a-dash"), "input-with-a-dash"},
		{"TestDasherize-5", Dasherize(""), ""},
	}.validate(t)
}

func TestHumanize(t *testing.T) {
	testList{
		{"TestHumanize-0", Humanize("the_humanize_method"), "The humanize method"},
		{"TestHumanize-1", Humanize("ThehumanizeMethod"), "Thehumanize method"},
		{"TestHumanize-2", Humanize("ThehumanizeMethod"), "Thehumanize method"},
		{"TestHumanize-3", Humanize("the  humanize  method  "), "The humanize method"},
		{"TestHumanize-4", Humanize("   capitalize dash-CamelCase_underscore trim  "), "Capitalize dash camel case underscore trim"},
		{"TestHumanize-5", Humanize(""), ""},
	}.validate(t)
}

func TestSwapCaseShouldSwapCaseOfCharacters(t *testing.T) {
	testList{
		{"TestSwapCase-0", SwapCase("AaBbCcDdEe"), "aAbBcCdDeE"},
		{"TestSwapCase-1", SwapCase("Hello World"), "hELLO wORLD"},
		{"TestSwapCase-2", SwapCase("Hello-World"), "hELLO-wORLD"},
		{"TestSwapCase-3", SwapCase(""), ""},
	}.validate(t)
}

func TestFormatNumber(t *testing.T) {
	testList{
		{"TestFormatNumber-0", FormatNumber(1000), "1,000"},
		{"TestFormatNumber-1", FormatNumber(100000), "100,000"},
		{"TestFormatNumber-2", FormatNumber(10000000), "10,000,000"},
		{"TestFormatNumber-3", FormatNumber(100000000), "100,000,000"},
	}.validate(t)
}

//end of file 
