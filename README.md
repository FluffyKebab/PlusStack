# PlusStack
Plus Stack is a interpretd langue witten in go.

##Memory in Plus Stack
The only form of memory avialable in Plus Stack is a stack. The stack can contain two data types: Integer arrays and functions.

##How soucre code is executed
The interpreter reads the source code from left to right and ignores most whitespace. If the current token is recegnised as a argument, the interpreter adds the current value to the list of arguments. If the current token is recegnised as a function the interpreter executes the function with the list of arguments.

Here is a code example:
```
10 30 +
```
When running this code the interpreter will first add 10 and 30 to the list of arguments. Secondly it will execute the plus function with 10 30 as arguments. The result will be the integer array [ 40 ]. Finaly the result will then be pushed to the memory stack.

##Popping the stack
There are two tokens that pops the stack: "A" and ".". The first poping token, "A", popes all elements in the stack and "." popes one. If the element at the top of the stack is a function, the function will be exectued, and if the element is an argument 

Here is a example 
```
30 . +
```
Firstly 30 will be added to the list of arguments. Then the element at the top of the stack gets popped. Assuming that this element is an integer array, it would also be added to the list of arguments. Then the two integer arrays will be added and the result pushed onto the stack.

##Functions


##Function definitons

###Plus +
If the plus function reciveces one integer array

###Minus -

###Multiplication * (Not implemended)
###Division : (Not implemended)

###Push ]

###Reduce /

###Delete D

###Map

###Ternary
