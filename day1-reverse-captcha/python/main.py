import itertools
import sys
import unittest


def reverse_captcha(value):
    value, lookahead = itertools.tee(map(int, value))
    lookahead = itertools.cycle(lookahead)
    next(lookahead)

    return sum(
        a
        for a, b in zip(value, lookahead)
        if a == b
    )


def reverse_captcha2(value):
    value = [int(c) for c in value]
    lookahead = len(value) // 2

    return sum(
        a
        for idx, a in enumerate(value)
        if a == value[(idx + lookahead) % len(value)]
    )


class TestCase(unittest.TestCase):

    def test_reverse(self):
        self.assertEqual(reverse_captcha('1122'), 3)
        self.assertEqual(reverse_captcha('1111'), 4)
        self.assertEqual(reverse_captcha('1234'), 0)
        self.assertEqual(reverse_captcha('91212129'), 9)

    def test_reverse2(self):
        self.assertEqual(reverse_captcha2('1212'), 6)
        self.assertEqual(reverse_captcha2('1221'), 0)
        self.assertEqual(reverse_captcha2('123425'), 4)
        self.assertEqual(reverse_captcha2('123123'), 12)
        self.assertEqual(reverse_captcha2('12131415'), 4)


if __name__ == '__main__':
    print(reverse_captcha2(sys.argv[1]))