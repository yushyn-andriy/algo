import sys
from collections import namedtuple


Info = namedtuple('Info', [
    'team',
    'letter',
    'submittion_time',
    'status',
])


stdin = sys.stdin


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
    data = {}
    logs = sorted(logs, key=lambda x: (x.submittion_time, x.status))

    max_num = 0    
    for info in logs:
        if info.team > max_num:
            max_num = info.team

        if info.team not in data:
            data[info.team] = Stat(info.team)
        
        stat = data[info.team] 
        if info.letter in stat.solved:
            continue

        if info.status == 1:
            if info.letter in stat.rejected:
                stat.total_time += 20 * stat.rejected[info.letter]
            
            stat.solved[info.letter] = 0
            stat.n_solved += 1
            stat.total_time += info.submittion_time
        else:
            stat.n_rejected += 1
            if info.letter not in stat.rejected:
                stat.rejected[info.letter] = 0
            
            stat.rejected[info.letter]+= 1

    for i in range(1, max_num+1):
        if i not in data:
            data[i] = Stat(i)

    score_board = sorted(
        data.values(), 
        key=lambda x: (-x.n_solved, x.total_time, x.team),
    )
    return score_board


def str_to_min(s):
    h, m = [int(x) for x in s.split(':')]
    return h * 60 + m

if __name__ == '__main__':
    n = int(stdin.readline().strip())
    stdin.readline()
    for j in range(n):
        logs = []
        for row in stdin:
            row = row.strip()
            if not row:
                break
            l_info = row.split()
            l_info[0] = int(l_info[0])
            l_info[2] = str_to_min(l_info[2])
            l_info[3] = 1 if l_info[3] == 'Y' else 0
            info = Info(*l_info)
            logs.append(info)

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
