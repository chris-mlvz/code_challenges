def longest(a1: str, a2: str) -> str:
    return ''.join(sorted(set(i for i in a1 + a2)))


if __name__ == '__main__':
    print(longest("aretheyhere", "yestheyarehere"))
    print(longest("loopingisfunbutdangerous", "lessdangerousthancoding"))
    print(longest("inmanylanguages", "theresapairoffunctions"))
