#include <iostream>
using namespace std;

void num_print(int x) {
  if (x == 0) {
    cout << "\n Done!" << endl;
    // return;
  } else {
    cout << x << " ";
    num_print(x - 1);
  }
}

int main() {
  int x = 10;
  num_print(x);
  return 0;
}
