import re

def do_rotate(character, base_a, key):
    return chr((((ord(character) - base_a) + key) % 26) + base_a)

def rotate(text, key):
    rotated = ""
    for i in range(0, len(text)):
        if re.match('[a-z]', text[i]):
            rotated += do_rotate(text[i], ord('a'), key)
        elif re.match('[A-Z]', text[i]):
            rotated += do_rotate(text[i], ord('A'), key)
        else:
            rotated += text[i]
        

    return rotated
