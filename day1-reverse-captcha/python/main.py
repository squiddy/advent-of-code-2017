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


class TestCase(unittest.TestCase):

    def test_reverse(self):
        self.assertEqual(reverse_captcha('1122'), 3)
        self.assertEqual(reverse_captcha('1111'), 4)
        self.assertEqual(reverse_captcha('1234'), 0)
        self.assertEqual(reverse_captcha('91212129'), 9)


if __name__ == '__main__':
    print(reverse_captcha(sys.argv[1]))