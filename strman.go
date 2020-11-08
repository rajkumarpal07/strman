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
//strman.go ... Created on :: 13-June-2020

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"html"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Appends strs to value
func Appends(value string, strs ...string) string {
	return AppendsArray(value, strs)
}

// AppendsArray an array of String to value
func AppendsArray(value string, appends []string) string {
	if appends == nil || len(appends) == 0 {
		return value
	}
	l := len(appends)
	slice := make([]string, (l+1)*2)
	slice[0] = value

	slice = append(slice, appends...)
	return strings.Join(slice, "")
}

// At :: Get the character at index. This method will take care of negative indexes.
// The valid value of index is between -(length-1) to (length-1).
// For values which don't fall under this range empty string will be returned.
func At(value string, index int) string {
	result := ""
	if isNilOrEmpty(value) {
		return result
	}
	length := len(value)
	if index < 0 {
		index = length + index
	}
	if index < length && index >= 0 {
		byteI := byte(index)
		result = value[byteI : byteI+1]
	}
	return result
}

// Between Returns an array with strings between start and end.
func Between(value string, start string, end string) ([]string, error) {
	parts := strings.Split(value, end)
	result := make([]string, len(parts)-1)

	if isNilOrEmpty(value) {
		return result, fmt.Errorf("strman:Between::value cannot be empty or nil")
	} else if isNilOrEmpty(start) || isNilOrEmpty(end) {
		return result, fmt.Errorf("strman:Between::start or end cannot be empty or nil")
	}

	var j int
	for j = 0; j < len(parts)-1; j++ {
		tempstr := parts[j]
		lIndex := strings.LastIndex(tempstr, start)
		result[j] = tempstr[(len(start) + lIndex):len(tempstr)]
	}
	return result, nil
}

//Chars Returns a String array consisting of the characters in the String.
func Chars(value string) ([]string, error) {

	//value = strings.Trim(value, " ")
	result := make([]string, len(value)-1)
	if isNilOrEmpty(value) {
		return result, fmt.Errorf("strman:Chars::value cannot be empty or nil")
	}

	result = strings.Split(value, "")
	return result, nil
}

//CollapseWhitespace Replace consecutive whitespace characters with a single space.
func CollapseWhitespace(value string) string {
	//\s matches a space, tab, new line, carriage return or form feed -- or does it???
	re := regexp.MustCompile(`\s+`)
	value = strings.TrimSpace(value)
	result := re.ReplaceAllString(value, " ")
	return result
}

//Contains Verifies that the needle is contained in the value. The search is case sensitive
func Contains(value string, needle string, caseSensitive bool) bool {
	if caseSensitive {
		return strings.Contains(value, needle)
	}
	return strings.Contains(strings.ToLower(value), strings.ToLower(needle))
}

//ContainsAll Verifies that all needles are contained in value. The search is case insensitive
//needs refactoring
func ContainsAll(value string, needles []string) bool {

	oneNotFound := true
	// Traversing the array
	for i := 0; i < len(needles); i++ {
		if !Contains(value, needles[i], false) {
			oneNotFound = false
		}
	}
	return oneNotFound
}

//ContainsAllCaseSensitive Verifies that all needles are contained in value. The search is case sensitive
func ContainsAllCaseSensitive(value string, needles []string, caseSensitive bool) bool {

	oneNotFound := true
	// Traversing the needles array
	for i := 0; i < len(needles); i++ {
		if !Contains(value, needles[i], true) {
			oneNotFound = false
		}
	}
	return oneNotFound
}

//ContainsAny Verifies that one or more of needles are contained in value. This is case insensitive
func ContainsAny(value string, needles []string) bool {

	return ContainsAnyCaseSensitive(value, needles, false)
}

