### **Basic calculator Golang module**

Basic calculator provides a primitive implementation of a calculator for evaluating simple expressions consisting of single digits and mathematical addition and subtraction operators.

You can give input string representing an expression to ```Eval()``` function and receive expected result.

For instance:

```golang
    input := "1+1"  
    Eval(input) 
    // Output: 2`
```  

Another example with a more complex expression:

```golang
    input := "2+1 -  2"
    Eval(input)  
    // Output: 1
```  


Given a string s representing an expression, implement a basic calculator to evaluate it.
Expression consists of single digits (like '2', '7' etc), `'+'`, `'-'` and `' '`

Example 1:
Input: s = "1 + 1"
Output: 2

Example 2:
Input: s = " 2-1 + 2 "
Output: 3
