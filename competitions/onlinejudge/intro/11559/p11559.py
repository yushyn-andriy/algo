import sys
from collections import namedtuple
import math


Travel = namedtuple('Travel', 'participants budget n_hotels n_weeks')


if __name__ == '__main__':
    while True:
        input_ = sys.stdin.readline()
        if input_ == '':
            break

        travel = Travel(*list(map(int, input_.strip().split())))
        min_price = math.inf
        for i in range(travel.n_hotels):
            price_per_participant = int(sys.stdin.readline().strip())
            available_beds = list(map(int, sys.stdin.readline().strip().split()))

            b = travel.participants * price_per_participant
            if b > travel.budget:
                continue
            
            is_vailable = False
            for av in available_beds:
                if av >= travel.participants:
                    is_vailable = True
                    break
            
            if is_vailable and b < min_price:
                min_price = b
        if min_price is not math.inf:
            print(min_price)
        else:
            print('stay home')            