//ContainsAnyCaseSensitive Verifies that one or more of needles are contained in value.
func ContainsAnyCaseSensitive(value string, needles []string, caseSensitive bool) bool {

	if caseSensitive {
		// Traversing the needles array
		for i := 0; i < len(needles); i++ {
			if strings.Contains(value, needles[i]) {
				return true
			}
		}
	} else {
		for i := 0; i < len(needles); i++ {
			if strings.Contains(strings.ToLower(value), strings.ToLower(needles[i])) {
				return true
			}
		}
	}
	return false
}

//CountSubstr Count the number of times substr appears in value- Case Sensitive
func CountSubstr(value string, substr string) int {
	return CountSubstring(value, substr, true, false)
}

//CountSubstring Count the number of times substr appears in value
func CountSubstring(value string, substr string, caseSensitive bool, allowOverlapping bool) int {
	if !caseSensitive {
		value = strings.ToLower(value)
		substr = strings.ToLower(substr)
	}
	return CountsSubstr(value, substr, allowOverlapping, 0)
}

//CountsSubstr Counts the number of times substr appears in value
func CountsSubstr(value string, substr string, allowOverlapping bool, count int) int {
	position := strings.Index(value, substr)
	if position == -1 {
		return count
	}
	offset := 0
	if !allowOverlapping {
		offset = position + len(substr)
	} else {
		offset = position + 1
	}
	count = count + 1
	return CountsSubstr(value[offset:], substr, allowOverlapping, count)
}

//EndsWith Test if value ends with search. The search is case sensitive.
func EndsWith(value string, search string) bool {
	return EndsWithCase(value, search, true)
}

//EndsWithCase Test if value ends with search. The search is case sensitive.
func EndsWithCase(value string, search string, caseSensitive bool) bool {

	if caseSensitive {
		return strings.HasSuffix(value, search)
	}
	return strings.HasSuffix(strings.ToLower(value), strings.ToLower(search))
}

//EnsureLeft Ensures that the value begins with prefix. If it doesn't exist, it's prepended.
//It is case sensitive.
func EnsureLeft(value string, prefix string) string {
	return EnsureLeftWithCase(value, prefix, true)
}

//EnsureLeftWithCase Ensures that the value begins with prefix. If it doesn't exist, it's prepended.
func EnsureLeftWithCase(value string, prefix string, caseSensitive bool) string {
	if caseSensitive {
		if strings.HasPrefix(value, prefix) {
			return value
		}
		return prefix + value
	}
	_value := strings.ToLower(value)
	_prefix := strings.ToLower(prefix)
	if strings.HasPrefix(_value, _prefix) {
		return _value
	}
	return _prefix + _value
}

//Base64Encode Encodes data with MIME base64.
func Base64Encode(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

//Base64Decode Decodes data encoded with MIME base64
func Base64Decode(value string) string {
	data, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return ""
	}
	return string(data)
}

//DecEncode Convert string chars to decimal unicode (16 digits)
func DecEncode(value string) string {
	var b []byte
	for _, c := range value {
		b = strconv.AppendInt(b, int64(c), 10)
	}
	return string(b)
}

//EnsureRight Ensures that the value ends with suffix. If it doesn't, it's appended.
//This operation is case sensitive.
func EnsureRight(value string, suffix string) string {

	return EnsureRightWithCase(value, suffix, true)
}

//EnsureRightWithCase Ensures that the value ends with suffix. If it doesn't, it's appended.
//This operation is case sensitive.
func EnsureRightWithCase(value string, suffix string, caseSensitive bool) string {
	if caseSensitive {
		if strings.HasSuffix(value, suffix) {
			return value
		}
		return value + suffix
	}
	_value := strings.ToLower(value)
	_suffix := strings.ToLower(suffix)
	if strings.HasSuffix(_value, _suffix) {
		return _value
	}
	return _value + _suffix
}

//First Returns the first 3 chars of String(ASCII)
func First(value string) string {
	return value[0:3]
}

