import re, sys
from PIL import Image, ImageDraw
from datetime import datetime

class LineObj:
    salt = []
    expected_salt = [42, 10, 2]
    good_salt = False

    joined = 0

    pixels = list()

    def __str__(self):
        s = ""
        if self.good_salt:
            s += '#'
        else:
            s += '!'

        unix = datetime.fromtimestamp(self.joined/1000)
        s += unix.strftime("%m%d-%H%M-")

        s += str(len(self.pixels))

        return s

    def __init__(self, matches : tuple):
        # self.salt = [int(i) for i in [matches[0], matches[1], matches[3]]]
        self.salt = [int(i) for i in [matches[0], matches[1], matches[2]]]
        if self.salt != self.expected_salt:
            print("WARNING: Salt does not match expected values!", file=sys.stderr)
            self.good_salt = False
        else:
            self.good_salt = True

        self.joined = int(matches[2])

        # coords = [int(i) for i in matches[4].split(',')]
        coords = [int(i) for i in matches[3].split(',')]
        # Thanks to stackoverflow.com/questions/44104729
        self.pixels = list(zip(*[iter(coords)]*2))

def view_line(obj : LineObj):
    im = Image.new('RGB', (600, 600), (255, 255, 255))
    d = ImageDraw.Draw(im)

    for pix in obj.pixels:
        x, y = pix[0], pix[1]
        d.ellipse((x, y, x+5, y+5), fill = 'black')

    im.show()

def save_list(obj_list : list):
    for i, j in enumerate(obj_list):
        im = Image.new('RGB', (600, 600), (255, 255, 255))
        d = ImageDraw.Draw(im)

        for pix in j.pixels:
            x, y = pix[0], pix[1]
            d.ellipse((x, y, x+5, y+5), fill = 'black')

        im.save("out/IGOR-{}.jpg".format(i))

def pretty_print(obj_list : list):
    for line in obj_list:
        print(line)

def decode(line_list : list) -> list:
    decoded_list = []

    # reg = "(\d*)\[(\d*),(\d*),\[(\d*),(.*)\]\]"
    reg = "(\d*)\[(\d*),\[(\d*),(.*)\]\]" # ITS DIFFERENT NOW???
    # I THINK ITS BECAUSE IT WAS SOMEONE ELSES DRAWING
    # SO IT DOESNT INCLUDE THE UNIX TIMESTAMP
    # FML
    for s in line_list:
        if s.startswith('#'): continue

        found = re.findall(reg, s)[0] # re.findall returns a list with a single tuple inside. idk.

        curr_obj = LineObj(found)
        decoded_list.append(curr_obj)

    return decoded_list

def read_file(path : str) -> list:
    encoded_list = list()

    fp = open(path, "r")
    while s := fp.readline(): # THEY FINALLY ADDED IT GUYS THE WALNUTS OPERATOR XD
        encoded_list.append(s)

    return decode(encoded_list)

def read_stdin() -> list:
    encoded_list = list()

    print("Running in stdin mode...")
    while (s := input("")) not in ("\x04" ""): # 0x04 is CTRL-D.
        encoded_list.append(s)

    return decode(encoded_list)

if __name__ == "__main__":
    if len(sys.argv) < 2:
        ret = read_stdin()
    else:
        ret = read_file(sys.argv[1])
    
    save_list(ret)
    sys.exit(0)
