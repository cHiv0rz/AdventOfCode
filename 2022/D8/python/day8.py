import os
import sys
import argparse
import re
from collections import defaultdict



def file_exist (string):
    if not os.path.exists(string):
        raise argparse.ArgumentTypeError('ERROR: input file does not exist.')
    return string

def visible_up(high,trees) -> bool:
    for each in trees:
        if each >= high:
            return False
    return True

def visible_down(high,trees) -> bool:
    for each in trees:
        if each >= high:
            return False
    return True

def visible_left(high,trees) -> bool:
    for each in trees:
        if each >= high:
            return False
    return True

def visible_right(high,trees) -> bool:
    for each in trees:
        if each >= high:
            return False
    return True


def part_one(rows,columns):
    visible = (len(rows) * 2) + (len(columns) * 2) - 4 # All trees in edge are visible
    for row in range(1,len(rows)-1):
        for column in range(1,len(columns[0])-1):
            value = rows[row][column]
            if (visible_up(value, columns[column][0:row]) or \
                visible_down(value, columns[column][row+1:])) or \
                    (visible_left(value, rows[row][0:column]) or  
                    visible_right(value, rows[row][column+1:])):
                        visible +=1
    print(visible)

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

    rows = [[int(h) for h in line] for line in data]
    columns = [[] for x in range(len(data))]

    for i in range(len(data[0])):
        temp = []
        for j in range(len(data[0])):
                temp = [*temp, int(data[j][i])]
        columns[i] = [*temp]     

    part_one(rows,columns)


    