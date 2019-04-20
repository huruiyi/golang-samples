package main

/*

#include <stdio.h>
void foo() {
	printf("this is c func foo...\n");
}

void bar(){
	printf("this is a func bar...\n");
}
*/
import "C"

func main() {
	C.foo()
}
