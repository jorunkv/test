
from threading import Thread


i = 0

def threadFunction1():
	global i 
	for n in range(1,1000000):
		i += 1

def threadFunction2():

	for n in range(1,1000000):
		global i
		i -= 1

def main():
	someThread1 = Thread(target = threadFunction1, args = (),)
	someThread1.start()

	someThread2 = Thread(target = threadFunction2, args = (),)
	someThread2.start()

	someThread1.join()
	someThread2.join()

	print i

main()
