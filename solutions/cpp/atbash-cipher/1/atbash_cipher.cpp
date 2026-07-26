#include "atbash_cipher.h"
#include <string>
#include <iostream>
using namespace std;

namespace atbash_cipher {

int get_altered_char(int target) {
    if ((target >= 'A' && target <= 'Z') || (target >= 'a' && target <= 'z')) {
        if (target < 'a') {
            target = target - 'A' + 'a';
        }
        int result = 'z' - target + 'a';
        return result;
    } else if (target >= '0' && target <= '9') {
        return target;
    } else {
        return ' ';
    }
}
    
string encode(string text) {
    string result{};
    int num_encoded_characters{0};
    for (char target : text) {
        char new_char = get_altered_char(target);
        if (new_char == ' ') {
            continue;
        }
        if (num_encoded_characters % 5 == 0 && num_encoded_characters > 0) {
            result += ' ';
        }
        result += new_char;
        num_encoded_characters++;
    }
    cout << "result: " << result << endl;
    return result;
}

string decode(string text) {
    string result{};
    int num_encoded_characters{0};
    for (char target : text) {
        char new_char = get_altered_char(target);
        if (new_char == ' ') {
            continue;
        }
        result += new_char;
        num_encoded_characters++;
    }
    return result;
}

}  // namespace atbash_cipher
