#include <stdint.h>
#include <stdio.h>

#define mask 0b11111111


uint32_t reverse_bits(uint32_t num) {
    return 0x00 | ((num & mask) << 24) | ((num & (mask << 24)) >> 24) | ((num & (mask<<8)) << 8) | ((num & (mask << 16)) >> 8);
}

int main() {
    int32_t num;
    uint32_t reversed;
    while(scanf("%d\n", &num) != EOF) {
        reversed = reverse_bits(num);
        if(((reversed &  (0x1 << 31)) != 0))
            printf("%d converts to %d\n", num, (int32_t)reversed);
        else
            printf("%d converts to %d\n", num, reversed);
    }

    return 0;
}
