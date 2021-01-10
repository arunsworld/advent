# Showcases

* Seperation of technical concern like parsing data from biz logic by creating a utility `parser.go`
* `bufio.Scanner` to parse text data (see `parser.go`)
    * Note the error handling when using scanner which can be easily missed
    * Remember to check buffer after existing for loop to avoid "off by one" errors
* More domain defined types to improve readability and object orientation
* Power of tests when Part 2 caused a chance in data structure
    * In first iteration group was a map (identical to answers) as it suited the purpose
    * In Part 2 the need to capture the size of group meant a significant change - but the presence of tests made it trivial