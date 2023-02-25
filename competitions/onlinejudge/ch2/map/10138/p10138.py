from collections import defaultdict
from typing import List, Tuple
import sys


stdin = sys.stdin


class Record:
    def __init__(self, month: int, day: int, hour: int, mins: int, loc: int, is_exit: bool):
        self.month = month
        self.day = day
        self.hour = hour
        self.mins = mins
        self.loc = loc
        self.is_exit = is_exit
        
    def get_time(self):
        return self.mins + self.hour * 60 + self.day * 24 * 60


def main() -> None:
    t = int(stdin.readline().strip())

    stdin.readline() # skip blank line
    while t > 0:
        fares = list(map(int, stdin.readline().strip().split()))
        license_to_record = defaultdict(list)
        for line in stdin:
            line = line.strip()
            if not line:
                break
            plate, month, day, hour, mins, command, loc = parse_record(line)
            r = Record(month, day, hour, mins, loc, command == "exit")
            license_to_record[plate].append(r)
        
        for plate, records in sorted(license_to_record.items()):
            records.sort(key=lambda r: r.get_time())
            total_cents = 200
            for i in range(len(records)):
                if not records[i].is_exit and i+1 < len(records) and records[i+1].is_exit:
                    dist = abs(records[i].loc - records[i+1].loc)
                    total_cents += dist * fares[records[i].hour]
                    total_cents += 100
            if total_cents != 200:
                print(f"{plate} ${total_cents / 100:.2f}")
        if t > 1:
            print()
        t -= 1

def parse_record(line: str) -> Tuple[str, int, int, int, int, str, int]:
    plate, time, command, loc = line.split()
    month, day, hour, mins = map(int, time.split(':'))
    loc = int(loc)
    return plate, month, day, hour, mins, command, loc

if __name__ == "__main__":
    main()
