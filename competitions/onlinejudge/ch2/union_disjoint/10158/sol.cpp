#include<bits\stdc++.h>
using namespace std;
int G[20005], n;
int find(int x){
	return G[x] == x ? x : G[x] = find(G[x]);
}
int main(){
	scanf("%d", &n);
	for (int i = 0; i < 2 * n; i++)
		G[i] = i;
	int a, b, c;
	while (scanf("%d%d%d", &c, &a, &b) == 3){
		if (a == 0 && b == 0 && c == 0)	break;
		int a1 = find(a), a2 = find(a + n);
		int b1 = find(b), b2 = find(b + n);
		if (c == 1){
			if (a1 == b2)
				puts("-1");
			else
				G[a1] = b1, G[a2] = b2;
		}
		if (c == 2){
			if (a1 == b1)
				puts("-1");
			else
				G[a1] = b2, G[a2] = b1;
		}
		if (c == 3)
			printf("%d\n", a1 == b1 ? 1 : 0);
		if (c == 4)
			printf("%d\n", a1 == b2 ? 1 : 0);
	}
	return 0;
}