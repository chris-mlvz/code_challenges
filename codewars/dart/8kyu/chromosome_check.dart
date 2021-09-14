void main(List<String> args) {
  print(chromosome_check('XY'));
  print(chromosome_check('XX'));
}

String chromosome_check(String sperm) {
  if (sperm.contains('Y')) {
    return "Congratulations! You're going to have a son.";
  }
  return "Congratulations! You're going to have a daughter.";
}
