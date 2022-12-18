## Go Revision

### Go characteristics 
* Fully Compiled Language 
* Fast Compilation than C++
* Strongly Typed Language 
* Concurrent (context switching) by default
* Inbuilt Garbage Collector

### Starting the project 
``` go mod init {{projectModuleName}}```

**projectModuleName** is usually github repo so that new packages can be downloaded from source when someone fetches the repo.  
This initializes the go.mod file with the import and version

```
  ===== main.go
  package main

  import "fmt"

  func main() {
  fmt.Println("This is testing the code")
  }
  ===== go.mod
  module github.com/ChampionTej05/GoWebapp

  go 1.15
  ```

### Primitive datatypes
Different ways of declaring vars in golang 
```go

  var username string = "this is string"
  fmt.Printf("Type of Username %T", username)

  var i int
  i = 42
  fmt.Println(" i number :", i)

  var f float32 = 3.2 //
  fmt.Println(" f number :", f)

  firstName := "Rakshit"
  fmt.Println(" first name :", firstName)
  // implicit type declaration is only allowed inside method. we can't have global vars declared as implicit

  isTrue := true //
  fmt.Println(" is true :", isTrue)

  var isFalse bool = false

  c := complex(3, 4) // complex numbers
  fmt.Println(" c number :", c)

  realN, imgN := real(c), imag(c)
  fmt.Println(" real : ", realN)
  fmt.Println(" imgN : ", imgN)
  ```

### Taking Inputs from the user 
We can search for any package in the go-lang on https://pkg.go.dev/.
We will be using *bufio* for taking user-input here 
```go
  import (
	"bufio"
	"fmt"
	"os"
  )

  func main() {

    var welcomeMessage string = "Welcome User"
    fmt.Println(welcomeMessage)

    // initialize the reader to read strings
    var reader = bufio.NewReader(os.Stdin)
    fmt.Printf("Enter Username: ")

    // read from stdin till "new line(/n)"
    input, err := reader.ReadString('\n')
    fmt.Println("Error if any", err)
    fmt.Printf("Welcome user : %v", input)
  }
```


### Conversion of strings to numbers 

```go
  import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
  )

  func main() {

    fmt.Println("Please rate the Pizza Rating")

    reader := bufio.NewReader(os.Stdin)

    input, _ := reader.ReadString('\n')
    fmt.Printf("Input given by User %v", input)
    fmt.Printf("Type of the input : %T", input)

    // convert strings

    numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
    if err != nil {
      fmt.Println("err ", err)
    } else {
      fmt.Println("New Rating ", numRating+5)
    }

  }
```

### Pointers 

 - (*) - Dereference opertor in golang (Assign/get value)
 - (&) - addressOf operator in golang (get the memory address)
 **Go doesn't support pointer arithmetic** 
```go
  var firstName *string                    // pointer to the string value
  fmt.Println(" first name : ", firstName) // Nil
  //firstName = "Rakshit" // error : can't assign string to pointer

  //  (*) derefrences the pointer to store the data
  *firstName = "Rakshit" // runtime error: invalid memory address or nil pointer dereference 
  // We have NIL value for firstName pointer so can't assign to NIL Pointer

  var lastName *string = new(string)
  fmt.Println("last name : ", lastName) // 0xc000188040 memory addres

  *lastName = "Rakshit"                  // assigns value to the pointer memory
  fmt.Println("last name : ", lastName)  // 0xc000188040 memory of same pointer
  fmt.Println("last name : ", *lastName) // actual value stored at the memory

  name := "Rakshit"
  namePtr := &name // fetches the memory address and assign to pointer
  // & is the addressOf operator in golang
  fmt.Println("namePtr :", *namePtr)

  name = "New Name "
  fmt.Println("namePtr :", *namePtr) // prints new name as ptr is pointing to memory
  ```
 

### Constants 
 - Needs to evaluated at COMPILE Time
 - Assign value during declaration only 
 - Can't assign value returned from function since it is runtime 

  #### Function level const 
 ```go
  const pi = 3 // needs to be available at compile time
  // so const can't be assigned value from functions
  fmt.Println(pi)

  // implicit conversion since no type of PI specified
  fmt.Println(pi + 3)    // int addition
  fmt.Println(pi + 3.14) // float addition

  const PI int = 3
  fmt.Println(PI + 3)             // int addition (Works)
  fmt.Println(PI + 3.14)          // float addition. Error : constant 3.14 truncated to integer
  fmt.Println(float32(PI) + 3.14) // works
  ```

  #### Package level const 
  ```go
  package main

  import "fmt"

  const pi = 3.14 // single const
  const (
    first  = 1
    second = "second"
  ) // const block
  func main() {
    fmt.Println(pi) // 3.14
    fmt.Println(first, second) // 1 second
  }
  ```

### Iota 
  - A counter which starts with zero
  - Increases by 1 after each line
  - Is only used with constant 
  - https://golangbyexample.com/iota-in-golang/

  ```go
  const (
    first  = iota // starts with 0
    second = iota
    third  = iota + 2 // constant expression
    fourth = 2 << iota
    fifth  // takes expression defined in earlier line which 2<<
    sixth
    seventh = iota // 6 as iota increases per line and it is line 6 now (start with 0) 
  ) // const block

  const (
    eight = iota // will be reset to 0 out of const block
  )

  func main() {
    fmt.Println(first, second, third, fourth, fifth, sixth, seventh, eight) // 0 1 4 16 32 64 6 0
  }
  ```


