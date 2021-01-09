# Showcases

* TDD approach
* Nested test cases for better story telling
* All logic in separate file (pwd_validator.go). Contains no I/O dependencies (eg. os package). Target: 100% test coverage.
    * Note the use of `io.Reader` in `NewPwdValidator` to abstract I/O
* CSV package to parse input data
* Use of regex to parse policy (parser defined as a global variable using `regexp.MustCompile`)
* Use of enums - parameter to identify policy standard
* Effecient conversion of strings to integers using `strconv.Atoi` 
