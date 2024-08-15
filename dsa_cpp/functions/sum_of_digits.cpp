#include <iostream>
using namespace std;

int sum(int num) {
  int sum = 0;
  while (num != 0) {
    sum += (num % 10);
    num = num / 10;
  }
  return sum;
}

int main() {
  int x;
  cout << " Enter no >>";
  cin >> x;
  cout << "Result :" << sum(x) << endl;
  return 0;
}
