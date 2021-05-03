import pickle

def save(high_score = 0):
	# a method for saving the highscore
	with open('data.bin', 'wb') as writer:
		pickle.dump(high_score, writer)

def load():
	# a method for loading the highscore
	try:
		with open('data.bin', 'rb') as reader:
			return pickle.load(reader)
	except (FileNotFoundError, ValueError, EOFError):
		return 0