package demo

import (
	"fmt"
)

type course struct {
	Id 				uint32 `gorm:"column:id;primary_key;AUTO_INCREMENT""`
	Name 			string `gorm:"column:name"`
	Description 	string `gorm:"column:description"`
}

//func queryMain() []course{
//	courses := process()
//	return courses
//}

//func processMain() {
//	fmt.Println(process1())
//}

// 获取一些结构体
func process2() []course{
	cour := make([]course,0,5)
	cour = append(cour, course{1,"数学","一等奖"})
	cour = append(cour, course{2,"英语","二等奖"})

	fmt.Println("query!!")
	return cour
}

// 获取一些结构体
func process1() []course{

	cour := make([]course,0,5)
	cour = append(cour, course{1,"数学","一等奖"})
	cour = append(cour, course{2,"英语","二等奖"})

	//fmt.Println("aaa!!")
	return cour
}

//// 过程性函数，do something
//func process1() []course{
//	//fmt.Println("do something"+ str)
//	cour := make([]course,0,5)
//	cour = append(cour, course{1,"数学","一等奖"})
//	return cour
//}


