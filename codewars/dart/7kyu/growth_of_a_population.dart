void main(List<String> args) {
  print(nbYear(1000, 2.0, 50, 1214));
  print(nbYear(1500000, 0.0, 10000, 2000000));
}

int nbYear(int p0, double percent, int aug, int p) {
  var years = 0;
  while (p0 < p) {
    p0 = (p0 + p0 * percent / 100 + aug).truncate();
    years++;
  }
  return years;
}
