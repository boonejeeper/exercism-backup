#include "hexadecimal.h"
#include <cmath>
#include <iostream>
using namespace std;

namespace hexadecimal {

// TODO: add your solution here
    int convert(std::string num_str) {
        int result{};
        
        int str_length(num_str.length());
        for (int i = 0; i < str_length; i++) {
            int current = num_str[str_length - i - 1] - '0';
            if (current > 9) {
                current = current - 'a' + '0';
                if (current < 0 || current > 5) {
                    return 0;
                }
                current += 10;
            }

            
            result += current * pow(16, i);
        }
        return result;
        
    }

}  // namespace hexadecimal