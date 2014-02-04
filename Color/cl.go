package Color

/*
#include <windows.h>
void Cl(int cl){
	SetConsoleTextAttribute(GetStdHandle(STD_OUTPUT_HANDLE), cl);
}
*/
import "C"

func SetColor(cl int) {
	C.Cl(C.int(cl))
}