//FirstNChars Returns the first n chars of String(ASCII)
func FirstNChars(value string, n int) string {
	return value[0:n]
}

//Format Formats a string using parameters
func Format(value string, params ...string) string {
	result := value
	return result
}

//Head Return the first char of String(ASCII)
func Head(value string) string {
	return FirstNChars(value, 1)
}

//HexDecode Convert hexadecimal unicode (4 digits) string to string chars
func HexDecode(value string) string {
	hs, _ := hex.DecodeString(value)
	return string(hs)
}

//HexEncode Convert string chars to hexadecimal unicode (4 digits)
func HexEncode(value string) string {
	b := []byte(value)
	s := hex.EncodeToString(b)
	return s
}

// IndexOf The indexOf() method returns the index within the calling String of the
// first occurrence of the specified value, starting the search at fromIndex.
// Returns -1 if the value is not found.
func IndexOf(value string, needle string, caseSensitive bool) int {
	if caseSensitive {
		return strings.Index(value, needle)
	}
	res1 := strings.Index(strings.ToLower(value), needle)
	return res1
}

//UnEqual Tests if two Strings are InEqual
func UnEqual(first string, second string) bool {
	equalToZero := strings.Compare(first, second)
	if equalToZero == 0 {
		return false
	}
	return true
}

// Insert Inserts 'substring' into the 'value' at the 'index' provided.
func Insert(value string, substring string, index int) string {
	res1 := ""
	if index > len(value) {
		return value
	}
	if len(value) == index { //after last element
		res1 = value[:index] + substring
		return res1
	}
	res1 = value[:index] + substring + value[index:]
	return res1
}

//IsLowerCase Verifies if String is lowercase
func IsLowerCase(value string) bool {
	val := []rune(value)
	for i := 0; i < len(val); i++ {
		if unicode.IsUpper(val[i]) == true {
			return false
		}
	}
	return true
}

//IsUpperCase Verifies if String is uppercase
func IsUpperCase(value string) bool {
	val := []rune(value)
	for i := 0; i < len(val); i++ {
		if unicode.IsLower(val[i]) == true {
			return false
		}
	}
	return true
}

//Last Return the last n chars of String
func Last(value string, n int) string {
	if n > len(value) {
		return value
	}
	l := len(value)
	res1 := value[l-n : l]
	return res1
}

//LeftPad Returns a new string of a given length such that
//the beginning of the string is padded.
func LeftPad(value string, padStr string, pLen int) string {
	if len(value) > pLen {
		return value
	}
	return strings.Repeat(padStr, pLen) + value
}

//RightPad2Len Right pad a string with padStr
// If the overallLen is shorter than the original string length
// the string will be shortened to this length (substr)
func RightPad2Len(value string, padStr string, overallLen int) string {
	var padCount int
	padCount = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = value + strings.Repeat(padStr, padCount)
	return retStr[:overallLen]
}

//LeftPad2Len Left pad a string with padStr
// If the overallLen is shorter than the original string length
// the string will be shortened to this length (substr)
func LeftPad2Len(value string, padStr string, overallLen int) string {
	var padCount int
	padCount = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCount) + value
	return retStr[(len(retStr) - overallLen):]
}

//IsString Checks whether parameter is String
func IsString(value interface{}) bool {
	res := reflect.TypeOf(value).Kind()
	if res == reflect.String {
		return true
	}
	return false
}

//LastIndexOf returns the index within the calling String object of the
// last occurrence of the specified value
// Returns -1 if the value is not found. The search starts from the end and case sensitive.
func LastIndexOf(value string, needle string) int {
	res1 := strings.LastIndex(value, needle)
	return res1
}

//LastIndexOfWithCase returns the index within the calling String object of the
//last occurrence of the specified value
//Returns -1 if the value is not found. The search starts from the end and case sensitive.
func LastIndexOfWithCase(value string, needle string, caseSensitive bool) int {
	res1 := 0
	if caseSensitive {
		res1 = strings.LastIndex(value, needle)
		return res1
	}
	_value := strings.ToLower(value)
	_needle := strings.ToLower(needle)
	return strings.LastIndex(_value, _needle)
}

