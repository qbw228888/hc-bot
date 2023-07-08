package api

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func RandN1dN2(n1 int, n2 int) int {
	var sum int = 0
	for i := 0; i < n1; i++ {
		sum += rand.Intn(n2) + 1
	}
	return sum
}

// GetCheckResult 得到检定结果，属性为N，从大成功到大失败分为1,2,3,4,5
func GetCheckResult(n int, groupId string) (int, int) {
	greatSuccess, greatFail := GetGreatSuccessAndGreatFail(groupId)
	rand := RandN1dN2(1, 100)
	if rand > greatFail {
		return 5, rand
	} else if rand > n {
		return 4, rand
	} else if rand > n/2 {
		return 3, rand
	} else if rand > n/5 {
		return 2, rand
	} else if rand > greatSuccess {
		return 1, rand
	} else {
		return 0, rand
	}
}

// GetLongRand model取值0,1,2分别对应大成功，普通，大失败三种情况
func GetLongRand(argStr string, model int) string {
	result := ""
	if argStr == "" {
		result += fmt.Sprintf("1d100=%v/100", RandN1dN2(1, 100))
	} else {
		split := strings.Split(argStr, "+")
		if len(split) < 2 {
			if len(strings.Split(argStr, "d")) == 1 {
				randNum, err := strconv.Atoi(argStr)
				if err != nil {
					return "错误的参数，请重新输入"
				}
				val := RandN1dN2(1, randNum)
				result += fmt.Sprintf("1d%v=%v/%v", randNum, val, randNum)
			} else {
				n1, n2, val, sum := getRdNdN(argStr)
				if n1 == -1 {
					return "错误的参数，请参考help中的正确格式捏"
				}
				result += fmt.Sprintf("%vd%v=%v/%v", n1, n2, val, sum)
			}
		} else if len(split) >= 2 {
			var v int = 0
			var s int = 0
			for i := 0; i < len(split); i++ {
				n1, n2, val, sum := getRdNdN(split[i])
				if n1 == -1 {
					return "错误的参数，请参考help中的正确格式捏"
				} else if n1 == 0 {
					v += val
					s += sum
					result += fmt.Sprintf("%v+", val)
				} else {
					if model == 0 {
						v += 1
						s += sum
					} else if model == 1 {
						v += val
						s += sum
					} else if model == 2 {
						v += sum
						s += sum
					}
					result += fmt.Sprintf("%vd%v+", n1, n2)
				}
			}
			result = strings.TrimRight(result, "+")
			result += fmt.Sprintf("=%v/%v", v, s)
		}
	}
	return result
}

// 返回ndn的随机值和总和
func getRdNdN(arg string) (int, int, int, int) {
	num, err := strconv.Atoi(arg)
	if err == nil {
		return 0, 0, num, num
	}
	args := strings.Split(arg, "d")
	if len(args) < 2 {
		return -1, -1, -1, -1
	}
	n1, err := strconv.Atoi(args[0])
	if err != nil {
		return -1, -1, -1, -1
	}
	n2, err := strconv.Atoi(args[1])
	if err != nil {
		return -1, -1, -1, -1
	}
	val := RandN1dN2(n1, n2)
	sum := n1 * n2
	return n1, n2, val, sum
}
