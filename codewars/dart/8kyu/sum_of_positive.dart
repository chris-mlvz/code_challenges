void main(List<String> args) {
  assert(positiveSum([1, -4, 7, 12]) == 20);
}

int positiveSum(List<int> arr) {
  return arr
      .where((element) => !element.isNegative)
      .fold(0, (prev, element) => prev + element);
}
