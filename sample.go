package main

import (
	"fmt"
)

func sample() {
	fmt.Println("This is testing the code")
	// declaring

	// var i int
	// i = 42
	// fmt.Println(" i number :", i)

	// var f float32 = 3.2 //
	// fmt.Println(" f number :", f)

	// firstName := "Rakshit"
	// fmt.Println(" first name :", firstName)

	// isTrue := true //
	// fmt.Println(" is true :", isTrue)

	// c := complex(3, 4) // complex numbers
	// fmt.Println(" c number :", c)

	// realN, imgN := real(c), imag(c)
	// fmt.Println(" real : ", realN)
	// fmt.Println(" imgN : ", imgN)

	// var firstName *string                    // pointer to the string value
	// fmt.Println(" first name : ", firstName) // Nil
	// //firstName = "Rakshit" // error : can't assign string to pointer

	// //  (*) derefrences the pointer to store the data
	// /* *firstName = "Rakshit" // runtime error: invalid memory address or nil pointer dereference */
	// // We have NIL value for firstName pointer so can't assign to NIL Pointer

	// var lastName *string = new(string)
	// fmt.Println("last name : ", lastName) // 0xc000188040 memory addres

	// /* *lastName = "Rakshit"     */        // assigns value to the pointer memory
	// fmt.Println("last name : ", lastName)  // 0xc000188040 memory of same pointer
	// fmt.Println("last name : ", *lastName) // actual value stored at the memory

	// name := "Rakshit"
	// namePtr := &name // fetches the memory address and assign to pointer
	// // & is the addressOf operator in golang
	// fmt.Println("namePtr :", *namePtr)
	// name = "New Name "
	// fmt.Println("namePtr :", *namePtr) // prints new name as ptr is pointing to memory

	// const pi = 3 // needs to be available at compile time
	// // so const can't be assigned value from functions
	// fmt.Println(pi)

	// // implicit conversion since no type of PI specified
	// fmt.Println(pi + 3)    // int addition
	// fmt.Println(pi + 3.14) // float addition

	// const PI int = 3
	// fmt.Println(PI + 3)             // int addition (Works)
	// fmt.Println(PI + 3.14)          // float addition. Error : constant 3.14 truncated to integer
	// fmt.Println(float32(PI) + 3.14) // works

	// // array of size of 3 with INT element
	// var arr [3]int
	// arr[0] = 1
	// arr[1] = 2
	// arr[2] = 3
	// fmt.Println(arr)

	// brr := [3]int{1, 2, 3}
	// fmt.Println(brr)

	// arr := [3]int{1, 2, 3}
	// slice := arr[:]    // begining to end
	// fmt.Println(slice) // [1 2 3]

	// // slice is kind of ptr, pointing to underlying array
	// arr[1] = 14
	// slice[2] = 27
	// fmt.Println(slice) //[1 14 27]

	// slice1 := []int{1, 2, 3} // compiler manages underlying array
	// fmt.Println(slice1)      // [1 2 3]
	// slice1 = append(slice1, 45)
	// slice1 = append(slice1, 6, 7, 8) //[1 2 3 45 6 7 8]
	// fmt.Println(slice1)

	// // [lower, upper]  upper is excluded
	// slice2 := slice1[1:3] // from idx 1 to 2, [2 3]
	// slice3 := slice1[1:]  // start at idx 1 , [2 3 45 6 7 8]
	// slice4 := slice1[:2]  // end at idx 1, [1 2]

	// fmt.Println(slice2, slice3, slice4) // [2 3] [2 3 45 6 7 8] [1 2]

	// // key -> String, value -> Int
	// m := map[string]int{"foo": 42}
	// fmt.Println(m)        // map[foo:42]
	// fmt.Println(m["foo"]) // 42
	// m["foo"] = 32
	// m["food"] = 1
	// fmt.Println(m) // map[foo:32 food:1]
	// delete(m, "food")
	// fmt.Println(m) // map[foo:32]

	// structs
	// heterogenous data structure with compile time fixed fields

	// type user struct {
	// 	firstName string
	// 	ID        int
	// 	lastName  string
	// }

	// var u user
	// fmt.Println(u) // { 0 } Initialized for zero value of all data types

	// u.ID = 1
	// u.firstName = "Rakshit"
	// u.lastName = "Kathawate"
	// fmt.Println(u) // {Rakshit 1 Kathawate}

	// u1 := user{
	// 	ID:        1,
	// 	firstName: "FirstName",
	// 	lastName:  "lastName", // needs , always
	// }
	// fmt.Println(u1) // {FirstName 1 lastName}

	// loops
	// all loops are for loop

	// loop with condition
	// var i int
	// for i < 5 {
	// 	println(i)
	// 	i++
	// 	if i == 3 {
	// 		break
	// 	}
	// }

	// // classic syantax
	// for j := 0; j < 5; j++ {
	// 	println(j)
	// }

	// infinite loop
	// var k int
	// for {
	// 	println(k)
	// }

	//traverse loop

	// slices := []int{1, 2, 3}
	// for j := 0; j < len(slices); j++ {
	// 	println(slices[j])
	// }

	// // slices := []int{1, 2, 3}
	// for i, v := range slices {
	// 	println(i, v)
	// }

	// maps := map[string]int{"1": 1, "2": 2, "3": 3}
	// for i, v := range maps {
	// 	println(i, v)
	// }

	// //just the keys

	// for k := range maps {
	// 	println(k) // ThreeOneTwo
	// }

	// // just values
	// for _, v := range maps {
	// 	println(v)
	// }

	// // panic : app can't proceed so we use panic
	// panic("Panic hit")
	// /*
	// 	panic: Panic hit

	// 	goroutine 1 [running]:
	// 	main.main()
	// 		/Users/rakshitkathawate/Desktop/PluralSight/Go-Course/Go_Revision/sample.go:185 +0x95
	// 	exit status 2
	// */

	// id1 := 1
	// id2 := 2

	// if id1 == id2 {
	// 	fmt.Println("IF")
	// } else if id1 == id2+1 {
	// 	fmt.Println("ELSE IF")
	// } else {
	// 	fmt.Println("ELSE")
	// }

}

// // func with params of different data types
// func sampleFunc(arg1 string, arg2 int) {
// 	fmt.Println(arg1, arg2)
// }

// // func with params of same data type
// func sampleSameData(arg1, arg2 string) {
// 	fmt.Println(arg1, arg2)
// }

// func errorFunc(arg1 string) error {
// 	fmt.Println(arg1)
// 	return errors.New("Hello error")
// }

// // multiple return params
// func multipleErrorFunc(arg1 string) (string, error) {
// 	fmt.Println(arg1)
// 	if 1 == 2 {
// 		return arg1, nil
// 	}
// 	return "", errors.New("Hello error")
// }
