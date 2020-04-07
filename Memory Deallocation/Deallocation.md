## Deallocating Memory in GO
- Variables refer to some data that is present somewehere in memory. So variables eventually have to be deallocated in memory.
- If we don't deallocate memory, every time the function with the variable is executed will use a new allocated memory, which can eventually lead to a memory leak. 
- Programming Languages such as C, you need to explicitly deallocate memory, but in Go, most part is done automatically.

## Garbage Collection in GO
- Garbage Collection is an automatic tool that deals with Deallocation.
- Garbage Collection in Interpreted Languages are done by the Interpreter. Example : JVM (Java Virtual Machine, Python Interpreter).
- Interpreters keep track of Pointers and determines when a variable is not in use anymore, i.e no pointers, no references. The garbage collector then Deallocates it.
- GO is a compiled language with Garbage Collector built in. Which is why GO is fast. 