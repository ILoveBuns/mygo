package demo

import (
	"bou.ke/monkey"
	"fmt"
	"log"
	"testing"
)

//func TestQueryMain(t *testing.T) {
//	cours := queryMain()
//	fmt.Println(cours)
//}

func TestQueryMain_2(t *testing.T) {
	guard := monkey.Patch(process2, func() []course {
		cour := make([]course,0,5)
		cour = append(cour, course{6,"政治","6等奖"})
		return cour
	})
	defer guard.Unpatch()
	a := process2()
	fmt.Println(a)
}

func TestQueryMain_1(t *testing.T) {
	guard := monkey.Patch(process1, func() []course {
		fmt.Println("t")
		cour := make([]course,0,5)
		cour = append(cour, course{6,"政治","6等奖"})
		return cour
	})
	defer guard.Unpatch()
	a := process1()

}


//func TestProcessMain_1(t *testing.T) {
//	guard := monkey.Patch(process1, func() []course {
//		cour := make([]course,0,5)
//		cour = append(cour, course{6,"政治","6等奖"})
//		return cour
//	})
//	defer guard.Unpatch()
//
//	a := process1()
//	fmt.Println(a)
//}