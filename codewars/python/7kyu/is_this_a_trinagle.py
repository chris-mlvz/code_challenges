def is_triangle(a: int, b: int, c: int) -> bool:
    sides = [a, b, c]
    if any(i <= 0 for i in sides):
        return False
    maxSide = max(sides)
    sides.remove(maxSide)
    if sum(sides) <= maxSide:
        return False
    return True


if __name__ == '__main__':
    print(is_triangle(1, 2, 2))
    print(is_triangle(7, 2, 2))
    print(is_triangle(1, 2, 3))
    print(is_triangle(1, 3, 2))
    print(is_triangle(3, 1, 2))
    print(is_triangle(5, 1, 2))
    print(is_triangle(1, 2, 5))
    print(is_triangle(2, 5, 1))
    print(is_triangle(4, 2, 3))
    print(is_triangle(5, 1, 5))
    print(is_triangle(2, 2, 2))
    print(is_triangle(-1, 2, 3))
    print(is_triangle(1, -2, 3))
    print(is_triangle(1, 2, -3))
    print(is_triangle(0, 2, 3))
