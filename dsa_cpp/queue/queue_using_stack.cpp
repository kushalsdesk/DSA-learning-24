#include <iostream>
#include <stack>
using namespace std;

int main() {

  int arr[] = {1, 2, 3, 4, 5};
  int size = sizeof(arr) / sizeof(arr[0]);
  stack<int> stk1;
  stack<int> stk2;

  for (int i = 0; i < size; i++) {
    stk1.push(arr[i]);
  }

  while (!stk1.empty()) {
    stk2.push(stk1.top());
    stk1.pop();
  }

  cout << "Queue Output using Stack >> ";
  while (!stk2.empty()) {
    cout << stk2.top() << " ";
    stk2.pop();
  }

  return 0;
}
