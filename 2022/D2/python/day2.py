import os
import sys
import argparse

def file_exist (string):
    if not os.path.exists(string):
        raise argparse.ArgumentTypeError('ERROR: input file does not exist.')
    return string

def part_one(list):
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
    
    for i in range(0,len(list)):
        your_score += you[list[i][1]] + plays[(list[i][0],list[i][1])]
    print(f"Part 1 score: {your_score}") 

def part_two(list):
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
          
    for i in range(0,len(list)):
        your_score += you[list[i][1]] + plays[(list[i][1],list[i][0])]
    print(f"Part 2 score: {your_score}")  
    


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='What Elf carry more calories?')
    parser.add_argument('filename', metavar='FILE', type=file_exist,
                    help='file doesnt exist')
    
    args = parser.parse_args() 
    values = []
    try:
        with open(args.filename, "r") as file:
            for num, line in enumerate(file,1):
                values.append(line.strip().split(" "))
    except IOError:
        print("ERROR: input file is not readable.")
        sys.exit(-1)
    
    part_one(values) 
    part_two(values)