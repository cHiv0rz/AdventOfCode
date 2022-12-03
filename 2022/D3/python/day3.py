import os
import sys
import argparse

types = [ "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", 
    "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E",
    "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U",
    "V", "W", "X", "Y", "Z"] 

def file_exist (string):
    if not os.path.exists(string):
        raise argparse.ArgumentTypeError('ERROR: input file does not exist.')
    return string


def part_one(list):
    priority = 0
    for each in list:
        compartment1 = each[:len(each)//2]
        compartment2 = each[len(each)//2:]
        for letter in compartment2:
            if letter in compartment1:
                priority += types.index(letter) + 1
                break
    print(priority)

def part_two(list):
    index = 0
    priority = 0
    while index < len(list):
        rucksack1 = list[index]
        rucksack2 = list[index+1]
        rucksack3 = list[index+2]
        for letter in rucksack1:
            if letter in rucksack2 and letter in rucksack3:
                priority += types.index(letter) + 1
                break 

        index += 3
    print(priority)


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='What Elf carry more calories?')
    parser.add_argument('filename', metavar='FILE', type=file_exist,
                    help='file doesnt exist')
    
    args = parser.parse_args() 
    values = []
    try:
        with open(args.filename, "r") as file:
            for num, line in enumerate(file,1):
                values.append(line.strip())
    except IOError:
        print("ERROR: input file is not readable.")
        sys.exit(-1)
    
    part_one(values)
    part_two(values)