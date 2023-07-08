# LEARN_BASICS_UNIT_TESTING
Learn basics of unit testing in golang

//Redirect to the test folder
cd test_folder
//Run tests for Go packages. It automatically identifies test files based on the naming convention, where test files end with _test.go. 
go test
// Adding the -v flag to the go test command enables verbose output. This means that in addition to the test results, it provides more detailed information about the individual tests being run. 
go test -v
//Along with the -v flag, adding the --cover flag enables code coverage analysis during testing. Code coverage measures how much of your code is being exercised by your tests. 
go test -v --cover
