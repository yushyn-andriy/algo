import sys


stdin = sys.stdin



def get_winner(P, grand_prix, scoring_systems):
    global_winners = []
    for scoring_system in scoring_systems:
        pilot_scores = [0] * P
        for i in range(len(grand_prix)):
            prix = grand_prix[i]
            for pilot_id, place in enumerate(prix):
                if place - 1 < len(scoring_system):
                    points_for_place = scoring_system[place - 1]
                    pilot_scores[pilot_id] += points_for_place
        

        best_score = max(pilot_scores)
        prix_winners = []
        for pilot_id, score in enumerate(pilot_scores):
            if score == best_score:
                prix_winners.append(pilot_id + 1)
        
        global_winners.append(prix_winners)

    return global_winners



if __name__ == '__main__':
    while True:
        G, P = [int(x) for x in stdin.readline().strip().split()]
        if G == 0:
            break
        
        scores = []
        grand_prix = []
        for _ in range(G):
            grand_prix.append([int(x) for x in stdin.readline().strip().split()])
        
        S = int(stdin.readline().strip())
        for _ in range(S):
            points = [int(x) for x in stdin.readline().strip().split()][1:]
            scores.append(points)
        
        winners_by_scoring_system = get_winner(P, grand_prix, scores)
        for winners in winners_by_scoring_system:
            print(' '.join([str(x) for x in winners]))

