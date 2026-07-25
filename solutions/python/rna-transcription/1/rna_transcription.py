DNA_TO_RNA = {
    "G": "C",
    "C": "G",
    "T": "A",
    "A": "U"
}

def to_rna(dna_strand):
    if len(dna_strand) == 0:
        return ""

    rna_strand = ""
    for i in range(0, len(dna_strand)):
        rna_strand += DNA_TO_RNA[dna_strand[i]]

    return rna_strand
