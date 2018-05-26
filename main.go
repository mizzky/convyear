package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)


func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	year, _ := strconv.Atoi(text)

	y2c(year)
}

func y2c(year int) {
	if year > 1988 {
		fmt.Println("平成：", (year-1988))
	}else if year > 1925 {
		fmt.Println("昭和：", (year-1925))
	}else if year > 1911 {
		fmt.Println("大正：", (year-1911))		
	}else if year > 1867 {
		fmt.Println("明治：", (year-1867))		
	}
}
