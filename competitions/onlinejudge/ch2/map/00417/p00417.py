import sys
import string




cin = sys.stdin




def _ordered_letters_permutations(s, p, letters, last_letter):
    if len(s) >= 5:
        return
    

    for i, ch in enumerate(letters[last_letter:]):
        tmp = s + ch
        p += [tmp]
        _ordered_letters_permutations(tmp, p, letters, last_letter + i + 1)


def generate_map():
    p = []
    _ordered_letters_permutations('', p, string.ascii_lowercase, 0)
    m = {}
    for i, word in enumerate(sorted(p, key=lambda x: (len(x), x))):
        m[word] = i+1 
    return m


word_map = generate_map()


def encode(word):
    return word_map.get(word, '0')


if __name__ == '__main__':
    for row in cin:
        word = row.strip()
        print(encode(word))
