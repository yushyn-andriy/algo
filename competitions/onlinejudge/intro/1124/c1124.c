#include <stdio.h>

int main() {
	char s[100];
	while(fgets(s, 100, stdin) != NULL) {
		printf("%s", s);
	}
	return 0;
}
