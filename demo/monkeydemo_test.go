package demo

import (
	"bou.ke/monkey"
	"fmt"
	"reflect"
	"testing"
)

//不进行打桩，传统方法测试有返回值
func TestHasReturnMain(t *testing.T) {
	hasReturnMain()
}

//使用monkey进行打桩，测试有返回值
func TestHasReturnMain_1(t *testing.T) {
	guard := monkey.Patch(hasReturn, func() []course {
		cour := make([]course,0,5)
		cour = append(cour, course{6,"政治","6等奖"})

		fmt.Println("this is mock fun!!")
		return cour
	})
	defer guard.Unpatch()
	hasReturnMain()
}


//不进行打桩，传统方法测试无返回值
func TestProcessMain(t *testing.T) {
	processMain()
}

//使用monkey进行打桩，测试无返回值
func TestProcessMain_1(t *testing.T) {
	guard := monkey.Patch(process, func() {
		cour := make([]course,0,5)
		cour = append(cour, course{6,"政治","6等奖"})

		fmt.Println("do other thing")
	})
	defer guard.Unpatch()

	processMain()
}

//不进行打桩，传统方法测试方法
func TestGetCourseDefMain(t *testing.T){
	getCourseDefMain()
}

//使用monkey进行打桩，测试方法
func TestGetCourseDefMain_1(t *testing.T){
	var c *course
	guard := monkey.PatchInstanceMethod(reflect.TypeOf(c), "GetCourseDef", func(_ *course) string {
		fmt.Println("this is mock get!!")

		return "test"
	})
	defer guard.Unpatch()

	getCourseDefMain()
}