import json
import base64
import collections

table = bytearray()

def score(guess, secret):
    # All characters that are not correct go into the usable pool.
    pool = collections.Counter(s for s, g in zip(secret, guess) if s != g)
    # Create a first tentative score by comparing char by char.
    total = 0
    i = 0
    for secret_char, guess_char in zip(secret, guess):
        if secret_char == guess_char:
            total += 2 * 3 ** i
        elif guess_char in secret and pool[guess_char] > 0:
            total += 1 * 3 ** i
            pool[guess_char] -= 1
        i += 1
    return total

with open("words.json", "rb") as rfile:
	words = json.load(rfile)[::-1]

for a in words:
	for b in words:
		table.append(score(a, b))

with open("table.dat", "wb") as wfile:
	wfile.write(base64.b64encode(table))

