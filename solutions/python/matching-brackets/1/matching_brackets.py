def is_paired(input_string):
    symbols = []

    for i in range(0, len(input_string)):
        char = input_string[i]
        if char in ['[','{','(']:
            symbols.append(char)
        if char in [']','}',')']:
            if len(symbols) > 0 and ((char == ']' and symbols[-1] == '[') or (char == '}' and symbols[-1] == '{') or (char == ')' and symbols[-1] == '(')):
                symbols.pop()
            else:
                return False

    return len(symbols) == 0
