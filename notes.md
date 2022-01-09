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
