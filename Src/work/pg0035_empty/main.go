// ///////////////////
// /
// / 標準入力
// /
// ////////////////////
package main

import "fmt"

//メイン関数
func main() {

	var ans string
	fmt.Println("あなたは誰")
	fmt.Scan(&ans)

	if len(ans) == 0 {
		fmt.Println("何か入れてください")
	} else {
		fmt.Printf("{}さんこにちｈわ")
		fmt.Println("")
	}

}
