bool isDivisible(int n, int x, int y) =>
    n % x == 0 && n % y == 0 ? true : false;

void main(List<String> args) {
  print(isDivisible(12, 3, 4));
}
