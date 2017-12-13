# tesht [![Build Status](https://travis-ci.org/edupo/tesht.svg?branch=master)](https://travis-ci.org/edupo/tesht)
Basic command line that checks the return value of any command and creates a test report (JUNIT for convenience) using the stdout and stderr to create the failed test result.

# Usage
```
tesht <command> <args> <pipes> <whatever> ...
```
This command will generate or add a line to a `results.xml` file in the same folder where `tesht` is called.

# TODO
 - ~~Parsing `result.xml` in~~
 - ~~Adding result to xml structures~~
 - ~~Writting `result.xml` back into disk~~
 - Flags and options ??

# Motivation
Just want to create a bunch of bash unit-tests that translates to JUnits so my Jenkins-Blueocean thing shows them beautifully to everybody. Because bash testing is beautiful!
