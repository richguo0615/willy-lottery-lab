package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	NONE_AWARD = iota
	HEAD_AWARD
	SECOND_AWARD
	THREE_AWARD
	FOUR_AWARD
)

var lotteryCount int64

var winRecords []*WinRecord

type WinRecord struct {
	Index  int64
	S1Nums []int
	S2Num  int
}

func main() {
	winS1 := map[int]bool{7: true, 8: true, 15: true, 19: true, 34: true, 38: true}
	winS2 := 8
	getOneLottery(winS1, winS2)
	fmt.Print("Done!")
}

func getOneLottery(winS1 map[int]bool, winS2 int) {

	lotteryCount++

	s1Nums := shuffle(newSection1Nums())
	s2Nums := shuffle(newSection2Nums())

	var myS1Nums []int
	for i := 0; i < 6; i++ {
		myS1Nums = append(myS1Nums, pickOneNum(s1Nums))
	}
	myS2Num := pickOneNum(s2Nums)

	mapCount := 0
	for _, num := range myS1Nums {
		if winS1[num] {
			mapCount++
		}
	}

	award := NONE_AWARD
	if mapCount >= 5 {
		if mapCount == 6 && myS2Num == winS2 {
			award = HEAD_AWARD
		} else if mapCount == 6 {
			award = SECOND_AWARD
		} else if mapCount == 5 && myS2Num == winS2 {
			award = THREE_AWARD
		} else if mapCount == 5 {
			award = FOUR_AWARD
		}

		record := &WinRecord{
			Index:  lotteryCount,
			S1Nums: myS1Nums,
			S2Num:  myS2Num,
		}
		winRecords = append(winRecords, record)

		fmt.Println("(Win!) lottery - award: ", award ,", index: ", record.Index, ", myS1Nums: ", record.S1Nums, ", myS2Num: ", record.S2Num)
	}

	if award != HEAD_AWARD {
		fmt.Println("lottery - index: ", lotteryCount, ", myS1Nums: ", myS1Nums, ", myS2Num: ", myS2Num)
		time.Sleep(50 * time.Millisecond)
		getOneLottery(winS1, winS2)
	}
}

func pickOneNum(nums []int) (aNum int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randIndex := r.Intn(len(nums))
	aNum = nums[randIndex]
	nums = append(nums[:randIndex], nums[randIndex+1:]...)
	return
}

func shuffle(vals []int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := len(vals)
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		randIndex := r.Intn(len(vals))
		ret[i] = vals[randIndex]
		vals = append(vals[:randIndex], vals[randIndex+1:]...)
	}
	return ret
}

func newSection1Nums() (nums []int) {
	for i := 1; i < 39; i++ {
		nums = append(nums, i)
	}
	return
}

func newSection2Nums() (nums []int) {
	for i := 1; i < 9; i++ {
		nums = append(nums, i)
	}
	return
}
