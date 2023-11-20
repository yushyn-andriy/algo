import sys


stdin = sys.stdin


'''
5  -> 0000 0101 -> Little Endian positive = 1010 0000 -> Big Endian positive
-5 -> 1111 1011 -> Little Endian negative = 1101 1111 -> Big Endian negative 
'''


def reverse_bits(num, bits_count=8):
    res = 0x00
    one = 0x01

    n = bits_count // 8 // 2
    # print(f'n={n}')
    cj = bits_count - 8
    ck = 0

    for b in range(n):
        # 0 - 7,   8 - 15,  17 - 23, 24 - 31
        # 24 - 31, 16 - 23,  8 - 15,  0 -  7

        j = cj
        k = ck
        # print(b, cj, ck)
        for _ in range(8):
            # print(format(res, '032b'))
            if j > 0:
                big_bit = (num & (one << j))
                litle_bit = (num & (one << k))
                diff = j - k            
                res = res | (big_bit >> diff) | (litle_bit << diff)
                j += 1
                k += 1

        ck += 8
        cj -= 8

    return res


def xor(a, b):
    return (a & ~b) | (b & ~a)


def binary(num):
    res = num
    return reverse_bits(res, 32)
    if num < 0:
        res = abs(num)
        res = xor(res, 0x01)
        # print(f'before', format(res, '032b'))
        res = reverse_bits(res, 32)
        # print(f'after', format(res, '032b'))
        return ~res
    else:
        if res & (0x1<<7) == 0:
            return reverse_bits(res, 32)

        # print(res)
        # print(f'before', format(res, '032b'))
        res = reverse_bits(~res, 32)
        res = xor(res, 0x1)
        # print(f'after', format(res, '032b'))
        return -res




if __name__ == '__main__':
    # print(format(reverse_bits(int('000110010011', base=2), 12), '04b'))
    # n = 123456789
    # print(n)
    # print('n', format(n, '032b'))

    # n = binary(n)

    # print('n', format(n, '032b'))
    # print(n)

    for line in stdin:
        num = int(line.strip())
        print(f'{num} converts to {binary(num)}')
