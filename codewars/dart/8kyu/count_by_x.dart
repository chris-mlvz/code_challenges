List<int> countBy(int x, int n) =>
    List.generate(n, (index) => (index + 1) * x);

void main(List<String> args) {
  print(countBy(3, 7));
}
