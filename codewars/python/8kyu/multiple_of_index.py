def multiple_of_index(arr: list):
    newList = []
    for i, value in enumerate(arr):
        if i > 0 and value % i == 0:
            newList.append(value)
    return newList


if __name__ == '__main__':
    print(multiple_of_index([22, -6, 32, 82, 9, 25]))
    print(multiple_of_index([68, -1, 1, -7, 10, 10]))
    print(multiple_of_index([11, -11]))
    print(multiple_of_index([-56, -85, 72, -26, -14, 76, -
          27, 72, 35, -21, -67, 87, 0, 21, 59, 27, -92, 68]))
    print(multiple_of_index([28, 38, -44, -99, -13, -54, 77, -51]))
    print(multiple_of_index([-1, -49, -1, 67, 8, -60, 39, 35]))
