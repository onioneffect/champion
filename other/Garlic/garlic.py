import re, sys, os, time
from PIL import Image, ImageDraw, ImageFont
from datetime import datetime

class LineObj:
    MY_DRAWING = 1
    SOMEONE_ELSES = 2

    salt = []
    expected_salt = [42, 10, 2]
    good_salt = False

    owner = 0
    joined = 0

    pixels = []

    def __init__(self, matches : tuple):
        self.salt = [int(i) for i in [matches[0], matches[2], matches[3]]]

        if self.salt != self.expected_salt:
            print("WARNING: Salt does not match expected values!", file=sys.stderr)
            self.good_salt = False
        else:
            self.good_salt = True

        if matches[1]:
            self.joined = int(matches[1])
            self.owner = LineObj.MY_DRAWING
        else:
            self.owner = LineObj.SOMEONE_ELSES

        coords = [int(i) for i in matches[4].split(',')]
        if len(coords) % 2:
            raise ValueError("Odd number of coordinates!")

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

    # Yes, I am using the meme font.
    draw_font = ImageFont.truetype('impact.ttf', 50)

    for i, j in enumerate(obj_list):
        l = len(j.pixels)

        for pix_index in range(1, l):
            start = j.pixels[pix_index-1]
            end = j.pixels[pix_index]

            d.line([start[0], start[1], end[0], end[1]], fill = 'black', width = 5)

            # Read these two lines and tell me I'm not a programming genius.
            d.text((20, 30), str(i-1), fill = 'white', font = draw_font)
            d.text((20, 30), str(i), fill = 'black', font = draw_font)

            if j.owner == LineObj.MY_DRAWING:
                d.ellipse((10, 10, 20, 20), fill = 'red')
            elif j.owner == LineObj.SOMEONE_ELSES:
                d.ellipse((10, 10, 20, 20), fill = 'blue')

            im.save("out/IGOR-{}.jpg".format(i))

def decode(line_list : list) -> list:
    decoded_list = []

    reg = "(\d+)\[(?:(\d+),)*(\d+),\[(\d+),(.+)\]\]"

    for s in line_list:
        if s.startswith('#'): continue

        found = re.findall(reg, s)[0]

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

    if not os.path.isdir("out"):
        os.mkdir("out")

    clean_files("out")

    sys.exit(0)
