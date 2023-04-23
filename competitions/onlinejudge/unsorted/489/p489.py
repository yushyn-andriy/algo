import sys



stdin = sys.stdin


def get_result(solution, guess):
    d = set()
    initial_solution_set = set(solution) 
    for ch in solution:
        d.add(ch)
    
    tried_wrong = set()
    tried_right = set()
    for ch in guess:

        if ch in initial_solution_set:
            if ch in d:
                d.remove(ch)
            if not d:
                return 'You win.'
            tried_right.add(ch)
        elif ch not in tried_wrong:
            tried_wrong.add(ch)
            if len(tried_wrong) >= 7:
                return 'You lose.'


    if len(tried_right) < len(initial_solution_set):
        return 'You chickened out.'

    return 'You lose.'


if __name__ == '__main__':
    for row in stdin:
        rnd = int(row.strip())
        if rnd == -1:
            break

        solution = stdin.readline().strip()
        guess = stdin.readline().strip()

        print(f'Round {rnd}')
        print(get_result(solution, guess))
