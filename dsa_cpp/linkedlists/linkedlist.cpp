#include <cstddef>
#include <cstdlib>
#include <iostream>
using namespace std;

struct Node {
  int data;
  Node *next;
};
Node *head = NULL;
void display_list() {
  Node *temp = head;
  cout << "::::::::...LIST...::::::::" << endl;

  // Check for empty list
  if (head == NULL) {
    cout << "List empty" << endl;
    return;
  }

  // Traverse the list
  while (temp != NULL) {
    cout << temp->data << "->";
    temp = temp->next;
  }
  cout << "\n::::::::::::::::::::::::::" << endl;
}

void insert_at_first(int data) {
  Node *ptr = new Node();
  ptr->data = data;
  ptr->next = head;
  head = ptr;
}
void delete_at_first() {
  if (head == NULL) {
    cout << "\n::::::::::::::::::::::::::" << endl;
    cout << "\nYou can't Delete whats not there bro >_<!" << endl;
    cout << "\n::::::::::::::::::::::::::" << endl;
    return;
  } else {
    cout << "\n::::::::::::::::::::::::::" << endl;
    cout << "\n >>>>>>>DELEDED 1st NODE <<<<<<<<" << endl;
    cout << "\n::::::::::::::::::::::::::" << endl;
    head = head->next;
  }
}
void insert_at_last(int data) {
  Node *ptr = new Node();
  ptr->data = data;
  ptr->next = NULL;
  Node *temp = head;
  if (head == NULL) {
    head = ptr;
  } else {
    while (temp->next != NULL) {
      temp = temp->next;
    }
    temp->next = ptr;
  }
}
void delete_at_last() {
  Node *temp, *prev_temp;
  if (head == NULL) {
    cout << "\n::::::::::::::::::::::::::" << endl;
    cout << "\nYou can't Delete whats not there bro >_<!" << endl;
    cout << "\n::::::::::::::::::::::::::" << endl;
    return;
  } else if (head->next == NULL) {
    cout << "\n::::::::::::::::::::::::::" << endl;
    cout << "\n >>>>>>>DELEDED THE ONLY NODE <<<<<<<<" << endl;
    cout << "\n::::::::::::::::::::::::::" << endl;
    temp = head;
    head = NULL;
    free(temp);
  } else {
    temp = head;
    while (temp->next != NULL) {
      prev_temp = temp;
      temp = temp->next;
    }
    prev_temp->next = NULL;
    cout << "\n::::::::::::::::::::::::::" << endl;
    cout << "\n >>>>>>>DELEDED THE Last NODE <<<<<<<<" << endl;
    cout << "\n::::::::::::::::::::::::::" << endl;
  }
}

int main() {
  int input, choice;
  while (1) {
    cout << "\n::::::Enter Choice:::::: \n 0-> Exit \n 1-> Insert at First \n "
            "2-> Insert at "
            "Last \n 3-> Display LinkedList\n 4-> Delete First Node\n 5-> "
            "Delete Last Node \n >>>>> ";
    cin >> choice;
    switch (choice) {
    case 1:
      cout << "Enter Data :: ";
      cin >> input;
      insert_at_first(input);
      break;

    case 2:
      cout << "Enter Data :: ";
      cin >> input;
      insert_at_last(input);
      break;
    case 3:
      display_list();
      break;
    case 0:
      cout << "Bye!" << endl;
      exit(0);
      break;
    case 4:
      delete_at_first();
      break;
    case 5:
      delete_at_last();
      break;
    default:
      cout << "Invalid Choice" << endl;
      break;
    }
  }
  return 0;
}
