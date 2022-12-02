import os
import sys
import argparse

def file_exist (string):
    if not os.path.exists(string):
        raise argparse.ArgumentTypeError('ERROR: input file does not exist.')
    return string

def part_one(file):
    your_score = 0
    you = {
        "X": 1, 
        "Y": 2, 
        "Z": 3
    }
    plays = {
        ("A","X"): 3,
        ("A","Y"): 6,
        ("A","Z"): 0,
        ("B","X"): 0,
        ("B","Y"): 3,
        ("B","Z"): 6,
        ("C","X"): 6,
        ("C","Y"): 0,
        ("C","Z"): 3,
    }
    try:
        with open(args.filename, "r") as file:
            for num, line in enumerate(file,1):
                values = line.strip().split(" ")
                your_score += you[values[1]] + plays[(values[0],values[1])]
            print(f"Part 1 score: {your_score}")  
    except IOError:
        print("ERROR: input file is not readable.")
        sys.exit(-1)

def part_two(file):
    your_score = 0
    you = {
        "X": 0, 
        "Y": 3, 
        "Z": 6
    }
    plays = {
        ("X","A"): 3,
        ("X","B"): 1,
        ("X","C"): 2,
        ("Y","A"): 1,
        ("Y","B"): 2,
        ("Y","C"): 3,
        ("Z","A"): 2,
        ("Z","B"): 3,
        ("Z","C"): 1,
    }
    try:
        with open(args.filename, "r") as file:
            for num, line in enumerate(file,1):
                values = line.strip().split(" ")
                your_score += you[values[1]] + plays[(values[1],values[0])]
            print(f"Part 2 score: {your_score}")  
    except IOError:
        print("ERROR: input file is not readable.")
        sys.exit(-1)


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='What Elf carry more calories?')
    parser.add_argument('filename', metavar='FILE', type=file_exist,
                    help='file doesnt exist')
    
    args = parser.parse_args() 
    
    part_one(args.filename) 
    part_two(args.filename)