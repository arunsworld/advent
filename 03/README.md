# Showcases

* Table Driven Testing in `Test_A_Bigger_Forest/is_able_to_count_trees_for_multiple_slopes`
    * Multiple use-cases are covered within the same test
* Domain specific types
    * Location and Slope are sets of 2 integers, but rather than calling them integers or a generic data structure
    * we're giving it specific names by defining types
* Functional Decomposition in `CountTrees`
    * Helps readability significantly by abstracting away the next level details into small functions
    * Allows for better composition
* AND YET - notice the unit tests tests CountTrees and not `IsLocationInsideForest` or `Relocate`
    * Unit testing isn't about testing the smallest unit
    * Instead, about testing the "public API" of the functionality/module
    * This avoids coupling of tests with code and facilitates refactoring
    * Focuses the tests on what we're trying to prove rather than how
* Implementing a set using map
    * `Trees` in a `Forest` identify the `Location` where a tree exists
    * This is therefore a set of locations where a tree exists