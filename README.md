# Champion
Generates arbitrary inputs for the game "Champ'd Up" from The Jackbox Party Pack.

## CONTEXT:
https://www.youtube.com/watch?v=cZ9gIoIy5No

## USAGE:
`./champion.exe [options] [arguments]`

#### Available options:
```
--help: Print help message and exit
--debug: Enable logging (more info)
--file: Specify where to write log output
```

## HOWTO:
#### NOTE: This only applies to the program files in the old/ directory
1. Run champ.go and point it to a simple-ish PNG file (up to 600x600 pixels, only 2 colors)
2. Run fix.py on the JSON output
3. Join a game of Champ'd Up using a proxy software of your choice (works best with BurpSuite)
4. Draw whatever you want on the screen and pick a random name for your champion
5. Turn the intercept function on in BurpSuite's control panel
6. Hit submit on your Chromium browser
7. Once BP intercepts the submit WebSocket request, you can replace the objects in the `lines` array with the output of fix.py
8. ???
9. Profit

## TODO:
* Rewrite entire program in a single language
* Write a proxy server to make the whole process easier
* Add different encoding for images (Gartic and TeeKO)
