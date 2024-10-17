package utilities

import "regexp"

const EmailRegex = "^[\\w\\.-]+@[a-zA-Z\\d\\.-]+\\.[a-zA-Z]{2,}$"
const PhoneRegexInternational = "^\\+?[1-9]\\d{1,14}$"
const PhoneRegexVietNam = "^(\\+84|0)(3|5|7|8|9)\\d{8}$"
const NumberRegex = "^-?\\d+(\\.\\d+)?$"
const ExceptSpecialCharacter = "^[^!@#$%^&*(),.?\":{}|<>]+$"
const OnlyLetter = "^[a-zA-Z]+$"

func CheckRegex(regex string, input string) bool {
	re := regexp.MustCompile(regex)
	return re.MatchString(input)
}
