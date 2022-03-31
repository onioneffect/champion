import re, sys, os, time
from PIL import Image, ImageDraw
from datetime import datetime

class LineObj:
    salt = []
    expected_salt = [42, 10, 2]
    good_salt = False

    joined = 0

    pixels = []

    def __init__(self, matches : tuple):
        """There's two different formats. See below."""
        # self.salt = [int(i) for i in [matches[0], matches[1], matches[3]]]
        self.salt = [int(i) for i in [matches[0], matches[1], matches[2]]]

        if self.salt != self.expected_salt:
            print("WARNING: Salt does not match expected values!", file=sys.stderr)
            self.good_salt = False
        else:
            self.good_salt = True

        self.joined = int(matches[2])

        """There's two different formats. See below."""
        # coords = [int(i) for i in matches[4].split(',')]
        coords = [int(i) for i in matches[3].split(',')]

        # Thanks to stackoverflow.com/questions/44104729
        self.pixels = list(zip(*[iter(coords)]*2))

def clean_files(dir : str):
    ls = os.listdir(dir)

    unix = int(time.time())
    out_dir = dir + '/' + str(unix)
    os.mkdir(out_dir)

    for filename in ls:
        if filename.startswith("IGOR-"):
            os.rename(dir + '/' + filename, out_dir + '/' + filename)

def save_list(obj_list : list):
    im = Image.new('RGB', (600, 600), (255, 255, 255))
    d = ImageDraw.Draw(im)

    for i, j in enumerate(obj_list):
        for pix in j.pixels:
            x, y = [*pix]
            d.ellipse((x, y, x+5, y+5), fill = 'black')

        im.save("out/IGOR-{}.jpg".format(i))

def decode(line_list : list) -> list:
    decoded_list = []

    # This regex matches the unix timestamp right after the 42
    # use it to tell if it's your own drawing or someone else's
    detector = "^(?:\d+)\[(\d+),10"

    # reg_me = "(\d+)\[(\d+),(\d+),\[(\d+),(.+)\]\]"
    reg_others = "(\d+)\[(\d+),\[(\d+),(.+)\]\]"

    for s in line_list:
        if s.startswith('#'): continue

        found = re.findall(reg_others, s)[0]

        curr_obj = LineObj(found)
        decoded_list.append(curr_obj)

    return decoded_list

def read_file(path : str) -> list:
    encoded_list = []

    fp = open(path, "r")
    while s := fp.readline(): # WALNUTS OPERATOR XD
        encoded_list.append(s)

    return decode(encoded_list)

def read_stdin() -> list:
    encoded_list = []

    print("Running in stdin mode...")
    while (s := input("")) not in ("\x04" ""): # CTRL-D or an empty line
        encoded_list.append(s)

    return decode(encoded_list)

if __name__ == "__main__":
    if len(sys.argv) < 2:
        ret = read_stdin()
    else:
        ret = read_file(sys.argv[1])
    
    save_list(ret)
    clean_files("out")

    sys.exit(0)
