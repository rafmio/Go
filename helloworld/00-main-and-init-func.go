main() and init() functions
https://www.geeksforgeeks.org/main-and-init-function-in-golang/ 

The Go reserve 2 functions for special purpose and the functions are:
- main()
- init() 

main() - the entry point of the executable programs. It doesn't take any arguments nor return anything.
Go automatically call main()  

init() function is present in every package and this function is called when the package is initialized. This function is declared implicitly, so you cannot reference it from anywhere and you are allowed to create multiple init() in the same program and the execute in the order they are created. 

init() executed before the main(), it does not depend on main()

The main purpose of the init() is to initialize the global variables that cannot be initialized in the global context.

