#include "eliuds_eggs.h"
#include <iostream>
#include <string>
using namespace std;

namespace chicken_coop {

int positions_to_quantity(unsigned int number) {
  string number_string{};
  // cout << number << endl;

  unsigned int mask{1};
  do {
    bool bit_on = number & mask;
    number_string = (bit_on ? "1" : "0") + number_string;
    mask = mask << 1;
  } while (number >= mask);

  int egg_count{};

  for (int i = 0; i < (int)number_string.length(); i++) {
    if (number_string[i] == '1') {
      egg_count++;
    }
  }

  return egg_count;
}

} // namespace chicken_coop
