int squareSum(List numbers) {
  return numbers.map((e) => e * e).fold(0, (prev, element) => element + prev);
}

void main(List<String> args) {
  print(squareSum([]));
}
