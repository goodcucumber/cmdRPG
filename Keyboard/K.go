package Keyboard

/*
	#include <conio.h>
	#include <stdio.h>
	#include <stdlib.h>
	int kbh(){
		return _kbhit();
	}
	int gch(){
		return _getch();
	}
	void clear(){
		fflush(stdin);
	}
*/
import "C"

const (
	KUP    int = 72 + 255
	KDOWN  int = 80 + 255
	KLEFT  int = 75 + 255
	KRIGHT int = 77 + 255
	KENTER int = 13
	KSPACE int = 32
	KS1    int = 115
	KS2    int = 83
	KL1    int = 108
	KL2    int = 76
	KQ1    int = 113
	KQ2    int = 81
	KESC   int = 27
)

func Hit() bool {
	if C.kbh() == 1 {
		return true
	}
	return false
}

func GetKeyCode() int {
	a := int(C.gch())
	if a == 224 {
		return int(C.gch()) + 255
	}
	return a
}

func ClearBuffer() {
	C.clear()
}
