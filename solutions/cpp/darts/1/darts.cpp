#include "darts.h"
#include <iostream>
#include <cmath>
namespace darts {

    float distance(float x1, float y1, float x2, float y2) {
        return sqrt(pow(x2 - x1, 2) + pow(y2 - y1, 2));
    }
// TODO: add your solution here
    int score(float x, float y) {
        float from_center = distance(x, y, 0.0f, 0.0f);
        //std::cout << from_center << std::endl;

        if (from_center <= 1) {
            return 10;
        } else if (from_center <= 5) {
            return 5;
        } else if (from_center <= 10) {
            return 1;
        }
        return 0;
    }

}  // namespace darts
