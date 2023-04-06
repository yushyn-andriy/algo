import sys
from enum import Enum


stdin = sys.stdin


ch_to_reverse_map = {
    'A': 'A',
    'E': '3',
    'H': 'H',
    'I': 'I',
    'J': 'L',
    'L': 'J',
    'M': 'M',
    'O': 'O',
    'S': '2',
    'T': 'T',
    'U': 'U',
    'V': 'V',

    'W': 'W',
    'X': 'X',
    'Y': 'Y',
    'Z': '5',
    
    '1': '1',
    '2': 'S',
    '3': 'E',
    '5': 'Z',
    '8': '8',
}



class SType(Enum):
    not_palindrome = '-- is not a palindrome.'
    regular_palindrome = '-- is a regular palindrome.'
    mirrored_string = '-- is a mirrored string.'
    mirrored_palindrome = '-- is a mirrored palindrome.'




def get_string_type(s):
    i, j = 0, len(s) - 1

    regular_palindrome = True
    mirrored_string = True
    while i <= j:
        if not regular_palindrome and not mirrored_string:
            return SType.not_palindrome.value
        
        if s[i] != s[j]:
            regular_palindrome = False

        ch1 = ch_to_reverse_map.get(s[i])
        ch2 = ch_to_reverse_map.get(s[j])

        if (
            ch1 is None or
            ch2 is None or
            ch1 != s[j] or
            ch2 != s[i]
        ):
            mirrored_string = False

        i+=1
        j-=1

    res = {
        (True, True): SType.mirrored_palindrome.value,
        (True, False): SType.regular_palindrome.value,
        (False, True): SType.mirrored_string.value,
        (False, False): SType.not_palindrome.value,
    }

    return res[(regular_palindrome, mirrored_string)]



if __name__ == '__main__':
    for row in stdin:
        row = row.strip()
        print(f'{row} {get_string_type(row)}')
        print()
