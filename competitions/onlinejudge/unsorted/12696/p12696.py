import sys
from decimal import getcontext, Decimal

context = getcontext()
context.prec = 8


stdin = sys.stdin


MAX_L = Decimal('56.00')
MAX_W = Decimal('45.00')
MAX_D = Decimal('25.00')
MAX_THREE_SIDE = Decimal('125.00')
MAX_WW = Decimal('7.00')


def is_allowed(l, w, d, ww):
    return  ww <= MAX_WW and (
            (l <= MAX_L and
             w <= MAX_W and
             d <=  MAX_D) or ((l + w + d) <=  MAX_THREE_SIDE))



if __name__ == '__main__':
    n = int(sys.stdin.readline().strip())
    n_allowed = 0
    for _ in range(n):
        l, w, d, ww = list(map(Decimal, stdin.readline().strip().split()))
        if is_allowed(l, w, d, ww):
            n_allowed += 1
            print(1)
        else:
            print(0)
    
    print(n_allowed)
