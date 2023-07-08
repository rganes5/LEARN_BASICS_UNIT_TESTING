# LEARN_BASICS_UNIT_TESTING
Learn basics of unit testing in golang

Redirect to the test folder
cd test_folder
Running tests for Go packages
go test - This command automatically identifies test files based on the naming convention, where test files end with _test.go.
Verbose output during test execution
go test -v - Adding the -v flag enables verbose output, providing more detailed information about the individual tests being run.
Code coverage analysis during testing
go test -v --cover - Adding the --cover flag, along with the -v flag, enables code coverage analysis. Code coverage measures how much of your code is being exercised by your tests.
