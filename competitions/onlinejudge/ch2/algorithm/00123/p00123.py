import sys
from collections import namedtuple
from copy import deepcopy
import os

KeywordInfo = namedtuple('KeywordInfo', 'word word_idx title_idx')


stdin = sys.stdin



def kwic(titles, stop_words):
    keywords_info = []
    index = []
    for title_idx, title in enumerate(titles):
        iterate_word = False
        end_idx = 0
        start_idx = 0
        for i in range(len(title)):
            if title[i] in [' ', os.linesep]:
                if iterate_word:
                    end_idx = i
                    word = title[start_idx:end_idx]
                    if word not in stop_words:
                        keywords_info.append(KeywordInfo(word, (start_idx, end_idx), title_idx))        
                    iterate_word = False
            elif iterate_word:
                continue
            else:
                iterate_word = True
                start_idx = i
        else:
            if iterate_word:
                end_idx = len(title)
                word = title[start_idx:end_idx]
                if word not in stop_words:
                    keywords_info.append(KeywordInfo(word, (start_idx, end_idx), title_idx))        
                iterate_word = False
    keywords_info.sort(key=lambda x: (x.word, x.title_idx, x.word_idx[0]))
    
    for info in keywords_info:
        title = list(deepcopy(titles[info.title_idx]))
        title[info.word_idx[0]:info.word_idx[1]] = list(info.word.upper())
        index.append(''.join(title))

    return index


if __name__ == '__main__':
    stop_words = set()
    titles = []
    reading_stop_words = True
    for line in stdin:
        line = line.strip(os.linesep).lower()
        if line == '::':
            reading_stop_words = False
            continue

        if reading_stop_words:
            stop_words.add(line)
        else:
            titles.append(line)
    index = kwic(titles, stop_words)
    for indexed in index:
        print(''.join(indexed))
