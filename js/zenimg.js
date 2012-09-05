/*
 * zenimg.js v1.0.0
 *
 * Copyright 2012 Ink Labs, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

var Zenimg = Zenimg || {};

Zenimg.get_image_url = function(image) {
	var render_location = 'http://i.zenimg.com';
	var file_options = new Array();

	if (typeof image.style === undefined) {
		throw 'Style is required';
	}

	if (typeof image.format === undefined) {
		throw 'Format is required';
	}

	file_options.push(image.style);

	if (image.style == 'CG') {
		if (typeof image.cg_style !== 'undefined' && image.cg_style == 'IW') {
			file_options.push(image.cg_style);
		}

		if (typeof image.cg_edge_color !== 'undefined' && image.cg_edge_color !== null && image.cg_edge_color !== '') {
			file_options.push('EC' + image.cg_edge_color);
		}

		if (typeof image.cg_depth !== 'undefined' && image.cg_depth !== null) {
			file_options.push('D' + image.cg_depth);
		}
	} else if (image.style == 'AL') {
		if (typeof image.al_edge_depth !== 'undefined' && image.al_edge_depth !== null) {
			file_options.push('ED' + image.al_edge_depth);
		}

		if (typeof image.al_rounded !== 'undefined' && image.al_rounded !== null) {
			file_options.push('RD' + image.al_rounded);
		}
	} else if (image.style == 'AC') {
		if (typeof image.ac_edge_depth !== 'undefined' && image.ac_edge_depth !== null) {
			file_options.push('ED' + image.ac_edge_depth);
		}
	}

	if (typeof image.frame_code !== 'undefined' && image.frame_code !== null) {
		file_options.push('F' + image.frame_code);
	}

	if (typeof image.wood_style !== 'undefined' && image.wood_style !== null) {
		file_options.push('W' + image.wood_style);
	}

	if (typeof image.background !== 'undefined' && image.background !== null) {
		file_options.push('BG' + image.background);
	}

	if (typeof image.background_texture !== 'undefined' && image.background_texture !== null) {
		file_options.push('BT' + image.background_texture);
	}

	if (typeof image.background_texture_color !== 'undefined' && image.background_texture_color !== null) {
		file_options.push('BTC' + image.background_texture_color);
	}

	if (image.shadow === true) {
		file_options.push('SHD');
	}

	if (typeof image.panoramic !== 'undefined' && image.panoramic === true) {
		file_options.push('PAN');
	} else {
		if (typeof image.pan !== 'undefined' && image.pan !== null) {
			file_options.push('P' + image.pan);
		}

		if (typeof image.tilt !== 'undefined' && image.tilt !== null) {
			file_options.push('T' + image.tilt);
		}

		if (typeof image.roll !== 'undefined' && image.roll !== null) {
			file_options.push('R' + image.roll);
		}
	}

	if (typeof image.actual_size !== 'undefined' && image.actual_size !== null) {
		file_options.push('A' + image.actual_size);
	}

	if (typeof image.size !== 'undefined' && image.size !== null) {
		file_options.push(image.size);
	}	

	// console.log(file_options);

	file_options = file_options.join('_') + '.' + image.format;

	if (typeof image.image_code !== 'undefined') {
		return render_location + '/v1/' + image.image_code + '_' + file_options;
	} else if (typeof image.url !== 'undefined') {
		var clean_url = encodeURIComponent(image.url);
		return render_location + '/v1/url/' + file_options + '?url=' + clean_url;
	}
};
