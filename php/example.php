<?php

require 'zenimg.php';

$url = 'http://i.imgur.com/ePJZv.jpg';
$clean_url = urlencode($url);

$params = array(
	'url' => $url,
	'style' => 'CG',
	'pan' => -25,
	'size' => '100x100',
	'format' => 'jpg',
);

echo Zenimg::get_img_url($params) . "\n";
echo Zenimg::get_img_code($clean_url) . "\n";

Zenimg::$cache_enabled = TRUE;
echo Zenimg::get_img_code($clean_url) . "\n";
