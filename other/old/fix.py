import json, sys

fptr = open(sys.argv[1])
json_resp = json.load(fptr)
a = list()

def filter(rgb : str):
    print(len(json_resp))
    for i, line in enumerate(json_resp):
        l = line["points"].split("|")

        if len(l) > 1 and line["color"] != rgb:
            json_resp[i]["points"] = "|".join(l[::len(l) - 1])
            a.append(json_resp[i])

    print(len(a))
    with open("filtered_{}.json".format(rgb), "w") as f:
        f.write(json.dumps(a))

# There's a lot more to this script than I cared to explain in the README

