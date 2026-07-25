#include <iostream>
#include <string>

namespace log_line {
    // "[<LEVEL>]: <MESSAGE>"
std::string message(std::string line) {
    int spaceIndex = line.find(" ");
    std::string message = line.substr(spaceIndex + 1);
    return message;
}

std::string log_level(std::string line) {
    // return the log level
    int openBracketIndex = line.find("[");
    int closeBracketIndex = line.find("]");
    std::string logLevel = line.substr(openBracketIndex + 1, closeBracketIndex - openBracketIndex - 1);
    return logLevel;
}

std::string reformat(std::string line) {
    // return the reformatted message
    return message(line) + " (" + log_level(line) + ")";
}
}  // namespace log_line
