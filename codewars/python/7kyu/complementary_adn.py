dna_dict = {'A': 'T', 'T': 'A', 'G': 'C', 'C': 'G'}


def DNA_strand(dna):
    return ''.join(dna_dict[i] for i in dna)


if __name__ == '__main__':
    print(DNA_strand("AAAA"))
    print(DNA_strand("ATTGC"))
    print(DNA_strand("GTAT"))
