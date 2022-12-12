import sys
import math

from collections import defaultdict, OrderedDict



stdin = sys.stdin


def best1_proposal(requirements, proposals):
    n_features = len(requirements)
    best_compliance_proposal = -1
    compliance_map = OrderedDict()    
    for proposal, (features_map, price) in proposals.items():
        n_comp_features = 0
        for feature in features_map.keys():
            if feature in requirements:
                n_comp_features += 1
        compliance = n_comp_features / float(n_features)
        if compliance > best_compliance_proposal:
            best_compliance_proposal = compliance
        compliance_map[proposal] = compliance

    #print('best', best_compliance_proposal)
    best_price = math.inf
    best_proposal = None
    for proposal, compliance in compliance_map.items():
        price = proposals[proposal][1]
        #print('curr compliance', compliance)
        if price < best_price and compliance == best_compliance_proposal:
            best_proposal = proposal

    return best_proposal


if __name__ == '__main__':
    n = 0
    while True:
        n += 1
        n_requirements, n_proposals = [int(x) for x in stdin.readline().strip().split()]
        if n_requirements == 0  and n_proposals == 0:
            break
        
        if n == 1:
            pass
        else:
            print()


        for j in range(n_requirements):
            stdin.readline()



        best_price = math.inf
        best_proposal = None
        compliance = -1
        for i in range(n_proposals):
            proposal = stdin.readline()
            price, n_features = stdin.readline().strip().split()
            price = float(price)
            n_features = int(n_features)

            for j in range(n_features):
                stdin.readline()

            if n_features > compliance or (n_features == compliance and price < best_price):
                compliance = n_features
                best_price = price
                best_proposal = proposal




        print(f'RFP #{n}')
        print(best_proposal, end='')
