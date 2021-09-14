import 'dart:math';

void main(List<String> args) {
  print(apple('50'));
  print(apple(2));
}

String apple(dynamic a) {
  if (a.runtimeType == String) {
    a = int.parse(a);
  }
  a = pow(a, 2);
  if (a > 1000) {
    return "It's hotter than the sun!!";
  }
  return 'Help yourself to a honeycomb Yorkie for the glovebox.';
}
