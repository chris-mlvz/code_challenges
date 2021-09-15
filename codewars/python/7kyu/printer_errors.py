def printer_error(s: str) -> str:
    allowedLetters = 'abcdefghijklm'
    errorsCount = 0
    for letter in s:
        if letter not in allowedLetters:
            errorsCount += 1
    return f'{errorsCount}/{len(s)}'


if __name__ == '__main__':
    s = "aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"
    print(printer_error(s))
