/*
#  Copyright 2012 Ink Labs, LLC
#
#  v.1.0.3
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

var Zenimg = Zenimg || {};

Zenimg.get_img_url = function(params) {
	var render_location = (typeof params.render_location !== 'undefined') ? params.render_location : 'http://i.zenimg.com';
	var format = (typeof params.format !== 'undefined') ? params.format : 'jpg';

	var file_options = new Array();

	if (typeof params.style === 'undefined') {
		throw 'Style is required';
	}

	file_options.push(params.style);

	if (params.style == 'CG') {
		if (params.cg_style == 'IW') {
			file_options.push(params.cg_style);
		}
	
		if (typeof params.cg_edge_color !== 'undefined') {
			file_options.push('EC' + params.cg_edge_color.toUpperCase());
		}
	
		if (typeof params.cg_depth !== 'undefined') {
			file_options.push('D' + params.cg_depth);
		}
	} else if (params.style == 'AL') {
		if (typeof params.al_edge_depth !== 'undefined') {
			file_options.push('ED' + params.al_edge_depth);
		}
	
		if (typeof params.al_rounded !== 'undefined') {
			file_options.push('RD' + params.al_rounded);
		}
	} else if (params.style == 'AC') {
		if (typeof params.ac_edge_depth !== 'undefined') {
			file_options.push('ED' + params.ac_edge_depth);
		}
	}

	if (typeof params.frame_code !== 'undefined') {
		file_options.push('F' + params.frame_code);
	}
	
	if (typeof params.wood_style !== 'undefined') {
		file_options.push('W' + params.wood_style);
	}
	
	if (typeof params.background !== 'undefined') {
		file_options.push('BG' + params.background);
	}
	
	if (typeof params.background_texture != 'undefined') {
		file_options.push('BT' + params.background_texture);
	}
	
	if (typeof params.background_texture_color !== 'undefined') {
		file_options.push('BTC' + params.background_texture_color.toUpperCase());
	}
	
	if (params.shadow == true) {
		file_options.push('SHD');
	}

	if (typeof params.pan !== 'undefined') {
		file_options.push('P' + params.pan);
	}
	
	if (typeof params.tilt !== 'undefined') {
		file_options.push('T' + params.tilt);
	}
	
	if (typeof params.roll !== 'undefined') {
		file_options.push('R' + params.roll);
	}

	if (typeof params.actual_size !== 'undefined') {
		file_options.push('A' + params.actual_size.toUpperCase());
	}
	
	if (typeof params.size !== 'undefined') {
		file_options.push(params.size.toUpperCase());
	}	

	// console.log(file_options);

	file_options = file_options.join('_') + '.' + format;
	
	if (typeof params.image_code !== 'undefined') {
		return render_location + '/v1/' + params.image_code + '_' + file_options;
	} else if (typeof params.url !== 'undefined') {
		var clean_url = encodeURIComponent(params.url);
		return render_location + '/v1/url/' + file_options + '?url=' + clean_url;
	}
};
