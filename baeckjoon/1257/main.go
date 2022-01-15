package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solve(money uint64, coins []int) (cnt uint64) {
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	for _, coin := range coins {
		div, mod := money/uint64(coin), money%uint64(coin)
		cnt += div
		if mod != 0 {
			money = mod
		} else {
			break
		}
	}
	return
}

func main() {
	var moneyStr, cntStr, coinLineStr string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fscan(reader, &moneyStr, &cntStr, &coinLineStr)
	money, _ := strconv.ParseUint(moneyStr, 10, 64)
	coinsStr := strings.Fields(coinLineStr)
	coins := []int{}
	for _, coinStr := range coinsStr {
		coin, _ := strconv.Atoi(coinStr)
		coins = append(coins, coin)
	}
	fmt.Fprintln(writer, Solve(money, coins))
}
