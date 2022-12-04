import os
import sys
import argparse

def file_exist (string):
    if not os.path.exists(string):
        raise argparse.ArgumentTypeError('ERROR: input file does not exist.')
    return string

def part_one(list):
    total = 0
    for each in list:
        assignment1, assignment2 = each.split(",")
        range1_1, range1_2 = assignment1.split("-")
        range2_1, range2_2 = assignment2.split("-")

        if int(range1_1) >= int(range2_1):
            if int(range1_2) <= int(range2_2):
                total +=1
                continue
        if int(range2_1) >= int(range1_1): 
            if int(range2_2) <= int(range1_2):
                total +=1
    print(f"Part 1: {total}")

def part_two(list):
    total = 0
    for each in list:
        assignment1, assignment2 = each.split(",")
        range1_1, range1_2 = assignment1.split("-")
        range2_1, range2_2 = assignment2.split("-")

        if int(range1_1) == int(range1_2):
            c1 = {int(range1_1)}
        else:
            c1 = { i for i in range(int(range1_1),int(range1_2)+1)}

        if int(range2_1) == int(range2_2):
            c2 = {int(range2_1)}
        else:
            c2 = { i for i in range(int(range2_1),int(range2_2)+1)}
    
        for number in c1:
            if number in c2:
                total += 1
                break
    print(f"Part 2: {total}")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='input file')
    parser.add_argument('filename', metavar='FILE', type=file_exist,
                    help="file doesn't exist")
    
    args = parser.parse_args() 
    values = []
    try:
        with open(args.filename, "r") as file:
            for line in file:
                values.append(line.strip())
    except IOError:
        print("ERROR: input file is not readable.")
        sys.exit(-1)

    part_one(values)
    part_two(values)