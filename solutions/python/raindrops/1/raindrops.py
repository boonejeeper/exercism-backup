def convert(number):
    result = ""
    if number % 3 == 0:
        result = f"{result}Pling"
    if number % 5 == 0:
        result = f"{result}Plang"
    if number % 7 == 0:
        result = f"{result}Plong"

    if result == "":
        result = f"{number}"
        
    return result
