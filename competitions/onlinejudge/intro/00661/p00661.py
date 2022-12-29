import sys



stdin = sys.stdin


def max_consuption(consuptions, switches):
    max_cons = -1
    current_consuption = 0
    switches_status = [False] * len(consuptions)
    for i in range(1, len(switches)):
        switch_idx = switches[i]
        switched_on = switches_status[switch_idx]
        if switched_on:
            current_consuption -= consuptions[switch_idx]
        else:
            current_consuption += consuptions[switch_idx]
            
        switches_status[switch_idx] = not switched_on 
        if current_consuption > max_cons:
            max_cons = current_consuption
    return max_cons


if __name__ == '__main__':
    sequance = 0
    first_input = True
    while True:
        sequance += 1
        n, m, c = [int(x) for x in stdin.readline().strip().split()]
        if n == 0 and m == 0 and c == 0:
            break
        
        if sequance == 1:
            pass
        else:
            pass #print()


        consumptions = [0] * (n + 1)
        switches = [0] * (m + 1)
        for i in range(1, n + 1):
            consumptions[i] = int(stdin.readline().strip())
        for i in range(1, m + 1):
            switches[i] = int(stdin.readline().strip())
        
        max_cons = max_consuption(consumptions, switches)
        print(f'Sequence {sequance}')
        if max_cons > c:
            print('Fuse was blown.')
        else:
            print('Fuse was not blown.')
            print(f'Maximal power consumption was {max_cons} amperes.')
        print()
