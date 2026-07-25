def equilateral(sides):
    a, b, c = sides
    return is_triangle(sides) and a == b and a == c

def isosceles(sides):
    a, b, c = sides
    return is_triangle(sides) and (a == b or a == c or b == c)

def scalene(sides):
    a, b, c = sides
    return is_triangle(sides) and a != b and a != c and b != c

def is_triangle(sides):
    a, b, c = sides
    return a + b >= c and a + c >= b and b + c >= a and a != 0 and b != 0 and c != 0