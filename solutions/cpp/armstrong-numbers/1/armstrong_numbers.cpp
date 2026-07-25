#include "armstrong_numbers.h"
#include <string>
#include <iostream>
#include <cmath>
using namespace std;

namespace armstrong_numbers {

// TODO: add your solution here
    bool is_armstrong_number(int num) {
        string num_str = to_string(num);
        int num_digits = num_str.length();

        int result{};
        for (int i = 0; i < num_digits; i++) {
            result += pow(num_str[i] - '0', num_digits);
        }
        
        return result == num;
    }
}  // namespace armstrong_numbers
