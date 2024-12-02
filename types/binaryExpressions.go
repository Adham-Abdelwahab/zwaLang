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

func (binary BinaryExpression) perform(lhs any, rhs any) any {
	switch lhs := lhs.(type) {
	case int:
		switch rhs := rhs.(type) {
		case int:
			switch binary.Operator {
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
			}
		case string:
			switch binary.Operator {
			case '+':
				return Itoa(lhs) + rhs
			case '*':
				return explode(rhs, lhs)
			}
		}
	case string:
		switch rhs := rhs.(type) {
		case int:
			switch binary.Operator {
			case '+':
				return lhs + Itoa(rhs)
			case '-':
				return cut(lhs, rhs)
			case '*':
				return explode(lhs, rhs)
			}
		case string:
			switch binary.Operator {
			case '+':
				return lhs + rhs
			case '-':
				return remove(lhs, rhs)
			}
		}
	}

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
