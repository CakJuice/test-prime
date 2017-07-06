"""
* author Cak Juice
"""

import time

current_milli_second = lambda: int(round(time.time() * 1000))

start_time = current_milli_second()

primes = []
limit = 25000

for n in range(2, limit + 1):
	is_prime = True
	for div in range(2, n):
		if n % div == 0:
			is_prime = False
			break

	if is_prime:
			primes.append(n)

end_time = current_milli_second()

print(primes)

print('Start Time:', start_time)
print('End Time:', end_time)
print('Selisih:', end_time - start_time, 'ms')