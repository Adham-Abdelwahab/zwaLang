package types

import "fmt"

func Atoi(str string) int {
	var result int
	fmt.Sscanf(str, "%d", &result)
	return result
}
func Itoa(num int) string {
	result := fmt.Sprintf("%d", num)
	return result
}
func Btoa(boolean string) bool {
	return boolean == "true"
}
func Atob(boolean bool) string {
	if boolean {
		return "true"
	}
	return "false"
}
func Btoi(boolean bool) int {
	if boolean {
		return 1
	}
	return 0
}

/* Int And ____ */
func IntAndInt(lhs int, operator rune, rhs int) int {
	switch operator {
	case '+':
		return lhs + rhs
	case '-':
		return lhs - rhs
	case '*':
		return lhs * rhs
	case '/':
		return lhs / rhs
	case '%':
		return lhs % rhs
	case '&':
		return min(lhs, rhs)
	case '|':
		return max(lhs, rhs)
	}
	return 0
}
func IntAndBool(lhs int, operator rune, rhs bool) any {
	if operator == '/' {
		operator = '*'
	}
	boolean := lhs > 0
	switch operator {
	case '+', '-', '*', '%':
		return IntAndInt(lhs, operator, Btoi(rhs))
	case '&':
		return boolean && rhs
	case '|':
		return boolean || rhs
	}
	return false
}
func IntAndString(lhs int, operator rune, rhs string) string {
	switch operator {
	case '+':
		return rhs + Itoa(lhs)
	case '-':
		return cut(rhs, lhs)
	case '*':
		return explode(rhs, lhs)
	}
	return "INVALID OPERATOR"
}

/* String And ____ */
func StringAndString(lhs string, operator rune, rhs string) string {
	switch operator {
	case '+':
		return lhs + rhs
	case '-':
		return remove(lhs, rhs)
	}
	return ""
}
func StringAndBool(lhs string, operator rune, rhs bool) any {
	boolean := lhs != ""
	switch operator {
	case '+':
		return lhs + Atob(rhs)
	case '*':
		if rhs {
			return lhs
		}
		return ""
	case '&':
		return boolean && rhs
	case '|':
		return boolean || rhs
	}
	return false
}

/* Bool And ____ */
func BoolAndBool(lhs bool, operator rune, rhs bool) bool {
	switch operator {
	case '&', '*', '/':
		return lhs && rhs
	case '|':
		return lhs || rhs
	}
	return false
}

func (binary BinaryExpression) perform(env env) any {
	lhs := binary.Lhs.Concrete(env)
	rhs := binary.Rhs.Concrete(env)
	operator := binary.Operator

	switch lhs := lhs.(type) {
	case int:
		switch rhs := rhs.(type) {
		case int:
			return IntAndInt(lhs, operator, rhs)
		case string:
			return IntAndString(lhs, operator, rhs)
		case bool:
			return IntAndBool(lhs, operator, rhs)
		}
	case string:
		switch rhs := rhs.(type) {
		case int:
			return IntAndString(rhs, operator, lhs)
		case string:
			return StringAndString(lhs, operator, rhs)
		case bool:
			return StringAndBool(lhs, operator, rhs)
		}
	case bool:
		switch rhs := rhs.(type) {
		case int:
			return IntAndBool(rhs, operator, lhs)
		case string:
			return StringAndBool(rhs, operator, lhs)
		case bool:
			return BoolAndBool(lhs, operator, rhs)
		}
	}

	fmt.Printf("Unable to process binary expression: %T %c %T\n", lhs, binary.Operator, rhs)
	return nil
}

func cut(from string, at int) string {
	if at >= 0 && at < len(from) {
		return from[0:at] + from[at+1:]
	}
	return from
}
func explode(with string, to int) string {
	if to == 0 {
		return ""
	} else if to == 1 {
		return with
	}
	return with + explode(with, to-1)
}
func remove(from string, substring string) string {
	length := len(from)
	lengthsubstr := len(substring)

	var modified = false
	var r string
	for i := 0; i < length; {
		if i <= length-lengthsubstr && from[i:i+lengthsubstr] == substring {
			i += lengthsubstr
			modified = true
		} else {
			r += string(from[i])
			i++
		}
	}

	if modified {
		return remove(r, substring)
	} else {
		return r
	}
}
