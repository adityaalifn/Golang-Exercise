package main

import ("fmt")

//func foo()  {
//	fmt.Println("The square root of 4 is", math.Sqrt(4))
//	//foo()
//}

//func add(x float64, y float64) float64 {
//	return x+y
//}
//
//func multiple(a, b string) (string,string) {
//	return a,b
//}

func main() {
	x := 15
	a := &x
	fmt.Println(a)
	fmt.Println(*a)

	*a = 5
	fmt.Println(x)

	*a = *a**a
	fmt.Println(x)
	fmt.Println(*a)
	//fmt.Println("A number from 1-100", rand.Intn(100))

	/*var num1,num2 float64 = 5.6,2.7
	var num2 float64 = 2.7*/

//	num1,num2  := 5.6,2.7
//	fmt.Println(add(num1,num2))
//
//	w1, w2 := "hey", "there"
//	fmt.Println(multiple(w1,w2))
//
//	var a int = 62
//	var b float64 = float64(a)
//	fmt.Println(b)
}