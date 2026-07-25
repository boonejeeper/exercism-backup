import math

def classify(number):
    """ A perfect number equals the sum of its positive divisors.

    :param number: int a positive integer
    :return: str the classification of the input integer
    """
    if number <= 0:
        raise ValueError("Classification is only possible for positive integers.")
        
    factors = set()
    for n in range(2, int(math.isqrt(number)) + 1):
        if number % n == 0:
            factors.add(n)
            factors.add(number // n)

    aliquot_sum = 1
    for factor in factors:
        aliquot_sum += factor
        
    print(f"factors: {factors}")
    print(f"aliquot_sum: {aliquot_sum}")

    if len(factors) == 0 or aliquot_sum < number:
        return "deficient"
    if aliquot_sum == number:
        return "perfect"
    else:
        return "abundant"
    
