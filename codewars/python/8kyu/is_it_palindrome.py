def is_palindrome(s: str):
    if s.lower()[::-1] == s.lower():
        return True
    return False


if __name__ == '__main__':
    print(is_palindrome('a'))
    print(is_palindrome('aba'))
    print(is_palindrome('Abba'))
    print(is_palindrome('malam'))
    print(is_palindrome('walter'))
    print(is_palindrome('kodok'))
    print(is_palindrome('Kasue'))
