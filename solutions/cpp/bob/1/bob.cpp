#include "bob.h"
#include <iostream>
#include <string>
using namespace std;

namespace bob {

    bool is_lowercase_letter(char target) {
        return target > 'a' && target < 'z';
    }
    
    bool is_uppercase_letter(char target) {
        return target > 'A' && target < 'Z';
    }

    bool is_whitespace(char target) {
        return target == ' ' || target == '\t' || target == '\n' || target == '\r';
    }
    
    bool is_alpha(char target) {
        return is_lowercase_letter(target) || is_uppercase_letter(target);
    }

    
// TODO: add your solution here
    string hey(string message) {
        bool has_alpha{false};
        bool is_all_caps{true};
        bool is_question{false};
        bool is_empty{true};
        for (size_t i = 0; i < message.length(); i++) {
            char target = message[i];
            cout << "target: " << target << endl;
            if (target == '?') {
                cout << "Found question mark" << endl;
                is_question = true;
            }
            if ((!has_alpha || is_question) && is_alpha(target)) {
                has_alpha = true;
                is_question = false;
                cout << "Found alpha, resetting question mark flag";
            }
            if (has_alpha && is_all_caps && !is_uppercase_letter(target) && is_lowercase_letter(target)) {
                cout << "Found a character that isn't a capital" << endl;
                is_all_caps = false;
            }
            if (is_empty && !is_whitespace(target)) {
                cout << "found a character that isn't whitespace" << endl;
                is_empty = false;
            }
        }

        is_all_caps = is_all_caps && has_alpha;

        cout << "Done parsing.\nis_empty:\t\t" << is_empty << "\nis_all_caps:\t\t" << is_all_caps << "\nis_question:\t\t" << is_question << endl;
        if (is_empty) {
            return "Fine. Be that way!";
        } else if (is_question && !is_all_caps) {
            return "Sure.";
        } else if (is_question && is_all_caps) {
            return "Calm down, I know what I'm doing!";
        } else if (is_all_caps) {
            return "Whoa, chill out!";
        } else {
            return "Whatever.";
        }
    }

}  // namespace bob
