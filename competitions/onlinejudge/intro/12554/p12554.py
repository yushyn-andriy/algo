import sys



stdin = sys.stdin
song = 'Happy birthday to you Happy birthday to you Happy birthday to Rujia Happy birthday to you'.split()




if __name__ == '__main__':
    n = int(stdin.readline().strip())
    names = []

    for _ in range(n):
        names.append(stdin.readline().strip())
    
    i = 0
    everyone_song = False
    while True:
        word_idx = i % len(song)
        name_idx = i % len(names)
        
        print(f'{names[name_idx]}: {song[word_idx]}')
        if i == len(names) - 1:
            everyone_song = True
        if everyone_song and word_idx == len(song) - 1:
            break
        i += 1
 