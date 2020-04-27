## Basic Data Types in GO :
### Integers 
There are four different types of signed and unsigned integers available based on their sizes.
 - int8	 : 8-bit signed integer
 - int16	 : 16-bit signed integer
 - int32	 : 32-bit signed integer
 - int64	 : 64-bit signed integer
 - uint8	 : 8-bit unsigned integer
 - uint16 : 16-bit unsigned integer
 - uint32 : 32-bit unsigned integer
 - uint64 :	64-bit unsigned integer
 - int : Both in and uint contain same size, either 32 or 64 bit.
 - uint : Both in and uint contain same size, either 32 or 64 bit.
 - rune : It is a synonym of int32 and also represent Unicode code points.
 - byte : It is a synonym of int8 .
 - uintptr : It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value.
  
### Floating Point 
 - float32	: 32-bit IEEE 754 floating-point number
 - float64	: 64-bit IEEE 754 floating-point number
### Strings 
   ```
    x:= "Hello World
    // Prints the String 'x' to Console
    fmt.Print(x)
```
