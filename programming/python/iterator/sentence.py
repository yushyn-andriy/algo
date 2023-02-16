import re
import reprlib
import collections


RE_WORD = re.compile(r'\w+')


class Sentence:
    def __init__(self, text):
        self.text = text
        self.words = RE_WORD.findall(text)
    

    def __getitem__(self, index):
        return self.words[index]

    def __len__(self):
        return len(self.words)


    def __repr__(self):
        return f'Sentence(text={reprlib.repr(self.text)}'



class Foo:
    def __iter__(self):
        pass


text = '''
interpreter works in the following order
1. iter if object doesnt realized tries find __getitem__ method and will create iterator which tries
to fetch elements in order starting from zero.


'''

s = Sentence(text)
print(s)
for w in s:
    print(w)




s = "ABC" # iterable
# under the hood doing the same as in the next example
for ch in s:
    print(ch)


ss = iter(s) # iterator
while True:
    try:
        r = next(ss)
        print(r)
    except StopIteration:
        del ss
        break



class Sentence1:
    def __init__(self, text):
        self.text = text
        self.words = RE_WORD.findall(text)


    def __iter__(self):
        return SentenceIterator(self.words)

    def __len__(self):
        return len(self.words)


    def __repr__(self):
        return f'Sentence(text={reprlib.repr(self.text)}'



class SentenceIterator:
    def __init__(self, words):
        self.words = words
        self.index = 0


    def __next__(self):
        try:
            word = self.words[self.index]
        except IndexError:
            raise StopIteration
        self.index += 1
        return word


    def __iter__(self):
        return self


s1 = Sentence1(text)
print(list(s1), list(s))


'''
Итерируемый обьект реализует метод __iter__ который возвращает итератор а итератор в свою
очередь реализует метод __iter__ который возвращает селф
а также __next__ который возвращает след обьект последовательности или возбуждает исключение 
StopIteration

Генератор - это итератор который порождает значение выражений переданных yield

'''




class Sentence2:
    def __init__(self, text):
        self.text = text
        self.words = RE_WORD.findall(text)


    def __iter__(self): # generators function when called returned object-generetor
        for w in self.words:
            yield w
        return 

    def __len__(self):
        return len(self.words)

    def __repr__(self):
        return f'Sentence(text={reprlib.repr(self.text)}'

print(list(Sentence2(text)))



def seq():
    yield 1
    yield 2
    return # генератор генерирует в таком случае StopIteration
    yield 3


print(list(seq()))

g = seq()

# print(next(g))
# print(next(g))
# print(next(g))
# print(next(g))


class Sentence3:
    def __init__(self, text):
        self.text = text


    def __iter__(self): # generators function when called returned object-generetor
        for m in RE_WORD.finditer(self.text):
            yield m.group()

    def __repr__(self):
        return f'Sentence(text={reprlib.repr(self.text)}'



print(list(Sentence3(text)))