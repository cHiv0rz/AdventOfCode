import os
import sys
import argparse

def file_exist (string):
    if not os.path.exists(string):
        raise argparse.ArgumentTypeError('ERROR: input file does not exist.')
    return string


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='What Elf carry more calories?')
    parser.add_argument('filename', metavar='FILE', type=file_exist,
                    help='file doesnt exist')

    args = parser.parse_args() 
    try:
        with open(args.filename, "r") as file:
            sum = 0
            elves = []
            temporal = []
            top = []
            for num, line in enumerate(file,1):
                if line != '\n':
                    temporal.append(int(line))
                else:
                    elves.append(temporal)
                    temporal = []
                
            for i in range(0,len(elves)):
                temporal = elves[i]
                for j in range(0, len(temporal)):
                    sum+=temporal[j]
                top.append(sum)
                sum = 0
            top.sort()
            
            print(f"The highest number of calories carried is {top[-1]} calories")
            top3_calories_sum = top[-1] + top[-2] + top[-3]
            print(f"Top 3 calories carried: {top[-3::]}, total sum: {top3_calories_sum}")
            
 
            
    except IOError:
        print("ERROR: input file is not readable.")
        sys.exit(-1)