# Showcases

* Somewhat advanced use of regular expressions while parsing bag rules
* Recursive algorithms to find contents and containers
    * NOTE: there are currently no checks to avoid circular loops causing stack overflow
    * Implementing this should however be trivial
* A linked-list style implementation with a bag identifying bags within it and the bags it is within
* In a normal linked list an object links to one head and one tail - this instead is a set of heads and tails
* One way would be to use the `bag` as a key to this set (Go map) however, that's not allowed since `bag` is not comparable due to the presence of a map
    * Note: unlike some other languages Go doesn't support a custom equality function to be used in a hashmap
    * However, the implementation is straightforward in building the map using the primary key (`color`)
    * This idea can be used to solve such problems: ie. define a Hashing function and use the output of that as the Key to the Go map