import re

def is_valid(isbn):
    no_hyphens = re.sub('-','',isbn)
    print(f"no_hyphens: {no_hyphens}")
    digits = re.sub('[^0-9X]','',isbn.upper())
    print(f"digits: {digits}")
    num_digits = len(digits)
    if num_digits != 10:
        print("num_digits != 10")
        return False
    if len(no_hyphens) != num_digits:
        print("len(no_hyphens) != num_digits")
        return False

    sum = 0
    for i in range(0,num_digits):
        sum += 10 if digits[i] == 'X' else int(digits[i]) * (num_digits - i)

    return sum % 11 == 0
