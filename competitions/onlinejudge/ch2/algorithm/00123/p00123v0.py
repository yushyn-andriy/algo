import sys
from collections import namedtuple
from copy import deepcopy

KeywordInfo = namedtuple('KeywordInfo', 'word word_idx title_idx')


stdin = sys.stdin



def kwic(titles, stop_words):
    keywords_info = []
    index = []
    for title_idx, title in enumerate(titles):
        for word_idx, word in enumerate(title):
            if word not in stop_words:
                keywords_info.append(KeywordInfo(word, word_idx, title_idx))        
    keywords_info.sort(key=lambda x: (x.word, x.title_idx, x.word_idx))
    
    for info in keywords_info:
        title = deepcopy(titles[info.title_idx])
        title[info.word_idx] = title[info.word_idx].upper()
        index.append(title)

    return index


if __name__ == '__main__':
    stop_words = set()
    titles = []
    reading_stop_words = True
    for line in stdin:
        if  '::' in line:
            reading_stop_words = False
            continue

        if reading_stop_words:
            stop_words.add(line)
        else:
            titles.append(line.lower().split())
    index = kwic(titles, stop_words)
    for indexed in index:
        print(' '.join(indexed))
