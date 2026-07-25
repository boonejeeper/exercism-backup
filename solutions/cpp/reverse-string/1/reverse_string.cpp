#include "reverse_string.h"
#include <string>
#include <iostream>
namespace reverse_string {

// TODO: add your solution here

    std::string reverse_string(std::string text) {
        std::string reversed{""};
        for (int i = text.length() - 1; i >= 0; i--) {
            reversed = reversed + text.substr(i, 1);
        }
        return reversed;
    }
}  // namespace reverse_string
