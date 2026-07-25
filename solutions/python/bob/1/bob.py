import re

def response(hey_bob):
    hey_bob = hey_bob.rstrip()
    print(f"hey_bob right stripped: {hey_bob}")
    is_question = len(hey_bob) > 0 and hey_bob[-1] == '?'
    print(f"is_question: {is_question}")
    is_all_caps = hey_bob.upper() == hey_bob and hey_bob.upper() != hey_bob.lower()
    print(f"is_all_caps: {is_all_caps}")
    is_empty = hey_bob.strip() == ''
    print(f"is_empty: {is_empty}")
    
    if is_question and not is_all_caps:
        return "Sure."
    if is_question and is_all_caps:
        return "Calm down, I know what I'm doing!"
    if is_all_caps and not is_question:
        return "Whoa, chill out!"
    if is_empty:
        return "Fine. Be that way!"
    
    return "Whatever."
