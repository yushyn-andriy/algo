import sys
import string



stdin = sys.stdin

string.ascii_lowercase


def get_digits_sum(s):
    if len(s) == 1:
        return float(s)
    
    cs = 0
    for ch in s:
        cs += int(ch)

    return get_digits_sum(str(cs))


def get_ratio(name, score_map):
    total_score = 0
    for ch in name:
        if ch in score_map:
            total_score += score_map[ch]
    return get_digits_sum(str(total_score))



if __name__ == '__main__':

    score_map = {}
    for i, ch in enumerate(string.ascii_lowercase):
        score_map[ch] = i+1
    while True:
        first_name = stdin.readline().strip().lower()
        if not first_name:
            break
        second_name = stdin.readline().strip().lower()


        r1 = get_ratio(first_name, score_map)
        r2 = get_ratio(second_name, score_map)

        ratio = (r1 / r2) * 100
        if ratio > 100:
            ratio = (r2/r1) * 100 

        print('{:.2f} %'.format(ratio))
    