import sys
from copy import deepcopy


stdin = sys.stdin




def next_permutatation(arr):
    
    for i in reversed(range(len(arr) - 1)):
        if arr[i] < arr[i + 1]:
            break
    else:
        return False

    j = next( j for j in reversed(range(i+1, len(arr))) if arr[i] < arr[j] )

    arr[i], arr[j] = arr[j], arr[i]
    arr[i+1:] = reversed(arr[i+1:])

    return True


def prev_permutatation(arr):
    
    for i in reversed(range(len(arr) - 1)):
        if arr[i] > arr[i + 1]:
            break
    else:
        return False

    arr[i+1:] = reversed(arr[i+1:])

    j = next( j for j in range(i+1, len(arr)) if arr[i] > arr[j] )

    arr[i], arr[j] = arr[j], arr[i]

    return True



def minimum_absolute_distance(arr):
    m = float('inf')
    for i in range(len(arr) - 1):
        diff = abs(ord(arr[i]) - ord(arr[i+1]))
        if diff < m:
            m = diff
    return m



if __name__ == '__main__':

    while True:
        row = stdin.readline().strip()
        if not row:
            break

        origin = list(row)

        min_dist = float('-inf')
        result = ''
        for i in range(10):
            if prev_permutatation(origin) is False:
                break

        for i in range(21):
            dist = minimum_absolute_distance(origin)
            if dist > min_dist:
                min_dist =  dist
                result = deepcopy(origin)
            if next_permutatation(origin) is False:
                break

        print(''.join(result) + str(min_dist))
