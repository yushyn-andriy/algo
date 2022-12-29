import sys
from collections import namedtuple, OrderedDict

NAME = BASE = 0
SIZE, DIM, BOUNDS = 1, 2, 3
stdin = sys.stdin

Bound = namedtuple('Bound', 'lower upper')
ArrayInfo = namedtuple('ArrayInfo', 'name base size dim bounds coef')
ArrayRef = namedtuple('ArrayRef', 'name indexes')


def read_list():
   return stdin.readline().strip().split()


def convert_to(l, converter):
    return [converter(x) for x in l]


def calc_coef(base, size, dim, bounds):
    CD = size
    coef = [0] * (len(bounds) + 1)
    coef[-1] = CD
    coef[0] = base
    for i in range(len(bounds) - 1, -1, -1):
        bound = bounds[i]
        ci = coef[i+1] * (bound.upper - bound.lower + 1)
        if i > 0:
            coef[i] = ci
    for i, c in enumerate(coef[1:]):
        coef[0] -= bounds[i].lower * c
    return coef


def array_info_from_list(l):
    name = l[NAME]
    nf = convert_to(l[1:], int)
    base = nf[BASE]
    size = nf[SIZE]
    dim = nf[DIM]
    bounds = []
    for i in range(BOUNDS, BOUNDS + dim * 2 - 1, 2):
        bounds.append(Bound(nf[i], nf[i+1]))
    coef = calc_coef(base, size, dim, bounds)
    info = ArrayInfo(
        name, base, size, dim, bounds, coef)
    return info


def array_ref_from_list(l):
    name = l[0]
    idxs = convert_to(l[1:], int)
    return ArrayRef(name, idxs)


def calculate_address(arr_map, ref):
    coef = arr_map[ref.name].coef
    address = coef[0]
    for i, c in enumerate(coef[1:]):
        address += c * ref.indexes[i]
    return address


if __name__ == '__main__':
    n_arr, n_ref = convert_to(read_list(), int)
    arr_map = {}
    refs = [None] * n_ref
    # read and save input
    for i in range(n_arr):
        info = array_info_from_list(read_list())
        arr_map[info.name] = info
    for j in range(n_ref):
        refs[j] = array_ref_from_list(read_list())
        r = refs[j]
        print(f'{r.name}[{", ".join(convert_to(r.indexes, str))}] = {calculate_address(arr_map, r)}')

