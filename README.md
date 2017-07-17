# Eigen
Eigen is a simple CLI that performs matrix and vector operations, modular arithmetic, and basic arithmetic. 
Eigen is currently a work in progress, so feel free to report any bugs or let me know if you have any suggestions for the program!

## Documentation
### Basic Commands
1. `exit` exits the interface.
2. `clear` clears the terminal window.
3. `clean` removes all variables you have declared so far. 

### Arithmetic
1. Use `+,-,/,*` to perform basic arithmetic in Eigen as you would normally.
  * Eigen supports arithmetic for floating point digits as well as integers
  * Whitespace does not matter: `2+2` is the same as `2 + 2`
  
### Variable assignment
1. Use `=` to assign a variable to a value.
2. Valid value types are: `int, float, matrix`
3. Variable names follow traditional rules:
  * The first character cannot be a number
  * All other characters can be either numbers or letters
  * Variables are case sensitive 
4. `exit, clear, clean` are all reserved key words and cannot be used as variables

### Matrix operations
Coming soon!

### Modular arithmetic operations
Coming soon!

## Install
1. [Install Go](https://golang.org/dl/) (if you have Homebrew on macOS, you can run `brew install go`)
2. Download this repository
3. Run `go run Main.go`
