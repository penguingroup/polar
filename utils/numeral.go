package utils

import (
	"fmt"
	"github.com/karldoenitz/Tigo/logger"
	"strconv"
)

// Decimal 将浮点数进行格式化
func Decimal(value float64, accuracy ...int) float64 {
	ac := 1
	if len(accuracy) > 0 {
		ac = accuracy[0]
	}
	formatStr := fmt.Sprintf("%%.%df", ac)
	var err error
	value, err = strconv.ParseFloat(fmt.Sprintf(formatStr, value), 64)
	if err != nil {
		logger.Error.Printf("format number failed: %s", value, err.Error())
	}
	return value
}

// TimeCompare 比较两个时间的长短，比如0.5h和1.2h
func TimeCompare(num1, num2 string) int {
	if len(num1) < 1 || len(num2) < 1 {
		return 0
	}
	num1 = num1[:len(num1)-1]
	num2 = num2[:len(num2)-1]
	n1, _ := strconv.ParseFloat(num1, 64)
	n2, _ := strconv.ParseFloat(num2, 64)
	if n1 < n2 {
		return -1
	} else if n1 > n2 {
		return 1
	} else {
		return 0
	}
}

//获取成绩等级
const (
	Excellent   = 1
	Good        = 2
	Qualified   = 3
	Unqualified = 4
)

func GetScoreGrade(score, total float32) int {
	if score*100/total >= 90 {
		return Excellent
	}
	if score*100/total >= 80 {
		return Good
	}
	if score*100/total >= 60 {
		return Qualified
	}
	return Unqualified
}

func ConvertSecondToHour(second int, formatStr ...string) string {
	hour := float64(second) / 3600
	if len(formatStr) > 0 {
		return fmt.Sprintf(formatStr[0], hour)
	}
	return fmt.Sprintf("%.2fh", hour)
}
