import sys

stdin = sys.stdin



def pig_latin(s):
    encoded = ''
    vowels = 'aeiou'
    flag = False
    is_vowel = False
    first_char = None
    for ch in s:
        if flag:
            if not ch.isalpha():
                if is_vowel:
                    encoded += 'ay' + ch
                if not is_vowel:
                    encoded += first_char + 'ay' + ch
                flag = False
            else:
                encoded += ch
        else:
            if ch.isalpha():
                if ch.lower() in vowels:
                    is_vowel = True
                    encoded += ch
                else:
                    is_vowel = False
                    first_char = ch
                flag = True
            else:
                encoded += ch

    return encoded


if __name__ == '__main__':
    for row in stdin:
        print(pig_latin(row), end='')
