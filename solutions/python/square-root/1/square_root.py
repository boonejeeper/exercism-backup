def square_root(number):
    if number == 1:
        return 1
    
    L = 0
    R = number

    while L != R - 1:
        M = int((L + R) / 2)

        if M * M <= number:
            L = M
        else:
            R = M

    return L
