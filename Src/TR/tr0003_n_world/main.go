// ///////////////////
// /
// /n進法
// /
// ////////////////////
package main

import (
	"fmt"
	"math"
	"strconv"
)

// メイン関数
func main() {

	const LEFT_CONDITION = "12"
	const RIGHT_CONDITION = "1331"

	fmt.Println("この設問には10進法以上の取り扱いの説明がないのでnは10以下と推測できる")
	fmt.Println("2のn乗は偶数であるが奇数となっている。これからnは奇数だと考えられる")
	fmt.Println("n=5とすると、下一桁が6のときしか満たさない。2の累乗が6になるのは4の倍数を累乗したときになるとき")
	fmt.Println("命題の累乗する数字12は奇数進法では奇数なので満たさないので,n=5はありえない。")

	//対象として絞ったもの7,9
	var targets = [2]int64{7, 9}

	for _, world := range targets {
		var right int64
		var left int64 = 1
		var power = getWorldNum(LEFT_CONDITION, world)
		left = left << power

		right = getWorldNum(RIGHT_CONDITION, world)

		fmt.Printf("左 %d 右 %d", left, right)
		fmt.Println("")
		if right == left {
			fmt.Printf("%d進法が条件を満たします", world)
		} else {
			fmt.Printf("%d進法が条件を満たしません", world)
		}
		fmt.Println("")
	}

}

// N進法から10進法への変換
func getWorldNum(num string, world int64) int64 {

	var keta = len(num)
	var ret int64 = 0

	for k := 0; k < keta; k++ {
		var pick = pickNum(num, int64(k))
		var p = math.Pow(float64(world), float64(keta-k-1))
		ret += int64(p) * pick
	}
	return ret
}

// /ｎ文字目を数値にして返す
func pickNum(num string, pos int64) int64 {
	var target = string(num[pos])
	var value int
	value, _ = strconv.Atoi(target)

	return int64(value)
}
