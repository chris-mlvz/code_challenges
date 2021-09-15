import 'dart:math';

void main(List<String> args) {
  print(expressionMatter(2, 1, 2));
  print(expressionMatter(2, 1, 1));
  print(expressionMatter(2, 2, 4));
  print(expressionMatter(3, 3, 3));
  print(expressionMatter(1, 1, 1));
  print(expressionMatter(1, 2, 3));
  print(expressionMatter(1, 3, 1));
  print(expressionMatter(2, 2, 2));
  print(expressionMatter(5, 1, 3));
  print(expressionMatter(3, 5, 7));
  print(expressionMatter(5, 6, 1));
  print(expressionMatter(1, 6, 1));
  print(expressionMatter(2, 6, 1));
  print(expressionMatter(6, 7, 1));
  print(expressionMatter(2, 10, 3));
  print(expressionMatter(1, 8, 3));
  print(expressionMatter(9, 7, 2));
  print(expressionMatter(1, 1, 10));
  print(expressionMatter(9, 1, 1));
  print(expressionMatter(10, 5, 6));
  print(expressionMatter(1, 10, 1));
}

int expressionMatter(a, b, c) {
  return <int>[a * b * c, a + b + c, (a + b) * c, a * (b + c)].reduce(max);
}
