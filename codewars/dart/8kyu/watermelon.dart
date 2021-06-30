bool divide(int w) => w.isEven && w > 2 ? true : false;

void main(List<String> args) {
  assert(divide(4) == true);
  assert(divide(2) == false);
  assert(divide(5) == false);
  assert(divide(88) == true);
  assert(divide(100) == true);
  assert(divide(67) == false);
  assert(divide(90) == true);
  assert(divide(10) == true);
  assert(divide(99) == false);
  assert(divide(32) == true);
}
