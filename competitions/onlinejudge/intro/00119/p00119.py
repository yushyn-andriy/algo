import sys
from collections import OrderedDict




stdin = sys.stdin


if __name__ == '__main__':
    input_ = stdin.readline().strip()    
    while True:
        n = int(input_)
        names =  stdin.readline().strip().split()
        amounts = OrderedDict( zip(
            names,
            [
                {'received': 0, 'gave': 0}
                for _ in range(len(names))
            ]
        ) )
        for i in range(n):
            data = stdin.readline().strip().split()
            data[1:3] = [int(x) for x in data[1:3]]
            
            name = data[0]
            amount = data[1]
            n_friends = data[2]
            friends = data[3:3+n_friends]

            if n_friends == 0 or amount == 0:
                continue

            per_friend = amount // n_friends
            gave = per_friend * n_friends
            from pprint import pprint

            amounts[name]['gave'] += gave
            for friend_name in friends:
                amounts[friend_name]['received'] += per_friend


        for key, value in amounts.items():
            print(key, value['received'] - value['gave']) 

        input_ = stdin.readline().strip()    
        if input_ == '':
            break

        print()
