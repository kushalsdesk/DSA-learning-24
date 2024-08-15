#include <iostream>
using namespace std;

void print_in_range(int start, int end) {
  for (int i = start; i <= end; i++) {
    cout << i << " ";
  }
}

int main() {
  int x, y;
  cout << "Enter Start Point: ";
  cin >> x;
  cout << "Enter End Point: ";
  cin >> y;
  print_in_range(x, y);
  return 0;
}
