package problem0020

func isValid(s string) bool {
	var bracketStack []rune

	for _, bracket := range s {
		switch bracket {
		case '(', '{', '[':
			bracketStack = append(bracketStack, bracket)
		case ')':
			if len(bracketStack) == 0 || bracketStack[len(bracketStack)-1] != '(' {
				return false
			}
			bracketStack = bracketStack[:len(bracketStack)-1]
		case '}':
			if len(bracketStack) == 0 || bracketStack[len(bracketStack)-1] != '{' {
				return false
			}
			bracketStack = bracketStack[:len(bracketStack)-1]
		case ']':
			if len(bracketStack) == 0 || bracketStack[len(bracketStack)-1] != '[' {
				return false
			}
			bracketStack = bracketStack[:len(bracketStack)-1]
		default:
			return false
		}
	}
	if len(bracketStack) > 0 {
		return false
	}

	return true
}
