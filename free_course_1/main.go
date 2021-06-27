package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

const l = 3

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func SeperateValue(val int) [l]int {
	return [l]int{val / 100, (val / 10) % 10, val % 10}
}

func IntInSlice(a int, nums *[l]int) bool {
	for _, v := range nums {
		if v == a {
			return true
		}
	}
	return false
}

func RandomIntValue(min, max int) int {
	return rand.Intn(max-min) + min
}

func CheckDuplicate(idx int, nums *[3]int) bool {
	r := RandomIntValue(1, 9)
	for {
		if IntInSlice(r, nums) {
			r = RandomIntValue(1, 9)
		} else {
			nums[idx] = r
			break
		}
	}
	return true
}

func InitializeGame() [l]int {
	rand.Seed(time.Now().UnixNano())
	nums := [l]int{RandomIntValue(1, 9)}

	for i := 1; i < len(nums); i++ {
		CheckDuplicate(i, &nums)
	}
	return nums
}

func CheckScore(sep_val [l]int, nums *[l]int) (int, int) {
	strike := 0
	ball := 0
	for j := 0; j < 3; j++ {
		if nums[j] == sep_val[j] {
			strike++
		} else if IntInSlice(sep_val[j], nums) {
			ball++
		}
	}
	return strike, ball
}

func main() {

OuterFor:
	for {
		nums := InitializeGame()
		fmt.Println(nums)

		cnt := 1
		for {
			fmt.Printf("숫자를 입력해주세요 : ")
			n, err := InputIntValue()
			if err != nil {
				fmt.Println("숫자만 입력하세요.")
				continue
			}
			fmt.Println("입력하신 숫자는 ", n, " 입니다.")

			sep_val := SeperateValue(n)
			fmt.Println("입력한 숫자는 ", sep_val)

			strike, ball := CheckScore(sep_val, &nums)
			fmt.Printf("스트라이크 %d 볼 %d \n", strike, ball)

			if strike == 3 {
				fmt.Printf("%d개의 숫자를 모두 맞히셨습니다! 게임 종료\n", l)
				fmt.Printf("총 %d번의 시도를 하셨습니다. \n", cnt)

				fmt.Printf("게임을 새로 시작하려면 1, 종료하려면 2를 입력하세요 : ")
				game, _ := InputIntValue()
				if game == 2 {
					break OuterFor
				}
				break
			}
			cnt++
		}
	}
}
