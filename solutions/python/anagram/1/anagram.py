def find_anagrams(word, candidates):
    sorted_target = sorted(word.lower())
    target_len = len(word)

    anagrams = []
    for candidate in candidates:
        if len(candidate) != target_len or word.lower() == candidate.lower():
            continue
        
        sorted_candidate = sorted(candidate.lower())

        if sorted_candidate == sorted_target:
            anagrams.append(candidate)

    return anagrams
