#include <stdio.h>
#include <stdbool.h>


int main() {
	char ch;
	bool flag = true;
	while((ch = getchar()) != EOF) {
		if(ch != '"') {
			putchar(ch);
			continue;
		}
		if(ch == '"' && flag) {
			flag = !flag;
			putchar('`');
			putchar('`');
			continue;
		} else
		if(ch == '"' && !flag) {
			putchar('\'');	
			putchar('\'');	
			flag = !flag;
		}
	}
	
	return 0;
}
