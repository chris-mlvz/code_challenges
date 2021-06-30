// My solution
String solution(String str) => str.split('').reversed.join();

// Best Practice
// dynamic solution(String str) =>
//     String.fromCharCodes(str.runes.toList().reversed);

void main(List<String> args) {
  print(solution('world'));
  // assert(solution('world') == 'dlrow');
}
