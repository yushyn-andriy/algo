from p11683 import Engraver


def test_engraver():
    engraver = Engraver(5, 5, [1, 3, 2, 4, 1])
    assert engraver.surface == [
        list('x'*5)
        for _ in range(5)
    ]
    engraver.engrave()
    assert engraver.surface == [
        list('ooooo'),
        list('oooxo'),
        list('oxoxo'),
        list('oxxxo'),
        list('xxxxx'),
    ]

