import sys
from collections import namedtuple
import math


stdin = sys.stdin
URLRelevance = namedtuple('URLRelevance', 'url relevance')


def im_feeling_good(relevances):
    max_relevance = -math.inf
    sites = []
    for r in relevances:
        if r.relevance > max_relevance:
            max_relevance = r.relevance
    for r in relevances:
        if r.relevance == max_relevance:
            sites.append(r.url)
    return sites


if __name__ == '__main__':
    n = int(stdin.readline().strip())
    for i in range(1, n + 1):
        relevances = []
        for j in range(10):
            line = stdin.readline().strip().split()
            url, relevance = line
            r = URLRelevance(url, int(relevance))
            relevances.append(r)

        sites = im_feeling_good(relevances)        
        print(f'Case #{i}:')
        for site in sites:
            print(site)
