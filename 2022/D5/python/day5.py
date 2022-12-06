import os
import sys
import argparse
import re

def file_exist (string):
    if not os.path.exists(string):
        raise argparse.ArgumentTypeError('ERROR: input file does not exist.')
    return string

def part_one(stacks, moves):
    print(stacks)
    result = ""
    for move in moves: 
        print(move)
        for number in range(1,move[0]+1): 
            for crate_origin in range(len(stacks[move[1]-1])-1,-1,-1): 
                if stacks[move[1]-1][crate_origin] != " ": 
                    top = stacks[move[1]-1][crate_origin] 
                    stacks[move[1]-1][crate_origin] = " " 
                    for crate_dest in range(len(stacks[move[2]-1])):
                        if crate_dest == len(stacks[move[2]-1])-1 and stacks[move[2]-1][crate_dest] != " ": 
                            stacks[move[2]-1].append(top)
                        elif stacks[move[2]-1][crate_dest] == " ":
                            stacks[move[2]-1][crate_dest] = top
                            break
                    break
                else:
                    continue

    for each in stacks:
        for index in range(len(each)-1,-1,-1):
            if each[index] != " ":
                result += each[index]
                break
    print(result)

        

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
            num_stacks = each.replace("\n","").replace("   ",",").replace(" ","").split(",")
            stacks = [ [] for x in range(len(num_stacks))]
        else:
            temp = []
            for element in each.replace("\n","").split(" "):
                if element.isdigit():
                    temp.append(int(element))
            moves.append(temp) 
    crates = crates[::-1]
    for crate in range(len(crates)):
        for stack in range(len(stacks)):
            stacks[stack].append(crates[crate][stack])

    part_one(stacks,moves)