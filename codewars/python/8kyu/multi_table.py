def multi_table(number: int):
    res = ''
    for i in range(1, 11):
        res += f'{i} * {number} = {i * number}\n'
    res = res.strip()
    return res


if __name__ == '__main__':
    print(multi_table(5))
    print(multi_table(1))
