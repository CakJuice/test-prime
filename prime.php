<?php

/*
 * author Cak Juice
 */

function currentMilliSecond() {
	return round(microtime(true) * 1000);
}

$startTime = currentMilliSecond();

$primes = [];
$limit = 25000;

for ($n=2;$n<=$limit;$n++) {
	$isPrime = true;
	for ($div=2;$div<$n;$div++) {
		if ($n % $div == 0) {
			$isPrime = false;
			break;
		}
	}

	if ($isPrime) {
		array_push($primes, $n);
	}
}

$endTime = currentMilliSecond();

print_r($primes);

echo 'Start Time: ' . strval($startTime) . "\n";
echo 'End Time: ' . strval($endTime) . "\n";
echo 'Selisih: ' .strval($endTime - $startTime) . "ms\n";