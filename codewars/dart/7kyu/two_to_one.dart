void main(List<String> args) {
  print(longest('aretheyhere', 'yestheyarehere'));
  print(longest('loopingisfunbutdangerous', 'lessdangerousthancoding'));
  print(longest('inmanylanguages', 'theresapairoffunctions'));
}

String longest(String a, String b) {
  var lettersList = [for (var rune in (a + b).runes) String.fromCharCode(rune)];
  lettersList.sort();
  return lettersList.toSet().join();
}
