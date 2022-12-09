import os
import sys
import argparse

def file_exist (string):
    if not os.path.exists(string):
        raise argparse.ArgumentTypeError('ERROR: input file does not exist.')
    return string

def is_visible(high,trees) -> bool:
    for each in trees:
        if each >= high:
            return False
    return True

def visibility(high,trees) -> int:
    num_trees = 0
    for each in trees:
        if each < high:
            num_trees +=1
        else:
            num_trees +=1
            break
    return num_trees

def part_one(rows,columns):
    visible = (len(rows) * 2) + (len(columns) * 2) - 4 # All trees in edge are visible
    for row in range(1,len(rows)-1):
        for column in range(1,len(columns[0])-1):
            value = rows[row][column]
            up = is_visible(value, columns[column][0:row])
            down = is_visible(value, columns[column][row+1:])
            left = is_visible(value, rows[row][0:column])
            right = is_visible(value, rows[row][column+1:])
            if (up or down) or (left or right):
                visible +=1
    print(f"Part one: {visible}")

def part_two(rows,columns):
    max_visibility = 0
    for row in range(1,len(rows)-1):
        for column in range(1,len(columns[0])-1):
            value = rows[row][column]
            up = visibility(value, columns[column][0:row][::-1])
            down = visibility(value, columns[column][row+1:])
            left = visibility(value, rows[row][0:column][::-1])
            right = visibility(value, rows[row][column+1:])
            prod = up * down * left * right
            if  prod > max_visibility:
                max_visibility = prod 
    print(f"Part two: {max_visibility}")

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
    
    rows = [[int(each) for each in line] for line in data]
    columns = [[] for column in range(len(data))]

    for idxRow in range(len(data[0])):
        temp = []
        for idxColumn in range(len(data[0])):
                temp = [*temp, int(data[idxColumn][idxRow])]
        columns[idxRow] = [*temp]     

    part_one(rows,columns)
    part_two(rows,columns)


    