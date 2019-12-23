#!/usr/bin/python3
import argparse
import sys

version = "0.0.1"

class Generator:
    """ Takes a file and applies a Permutator to each entry. """
    def __init__(self, infile, permutations=""):
        output = []
        self.keyboard = []

        with open("./keyboard_chars.txt", 'r') as f:
            for line in f:
                self.keyboard += [line.rstrip()]
        # print(self.keyboard)

        with open(infile, 'r') as f:
            for line in f:
                p = Permutator(line, permutations=permutations)
                output += p.output
        
        # Remove pesky duplicates.
        output = list(set(output))

        # Print - handle output with bash.
        for line in output:
            print(line)
        
class Permutator:
    """ Takes a password and permutes it according to rules. """
    def __init__(self, word, permutations=""):
        self.output = [word.rstrip()]
        for perm in permutations.split(","):
            perm = perm.rstrip()
            self.run_permutation(perm)
    
    PERMUTATIONS = [
        "ascii",
        "case",
        "omit",
        "add",
        "swap_two",
    ]
    
    def run_permutation(self, arg):
        if arg == "ascii":
            self.permute_ascii()
        if arg == "case":
            self.permute_case()
        if arg == "omit":
            self.permute_omit()
        if arg == "add":
            self.permute_add()
        if arg == "swap_two":
            self.permute_swap_two()
    
    # Change all chars case.
    def permute_case(self):
        generated = []
        for w in self.output:
            nw = list(w)
            for i, ch in enumerate(nw):
                if ch.isupper():
                    nw[i] = ch.lower()
                else:
                    nw[i] = ch.upper()
            generated += ["".join(nw)]
        self.output += generated

    def permute_ascii(self):
        generated = []
        for w in self.output:
            for i, ch in enumerate(w):
                nw = list(w)
                # Loop ascii characters from SPACE to ~.
                for asc in range(32, 126):
                    nw[i] = chr(asc)
                    generated += ["".join(nw)]
        self.output += generated
    
    def permute_omit(self):
        generated = []
        for w in self.output:
            for i, ch in enumerate(w):
                nw = list(w)
                omitted = nw[:i] + nw[i+1:]
                generated += ["".join(omitted)]
        self.output += generated
    
    # Add ascii at the beginning or end.
    def permute_add(self):
        generated = []
        for w in self.output:
            for asc in range(32, 126):
                generated += [w + chr(asc)]
                generated += [chr(asc) + w]
        self.output += generated
    
    def permute_swap_two(self):
        generated = []
        for w in self.output:
            # Convert string to list of chars.
            lw = list(w)
            for i in range(0, len(lw) - 1):
                for j in range(i + 1, len(lw)):
                    # Make a copy if the original list.
                    permuted = lw[:]
                    tmp = permuted[i]
                    permuted[i] = permuted[j]
                    permuted[j] = tmp
                    generated += ["".join(permuted)]
        self.output += generated

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Wordlist generator for password cracking.")
    parser.add_argument(
        "infile", 
        type=str, 
        help="Text file with a password per line. New password permutations will be created out of each line.", 
    )
    parser.add_argument(
        "--permutations",
        type=str,
        default="",
        required=True,
        help="List of permutations: {}".format(Permutator.PERMUTATIONS),
    )
    args = parser.parse_args()

    # print("Password generator: ", version)
    Generator(args.infile, permutations=args.permutations)
