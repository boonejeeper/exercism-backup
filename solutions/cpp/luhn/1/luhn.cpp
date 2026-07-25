#include "luhn.h"
#include <string>
#include <iostream>
using namespace std;

namespace luhn {

// TODO: add your solution here
    bool valid(string num_str) {

        int running_total{};
        int str_length(num_str.length());
        bool should_double{};
        int non_trivial_digits{};
        for (int i = 0; i < str_length; i++) {
            char target_char = num_str[str_length - i - 1];
            if (target_char == ' ') { 
                cout << "\t\t\tskipping space" << endl;
                continue;
            }
            int target_digit = target_char - '0';
            if (target_digit < 0 || target_digit > 9) {
                cout << "\t\t\tnot a digit" << endl;
                return false;
            }
            non_trivial_digits++;
            cout << "i: " << i << "\tdigit: " << target_digit << endl;
            if (should_double) {
                cout << "\t\t\tdoubling the digit" << endl;
                target_digit *= 2;
                if (target_digit > 9) {
                    target_digit -= 9;
                }
                cout << "i: " << i << "\tdigit: " << target_digit << endl;
            }
            cout << "i: " << i << "\tdigit: " << target_digit << endl;
            running_total += target_digit;
            should_double = !should_double;
        }
        
        return !(running_total == 0 && non_trivial_digits == 1) && running_total % 10 == 0;
    }

}  // namespace luhn
