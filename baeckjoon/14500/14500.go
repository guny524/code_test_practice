package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coor [2]int

func (c *Coor) isOutRange() bool {
	for i := 0; i < 2; i++ {
		if c[i] < 0 || c[i] >= size[i] {
			return true
		}
	}
	return false
}

func (c *Coor) isCollide(picked []Coor) bool {
	for _, p := range picked {
		if c[0] == p[0] && c[1] == p[1] {
			return true
		}
	}
	return false
}

func CalMax(arr *[][]int, picked []Coor, last Coor) (max int) {
	switch len(picked) {
	case 4:
		for _, p := range picked {
			max += (*arr)[p[0]][p[1]]
		}
		return max
	case 3:
		flag := []bool{true, true}
		sum := []int{0, 0}
		index := 0
		for i := 0; i < 2; i++ {
			tmp := picked[0][i]
			for j := 0; j < 3; j++ {
				sum[i] += picked[j][i]
				if tmp != picked[j][i] {
					flag[i] = false
					index = 1 - i
				}
			}
			sum[i] /= 3
		}
		if flag[index] {
			for _, add := range []int{-1, 1} {
				newCoor := Coor{sum[0] + (1-index)*add, sum[1] + index*add}
				if !(newCoor.isOutRange() || newCoor.isCollide(picked)) {
					ret := CalMax(arr, append(picked, newCoor), newCoor)
					if ret > max {
						max = ret
					}
				}
			}
		}
	}
	for _, c := range [][2]int{Coor{-1, 0}, Coor{1, 0}, Coor{0, -1}, Coor{0, 1}} {
		newCoor := Coor{last[0] + c[0], last[1] + c[1]}
		if !(newCoor.isOutRange() || newCoor.isCollide(picked)) {
			ret := CalMax(arr, append(picked, newCoor), newCoor)
			if ret > max {
				max = ret
			}
		}
	}
	return max
}

var size Coor

func Solve(arr [][]int, col, row int) (max int) {
	size = Coor{col, row}
	for i := 0; i < size[0]; i++ {
		for j := 0; j < size[1]; j++ {
			newCoor := Coor{i, j}
			ret := CalMax(&arr, []Coor{newCoor}, newCoor)
			if ret > max {
				max = ret
			}
		}
	}
	return max
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	bytes, _, _ := reader.ReadLine()
	line := string(bytes)
	nums := strings.Fields(line)
	col, _ := strconv.Atoi(nums[0])
	row, _ := strconv.Atoi(nums[1])
	arr := make([][]int, col)
	for i := 0; i < col; i++ {
		arr[i] = make([]int, row)
		bytes, _, _ = reader.ReadLine()
		line = string(bytes)
		nums = strings.Fields(line)
		for j := 0; j < row; j++ {
			num, _ := strconv.Atoi(nums[j])
			arr[i][j] = num
		}
	}
	max := Solve(arr, col, row)
	fmt.Println(max)
}
