void main(List<String> args) {
  print(isPalindrome('a'));
  print(isPalindrome('aba'));
  print(isPalindrome('Abba'));
  print(isPalindrome('hello'));
  print(isPalindrome('Bob'));
  print(isPalindrome('Madam'));
  print(isPalindrome('AbBa'));
  print(isPalindrome(''));
}

bool isPalindrome(String str) {
  final reversedStr =
      String.fromCharCodes(str.toLowerCase().runes.toList().reversed);
  if (reversedStr == str.toLowerCase()) {
    return true;
  }
  return false;
}
