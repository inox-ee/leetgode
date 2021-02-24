package problem0008

import (
	"math"
	"strconv"
	"strings"
)

func myAtoi(s string) int {
	whiteSpace := " "
	signs := "+-"
	numbers := "0123456789"

	readin := false
	tmp := ""
	sign := 1
	nextSign := false

	strs := strings.Split(s, "")
	for _, str := range strs {
		if readin {
			if strings.Contains(numbers, str) {
				tmp += str
			} else {
				break
			}
		} else {
			if nextSign {
				if strings.Contains(numbers, str) {
					readin = true
					tmp += str
				} else {
					break
				}
			} else if str == whiteSpace {
				continue
			} else if strings.Contains(signs, str) {
				if str == "-" {
					sign = -1
				}
				nextSign = true
			} else if strings.Contains(numbers, str) {
				readin = true
				tmp += str
			} else {
				break
			}
		}
	}

	result, ok := strconv.Atoi(tmp)
	if result == 0 {
		return result
	} else if ok != nil || result >= int(math.Pow(2, 31)) {
		if math.Signbit(float64(sign)) {
			return int(math.Pow(2, 31) * float64(sign))
		}
		return int(math.Pow(2, 31) - 1)
	}

	return result * sign
}

/* func myAtoi2(s string) int {
	return convert(clean(s))
}

func clean(s string) (sign int, abs string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}

	switch s[0] { // switch を候補として思い浮かべられるように。
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9': // シングルクオートで比較可能。
		sign, abs = 1, s
	case '+':
		sign, abs = 1, s[1:]
	case '-':
		sign, abs = -1, s[1:]
	default:
		abs = ""
		return
	}

	for i, b := range abs {
		if b < '0' || '9' < b { // 当然 rune の不等式評価も可能
			abs = abs[:i]
			break
		}
	}

	return
}

func convert(sign int, absStr string) int {
	abs := 0

	for _, b := range absStr {
		abs = abs*10 + int(b-'0') // これで擬似的に atoi が可能。
		switch {
		case sign == 1 && abs > math.MaxInt32: // `math.MaxInt32`、便利
			return math.MaxInt32
		case sign == -1 && sign*abs < math.MaxInt32:
			return math.MinInt32
		}
	}

	return sign * abs
}
 */