### Array 
  Static sized in Golang 
  ```go
    // array of size of 3 with INT element
    var arr [3]int
    arr[0] = 1
    arr[1] = 2
    arr[2] = 3
    fmt.Println(arr)

    brr := [3]int{1, 2, 3}
    fmt.Println(brr)
  ```

### Slices
  Dynamic sized array-like structure, pointing to underlying arrays.
  ```go
    arr := [3]int{1, 2, 3}
    slice := arr[:]    // begining to end
    fmt.Println(slice) // [1 2 3]

    // slice is kind of ptr, pointing to underlying array
    arr[1] = 14
    slice[2] = 27
    fmt.Println(slice) //[1 14 27]

    slice1 := []int{1, 2, 3} // compiler manages underlying array
    fmt.Println(slice1)      // [1 2 3]
    slice1 = append(slice1, 45)
    slice1 = append(slice1, 6, 7, 8) //[1 2 3 45 6 7 8]
    fmt.Println(slice1)

    // [lower, upper]  upper is excluded
    slice2 := slice1[1:3] // from idx 1 to 2, [2 3]
    slice3 := slice1[1:]  // start at idx 1 , [2 3 45 6 7 8]
    slice4 := slice1[:2]  // end at idx 1, [1 2]

    fmt.Println(slice2, slice3, slice4) // [2 3] [2 3 45 6 7 8] [1 2]
  ```
    
### Maps 
  Key Value Pair Data Structure of fixed data type at compile time
  ```go
    // key -> String, value -> Int
    m := map[string]int{"foo": 42}

    fmt.Println(m)        // map[foo:42]
    fmt.Println(m["foo"]) // 42

    m["foo"] = 32
    m["food"] = 1
    fmt.Println(m) // map[foo:32 food:1]

    delete(m, "food")
    fmt.Println(m) // map[foo:32]
  ```

### Structs 
  Heterogenous data structure with compile time fixed fields
```go
    type user struct {
      firstName string
      ID        int
      lastName  string
    }

    var u user
    fmt.Println(u) // { 0 } Initialized for zero value of all data types

    u.ID = 1
    u.firstName = "Rakshit"
    u.lastName = "Kathawate"
    fmt.Println(u) // {Rakshit 1 Kathawate}

    u1 := user{
        ID:        1,
        firstName: "FirstName",
        lastName:  "lastName", // needs [,] always
      }
    fmt.Println(u1) // {FirstName 1 lastName}
  ```
    
### Functions 
```go
    // func with params of different data types
    func sampleFunc(arg1 string, arg2 int) {
      fmt.Println(arg1, arg2)
    }

    // func with params of same data type
    func sampleSameData(arg1, arg2 string) {
      fmt.Println(arg1, arg2)
    }

    func errorFunc(arg1 string) error {
      fmt.Println(arg1)
      return errors.New("Hello error")
    }

    // multiple return params
    func multipleErrorFunc(arg1 string) (string, error) {
      fmt.Println(arg1)
      if 1 == 2 {
        return arg1, nil
      }
      return "", errors.New("Hello error")
    }
  ```

### Methods 
  - Methods are special type of functions which allows to add Object oriented behaviour in golang
  - Every method is tied to some or other object through which it can be called
```go 
    // method is tied to the controller
    func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
      w.Write([]byte("Hello Rakshit"))
    }

    // method call 
    uc.ServeHTTP(w,r)
```


### Loops
  - All loops in go are of the for type only 
  ```go

    // loop with condition
    var i int
    for i < 5 {
      println(i)
      i++
      if i == 3 {
        break
      }
    }

    // classic syantax
    for j := 0; j < 5; j++ {
      println(j)
    }

    // infinite loop
    var k int
    for {
      println(k)
    }


    //traverse loop
    slices := []int{1, 2, 3}
    for j := 0; j < len(slices); j++ {
      println(slices[j])
    }

    for i, v := range slices {
      println(i, v)
    }

    maps := map[string]int{"1": 1, "2": 2, "3": 3}
    for i, v := range maps {
      println(i, v)
    }

    
    //just the keys
    for k := range maps {
      println(k) // ThreeOneTwo
    }

    // just values
    for _, v := range maps {
      println(v)
    }
  ```

### Panic
  - Panic in golang is way to tell application that it has entered a state from where it can't proceed further
  - It sends high priority signal to app and adds verbose regarding the same 
  
  ```go

    // panic : app can't proceed so we use panic
    panic("Panic hit")
    /*
      panic: Panic hit

      goroutine 1 [running]:
      main.main()
        /Users/rakshitkathawate/Desktop/PluralSight/Go-Course/Go_Revision/sample.go:185 +0x95
      exit status 2
    */
  ```

### Branching Construct
  - Branching construct of golang is same as any other language
  ```go 
    id1 := 1
    id2 := 2

    if id1 == id2 {
      fmt.Println("IF")
    } else if id1 == id2+1 {
      fmt.Println("ELSE IF")
    } else {
      fmt.Println("ELSE")
    }
  ```

### Switches 
  - ":" define the case block in switch 
  ```go 

    r := HTTPRequest{Method: "POST"}

    switch r.Method {
      case "GET":
        fmt.Println("Get request")
      case "POST":
        fmt.Println("POST request")
      default:
        fmt.Println("Get request")

    }
  ```

