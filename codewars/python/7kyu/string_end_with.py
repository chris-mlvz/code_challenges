def solution(string: str, ending: str) -> bool:
    return string.endswith(ending)


if __name__ == '__main__':
    print(solution('abcde', 'cde'))
    print(solution('abcde', 'abc'))
    print(solution('abcde', ''))
