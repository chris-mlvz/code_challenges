num sum(List<num> arr) => arr.fold(0, (prev, element) => prev + element);

void main(List<String> args) {
  print(sum([1, 2, 3, 4]));
}
