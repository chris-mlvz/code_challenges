int stringToNumber(String str) => int.parse(str);

void main(List<String> args) {
  print(stringToNumber('7'));
  assert(stringToNumber('7') == 7);
}
