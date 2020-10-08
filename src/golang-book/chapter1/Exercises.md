# Exercises
1. What is whitespace?

	- Whitespace is the space between characters which is made of space, tab and newline characters.

2. What is a comment? What are the two ways of writing a comment?
	- Comment is a section of code which is ignored by compiler thier are only used for devloper to make notes. The two types of the comment are `//` or block comments `/* */`.

3. Our program began with package main. What would the files in the fmt package begin with?
	- the fmt package file would begin with fmt,

4. We used the Println function defined in the fmt package. If you wanted to use the Exit function from the os package, what would you need to do?
	- In order to exit the program you would need to import os package (libary) and then we would need to evnoke the os.Exit() Method.

5. Modify the program we wrote so that instead of printing Hello, World it prints Hello, my name is followed by your name.

```
package main

import "fmt"

func main(){
fmt.Println(Hello, my name is Hamza)
}
```
