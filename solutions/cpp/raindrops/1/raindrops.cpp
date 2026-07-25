#include "raindrops.h"
#include <string>

namespace raindrops {

// TODO: add your solution here
    std::string convert(int dropNumber) {
        std::string result{""};
        if (dropNumber % 3 == 0) {
            result += "Pling";
        }
        if (dropNumber % 5 == 0) {
            result += "Plang";
        }
        if (dropNumber % 7 == 0) {
            result += "Plong";
        }
        if (result.length() == 0) {
            result = std::to_string(dropNumber);
        }
        return result;
    }

}  // namespace raindrops
