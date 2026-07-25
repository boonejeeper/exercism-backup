#include "trinary.h"
#include <string>
#include <iostream>
#include <cmath>
using namespace std;
namespace trinary {

// TODO: add your solution here
    int to_decimal(std::string num_str) {
        int result{};
        // for (int i = num_str.length() - 1; i >= 0; i--) {
        int str_length(num_str.length());
        for (int i = 0; i < str_length; i++) {
            int current = num_str[str_length - i - 1] - '0';
            if (current < 0 || current > 2) { return 0; }

            result += current * pow(3, i);
        }
        return result;
    }
}  // namespace trinary
