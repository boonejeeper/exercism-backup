def reverse(text):
    reversed = ""
    for char in range(len(text) - 1, -1, -1):
        reversed += text[char]

    return reversed
