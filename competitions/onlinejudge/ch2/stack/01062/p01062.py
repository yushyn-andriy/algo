import sys


cin = sys.stdin


def solve(cs):
    stacks = []
    for i, ch in enumerate(cs):
        if not stacks:
            stacks.append([ch])
            continue


        found = False
        mini = -1
        for j in range(len(stacks)):
            if ch <= stacks[j][-1]:
                found = True
                if mini == -1:
                    mini = j
                elif stacks[mini][-1] > stacks[j][-1]:
                    mini = j


        if not found:
            stacks.append([ch])
        else:
            stacks[mini].append(ch)

    return len(stacks)



if __name__ == '__main__':
    for i, row in enumerate(cin):
        cs = row.strip()
        if cs == 'end':
            break
        s = 'Case {i}: {answer}'.format(i=i+1, answer=solve(cs))
        print(s)
