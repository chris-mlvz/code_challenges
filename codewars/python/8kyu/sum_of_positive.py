def positive_sum(arr):
    return sum(i for i in arr if i > 0)


if __name__ == "__main__":
    assert(positive_sum([-1, 2, 3, 4, -5]) == 9)
