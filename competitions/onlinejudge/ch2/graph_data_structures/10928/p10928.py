import sys



stdin = sys.stdin



if __name__ == '__main__':
    t = int(stdin.readline().strip())
    for _ in range(t):
        n = int(stdin.readline().strip())
        

        nbrs = {}
        place = 1
        for _, row in enumerate(stdin):
            row = row.strip()
            if not row:
                break

            nbrs[place] = len(row.split())
            place += 1


        items = list(sorted(nbrs.items(), key=lambda x: x[1]))
        place1 = items[0]
        result = [str(place1[0])]
        for p in items[1:]:
            if p[1] == place1[1]:
                result.append(str(p[0]))
        print(' '.join(result))
