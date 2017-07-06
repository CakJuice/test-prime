/*
 * author Cak Juice
 */

function currentMilliSecond() {
	return new Date().getTime();
}

var startTime = currentMilliSecond();

var primes = [];
var limit = 25000;

for (var n=2;n<=limit;n++) {
	var isPrime = true;

	for (var div=2;div<n;div++) {
		if (n % div == 0) {
			isPrime = false;
			break;
		}
	}

	if (isPrime) {
		primes.push(n);
	}
}

var endTime = currentMilliSecond();

console.log(primes);

console.log('Start Time:', startTime);
console.log('End Time:', endTime);
console.log('Selisih:', endTime - startTime, 'ms');