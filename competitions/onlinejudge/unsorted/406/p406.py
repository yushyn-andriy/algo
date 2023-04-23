import sys


stdin = sys.stdin


def list_str(l):
    return ' '.join(map(str, l))


if __name__ == '__main__':
    upper_bound = 1001

    L = [i for i in range(0, upper_bound)]
    for i in range(2, upper_bound):
        for j in range(2, upper_bound):
            if i*j < len(L):
                L[i*j] = 0
            else:
                break
    

    for row in stdin:
        a, b = list(map(int, row.split()))
        prime_numbers = []
        for v in L[:a+1]:
            if v > a:
                break
            if v != 0:
                prime_numbers.append(v)

        if len(prime_numbers) % 2 == 1 and 2 * b - 1 > len(prime_numbers):
            print(f"{a} {b}:", list_str(prime_numbers))
        elif len(prime_numbers) % 2 == 0 and 2 * b > len(prime_numbers):
            print(f"{a} {b}:", list_str(prime_numbers))
        elif len(prime_numbers) % 2 == 1:
            mid = len(prime_numbers) // 2
            start = mid - b + 1
            stop = start + (2 * b - 1)
            print(f"{a} {b}:", list_str(prime_numbers[start:stop]))
        else:
            mid = len(prime_numbers) // 2
            start = mid - b
            stop = start + 2 * b
            print(f"{a} {b}:", list_str(prime_numbers[start:stop]))
        print()