import sys


stdin = sys.stdin


def count_moves(S, T):
    s_zero_qty = s_unit_qty = s_question_qty = 0
    t_zero_qty = t_unit_qty = 0

    S = list(S)
    T = list(T)
    for i in range(len(S)):
        if S[i] == '0':
            s_zero_qty += 1
        elif S[i] == '1':
            s_unit_qty += 1
        elif S[i] == '?':
            s_question_qty += 1
        else:
            raise ValueError
        
        if T[i] == '0':
            t_zero_qty += 1
        elif T[i] == '1':
            t_unit_qty += 1
        else:
            raise ValueError

    assert s_zero_qty + s_unit_qty + s_question_qty == len(S)
    assert t_zero_qty + t_unit_qty == len(T)


    zeros_diff = s_zero_qty - t_zero_qty
    unit_diff = s_unit_qty - t_unit_qty
    n_edit = 0
    for i in range(len(S)):
        sch = S[i]
        tch = T[i]

        if sch == '?':
            if tch == '0' and zeros_diff < 0:
                S[i] = '0'
                zeros_diff += 1
                n_edit += 1
                s_question_qty -= 1
            elif tch == '1' and unit_diff < 0:
                S[i] = '1'
                unit_diff += 1
                n_edit += 1
                s_question_qty -= 1
            elif zeros_diff < 0:
                S[i] = '0'
                n_edit += 1
                zeros_diff += 1
                s_question_qty -= 1
            elif unit_diff < 0:
                S[i] = '1'
                n_edit += 1
                unit_diff += 1
                s_question_qty -= 1
        elif zeros_diff > 0 and T[i] == '1' and S[i] == '0':
            S[i] = '1'
            n_edit += 1
            zeros_diff -= 1


    s_zero_qty = s_unit_qty = 0
    t_zero_qty = t_unit_qty = 0
    n_mismatch = 0
    for i in range(len(S)):
        sch = S[i]
        tch = T[i]
        if sch != tch:
            n_mismatch += 1
        
        if sch == '0':
            s_zero_qty += 1
        elif sch == '1':
            s_unit_qty += 1
        
        if tch == '0':
            t_zero_qty += 1
        elif tch == '1':
            t_unit_qty += 1

    if s_zero_qty != t_zero_qty or s_unit_qty != t_unit_qty:
        return -1

    return n_mismatch // 2 + n_edit


if __name__ == '__main__':
    n = int(stdin.readline().strip())
    for i in range(1, n+1):
        S = stdin.readline().strip()
        T = stdin.readline().strip()
        print(f'Case {i}: {count_moves(S, T)}')
