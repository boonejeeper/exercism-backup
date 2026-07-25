COLORS = [
    "black",
    "brown",
    "red",
    "orange",
    "yellow",
    "green",
    "blue",
    "violet",
    "grey",
    "white"
]

MAGNITUDES = [
    "",
    "kilo",
    "mega",
    "giga"
]

def label(colors):
    digits = 0

    for i in range(0, 2):
        digits += COLORS.index(colors[i]) * pow(10, 1-i)

    digits *= pow(10, COLORS.index(colors[2]))

    magnitude = 0
    while digits >= 1000:
        digits /= 1000
        magnitude += 1

    return str(int(digits)) + " " + MAGNITUDES[magnitude] + "ohms"
