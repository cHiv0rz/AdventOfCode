import os
import sys
import argparse
import re
from collections import defaultdict

def file_exist (string):
    if not os.path.exists(string):
        raise argparse.ArgumentTypeError('ERROR: input file does not exist.')
    return string

def calculate_paths_size(data) -> defaultdict:
    sizes = defaultdict(int)
    current_path = []
    for command in data:
        if folder := re.match(r"\$ cd (.+)", command):
            if folder.group(1) == "..":
                current_path = current_path[:-1] 
            else:
                current_path.append(folder.group(1))
        elif size := re.match(r"(\d+) .+", command):
            sizes[tuple(current_path)] += int(size.group(1)) 
            for i in range(1, len(current_path)):
                sizes[tuple(current_path[:-i])] += int(size.group(1))
    return(sizes)

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='D5 Input')
    parser.add_argument('filename', metavar='FILE', type=file_exist,
                    help='file doesnt exist')
    
    args = parser.parse_args() 
    try:
        with open(args.filename, "r") as file:
            data = file.read().splitlines()

    except:
        print("ERROR: input file is not readable.")
        sys.exit(-1) 
        
    path_sizes = calculate_paths_size(data)

    # Part one
    sum = 0
    for each in path_sizes.values():
        if each <= 100000:
            sum += each
    print(f"Part one: {sum}")

    # Part two
    disk_in_use = sorted(path_sizes.values())[-1]
    space_needed = abs(30_000_000 - (70_000_000 - sorted(path_sizes.values())[-1]))
    result = []
    for each in sorted(path_sizes.values()):
        if each >= space_needed:
            result.append(each)
    print(f"Part two: {sorted(result)[0]}")