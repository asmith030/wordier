import sys, json

with open("words.json", "rb") as rfile:
	words = json.load(rfile)[::-1]

print(words[int(sys.argv[1])])