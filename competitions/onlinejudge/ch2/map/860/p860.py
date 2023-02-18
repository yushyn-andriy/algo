import sys
import re
import math


stdin = sys.stdin


END_OF_TEXT = 'END_OF_TEXT'
END_OF_INPUT = 'END_OF_INPUT'
RE_WORD = re.compile(r'[^,.:;!?"()\s]+')



def calculate(text):
    freq = {}
    for lm, word in enumerate(RE_WORD.finditer(text)):
        word = word.group()
        if word not in freq:
            freq[word] = 0
        freq[word] += 1
    lm +=1


    # calculating entropy
    Et = 0.0
    for _, p in freq.items():
        Et += p * (math.log10(lm) - math.log10(p))
    Et = Et / lm

    # calculating max entropy
    Emax = math.log10(lm)

    # calculating relative entropy
    Erel = (round(Et / Emax, 2)) * 100

    return lm, round(Et, 1), int(Erel)


if __name__ == '__main__':
    for r1 in stdin:
        text = ''
        if END_OF_INPUT in r1:
            break 

        text += r1.lower()
        for r2 in stdin:
            if END_OF_TEXT in r2:
                break
            text += r2.lower()

        a, b, c = calculate(text)

        # print(a, b, c)
        print('%d %.1f %d' % (a, b, c))
