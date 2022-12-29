import sys
from collections import namedtuple


Interval = namedtuple('Interval', 'begin end')



stdin = sys.stdin



def readlist(converter):
    return [converter(x) for x in stdin.readline().strip().split()]


def all_green(ct, time):
    for t in time:
        if ct % (t * 2) >(t - 5):
            return False
    return True


if __name__ == '__main__':
    max_seconds = 3600
    idx = 1
    while True:
        times = readlist(int)
        if not times:
            break
        ct = min(times) * 2
        while not all_green(ct, times) and ct <= max_seconds:
            ct+=1

        if ct >= max_seconds + 1:
            print(f"Set {idx} is unable to synch after one hour.")
        else:
            print(f"Set {idx} synchs again at  {ct // 60} minute(s) and {ct % 60} second(s) after all turning green.")

        idx += 1
