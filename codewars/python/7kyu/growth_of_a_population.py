def nb_year(p0, percent, aug, p):
    years = 0
    while p0 < p:
        p0 = int(p0 + p0 * percent/100 + aug)
        years += 1
    return years


if __name__ == '__main__':
    print(nb_year(1000, 2.0, 50, 1214))
    print(nb_year(1500000, 0.0, 10000, 2000000))
