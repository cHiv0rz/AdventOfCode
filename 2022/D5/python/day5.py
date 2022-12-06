import os
import sys
import argparse
import re

def file_exist (string):
    if not os.path.exists(string):
        raise argparse.ArgumentTypeError('ERROR: input file does not exist.')
    return string

def part_one(crates, moves):
    print(crates)
    print(moves)
    
    crates[2].pop(-1)
    print(crates)
    # for each in moves:
    #     for movement in range(1,each[0]+1):
    #         if crates[each[1]] == " ":
    #             print("yololo")           
        

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='D5 Input')
    parser.add_argument('filename', metavar='FILE', type=file_exist,
                    help='file doesnt exist')
    
    args = parser.parse_args() 
    try:
        with open(args.filename, "r") as file:
            data = file.readlines()
            data.remove("\n")
    except:
        print("ERROR: input file is not readable.")
        sys.exit(-1) 

    crates = []
    moves = []
    temp_crates = []
    for each in data:
        if "[" in each:
            line = each.replace("\n","")
            temp = []
            for character in range(1,len(line)+1,4):
                temp.append(line[character])
            crates.append(temp)
        elif each.startswith(" "):
            stacks = each.replace("\n","").replace("   ",",").replace(" ","").split(",")
        else:
            temp = []
            for element in each.replace("\n","").split(" "):
                if element.isdigit():
                    temp.append(int(element))
            moves.append(temp) 
        
    part_one(crates, moves)