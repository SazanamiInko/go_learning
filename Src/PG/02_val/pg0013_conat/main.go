// ///////////////////
// /
// / 定数
// /
// ////////////////////
package main

//メイン関数
func main() {
	const tax_rate = 0.10
	var price = 1000.00
	var in_tax = price * (1.00 + tax_rate)
	println("税込み価格は", in_tax)
}
