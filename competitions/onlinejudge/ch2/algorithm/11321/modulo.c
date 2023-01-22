#include <stdio.h>



int main() {
    for (int i = -10000; i < 10000; i++)
        printf("%d\n", i%3);
    return 0;
}