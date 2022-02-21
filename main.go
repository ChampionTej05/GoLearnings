package main

import "fmt"

func main() {
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
}
