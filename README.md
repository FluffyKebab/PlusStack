# PlusStack
Plus Stack is a interpretd langue witten in go.

## Memory in Plus Stack
The only form of memory avialable in Plus Stack is a stack. The stack can contain two data types: Integer arrays and functions.

## How soucre code is executed
The interpreter reads the source code from left to right and ignores most whitespace. If the current symbol is recegnised as an argument, the interpreter adds the current value to the list of arguments. If the current token is recegnised as a function the interpreter executes the function with the list of arguments.

Here is a code example:
```
10 30 +
```
When running this code the interpreter will first add 10 and 30 to the list of arguments. Secondly it will execute the plus function with 10 30 as arguments. The result will be the integer array [ 40 ]. Finaly the result will then be pushed to the memory stack.

## Popping the stack
There are two tokens that pops the stack: "A" and ".". The first poping token, "A", popes all elements in the stack and "." popes one. If the element at the top of the stack is a function, the function will be exectued, and if the element is an argument 

Here is a example 
```
30 . +
```
Firstly 30 will be added to the list of arguments. Then the element at the top of the stack gets popped. Assuming that this element is an integer array, it would also be added to the list of arguments. Then the two integer arrays will be added and the result pushed onto the stack.

## Functions


## Function definitons

### Plus +
If the plus function reciveces one integer array, the function will for every integer in the array push a new array with the lenght of the integer. 
```
2 3 1]
.+

Output:
0) [ 0 ]
1) [ 0 0 0 ]
2) [ 0 0 ]
```

If the plus function reviveces multipe intiger arrays with the same lenghts, the elements at each posistion will be added.
```
1 2 3]
3 2 1]
..+

Output:
0) [ 4 4 4 ]
```

If the function recivces 

```
0 0 0]
1 1 1 1]
1 5 A+

Output:
0) [ 7 7 7 7 ] 
1) [ 6 6 6 ]
```

### Minus -

### Multiplication * (Not implemended)
### Division : (Not implemended)

### Reduce /

### Delete D
Takes any input and does nothing

### Map


### Ternary (?)
The ternary function takes an array of lenght 1 with the value of 1 or 0 and to functions. If the value given is 1, the first function will be pushed to the memory stack. If the value is 0, the second function will be pushed.

```
1 (A 10 +) (A 20 +) ?
30 .
```
Here the first function will be pushed to the stack. Then the function is executed with 30 as an argument. The result, 40, will then be pushed to the memory stack.

