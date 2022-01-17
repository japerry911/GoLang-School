- packages
  - imported in
  - programs start running in `main` package
  - By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.
  - Note: The environment in which these programs are executed is deterministic, so each time you run the example program rand.Intn will return the same number.
  - (To see a different number, seed the number generator; see rand.Seed. Time is constant in the playground, so you will need to use something else as the seed.)
- Exported names
  - In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.
  - pizza and pi do not start with a capital letter, so they are not exported.
  - When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.
- types
  - bool
  - string
  - int int8 int16 int32 int64
  - uint uint8 uint16 uint32 uint64 uintptr
  - byte // alias for uint8
  - rune // alias for int32
    // represents a Unicode code point
  - float32 float64
  - complex64 complex128
- Zero values
  - Variables declared without an explicit initial value are given their zero value.
  - The zero value is:
    - 0 for numeric types,
    - false for the boolean type, and
    - "" (the empty string) for strings.
- Type conversions
  - example
  ```go
  var i int = 42
  var f float64 = float64(i)
  var u uint = uint(f)
  ```
  - more simply put
  ```go
  i := 42
  f := float64(i)
  u := uint(f)
  ```
- constants
  - can't be declared with `:=` syntax
  - example
  ```go
  const value = true
  ```
- for loop
  - three components
    1. init statement, executed before first iteration
    2. condition expression, evaluated before every iteration
    3. post statement, executed at the end of every iteration
    - init and post statements are optional
      - example
      ```go
      for ; sum < 1000; {
        sum += sum
      }
      ```
      - doesn't need semi-colons either, turns into `while` essentially
    - if it's just `for {}` then it will run forever
- Defer

  - A defer statement defers the execution of a function until the surrounding function returns.
  - The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
  - Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

- pointers

  - `var p *int`
    - pointer to integer value, zero-value = `nil`
  - the `*` operator denotes the pointer's underlying value
    ```go
      fmt.Println(*p) // read i through the pointer p
      *p = 21         // set i through the pointer p
    ```
    - known as derefrencing or indirecting

- arrays

  - The type [n]T is an array of n values of type T.
  - var a [10]int
    - declares a variable a as an array of ten integers.
  - An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.

- slices

  - dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.
  - The type []T is a slice with elements of type T.
  - A slice is formed by specifying two indices, a low and high bound, separated by a colon:
    - a[low : high]
      - This selects a half-open range which includes the first element, but excludes the last one.
  - length & capacity
    - length
      - number of elements it contains
      - `len()`
    - capacity
      - the number of elements in the underlying array, counting from the first element in the slice.
      - `cap()`
  - The zero value of a slice is `nil`
  - using `make` command

    - this is how you create dynamically-sized arrays.
    - The make function allocates a zeroed array and returns a slice that refers to that array:
      ```go
        a := make([]int, 5)  // len(a)=5
      ```
    - To specify a capacity, pass a third argument to make:

      ```go
        b := make([]int, 0, 5) // len(b)=0, cap(b)=5

        b = b[:cap(b)] // len(b)=5, cap(b)=5
        b = b[1:]      // len(b)=4, cap(b)=4
      ```

  - `append` command
    - It is common to append new elements to a slice, and so Go provides a built-in append function. The documentation of the built-in package describes append.
      - `func append(s []T, vs ...T) []T`
    - The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.
    - The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
    - If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

- maps
  - A map maps keys to values.
  - The zero value of a map is nil. A nil map has no keys, nor can keys be added.
  - The make function returns a map of the given type, initialized and ready for use.
