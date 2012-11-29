<?php
require 'zenimg.php';

$test_path = dirname(__DIR__) . '/tests.json';
$tests = json_decode(file_get_contents($test_path), TRUE);

$status = 0;
foreach ($tests as $test) {
	$params = $test[0];
	$expected = $test[1];
	
	$result = Zenimg::get_img_url($params);

	if ($expected == $result) {
		echo "\t[OK]\t" . $result . "\n";
	} else {
		echo "\t[FAIL]\t" . $result . "\n";
		echo "\tEXP-->\t" . $expected . "\n";
		$status = 1;
	}
}

exit($status);
