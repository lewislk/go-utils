package strs

import "unicode"

func IsBlank(s string) bool {
	if s == "" {
		return true
	}
	for _, v := range s {
		if !unicode.IsSpace(v) {
			return false
		}
	}
	return true
}

func IsPtrBlank(s *string) bool {
	return s == nil || IsBlank(*s)
}

func IsEmpty(s string) bool {
	return s == ""
}

func IsPtrEmpty(s *string) bool {
	return s == nil || IsEmpty(*s)
}

func IsNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, v := range s {
		if !unicode.IsNumber(v) {
			return false
		}
	}
	return true
}

func IsPtrNumeric(s *string) bool {
	return s != nil && IsNumeric(*s)
}
