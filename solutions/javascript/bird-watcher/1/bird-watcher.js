// @ts-check
//
// The line above enables type checking for this file. Various IDEs interpret
// the @ts-check directive. It will give you helpful autocompletion when
// implementing this exercise.

/**
 * Calculates the total bird count.
 *
 * @param {number[]} birdsPerDay
 * @returns {number} total bird count
 */
export function totalBirdCount(birdsPerDay) {
  var count = 0;
  for (let day = 0; day < birdsPerDay.length; day++) {
    count += birdsPerDay[day];
  }
  return count;
}

/**
 * Calculates the total number of birds seen in a specific week.
 *
 * @param {number[]} birdsPerDay
 * @param {number} week
 * @returns {number} birds counted in the given week
 */
export function birdsInWeek(birdsPerDay, week) {
  var count = 0;
  for (let day = 7 * (week-1); day < week * 7; day++) {
    count += birdsPerDay[day];
  }
  return count;
}

/**
 * Fixes the counting mistake by increasing the bird count
 * by one for every second day.
 *
 * @param {number[]} birdsPerDay
 * @returns {number[]} corrected bird count data
 */
export function fixBirdCountLog(birdsPerDay) {
  for (let day = 0; day < birdsPerDay.length; day++) {
    if (day % 2 == 0) {
      birdsPerDay[day] += 1;
    }
  }
  return birdsPerDay;
}
