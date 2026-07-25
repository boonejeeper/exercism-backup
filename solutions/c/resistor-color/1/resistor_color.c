#include "resistor_color.h"
#include <stdlib.h>
#include <stddef.h>
#include <string.h> // for memcpy

int color_code(resistor_band_t band_color) {
  return band_color;
}

resistor_band_t* colors() {
  resistor_band_t *colorList = malloc(10 * sizeof(resistor_band_t));
  memcpy(colorList, (resistor_band_t[]){ BLACK, BROWN, RED, ORANGE, YELLOW, GREEN, BLUE, VIOLET, GREY, WHITE }, 10 * sizeof(resistor_band_t));
  return colorList;
}
