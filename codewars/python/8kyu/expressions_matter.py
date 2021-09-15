import math


def expression_matter(a: int, b: int, c: int) -> int:
    x = a + b + c
    y = a * b * c
    z = (a + b) * c
    w = a * (b + c)
    return max(x, y, z, w)


if __name__ == '__main__':
    print(expression_matter(2, 1, 2))
    print(expression_matter(2, 1, 1))
    print(expression_matter(2, 2, 4))
    print(expression_matter(3, 3, 3))
    print(expression_matter(1, 1, 1))
    print(expression_matter(1, 2, 3))
    print(expression_matter(1, 3, 1))
    print(expression_matter(2, 2, 2))
    print(expression_matter(5, 1, 3))
    print(expression_matter(3, 5, 7))
    print(expression_matter(5, 6, 1))
    print(expression_matter(1, 6, 1))
    print(expression_matter(2, 6, 1))
    print(expression_matter(6, 7, 1))
    print(expression_matter(2, 10, 3))
    print(expression_matter(1, 8, 3))
    print(expression_matter(9, 7, 2))
    print(expression_matter(1, 1, 10))
    print(expression_matter(9, 1, 1))
    print(expression_matter(10, 5, 6))
    print(expression_matter(1, 10, 1))
