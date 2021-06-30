int multiply(int a, int b) {
  return a * b;
}

// int multiply(int a, int b) => a * b;

void main(List<String> args) {
  print(multiply(3, 4));
  assert(multiply(3, 4) == 12);
}
