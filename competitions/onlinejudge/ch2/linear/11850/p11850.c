#include<stdio.h>
#include <stdlib.h>

#define TOTAL_MILES 1422

int cmpfunc (const void * a, const void * b) {
   return ( *(short int*)a - *(short int*)b );
}

int main() {
    short int n;
    while (1) {
        scanf("%hd", &n);
        if (n == 0) break;
        short int st[n];
        for(short int i = 0; i<n; i++)
            scanf("%hd", &st[i]);

        char possible = 1;
        qsort(st, n, sizeof(short int), cmpfunc);
        for(short int i = 0; i< n -1; i++) {
            short int diff = st[i+1] - st[i];

            if (diff <= 200) {
                continue;
            }
            else possible = 0;
        }
        if (possible == 0 || ((TOTAL_MILES - st[n-1])*2 > 200)) {
           printf("IMPOSSIBLE\n");
        } else {
           printf("POSSIBLE\n");
        }
    }

    return 0;
}