//LeftTrim Removes all spaces on left
func LeftTrim(value string) string {
	return strings.TrimLeft(value, " ")
}

//Length Returns length of String. Delegates to len method.
func Length(value string) int {
	return len(value)
}

//Prepend Return a new String starting with prepends
func Prepend(value string, prepends ...string) string {
	return PrependArray(value, prepends)
}

//PrependArray Return a new String starting with prepends
func PrependArray(value string, prepends []string) string {
	if prepends == nil || len(prepends) == 0 {
		return value
	}
	result := strings.Join(prepends[:], "")
	return result + value
}

//RemoveEmptyStrings Remove empty Strings from string array
func RemoveEmptyStrings(stringsArr []string) ([]string, error) {
	if stringsArr == nil {
		return stringsArr, fmt.Errorf("strman:RemoveEmptyString:: Input array should not be nil")
	}
	result := []string{}
	for i := 0; i < len(stringsArr); i++ {

		if len(strings.TrimSpace(stringsArr[i])) != 0 {
			result = append(result, stringsArr[i])
		}
	}
	return result, nil
}

//RemoveLeft Returns a new String with the prefix removed, if present. This is case sensitive.
func RemoveLeft(value string, prefix string) string {
	return RemoveLeftWithCase(value, prefix, true)
}

//RemoveLeftWithCase Returns a new String with the prefix removed, if present.
func RemoveLeftWithCase(value string, prefix string, caseSensitive bool) string {
	result := value
	if caseSensitive {
		if strings.HasPrefix(value, prefix) {
			result = value[len(prefix):]
			return result
		}
	}
	_value := strings.ToLower(value)
	_prefix := strings.ToLower(prefix)

	if strings.HasPrefix(_value, _prefix) {
		result = _value[len(prefix):]
	}
	return result
}

//RemoveNonWords Remove all non word characters.
//ACCENTED CHARS are removed including whitespace
func RemoveNonWords(value string) string {
	result := value
	pattern := "[^a-zA-Z0-9]+"
	re := regexp.MustCompile(pattern)
	result = re.ReplaceAllString(value, "")
	return result
}

//RemoveRight Returns a new string with the 'suffix' removed, if present. Search is case sensitive.
func RemoveRight(value string, suffix string) string {
	return RemoveRightWithCase(value, suffix, true)
}

//RemoveRightWithCase Returns a new string with the 'suffix' removed, if present.
func RemoveRightWithCase(value string, suffix string, caseSensitive bool) string {
	result := value
	ends := EndsWithCase(value, suffix, caseSensitive)
	if ends {
		_value := strings.ToLower(value)
		_suffix := strings.ToLower(suffix)
		_index := LastIndexOf(_value, _suffix)
		result = _value[0:_index]
	}
	return result
}

//RemoveSpaces Removes all spaces
func RemoveSpaces(value string) string {
	re1, _ := regexp.Compile("[\\s]+")
	result := re1.ReplaceAllString(value, "")
	return result
}

//Repeat Returns a repeated string given a multiplier.
func Repeat(value string, multiplier int) string {
	var sb strings.Builder
	bufsize := multiplier * len(value)
	sb.Grow(bufsize)
	for i := 0; i < multiplier; i++ {
		sb.WriteString(value)
	}
	return sb.String()
}

//Replace Replace all occurrences of 'search' value to 'newvalue'. Uses ReplaceAllString method.
func Replace(value string, search string, newValue string, caseSensitive bool) string {
	result := ""
	if caseSensitive {
		return strings.ReplaceAll(value, search, newValue)
	}
	_value := strings.ToLower(value)
	_search := strings.ToLower(search)
	pattern := "[(?i)" + _search + "]+"
	re := regexp.MustCompile(pattern)
	result = re.ReplaceAllString(_value, newValue)
	return result
}

