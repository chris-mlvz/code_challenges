void main(List<String> args) {
  var s = 'aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz';
  print(printerError(s));
}

String printerError(String s) {
  var allowedLetters = 'abcdefghijklm'.runes;
  var errorsCount = 0;
  s.runes.forEach((rune) {
    if (!allowedLetters.contains(rune)) {
      errorsCount++;
    }
  });
  return '$errorsCount/${s.length}';
}
