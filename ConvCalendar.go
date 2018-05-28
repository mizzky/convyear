package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	cal := text

	year, _ := strconv.Atoi(text)

	// y2c(year)
	c2y(cal, year)
}

func y2c(year int) {
	if year > 1988 {
		fmt.Println("平成：", (year - 1988))
	} else if year > 1925 {
		fmt.Println("昭和：", (year - 1925))
	} else if year > 1911 {
		fmt.Println("大正：", (year - 1911))
	} else if year > 1867 {
		fmt.Println("明治：", (year - 1867))
	} else {
		fmt.Println("this param is too old..")
	}
}

func c2y(calendar string, num int) {
	if calendar == "平成" && num > 0 {
		fmt.Println(num + 1988)
	}
}
