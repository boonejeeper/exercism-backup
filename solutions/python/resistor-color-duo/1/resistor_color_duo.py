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

def value(colors):
    value = 0
    
    for i in range(0,2):
        value += COLORS.index(colors[i]) * pow(10,2-i-1)
    
    return value


