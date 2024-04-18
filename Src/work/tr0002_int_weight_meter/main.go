// ///////////////////
// /
// / int体重計
// /
// ////////////////////
package main

import (
	"fmt"
	"math"
)

// 生徒構造体
type TStudent struct {
	name        string
	idealWeight int
	weight      float32
	records     []TRecord
}

// 記録用紙
type TRecord struct {
	weight  float32
	isIdeal bool
	tooFat  bool
	paper   TWeightPaper
}

// 体重計の示す体重
type TWeightPaper struct {
	displayWeight int32
}

// メイン関数
func main() {
	//A君
	var studenA = CreateStudent("A", 63)
	//B君
	var studenB = CreateStudent("B", 53)
	//C君
	var studenC = CreateStudent("C", 60)
	//D君
	var studenD = CreateStudent("D", 30)

	var students = [4]*TStudent{&studenA, &studenB, &studenC, &studenD}

	for _, student := range students {
		recordingWeight(student)
		fatStudent(student)

	}
	for _, student := range students {
		fmt.Print(student.name, student.weight, student.records[0].paper.displayWeight)
		fmt.Println("")

	}
}

// 生徒を作る
func CreateStudent(name string, idealWeight int) TStudent {
	var student TStudent
	student.name = name
	student.idealWeight = idealWeight
	student.weight = float32(idealWeight - 2)
	return student
}

// 生徒を太らせる
func fatStudent(student *TStudent) {
	*&student.weight += 0.1
}

// 体重を記録する
func recordingWeight(student *TStudent) {
	var record TRecord
	record.weight = student.weight
	record.paper = takeHelthMeter(student)
	var s = student.records
	append(s, record)

}

// 体重計に乗る
func takeHelthMeter(student *TStudent) TWeightPaper {
	var paper TWeightPaper
	paper.displayWeight = int32(math.Round(float64(student.weight/10)) * 10)
	return paper
}
