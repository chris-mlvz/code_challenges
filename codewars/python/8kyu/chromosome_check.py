def chromosome_check(sperm: str):
    if 'Y' in sperm:
        return "Congratulations! You're going to have a son."
    return "Congratulations! You're going to have a daughter."

if __name__ == '__main__':
    print(chromosome_check('XY'))
    print(chromosome_check('XX'))
