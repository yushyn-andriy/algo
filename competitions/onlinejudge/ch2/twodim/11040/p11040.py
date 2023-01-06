import sys
import os


stdin = sys.stdin
N = 9


def print_m(m):
    m.reverse()
    print(f'{os.linesep}'.join([' '.join([str(x) for x in row]) for row in m]))


def read_matrix():
    m = [[0] * N for _ in range(N)]
    start_idx = 4
    end_idx = 4
    for r in range(8, -1, -2):
        row = sys.stdin.readline().strip()
        seq = [int(x) for x in row.split()]

        c = start_idx
        for num in seq:
            m[r][c] = num
            c += 2

        start_idx -= 1
        end_idx += 1   
    return m


def solve(m):
    # indexes known elements 
    # on the current sliding window
    f_idx, s_idx = 0, 2

    # intitial sliding window steps and result matrix
    ws = 4
    result = [[] for _ in range(9)]
    
    res_idx = 0
    last_element_upper_row = None
    for r in range(0, 8, 2):
        cf_idx, cs_idx = f_idx, s_idx
        last_element_bottom_row = None
        for _ in range(ws):
            # getting known vars
            fn = m[r][cf_idx]
            sn = m[r][cs_idx]
            thn = m[r+2][cf_idx+1]

            # solving equation            
            x = (thn - (fn + sn)) // 2
            y1 = fn + x
            y2 = sn + x

            # form result matrix
            result[res_idx] += [fn, x]
            result[res_idx + 1] += [y1, y2]
    
            # save for append later
            # in order to avoid duplicates 
            last_element_bottom_row = sn
            last_element_upper_row = thn

            # moving sliding window
            cf_idx = cs_idx
            cs_idx += 2

        # insert last element of the row
        result[res_idx].append(last_element_bottom_row)

        # change level on the result 
        res_idx += 2

        # moving sliding windows on upper 
        # stair one to the right
        f_idx += 1
        s_idx += 1

        # decrement sliding window steps
        ws -= 1
    
    # insert top element of the pyramid
    result[res_idx].append(last_element_upper_row)

    return result

if __name__ == '__main__':
    cases = int(sys.stdin.readline().strip())
    for _ in range(cases):
        m = read_matrix()
        r = solve(m)
        print_m(r)
