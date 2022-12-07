import os
import sys
import argparse
from typing import List

def part_one(data) -> List:
    n = 4
    for each in data:
        marker = []
        for idx, char in enumerate(each):
            marker.append(char)
            if len(set(marker[-n:])) == n:
            # note: we can change last to lines by:
            # marker = [*marker, char][-4:]
            # if len(set(marker)) == 4:
            # it will be less consumption in memory, but the cost in CPU will be higher,
            # since we it will be creating, copying and freezing memory any time, which produces
            # memory frangmentation. But it's cool :D
                print(idx+1)
                break
            
def part_two(data) -> List:
    n = 14
    for each in data:
        marker = []
        for idx, char in enumerate(each):
            marker.append(char)
            if len(set(marker[-n:])) == n:
                print(idx+1)
                break

def file_exist (string):
    if not os.path.exists(string):
        raise argparse.ArgumentTypeError('ERROR: input file does not exist.')
    return string

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='D5 Input')
    parser.add_argument('filename', metavar='FILE', type=file_exist,
                    help='file doesnt exist')
    
    args = parser.parse_args() 
    try:
        with open(args.filename, "r") as file:
            data = [x.strip() for x in file.readlines()]

    except:
        print("ERROR: input file is not readable.")
        sys.exit(-1) 
        
    part_one(data)
    part_two(data)