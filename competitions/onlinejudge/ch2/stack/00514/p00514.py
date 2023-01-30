import sys



stdin = sys.stdin


def is_possible(perm, n):
    coaches = [x for x in range(n, 0, -1)]
    stack = []

    for pc in perm:
        if not stack and not coaches:
            return 'No'
        elif stack and coaches:
            if stack[-1] == pc:
                stack.pop()
                continue
            elif coaches[-1] == pc:
                coaches.pop()
                continue
            else:
                while coaches:
                    cc = coaches.pop()
                    if cc == pc:
                        break
                    stack.append(cc)
                if cc != pc:
                    return 'No'
        elif not stack and coaches:
            while coaches:
                cc = coaches.pop()
                if cc == pc:
                    break
                stack.append(cc)
        elif stack and not coaches:
            while stack:
                sc = stack.pop()
                if sc == pc:
                    break
                else:
                    return 'No'
    return 'Yes' 



if __name__ == '__main__':
    for nr in stdin:
        n = int(nr.strip())
        if n == 0:
            break
        for row in stdin:
            row = row.strip()
            if row == '0':
                break
            print(is_possible([int(x) for x in row.split()], n))

        print()
