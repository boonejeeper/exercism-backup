#include "difference_of_squares.h"
#include <cmath>
#include <iostream>
using namespace std;

namespace difference_of_squares {

// TODO: add your solution here
    int square_of_sum(int num) {
        int result{};
        for (int i = 1; i <= num; i++) {
            result += i;
        }
        return result * result;
    }
    
    int sum_of_squares(int num) {
        int result{};
        for (int i = 1; i <= num; i++) {
            result += i * i;
        }
        return result;
    }
    
    int difference(int num) {
        return square_of_sum(num) - sum_of_squares(num);
    }

}  // namespace difference_of_squares
