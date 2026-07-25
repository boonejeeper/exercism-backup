def is_armstrong_number(number):
    index = 0
    working_number = number
    digits = []
    armstrong = 0
    while working_number > 0:
        digits.append(working_number % 10)
        working_number = int(working_number / 10)
        index += 1

    for index in range(len(digits)):
        armstrong += pow(digits[index], len(digits))
        
    return armstrong == number
