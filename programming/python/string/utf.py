import encodings
from pathlib import Path
import codecs

folder = Path('./encodings')

s = """
Jeg hører deg I\'ve been living her for a long time,
but it seems like I want to go home.
It is not my country, I understood it yesterday for sure.
Привет мир
"""


if __name__ == '__main__':
    if not folder.exists():
        folder.mkdir()
    for name, enc in encodings._aliases.items():
        filename = f'{name}__{enc}'
        path = folder / Path(filename)
        # encoded = s.encode(enc, errors='ignore')
        try:
            encoded = codecs.encode(s, enc, errors='replace')
            path.write_bytes(encoded)
        except Exception as e:
            print(e)
