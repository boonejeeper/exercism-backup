
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

TOLERANCES = {
    "grey": "0.05%",
    "violet": "0.1%",
    "blue": "0.25%",
    "green": "0.5%",
    "brown": "1%",
    "red": "2%",
    "gold": "5%",
    "silver": "10%"
}

MAGNITUDES = [
    "",
    "kilo",
    "mega",
    "giga"
]

def resistor_label(colors):
    if len(colors) == 1:
        return f"{COLORS.index(colors[0])} ohms"
    
    digits = 0

    tolerance = f" ±{TOLERANCES[colors.pop()]}"
    magnitude_band = colors.pop()

    print(f"colors: {colors}")
    for i in range(0, len(colors)):
        digits += COLORS.index(colors[i]) * pow(10, len(colors)-i-1)
        print(f"digits: {digits}")

    print(f"digits: {digits}  -- magnitude band: {magnitude_band}")
    digits *= pow(10, COLORS.index(magnitude_band))

    magnitude = 0
    while digits >= 1000:
        digits /= 1000
        magnitude += 1

    if digits % 1 == 0:
        digits = int(digits)
    
    return str(digits) + " " + MAGNITUDES[magnitude] + "ohms" + tolerance
