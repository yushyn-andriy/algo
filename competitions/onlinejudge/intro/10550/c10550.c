#include <stdio.h>
#include <stdbool.h>


short turn_until(short cursor, short stop, bool clockwise) {
	short total = 0;
	while(cursor != stop) {
		total += 9;
		if(clockwise) {
			cursor = (cursor - 1) % 40;
			if(cursor<0) cursor = 39;
		} 
		else
		cursor = (cursor +1) % 40;
	}
	return total;
}


short get_degree(short a, short b, short c, short d) {
	short total = 1080;
	total += turn_until(a, b, true);
	total += turn_until(b, c, false);
	total += turn_until(c, d, true);

	return total;
}

int main() {
	short a, b, c, d;
	short degrees = 1080;
	while(scanf("%hi %hi %hi %hi", &a, &b, &c, &d) != EOF) {
		if (a==0 && b==0 && c==0 && d==0) break;	
		printf("%hi\n", get_degree(a, b, c, d));
	}

	return 0;
}
