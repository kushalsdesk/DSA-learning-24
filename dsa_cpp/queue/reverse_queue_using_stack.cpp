#include <iostream>
#include <queue>
#include <stack>
using namespace std;

int main() {

  int arr[] = {1, 2, 3, 4, 5};
  int size = sizeof(arr) / sizeof(arr[0]);
  queue<int> q;
  stack<int> stk;
  for (int i = 0; i < size; i++) {
    q.push(arr[i]);
  }

  while (!q.empty()) {
    stk.push(q.front());
    q.pop();
  }
  cout << "Queue Output using Stack >> ";
  while (!stk.empty()) {
    cout << stk.top() << " ";
    stk.pop();
  }

  return 0;
}
