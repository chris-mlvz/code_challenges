const dict_grade = <int, String>{90: 'A', 80: 'B', 70: 'C', 60: 'D', 0: 'F'};

String getGrade(int a, int b, int c) {
  var average = (a + b + c) / 3;
  for (var key in dict_grade.keys) {
    if (average >= key) {
      return dict_grade[key]!;
    }
  }
  throw Exception();
}

void main(List<String> args) {
  print(getGrade(80, 90, 75));
}
