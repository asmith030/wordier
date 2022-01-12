import sys, json, base64

if len(sys.argv) < 3:
	print("what?")

with open("table.dat", "rb") as rfile:
	data = base64.b64decode(rfile.read())

with open("words.json", "rb") as rfile:
	words = json.load(rfile)[::-1]

guess = words.index(sys.argv[1])
answer = words.index(sys.argv[2])

print(data[guess * len(words) + answer])

