// ///////////////////
// /
// / if文
// /
// ////////////////////
package main

import (
	"fmt"
	"strconv"
)

// メイン関数
func main() {
	var ans string
	fmt.Println("数値を入力して")
	fmt.Scan(&ans)
	var res, _ = strconv.Atoi(ans)

	if res%2 == 1 {
		fmt.Println("半が出やした。")
	}

	fmt.Println("イカサマじゃねえか")

}
