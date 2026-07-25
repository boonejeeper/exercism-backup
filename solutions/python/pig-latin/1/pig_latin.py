import re

def applyRules(text):
    print(f"starting applyRules: {text}")
    rule_3_results = re.search(r"^[^aeiou\d]*qu",text)
    if rule_3_results:
        print( f"{text[rule_3_results.end():]}{text[:rule_3_results.end()]}ay")
        return f"{text[rule_3_results.end():]}{text[:rule_3_results.end()]}ay"
    rule_4_results = re.search(r"^([^aeiou\d]+)(y)", text)
    if rule_4_results:
        print(f"{rule_4_results}")
        print(f"{text[rule_4_results.end()-1:]}{text[:rule_4_results.end()-1]}ay")
        return f"{text[rule_4_results.end()-1:]}{text[:rule_4_results.end()-1]}ay"
    if re.search("^[aeiou]|^xr|^yt", text):
        print(f"{text}ay")
        return f"{text}ay"
    rule_2_results = re.search(r"^[^aeiou\d]+", text)
    if rule_2_results:
        print(f"{text[rule_2_results.end():]}{text[:rule_2_results.end()]}ay")
        return f"{text[rule_2_results.end():]}{text[:rule_2_results.end()]}ay"

    return ""


def translate(text):
    result = ""
    text_pieces = text.split()
    for piece in text_pieces:
        result = result + applyRules(piece) + " "

    return result.strip()