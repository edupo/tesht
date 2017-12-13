# tesht [![Build Status](https://travis-ci.org/edupo/tesht.svg?branch=master)](https://travis-ci.org/edupo/tesht)
Basic command line that checks the return value of any command and creates a test report (JUNIT for convenience) using the stdout and stderr to create the failed test result.

# Usage
```
tesht <command> <args> <pipes> <whatever> ...
```
This command will generate or add a line to a `results.xml` file in the same folder where `tesht` is called.

# TODO
 - _Parsing `result.xml` in_
 - _Adding result to xml structures_
 - _Writting `result.xml` back into disk_
 - Flags and options ??

# Motivation
Just want to create a bunch of bash unit-tests that translates to JUnits so my Jenkins-Blueocean thing shows them beautifully to everybody. Because bash testing is beautiful!
