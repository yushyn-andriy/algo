import sys
from collections import OrderedDict, deque



cin = sys.stdin


def enqueue(mate, tq, tmt):
    group_id = tmt[mate]
    if group_id not in tq:
        tq[group_id] = deque([])
    tq[group_id].append(mate)


def dequeue(tq):
    for group_id in tq.keys():
        mate = tq[group_id].popleft()
        if not tq[group_id]:
            del tq[group_id]
        break
    return mate    


if __name__ == '__main__':
    for i, row in enumerate(cin):
        n_teams = int(row.strip())
        if n_teams == 0:
            break
        
        # if i != 0:
        #     print()

        teammate_team_map = {}
        for team_id in range(n_teams):
            teammates = [int(x) for x in cin.readline().strip().split()][1:]
            for mate in teammates:
                teammate_team_map[mate] = team_id

        
        print(f'Scenario #{i+1}')

        team_queue = OrderedDict()
        for row in cin:
            args = row.strip().split()
            cmd = args[0]
            if cmd == 'STOP':
                break
            if cmd == 'ENQUEUE':
                mate = int(args[1])
                enqueue(mate, team_queue, teammate_team_map)
            elif cmd == 'DEQUEUE':
                mate = dequeue(team_queue)
                print(mate)
        print()