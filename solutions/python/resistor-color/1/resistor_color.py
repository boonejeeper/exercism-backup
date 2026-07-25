def color_code(color):
    try:
        return colors().index(color)
    except:
        return -1


def colors():
    return [
        "black",
        "brown",
        "red",
        "orange",
        "yellow",
        "green",
        "blue",
        "violet",
        "grey",
        "white",
    ]
