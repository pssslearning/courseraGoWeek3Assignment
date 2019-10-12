# pssslearning courseraGoWeek3Assignment
Assigment Week 3 - Course  https://www.coursera.org/learn/golang-getting-started/peer/sLPZg/module-3-activity-slice-go
## Coursera course: Programming with Google Go (University of California, Irvine) 

[Coursera - Getting Started with Go](https://www.coursera.org/learn/golang-getting-started/home/welcome)

### Week 3 - Assignment 1 - `slice.go`

#### Goals

The goal of this assignment is to practice working with integers, slices and loops in Go.

#### Assignment directions

Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

Submit your source code for the program, “slice.go”.

#### My assignment proposal is placed at
```sh
slice/slice.go
```

#### A sample log for its compiling a testing cycle:

```sh
devel1@vbxdeb10mate:~$ cd $GOPATH/src/github.com/pssslearning/courseraGoWeek3Assignment/slice/
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoWeek3Assignment/slice$ go build slice.go 
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoWeek3Assignment/slice$ ./slice

Welcome user to the Assignment program for Week 3,

--------------------------------------------------------------
Please enter a string that can be interpreted as INTEGER ... 
      (press <CTRL+C> or 'X' to exit.) 
--------------------------------------------------------------
45
--------------------------------------------------------------
The ordered list of integers introduced by now are:
Position #0 - Value[45]
--------------------------------------------------------------
Please enter a string that can be interpreted as INTEGER ... 
      (press <CTRL+C> or 'X' to exit.) 
--------------------------------------------------------------
-3
--------------------------------------------------------------
The ordered list of integers introduced by now are:
Position #0 - Value[-3]
Position #1 - Value[45]
--------------------------------------------------------------
Please enter a string that can be interpreted as INTEGER ... 
      (press <CTRL+C> or 'X' to exit.) 
--------------------------------------------------------------
80
--------------------------------------------------------------
The ordered list of integers introduced by now are:
Position #0 - Value[-3]
Position #1 - Value[45]
Position #2 - Value[80]
--------------------------------------------------------------
Please enter a string that can be interpreted as INTEGER ... 
      (press <CTRL+C> or 'X' to exit.) 
--------------------------------------------------------------
234
--------------------------------------------------------------
The ordered list of integers introduced by now are:
Position #0 - Value[-3]
Position #1 - Value[45]
Position #2 - Value[80]
Position #3 - Value[234]
--------------------------------------------------------------
Please enter a string that can be interpreted as INTEGER ... 
      (press <CTRL+C> or 'X' to exit.) 
--------------------------------------------------------------
-2345
--------------------------------------------------------------
The ordered list of integers introduced by now are:
Position #0 - Value[-2345]
Position #1 - Value[-3]
Position #2 - Value[45]
Position #3 - Value[80]
Position #4 - Value[234]
--------------------------------------------------------------
Please enter a string that can be interpreted as INTEGER ... 
      (press <CTRL+C> or 'X' to exit.) 
--------------------------------------------------------------
FAKE
ENTRY REJECTED [FAKE]. Cannot be parsed as an integer.
--------------------------------------------------------------
Please enter a string that can be interpreted as INTEGER ... 
      (press <CTRL+C> or 'X' to exit.) 
--------------------------------------------------------------
3456
--------------------------------------------------------------
The ordered list of integers introduced by now are:
Position #0 - Value[-2345]
Position #1 - Value[-3]
Position #2 - Value[45]
Position #3 - Value[80]
Position #4 - Value[234]
Position #5 - Value[3456]
--------------------------------------------------------------
Please enter a string that can be interpreted as INTEGER ... 
      (press <CTRL+C> or 'X' to exit.) 
--------------------------------------------------------------
-34
--------------------------------------------------------------
The ordered list of integers introduced by now are:
Position #0 - Value[-2345]
Position #1 - Value[-34]
Position #2 - Value[-3]
Position #3 - Value[45]
Position #4 - Value[80]
Position #5 - Value[234]
Position #6 - Value[3456]
--------------------------------------------------------------
Please enter a string that can be interpreted as INTEGER ... 
      (press <CTRL+C> or 'X' to exit.) 
--------------------------------------------------------------
X

Goodbye !!!

```