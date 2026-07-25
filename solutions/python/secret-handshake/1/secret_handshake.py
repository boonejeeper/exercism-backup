def commands(binary_str):
    handshake = []
    sequence = int(binary_str, 2)
    if sequence & 1:
        handshake.append("wink")

    if sequence & 2:
        handshake.append("double blink")

    if sequence & 4:
        handshake.append("close your eyes")

    if sequence & 8:
        handshake.append("jump")

    if sequence & 16:
        handshake.reverse()

    return handshake
