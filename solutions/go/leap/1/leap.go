/* leap package contains a utility function that returns true if the int passed is a leap year, false otherwise.
 *
 */
package leap

/* IsLeapYear determines if a year passed as an int is a leap year.
 * leap years occur every year that is divisible by 4. EXCEPT when
 * the year is also divisible by 100 but not 400.
 */
func IsLeapYear(year int) bool {
    if year % 100 == 0 && year % 400 != 0 || year % 4 != 0 {
        return false
    }
    return true
}
