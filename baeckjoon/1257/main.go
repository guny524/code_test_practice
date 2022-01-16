package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Node struct {
	price uint64 // 현재 까지 선택한 동전 조합의 총 금액
	cnt   uint64 // cnt == 트리의 depth
	bound uint64 // 최소 cnt의 한계치, 항상 minCnt <= bound
}

func GetMaxCoin(coins []uint64) (max uint64) {
	for _, coin := range coins {
		if max < coin {
			max = coin
		}
	}
	return
}

// Solve 멈추는 조건: 금액이 money를 넘어가거나, 현재까지 최소 갯수보다 바운드가 크면 멈춤
func Solve(money uint64, coins []uint64) uint64 {
	maxCoin := GetMaxCoin(coins)
	var minCnt uint64 = math.MaxUint64
	var minBound uint64 = math.MaxUint64
	// getBound 전체 money에서 현재 선택한 조합으로 만들어진 price를 뺀 남은 금액을
	// 가장 큰 동전과 1의 조합으로 메꿨을 때 cnt 값 계산
	// 이 조합 보다 동전 갯수가 많으면 답이 없다 이런 뜻
	getBound := func(price, cnt uint64) (bool, uint64) {
		bound := cnt + (money-price)/maxCoin + (money-price)%maxCoin
		if bound > minBound {
			return true, uint64(0) // 가지치기
		} else {
			minBound = bound
			return false, bound
		}
	}
	getBound(0, 0)
	q := []Node{}
	q = append(q, Node{0, 0, 0})
	for len(q) != 0 {
		// fmt.Println(len(q))
		n := q[0]
		if n.bound > minBound {
			// 큐에 넣을 땐 minBound 보다 작을 수 있지만
			// 다른 곳에서 minBound가 업데이트 될 수 있으니, 큐에서 꺼낼 때도 검사
			q = q[1:]
			continue
		}
		for _, coin := range coins { // 자식 노드들 만들기
			if price := n.price + coin; price > money { // 원하는 금액을 넘어서면
				continue
			} else if cnt := n.cnt + 1; price == money && cnt < minCnt { // 원하는 금액이 되면 cnt 업데이트
				minCnt = cnt
				continue
			} else if prune, bound := getBound(price, cnt); prune || bound > minCnt {
				continue
			} else {
				q = append(q, Node{price, cnt, bound})
			}
		}
		q = q[1:]
	}
	return minCnt
}

func main() {
	var moneyStr, cntStr, coinStr string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fscan(reader, &moneyStr, &cntStr)
	money, _ := strconv.ParseUint(moneyStr, 10, 64)
	cnt, _ := strconv.Atoi(cntStr)
	coins := []uint64{}
	for i := 0; i < cnt; i++ {
		fmt.Fscan(reader, &coinStr)
		coin, _ := strconv.ParseUint(coinStr, 10, 64)
		coins = append(coins, coin)
	}
	fmt.Fprint(writer, Solve(money, coins))
}
