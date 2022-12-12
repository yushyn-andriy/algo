from collections import namedtuple
import sys


record = namedtuple('Record', 'square animals reward')


def burden(rec):
    return rec.square * rec.reward


def total(records):
    total = 0
    for rec in records:
        total += burden(rec)
    return total


if __name__ == '__main__':
    cases = int(sys.stdin.readline().strip())
    for i in range(cases):
        farmers = int(sys.stdin.readline().strip())
        records = []
        for j in range(farmers):
            rec = record._make([
                int(x)
                for x in sys.stdin.readline().strip().split()
            ])
            records.append(rec)
        print(total(records))
