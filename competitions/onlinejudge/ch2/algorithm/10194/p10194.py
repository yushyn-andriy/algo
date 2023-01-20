import sys
from collections import namedtuple


stdin = sys.stdin


Match = namedtuple('Match', 'com1 score1 com2 score2')

class Stat:
    __slots__ = (
        'name',
        'team_rank',
        'total_points_earned',
        'games_played',
        'wins',
        'ties',
        'losses',
        'goal_difference',
        'goals_scored',
        'goals_against',
    )

    def __init__(self, name):
        self.name = name

        self.team_rank = 0
        self.total_points_earned = 0
        self.games_played = 0
        self.wins = 0
        self.ties = 0
        self.losses = 0
        self.goal_difference = 0
        self.goals_scored = 0
        self.goals_against = 0


def parse_match(row):
    p1, p2 = row.split('@')
    com1, score1 = p1.split('#')
    score2, com2  = p2.split('#')
    score1, score2 = int(score1), int(score2)
    return Match(com1, score1, com2, score2)


def update_stats(stats, m):
    com1, score1, com2, score2 = m
    s1, s2 = stats[com1], stats[com2]
    
    # games played
    s1.games_played += 1
    s2.games_played += 1

    # scored
    s1.goals_scored += score1
    s2.goals_scored += score2

    # against
    s1.goals_against += score2
    s2.goals_against += score1

    # goals difference
    s1.goal_difference = s1.goals_scored - s1.goals_against
    s2.goal_difference = s2.goals_scored - s2.goals_against

    if score1 > score2:
        s1.wins += 1
        s2.losses += 1
        
        s1.total_points_earned += 3
    elif score1 < score2:
        s2.wins += 1
        s1.losses += 1

        s2.total_points_earned += 3
    else:
        s1.total_points_earned += 1
        s2.total_points_earned += 1
        
        s1.ties += 1
        s2.ties += 1
    

    



if __name__ == '__main__':
    n = int(stdin.readline().strip())
    for j in range(n):
        tournament_name = stdin.readline().strip()
        
        n_teams = int(stdin.readline().strip())
        stats = {}
        for _ in range(n_teams):
            name = stdin.readline().strip()
            stats[name] = Stat(name)


        n_played = int(stdin.readline().strip())
        for _ in range(n_played):
            match = parse_match(stdin.readline().strip())
            update_stats(stats, match)


        data = sorted(stats.values(), key=lambda x: (
            -x.total_points_earned,
            -x.wins,
            -x.goal_difference,
            -x.goals_scored,
            x.games_played,
            x.name.lower(),
        ))


        print(tournament_name)
        for i, r in enumerate(data):
            rank = i + 1
            s = '{rank}) {name} {b}p, {c}g ({d}-{e}-{f}), {g}gd ({h}-{i})'.format(
                rank=rank,
                name=r.name,
                b=r.total_points_earned,
                c=r.games_played,
                d=r.wins,
                e=r.ties,
                f=r.losses,
                g=r.goal_difference,
                h=r.goals_scored,
                i=r.goals_against,
            )
            print(s)
        
        if j < n - 1:
            print()