//Reverse Reverse the input String
func Reverse(value string) string {
	result := ""
	for _, v := range value {
		result = string(v) + result
	}
	return result
}

//RightPad Returns a new string of a given length such that the ending of the string is padded.
func RightPad(value string, pad string, length int) string {
	if len(value) > length {
		return value
	}
	mySlice1 := make([]string, 0)
	mySlice1 = append(mySlice1, value)
	result := append(mySlice1, Repeat(pad, length-len(value)))
	return strings.Join(result, "")
}

//RightTrim Remove all spaces on right.
func RightTrim(value string) string {
	return strings.TrimRight(value, " ")
}

//SafeTruncate Truncate the string securely, not cutting a word in half. It always returns the last full word.
func SafeTruncate(value string, length int, filler string) string {
	if length == 0 {
		return ""
	}
	if length >= len(value) {
		return value
	}
	words := Words(value)
	spaceCount := 0
	tresult := ""
	for i, word := range words {
		if len(tresult)+len((word))+len(filler)+spaceCount > length {
			break
		} else {
			if i == 0 {
				tresult = tresult + word
				spaceCount++
			} else {
				tresult = tresult + " " + word
				spaceCount++
			}
		}
	}
	return tresult + filler
}

//Split Alias for String split function. Defined only for completeness.
func Split(value string, delimiterregex string) []string {
	a := regexp.MustCompile(delimiterregex)
	temp := a.Split(value, -1)
	result := []string{}
	for i := range temp {
		result = append(result, temp[i])
	}
	return result
}

//Truncate Truncates the unsecured form string, cutting the independent string of required position.
func Truncate(value string, length int, filler string) string {
	if length == 0 {
		return ""
	}
	if length >= len(value) {
		return value
	}
	result := value[0:(length-len(filler))] + filler
	return result
}

//HTMLDecode Converts all HTML entities to applicable characters.
//UnescapeString unescapes entities like "&lt;" to become "<".
// &aacute; unescapes to "á"
func HTMLDecode(encodedHTML string) string {
	unescaped := html.UnescapeString(encodedHTML)
	return unescaped
}

//HTMLEncode Convert all applicable characters to HTML entities.
//EscapeString escapes special characters like "<" to become "&lt;"
//NOTE:: "á" doesnt return currently "&aacute;" but returns "á"
func HTMLEncode(plainHTML string) string {
	escaped := html.EscapeString(plainHTML)
	return escaped
}

//Shuffle It returns a string with its characters in random order.
func Shuffle(value string) string {
	rand.Seed(time.Now().Unix())
	arrRune := []rune(value)
	rand.Shuffle(len(arrRune), func(i, j int) {
		arrRune[i], arrRune[j] = arrRune[j], arrRune[i]
	})
	return string(arrRune)
}

//Slice A substring method
func Slice(value string, begin int, end int) string {
	return value[begin:end]
}

//Slugify Convert a String to a slug
func Slugify(value string) string {
	transliterated := Transliterate(CollapseWhitespace(strings.ToLower(strings.Trim(value, " "))))
	transliterated = Replace(transliterated, "&", "-and-", false)
	res1 := Words(transliterated)
	res1, _ = RemoveEmptyStrings(res1)
	res2 := strings.Join(res1, "-")
	return res2
}

//Transliterate Remove all non valid characters.
func Transliterate(value string) string {
	LoadMapArrays()
	res1 := value
	for _, char1 := range res1 {
		for key, innerval := range Hm {
			for _, v := range innerval {
				if strings.Contains(string(char1), v) {
					res1 = strings.ReplaceAll(res1, string(char1), key)
				}
			}
		}
	}
	return res1
}

