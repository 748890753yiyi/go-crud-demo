package main

import (
	"time"
	"strconv"
	"fmt"
)

func main() {

	var mapmou = map[string]int{"January": 1, "february": 2, "March": 3, "April": 4, "May": 5, "June": 6, "July": 7, "August": 8, "September": 9, "October": 10, "November": 11, "December": 12}
	idCard := "140603199010155430" //身份证
	now := time.Now()
	now_year := now.Year()                 // 年
	now_mo := mapmou[now.Month().String()] // 月
	now_day := now.Day()                   // 日
	fmt.Println(now_year, now_mo, now_day)
	idcard_year, _ := strconv.Atoi(Substr(idCard, 6, 4)) // 年
	idcard_mo, _ := strconv.Atoi(Substr(idCard, 10, 2))  // 月
	idcard_day, _ := strconv.Atoi(Substr(idCard, 12, 2)) // 日
	fmt.Println(idcard_year, idcard_mo, idcard_day)
	fmt.Println("idCard:" + idCard)
	age := now_year - idcard_year // 如果计算虚岁需这样：age := now_year - idcard_year+1
	if now_year < idcard_year {
		age = 0
	} else {
		if now_mo < idcard_mo {
			age = age - 1
		} else {
			if now_day < idcard_day {
				age = age - 1
			}
		}
	}
	fmt.Println("age:", age)
}

func Substr(str string, start, length int) string {

	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}