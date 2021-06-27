package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func SeperateValue(val int) [3]int {
	return [3]int{val / 100, (val / 10) % 10, val % 10}
}

func IntInSlice(a int, nums [3]int) bool {
	for _, v := range nums {
		if v == a {
			return true
		}
	}
	return false
}

func main() {
	rand.Seed(time.Now().UnixNano())
	nums := [3]int{}

	nums[0] = rand.Intn(9) + 1
	i := rand.Intn(9) + 1
	for {
		if i != nums[0] {
			nums[1] = i
			break
		}
		i = rand.Intn(9) + 1
	}
	for {
		if i != nums[0] && i != nums[1] {
			nums[2] = i
			break
		}
		i = rand.Intn(9) + 1
	}
	fmt.Println(nums)

	cnt := 1
	for {
		fmt.Printf("숫자를 입력해주세요 : ")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else {
			fmt.Println("입력하신 숫자는 ", n, " 입니다.")

			sep_val := SeperateValue(n)
			fmt.Println("입력한 숫자는 ", sep_val)
			strike := 0
			ball := 0
			for j := 0; j < 3; j++ {
				if nums[j] == sep_val[j] {
					strike++
				} else {
					if IntInSlice(sep_val[j], nums) {
						ball++
					}
				}
			}
			fmt.Printf("스트라이크 %d 볼 %d", strike, ball)
			if strike == 3 {
				fmt.Println("3개의 숫자를 모두 맞히셨습니다! 게임 종료")
				break
			}
			cnt++
		}

	}
}
