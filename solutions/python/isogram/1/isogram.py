def is_isogram(string):

    string = string.lower()
    letters = set()
    for char in string:
        if char in letters and char != " " and char != "-":
            return False
        letters.add(char)
        
    return True
