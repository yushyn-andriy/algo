#include <iostream>
#include <cstdio>

using namespace std;

int main() {
	unsigned int x = 0x5;
	unsigned int y = 0x7;

	printf("x - %.8b, y - %.8b\n", x, y);
	printf("x|y %8.b\n", x|y);
	printf("x&y %.8b\n", x&y);
	printf("x^y %.8b\n", x^y);
	printf("x<<1 %.8b\n", x<<1);
	printf("x>>1 %.8b\n", x>>1);
	
	return 0;
}