//Surround Surrounds a 'value' with the given 'prefix' and 'suffix'.
func Surround(value string, prefix string, suffix string) string {
	_prefix := prefix
	_suffix := suffix
	if isNilOrEmpty(prefix) {
		_prefix = ""
	}
	if isNilOrEmpty(suffix) {
		_suffix = ""
	}
	return (_prefix + value + _suffix)
}

//ToCamelCase Transform to camelCase
func ToCamelCase(value string) string {
	if isNilOrEmpty(value) {
		return ""
	}
	studlyStr := ToStudlyCase(value)
	result := strings.ToLower(studlyStr[:1]) + studlyStr[1:]
	return result
}

//ToStudlyCase Transform to StudlyCaps.
func ToStudlyCase(value string) string {
	result := ""
	twords := CollapseWhitespace(strings.Trim(value, " "))
	words := Split(twords, "\\s*(_|-|\\s)\\s*")

	resultarr := []string{}
	for i := range words {
		words[i] = strings.Title(words[i])
		resultarr = append(resultarr, words[i])
	}
	result = strings.Join(resultarr, "")
	return result
}

//Tail Return tail of the String
func Tail(value string) string {
	return Last(value, len(value)-1)
}

//ToDecamelize Decamelize String
func ToDecamelize(value string, chr string) string {
	if isNilOrEmpty(chr) {
		chr = " "
	}
	camelCasedString := ToCamelCase(value)
	re := regexp.MustCompile(`([A-Z][a-z]*)|([a-z]*)`)
	words := re.FindAllString(camelCasedString, -1)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	result := strings.Join(words, chr)
	return result
}

//ToKebabCase Transform to kebab-case.
func ToKebabCase(value string) string {
	return ToDecamelize(value, "-")
}

//ToSnakeCase Transform to snake_case.
func ToSnakeCase(value string) string {
	return ToDecamelize(value, "_")
}

//Join concatenates all the elements of the strings array into a single String.
//The separator string is placed between elements in the resulting string.
func Join(strarr []string, separator string) string {
	result := strings.Join(strarr, separator)
	return result
}

//Capitalize Converts the first character of string to upper case and the remaining to lower case.
func Capitalize(value string) string {
	if len(value) == 0 {
		return ""
	}
	result := strings.ToLower(value)
	runes := []rune(result)
	if unicode.IsLower(runes[0]) {
		runes[0] = unicode.ToUpper(runes[0])
		result = string(runes)
	}
	return result
}

//LowerFirst Converts the first character of string to lower case.//
func LowerFirst(value string) string {
	if len(value) == 0 {
		return ""
	}
	result := value
	runes := []rune(value)
	if unicode.IsUpper(runes[0]) {

		runes[0] = unicode.ToLower(runes[0])
		result = string(runes)
	}
	return result
}

//UpperFirst Converts the first character of string to upper case.
func UpperFirst(value string) string {
	if len(value) == 0 {
		return ""
	}
	result := value
	runes := []rune(value)
	if unicode.IsLower(runes[0]) {
		runes[0] = unicode.ToUpper(runes[0])
		result = string(runes)
	}
	return result
}

//IsEnclosedBetween Verifies whether String is enclosed by same encloser
func IsEnclosedBetween(value string, encloser string) bool {
	result := IsEnclosedBetweenTwo(value, encloser, encloser)
	return result
}

//IsEnclosedBetweenTwo Verifies whether String is enclosed by 2 different enclosers
func IsEnclosedBetweenTwo(value string, leftEncloser string, rightEncloser string) bool {
	result := strings.HasPrefix(value, leftEncloser) && strings.HasSuffix(value, rightEncloser)
	return result
}

//TrimStartSpaces Removes leading whitespace from string.
func TrimStartSpaces(value string) string {
	return TrimStart(value, " ")
}

//TrimStart Removes leading chars from string.
func TrimStart(value string, chars string) string {
	return strings.TrimLeft(value, chars)
}

