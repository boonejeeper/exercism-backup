EXPECTED_BAKE_TIME = 40
PREPARATION_TIME = 2

def bake_time_remaining(actual_time_in_oven):
    """Calculate the bake time remaining.

    :param elapsed_bake_time: int - baking time already elapsed.
    :return: int - remaining bake time (in minutes) derived from 'EXPECTED_BAKE_TIME'.

    Function that takes the actual minutes the lasagna has been in the oven as
    an argument and returns how many minutes the lasagna still needs to bake
    based on the `EXPECTED_BAKE_TIME`.
    """
    return EXPECTED_BAKE_TIME - actual_time_in_oven

def preparation_time_in_minutes(number_of_layers):
    """Calculate the preparation time in minutes.

    :param number_of_layers: int - the number of layers to prepare.
    :return: int - number of minutes to prepare the lasagna. derived from 'PREPARATION_TIME'.

    Function that takes the number of layers desired in the lasagna as
    an argument and returns how many minutes it will take to prepare the lasagna.
    based on the `PREPARATION_TIME`.
    """
    return number_of_layers * PREPARATION_TIME

def elapsed_time_in_minutes(number_of_layers, elapsed_bake_time):
    """Calculate the preparation time in minutes.

    :param number_of_layers: int - the number of layers to prepare.
    :param elapsed_bake_time: int - baking time already elapsed.
    :return: int - the total minutes you have been in the kitchen cooking

    This function should return the total minutes you have been in the kitchen
    cooking — your preparation time layering + the time the lasagna has spent 
    baking in the oven
    """
    return preparation_time_in_minutes(number_of_layers) + elapsed_bake_time
