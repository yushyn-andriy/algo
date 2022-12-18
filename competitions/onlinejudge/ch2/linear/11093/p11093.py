import sys



stdin = sys.stdin


def get_available_station(n, petrol, consuptions):
    for i in range(n):
        j = (i + 1) % n
        avail_petrol = petrol[i] - consuptions[i]
        while j != i and avail_petrol >= 0:
            avail_petrol += petrol[j]
            avail_petrol -= consuptions[j]
            j = (j + 1) % n
        if i == j and avail_petrol >= 0:
            return i + 1, True
    return 0, False



if __name__ == '__main__':
    k = stdin.readline().strip()
    if not k:
        sys.exit(0)
    k = int(k)
    for i in range(k):
        l = stdin.readline().strip()
        if not l:
            break
        n = int(l)
        petrol = [int(x) for x in stdin.readline().strip().split()[:n]]
        consuptions = [int(x) for x in stdin.readline().strip().split()[:n]]
        station, ok = get_available_station(n, petrol, consuptions)
        if ok:
            print("Case %s: Possible from station %s" % (i+1, station))
        else:
            print("Case %s: Not possible" % (i+1))

