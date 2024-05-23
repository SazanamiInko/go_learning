// ///////////////////
// /
// / 標準入力
// /
// ////////////////////
package main

import (
	"fmt"
	"strconv"
)

// メイン関数
func main() {
	itinen_teachers := []string{"西", "南", "カー子", "平"}

	var ans string
	fmt.Println("何組の先生が知りたいですか	？")
	fmt.Scan(&ans)

	var kumi, _ = strconv.Atoi(ans)
	var teacher = itinen_teachers[kumi-1]
	fmt.Printf("%d組の先生は%sです", kumi, teacher)

}
