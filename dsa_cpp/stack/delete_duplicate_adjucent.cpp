#include <iostream>
#include <stack>
#include <string>
using namespace std;

int main() {
  stack<char> stack1;
  stack<char> stack2;
  string input = "kunnaal";
  string ans = "";

  for (char c : input) {
    if (!stack1.empty() && stack1.top() == c) {
      stack1.pop();
    } else {
      stack1.push(c);
    }
  }
  char c;
  while (!stack1.empty()) {
    c = stack1.top();
    stack1.pop();
    stack2.push(c);
  }
  while (!stack2.empty()) {
    ans += stack2.top();
    stack2.pop();
  }
  cout << "Output :: " << ans << endl;
  return 0;
}
