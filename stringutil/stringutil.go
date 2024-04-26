package stringutil


import (

	"strings"
)

type StringFunction func(string) string

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func ToUpper(s string) string {
    return strings.ToUpper(s)
}

func Concat(strs ...string) string {
    return strings.Join(strs, "")
}

func IsEmpty(s string) bool {
    return len(s) == 0
}

func FirstRune(s string) rune {
    if len(s) == 0 {
        return 0
    }
    return []rune(s)[0]
}

func GetStringFunction(funcType string) StringFunction {
    switch funcType {
    case "reverse":
        return Reverse
    case "toUpper":
        return ToUpper
   
    case "isEmpty":
        return func(s string) string {
            if IsEmpty(s) {
                return "The string is empty"
            }
            return "The string is not empty"
        }
    case "firstRune":
        return func(s string) string {
            if len(s) == 0 {
                return "The string is empty"
            }
            return string(FirstRune(s))
        }
    default:
        return func(s string) string {
            return "Invalid function type"
        }
    }
}


