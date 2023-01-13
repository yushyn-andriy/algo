import sys
from collections import namedtuple


Info = namedtuple('Info', [
    'team',
    'letter',
    'submittion_time',
    'status',
])


stdin = sys.stdin

ACCEPTED = 1
REJECTED = 0
PENALTY = 20


class Stat:
    def __init__(self, team):
        self.solved ={}
        self.rejected = {}
        self.n_solved = 0
        self.n_rejected = 0
        self.total_time = 0
        self.team = team

    def __repr__(self) -> str:
        return f'S(team={self.team}, solved={self.n_solved}, total_time={self.total_time})'


def solve(logs):
    # statistics data from log file
    data = {}

    # should be sorted in order to get right results
    # 
    logs = sorted(logs, key=lambda x: (x.submittion_time, x.status))

    last_team_number = 0    
    for info in logs:
        # looking for last team number
        if info.team > last_team_number:
            last_team_number = info.team

        # add to stat to data if not presented
        if info.team not in data:
            data[info.team] = Stat(info.team)
        
        stat = data[info.team] 

        # ignore if already presented in solved
        if info.letter in stat.solved:
            continue
        
        # calculate total time
        # accounting rejected submissions
        if info.status == ACCEPTED:
            if info.letter in stat.rejected:
                stat.total_time += PENALTY * stat.rejected[info.letter]

            # mark as solved
            stat.solved[info.letter] = 0

            # increase solved count
            stat.n_solved += 1

            # increase total time
            stat.total_time += info.submittion_time
        else:
            # incerease rejected count
            stat.n_rejected += 1


            if info.letter not in stat.rejected:
                stat.rejected[info.letter] = 0
            
            stat.rejected[info.letter] += 1


    # add teams than not presented in data
    for i in range(1, last_team_number+1):
        if i not in data:
            data[i] = Stat(i)

    # sorting according requrements
    score_board = sorted(
        data.values(), 
        key=lambda x: (-x.n_solved, x.total_time, x.team),
    )
    return score_board


def str_to_min(s):
    h, m = [int(x) for x in s.split(':')]
    return h * 60 + m

def info_from_row(row):
    row = row.split()
    row[0] = int(row[0])
    row[2] = str_to_min(row[2])
    row[3] = 1 if row[3] == 'Y' else 0
    return Info(*row)

if __name__ == '__main__':
    n = int(stdin.readline().strip())
    stdin.readline()
    for j in range(n):
        logs = []
        for row in stdin:
            row = row.strip()
            if not row:
                break
            
            logs.append(info_from_row(row))

        results = solve(logs)

        print('RANK TEAM PRO/SOLVED TIME')
        rank = 0
        looser_rank = 0
        prev = None
        for i, r in enumerate(results):
            if r.n_solved:
                looser_rank += 1
                if i > 0  and prev.n_solved == r.n_solved and prev.total_time == r.total_time:
                    pass
                else:
                    rank = i + 1
                print(f'{rank:>4} {r.team:>4} {r.n_solved:>4} {r.total_time:>10}')
            else:
                print(f'{looser_rank+1:>4} {r.team:>4}')
            prev = r
    
        if j == n - 1:
            continue
        print()
