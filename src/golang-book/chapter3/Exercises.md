# Exercises


1. What are two ways to create a new variable?

   - The two ways to create a new variable are following:
        
        -  ```
           var x int = 5
           ```
        - ```
          x:=50
          ```
    
2. What is the value of x after running x := 5; x += 1?

    - The value of x would be 6.
    
3. What is scope? How do you determine the scope of a variable in Go?

    - Scope is where variable can be used. The scope in Go is determined lexically using blocks, so look for the nearest curly braces.

4. What is the difference between var and const?
    
    -  The different between var and const is that var has the ability to update it value during runtime where as const can only be set once and can be change.
    
5. Using the example program as a starting point, write a program that converts from Fahrenheit into Celsius (C = (F − 32) * 5/9).
    
     ```
      package main 
      
      import "fmt"
      
      func main(){
           	fmt.Println("Enter a number: ")
            var input float64
            fmt.Scanf("%f", &input)
            output:= (input - 32) * 5/9
            fmt.Println(output)
      }
      ```
      
6. Write another program that converts from feet into meters (1 ft = 0.3048 m).

     ```
      package main 
      
      import "fmt"
      
      func main(){
           	fmt.Println("Enter a number: ")
            var input float64
            fmt.Scanf("%f", &input)
            output:= input  *  0.3048
            fmt.Println(output)
      }
      ```
      