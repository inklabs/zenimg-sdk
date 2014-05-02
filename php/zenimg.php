<?php

/*
#  Copyright 2012 Ink Labs, LLC
#
#  v.1.0.8
# 
#  Licensed under the Apache License, Version 2.0 (the "License"); you may
#  not use this file except in compliance with the License. You may obtain
#  a copy of the License at
# 
#  http://www.apache.org/licenses/LICENSE-2.0
# 
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
#  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
#  License for the specific language governing permissions and limitations
#  under the License.
*/

if ( ! class_exists('Arr')) {
	Class Arr {
		public static function get($array, $key, $default = NULL)
		{
		    return isset($array[$key]) ? $array[$key] : $default;
		}
	}
}

class Zenimg {

	public static $render_location = 'http://i.zenimg.com';
	public static $cache_enabled = FALSE;
	public static $cache_time = 4294967295;
	public static $cache_path = '/tmp/';
	public static $pan_angles = array(
		0,
		15,
		30,
		45,
		60,
		75,
		90,
		105,
		120,
		135,
		150,
		165,
		180,
		-165,
		-150,
		-135,
		-120,
		-105,
		-90,
		-75,
		-60,
		-45,
		-30,
		-15,
	);

	public static function get_img_url($params, $prefer_cached_urlc = FALSE) {

		$image_code = Arr::get($params, 'image_code');
		$url = Arr::get($params, 'url');

		$style = Arr::get($params, 'style');
		$cg_style = Arr::get($params, 'cg_style');
		$cg_edge_color = Arr::get($params, 'cg_edge_color');
		$cg_depth = Arr::get($params, 'cg_depth');
		$al_edge_depth = Arr::get($params, 'al_edge_depth');
		$al_rounded = Arr::get($params, 'al_rounded');
		$ac_edge_depth = Arr::get($params, 'ac_edge_depth');
		$frame_code = Arr::get($params, 'frame_code');
		$mat_width = Arr::get($params, 'mat_width');
		$mat_color = Arr::get($params, 'mat_color');
		$curl = Arr::get($params, 'curl');
		$curl_width = Arr::get($params, 'curl_width');
		$wood_style = Arr::get($params, 'wood_style');
		$ald_style = Arr::get($params, 'ald_style');
		$forex_edge_color = Arr::get($params, 'forex_edge_color');
		$background = Arr::get($params, 'background');
		$background_texture = Arr::get($params, 'background_texture');
		$background_texture_color = Arr::get($params, 'background_texture_color');
		$shadow = Arr::get($params, 'shadow');
		$pan = Arr::get($params, 'pan');
		$tilt = Arr::get($params, 'tilt');
		$roll = Arr::get($params, 'roll');
		$actual_size = Arr::get($params, 'actual_size');
		$size = Arr::get($params, 'size');
		$format = Arr::get($params, 'format', 'jpg');

		$file_options = array();

		$file_options[] = $style;

		switch ($style) {
			case 'CG':
			case 'CG2':
				if ($cg_style == 'IW') {
					$file_options[] = $cg_style;
				}

				if ($cg_edge_color !== NULL) {
					$file_options[] = 'EC' . strtoupper($cg_edge_color);
				}
		
				if ($cg_depth !== NULL) {
					$file_options[] = 'D' . $cg_depth;
				}
			break;

			case 'CF':
				if ($frame_code !== NULL) {
					$file_options[] = 'F' . $frame_code;
				}
			break;

			case 'P':
				if ($curl !== NULL) {
					$file_options[] = 'C' . $curl;

					if ($curl_width != NULL) {
						$file_options[] = 'CW' . $curl_width;
					}
				} else {
					if ($frame_code !== NULL) {
						$file_options[] = 'F' . $frame_code;
					}

					if ($mat_width !== NULL) {
						$file_options[] = 'MW' . $mat_width;
					}

					if ($mat_color !== NULL) {
						$file_options[] = 'MC' . strtoupper($mat_color);
					}
				}
			break;

			case 'AC':
				if ($ac_edge_depth !== NULL) {
					$file_options[] = 'ED' . $ac_edge_depth;
				}
			break;

			case 'AL':
				if ($al_edge_depth !== NULL) {
					$file_options[] = 'ED' . $al_edge_depth;
				}

				if ($al_rounded !== NULL) {
					$file_options[] = 'RD' . $al_rounded;
				}
			break;

			case 'WD':
				// No options for Wood
			break;

			case 'ALD':
				if ($ald_style == 1) {
					$file_options[] = 'S1';
				}
			break;

			case 'ACD':
				// No options for Acrylic Dibond
			break;

			case 'F':
				if ($forex_edge_color !== NULL) {
					$file_options[] = 'EC' . strtoupper($forex_edge_color);
				}
			break;

			default:
				return '';
			break;
		}

		if ($wood_style !== NULL) {
			$file_options[] = 'W' . $wood_style;
		}

		if ($background !== NULL) {
			$file_options[] = 'BG' . $background;
		}

		if ($background_texture !== NULL) {
			$file_options[] = 'BT' . $background_texture;
		}

		if ($background_texture_color !== NULL) {
			$file_options[] = 'BTC' . strtoupper($background_texture_color);
		}

		if ($shadow === TRUE) {
			$file_options[] = 'SHD';
		}

		if ($pan !== NULL) {
			$file_options[] = 'P' . $pan;
		}

		if ($tilt !== NULL) {
			$file_options[] = 'T' . $tilt;
		}

		if ($roll !== NULL) {
			$file_options[] = 'R' . $roll;
		}

		if ($actual_size !== NULL) {
			$file_options[] = 'A' . strtoupper($actual_size);
		}

		if ($size !== NULL) {
			$file_options[] = strtoupper($size);
		} else {
			return '';
		}
		
		$file_options = implode('_', $file_options) . '.' . $format;

		if ($image_code !== NULL) {
			return self::$render_location . '/v1/' . $image_code . '_' . $file_options;
		} elseif ($url !== NULL) {
			$clean_url = urlencode($url);

			if ($prefer_cached_urlc) {
				$image_code = self::get_img_code($clean_url);
				
				if ($image_code !== NULL) {
					return self::$render_location . '/v1/' . $image_code . '_' . $file_options;
				}
			}
			
			return self::$render_location . '/v1/url/' . $file_options . '?url=' . $clean_url;
		}

		return '';
	}

	public static function get_img_code($clean_url) {
		if (self::$cache_enabled) {
			require_once 'simpleCache.php';

			$cache_key = sha1($clean_url);

			$cache = new SimpleCache();
			$cache->cache_path = self::$cache_path;
			$cache->cache_time = self::$cache_time;
			$data = $cache->get_cache($cache_key);

			if ($data !== FALSE AND $data !== '') {
				return $data;
			}
		}

		$image_url = self::$render_location . '/v1/urlc?url=' . $clean_url;
		$data = file_get_contents($image_url);

		if ($data !== '!!' AND $data !== '') {
			if (self::$cache_enabled) {
				$cache->set_cache($cache_key, $data);
			}
		} else {
			$data = NULL;
		}

		return $data;
	}
}
