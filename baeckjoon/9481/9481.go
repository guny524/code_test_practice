package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Confirm(candis [][]bool, ans []int, num, idx, n int) {
	for i := 0; i < n; i++ {
		// 확정된 위치 후보 모두 지우기
		candis[idx][i] = true
		// num 다른 위치들에서 지우기
		candis[i][num-1] = true
	}
	ans[idx] = num
}

// l2r 왼쪽에서 -> 오른쪽 방향으로 커지는가?
func AllConfirm(candis [][]bool, ans []int, l2r bool, n int) {
	num := 1
	switch l2r {
	case true:
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				candis[i][j] = true
			}
			ans[i] = num
			num++
		}
	case false:
		for i := n - 1; i >= 0; i-- {
			for j := 0; j < n; j++ {
				candis[i][j] = true
			}
			ans[i] = num
			num++
		}
	}
}

// 후보배제 clue 1인 경우는 안 들어옴
func RemoveCandi(candi [][]bool, idx, clue, n int) [][]bool {
	for i := 0; i < clue-1; i++ {
		candi[idx][n-1-i] = true
	}
	return candi
}

func isFinish(ans []int) bool {
	for _, a := range ans {
		if a == 0 {
			return false
		}
	}
	return true
}

// clue를 만족하냐
func isFit(ans []int, n, l, r int) bool {
	max := 0
	cnt_l, cnt_r := 0, 0
	for i := 0; i < n; i++ {
		if max < ans[i] {
			max = ans[i]
			cnt_l++
			if max == n {
				break
			}
		}
	}
	max = 0
	for i := n - 1; i >= 0; i-- {
		if max < ans[i] {
			max = ans[i]
			cnt_r++
			if max == n {
				break
			}
		}
	}
	return (l == cnt_l) && (r == cnt_r)
}

func BackTrack(candi [][]bool, ans []int, idx, n, l, r int) (ret [][]int) {
	// // 후보 수가 제일 적은 위치 찾기
	// for i := 0; i < n; i++ {

	// }
	// 확정 안 된 시작위치 찾기
	for i := idx; i < n; i++ {
		if ans[i] == 0 {
			idx = i
			break
		}
	}
	//# 경우의 수가 적은 부분 선택하고, 후보배제 적극 활용해야함
	for j, can := range candi[idx] {
		if !can {
			candi_tmp := make([][]bool, n)
			copy(candi_tmp, candi)
			for i := 0; i < n; i++ {
				candi_tmp[i] = make([]bool, n)
				copy(candi_tmp[i], candi[i])
			}
			ans_tmp := make([]int, n)
			copy(ans_tmp, ans)

			Confirm(candi_tmp, ans_tmp, j+1, idx, n)
			if isFinish(ans_tmp) {
				if isFit(ans_tmp, n, l, r) {
					ret = append(ret, ans_tmp)
				} else {
					return ret
				}
			} else {
				ret = append(ret, BackTrack(candi_tmp, ans_tmp, idx, n, l, r)...)
			}
		}
	}
	return ret
}

func Solve(n, l, r int) int {
	ans := make([]int, n)
	// 1로 set 되면 불가능한 후보, bool default 값이 false니까 모든 후보 가능
	// cadis[pos][candi]
	candiNotPossible := make([][]bool, n)
	for i := 0; i < n; i++ {
		candiNotPossible[i] = make([]bool, n)
	}
	// 확정지을 수 있는 거 확정 짓기
	switch l {
	case 1:
		Confirm(candiNotPossible, ans, n, 0, n)
	case n:
		AllConfirm(candiNotPossible, ans, true, n)
	}
	switch r {
	case 1:
		Confirm(candiNotPossible, ans, n, n-1, n)
	case n:
		AllConfirm(candiNotPossible, ans, false, n)
	}
	if r+l == n+1 {
		Confirm(candiNotPossible, ans, n, l-1, n)
	}
	// 후보배제
	candiNotPossible = RemoveCandi(candiNotPossible, 0, l, n)
	candiNotPossible = RemoveCandi(candiNotPossible, n-1, r, n)
	// 한 칸만 남은 후보 있으면 확정
	for i := 0; i < n; i++ {
		idx := 0
		sum := 0
		for j := 0; j < n; j++ {
			if !candiNotPossible[i][j] {
				sum++
				idx = j
			}
		}
		// 확정 num == idx + 1
		if sum == 1 {
			Confirm(candiNotPossible, ans, idx+1, i, n)
		}
	}

	candi_tmp := make([][]bool, n)
	copy(candi_tmp, candiNotPossible)
	for i := 0; i < n; i++ {
		candi_tmp[i] = make([]bool, n)
		copy(candi_tmp[i], candiNotPossible[i])
	}
	ans_tmp := make([]int, n)
	copy(ans_tmp, ans)

	ret := BackTrack(candi_tmp, ans_tmp, 0, n, l, r)
	// for _, re := range ret {
	// 	fmt.Println(re)
	// }
	return len(ret) % 1000000007
}

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		bytes, _, _ := reader.ReadLine()
		line := string(bytes)
		strs := strings.Fields(line)
		nums := [3]int{}
		for i := 0; i < 3; i++ {
			nums[i], _ = strconv.Atoi(strs[i])
		}
		if nums[0] == nums[1] && nums[1] == nums[2] && nums[2] == 0 {
			break
		}
		fmt.Println(Solve(nums[0], nums[1], nums[2]))
	}
}
