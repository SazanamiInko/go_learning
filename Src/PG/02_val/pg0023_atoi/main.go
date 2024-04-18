// ///////////////////
// /
// / 数値変換
// /
// ////////////////////
package main

import (
	"fmt"
	"strconv"
)

// メイン関数
func main() {
	const tax_rate = 0.10
	var ans string

	fmt.Println("パンはいくらですか?")
	fmt.Scan(&ans)

	var price, _ = strconv.Atoi(ans)

	var in_tax = int32(float64(price)*(1.00+tax_rate) + 0.5)
	fmt.Print("パンの税込み価格は")
	fmt.Print(in_tax)
	fmt.Print("円")
}
