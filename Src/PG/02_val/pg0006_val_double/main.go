// ///////////////////
// /
// / 少数型変数
// /
// ////////////////////
package main

import "fmt"

//メイン関数
func main() {

	var ans string
	var res string = "オウム"
	var age int32 = 5
	var weight float32 = 150.2

	fmt.Println("あなたは誰")
	fmt.Scan(&ans)
	fmt.Print("僕は")
	fmt.Print(res)
	fmt.Print("です")
	fmt.Print(age)
	fmt.Println("歳です")
	fmt.Print(weight)
	fmt.Print("gです")

}
