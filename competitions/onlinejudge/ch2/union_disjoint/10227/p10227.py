import sys
from collections import defaultdict


stdin = sys.stdin


if __name__ == '__main__':
    T = int(stdin.readline().strip())
    stdin.readline()

    for i in range(T):
        PERSON, TREE = list(map(int, stdin.readline().strip().split()))
        pairs = []
        registry = set()
        pair_reg = set()
        for row in stdin:
            row = row.strip()
            if len(row) == 0:
                break

            person, tree =  list(map(int, row.split()))
            new_pair = (person, tree)
            if new_pair in pair_reg:
                continue

            pairs.append([person, tree])
            registry.add(person)
            pair_reg.add(new_pair)

        for idx in range(1, PERSON+1):
            if idx not in registry:
                pairs.append((idx, -1))

        pairs = sorted(pairs, key=lambda x: x[1])
        d = defaultdict(list)
        for p, t in pairs:
            d[p].append(t)
        
        unique = defaultdict(list)
        for key, value in d.items():
            unique[tuple(value)].append(key)
        
        print(len(unique))
        if i == T -1:
            continue
        print()






