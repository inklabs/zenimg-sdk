/*
 * zenimg.js v1.0.2
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

Zenimg.get_img_url = function(params) {
	var render_location = params.render_location || 'http://i.zenimg.com';
	var image_code = params.image_code || null;
	var url = params.url || null;
	var style = params.style || null;
	var cg_style = params.cg_style || null;
	var cg_edge_color = params.cg_edge_color || null;
	var cg_depth = params.cg_depth || null;
	var al_edge_depth = params.al_edge_depth || null;
	var al_rounded = params.al_rounded || null;
	var ac_edge_depth = params.ac_edge_depth || null;
	var frame_code = params.frame_code || null;
	var wood_style = params.wood_style || null;
	var background = params.background || null;
	var background_texture = params.background_texture || null;
	var background_texture_color = params.background_texture_color || null;
	var shadow = params.shadow || null;
	var pan = params.pan || null;
	var tilt = params.tilt || null;
	var roll = params.roll || null;
	var actual_size = params.actual_size || null;
	var size = params.size || null;
	var format = params.format || 'jpg'

	var file_options = new Array();

	if (style == null) {
		throw 'Style is required';
	}
	
	if (format == null) {
		throw 'Format is required';
	}

	file_options.push(style);

	if (style == 'CG') {
		if (cg_style == 'IW') {
			file_options.push(cg_style);
		}
	
		if (cg_edge_color != null) {
			file_options.push('EC' + cg_edge_color);
		}
	
		if (cg_depth != null) {
			file_options.push('D' + cg_depth);
		}
	} else if (style == 'AL') {
		if (al_edge_depth != null) {
			file_options.push('ED' + al_edge_depth);
		}
	
		if (al_rounded != null) {
			file_options.push('RD' + al_rounded);
		}
	} else if (style == 'AC') {
		if (ac_edge_depth != null) {
			file_options.push('ED' + ac_edge_depth);
		}
	}

	if (frame_code != null) {
		file_options.push('F' + frame_code);
	}
	
	if (wood_style != null) {
		file_options.push('W' + wood_style);
	}
	
	if (background != null) {
		file_options.push('BG' + background);
	}
	
	if (background_texture != null) {
		file_options.push('BT' + background_texture);
	}
	
	if (background_texture_color != null) {
		file_options.push('BTC' + background_texture_color);
	}
	
	if (shadow == true) {
		file_options.push('SHD');
	}

	if (pan != null) {
		file_options.push('P' + pan);
	}
	
	if (tilt != null) {
		file_options.push('T' + tilt);
	}
	
	if (roll != null) {
		file_options.push('R' + roll);
	}

	if (actual_size != null) {
		file_options.push('A' + actual_size.toUpperCase());
	}
	
	if (size != null) {
		file_options.push(size.toUpperCase());
	}	

	// console.log(file_options);

	file_options = file_options.join('_') + '.' + format;

	if (image_code != null) {
		return render_location + '/v1/' + image_code + '_' + file_options;
	} else if (url != null) {
		var clean_url = encodeURIComponent(url);
		return render_location + '/v1/url/' + file_options + '?url=' + clean_url;
	}
};
