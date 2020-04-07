## Deallocating Memory in GO
- Variables refer to some data that is present somewehere in memory. So variables eventually have to be deallocated in memory.
- If we don't deallocate memory, every time the function with the variable is executed will use a new allocated memory, which can eventually lead to a memory leak. 
- Programming Languages such as C, you need to explicitly deallocate memory, but in Go, most part is done automatically.