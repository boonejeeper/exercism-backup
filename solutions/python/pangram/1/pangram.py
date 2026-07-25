import re

def is_pangram(sentence):
    letters_only = re.sub('[^a-z]','',sentence.lower())
    letters_used = set()
    for i in range(0,len(letters_only)):
        letters_used.add(letters_only[i])

    print(letters_only)
    print(letters_used)

    if len(letters_used) == 26:
        return True

    return False