//TrimEndSpaces Removes trailing whitespace from string.
func TrimEndSpaces(value string) string {
	return TrimEnd(value, " ")
}

//TrimEnd Removes trailing chars from string.
func TrimEnd(value string, chars string) string {
	return strings.TrimRight(value, chars)
}

//CharsCount Counts the number of occurrences of each character in the string
func CharsCount(invalue string) map[string]int {
	setmap := make(map[string]int)
	runes := []rune(invalue)
	for i := 0; i < len(runes); i++ {
		if x, found := setmap[string(runes[i])]; found {
			setmap[string(runes[i])] = x + 1
		} else {
			setmap[string(runes[i])] = 1
		}
	}
	return setmap
}

//IsBlank Checks if string is empty.
func IsBlank(value string) bool {
	return len(value) == 0
}

//Underscored Changes passed in string to all lower case and adds underscore between words.
func Underscored(value string) string {
	camelCasedString := ToCamelCase(value)
	re := regexp.MustCompile(`([A-Z][a-z]*)|([a-z]*)`)
	words := re.FindAllString(camelCasedString, -1)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	result := strings.Join(words, "_")
	return result
}

//Words Splits a String to words
func Words(value string) []string {
	return WordsDelimitedBy(value, "\\W+")
}

//WordsDelimitedBy Splits a String to words
func WordsDelimitedBy(value string, delimiter string) []string {
	return Split(value, delimiter)
}

//StrTuple is a string tuple for use in the Zip function.
type StrTuple struct {
	a, b string
}

//Zip Aggregates the contents of 2 strings arrays into a single list of tuples.
func Zip(a, b []string) ([]StrTuple, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("zipzoom: arguments must be of same length")
	}
	r := make([]StrTuple, len(a), len(a))
	for i, e := range a {
		r[i] = StrTuple{e, b[i]}
	}
	return r, nil
}

//Lines Split lines to an array
func Lines(input string) []string {
	pass1 := strings.Replace(input, "\r\n", "\n", -1)
	pass2 := strings.Replace(pass1, "\r", "\n", -1)
	lines2 := strings.Split(pass2, "\n")
	return lines2
}

//Dasherize Converts a underscored or camelized string into an dasherized one.
func Dasherize(input string) string {
	return ToKebabCase(input)
}

//Humanize Converts an underscored, camelized, or dasherized string into a humanized one.
//Also removes beginning and ending whitespace.
func Humanize(input string) string {
	result := UpperFirst(Underscored(input))
	return strings.ReplaceAll(result, "_", " ")
}

//SwapCase Returns a copy of the string in which all the case-based characters have had their case swapped.
func SwapCase(input string) string {
	if len(input) == 0 {
		return ""
	}
	runes := []rune(input)
	for i := 0; i < len(runes); i++ {
		if unicode.IsUpper(runes[i]) {
			runes[i] = unicode.ToLower(runes[i])
		} else {
			runes[i] = unicode.ToUpper(runes[i])
		}
	}
	return string(runes)
}

//FormatNumber  Returns a string representation of the number passed in where groups of three digits are delimited by comma
func FormatNumber(number int64) string {
	sign := ""
	if number < 0 {
		sign = "-"
		number = 0 - number
	}
	parts := []string{"", "", "", "", "", "", ""}
	j := len(parts) - 1

	for number > 999 {
		parts[j] = strconv.FormatInt(number%1000, 10)
		switch len(parts[j]) {
		case 2:
			parts[j] = "0" + parts[j]
		case 1:
			parts[j] = "00" + parts[j]
		}
		number = number / 1000
		j--
	}
	parts[j] = strconv.Itoa(int(number))
	return sign + strings.Join(parts[j:], ",")
}

//isNilOrEmpty is nil or empty string is supplied
func isNilOrEmpty(empstr string) bool {
	empstr = strings.Trim(empstr, " ")
	return len(empstr) == 0
}
