package demo

import (
	"fmt"
)


//结构体定义
type course struct {
	Id 				uint32
	Name 			string
	Description 	string
}

/**
 *	调用有返回值的函数
 */
func hasReturnMain() {
	str := hasReturn()
	fmt.Println(str)
}

/**
 *	调用过程性函数
 */
func processMain(){
	process()
}

/**
 *	调用方法
 */
func getCourseDefMain(){
	c := course{1,"数学","一等奖"}
	def := c.GetCourseDef()
	fmt.Println(def)
}


//返回一些结构体
func hasReturn() []course{

	cour := make([]course,0,5)
	cour = append(cour, course{1,"数学","一等奖"})
	cour = append(cour, course{2,"英语","二等奖"})

	fmt.Println("this is origin fun!!")
	return cour
}

//过程性函数，do something
func process() {
	cour := make([]course,0,5)
	cour = append(cour, course{1,"数学","一等奖"})

	fmt.Println("do something")
}

//方法
func (c *course) GetCourseDef() string{
	def := c.Name + ";" + "获得" + c.Description

	fmt.Println("this is origin get!!")
	return def
}
