#include <iostream>
using namespace std;

void num_print(int x, int var) {
  if (var == x + 1) {
    cout << "\n Done!" << endl;
    // return;
  } else {
    cout << var << " ";
    num_print(x, var + 1);
  }
}

int main() {
  int x = 10;
  int var = 1;
  num_print(x, var);
  return 0;
}
