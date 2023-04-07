import sys


stdin = sys.stdin


mmap = {
    0: 0,
    1: 1,
}


def fib(n):    
    return mmap[n]



if __name__ == '__main__':
    for i in range(2, 5000+1):
        mmap[i] = mmap[i-1] + mmap[i-2]

    for row in stdin:
        n = int(row.strip())
        print(f'The Fibonacci number for {n} is {fib(n)}')
