def square(number):
    if number < 1 or number > 64:
        raise ValueError("square must be between 1 and 64")
    return pow(2, number - 1)

def total():
    sum = 0
    for index in range(64):
        sum = sum + square(index+1)
        print(f"square: {index+1} index: {index}, sum: {sum}")

    return sum
