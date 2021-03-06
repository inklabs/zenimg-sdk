#  Copyright 2012 Ink Labs, LLC
#
#  v.1.0.9
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

__version__ = '1.0.0'
__render_location__ = 'http://i.zenimg.com'

import urllib

def get_img_url(params):
	image_code = params.get('image_code')
	url = params.get('url')
	style = params.get('style')
	cg_style = params.get('cg_style')
	cg_edge_color = params.get('cg_edge_color')
	cg_depth = params.get('cg_depth')
	al_edge_depth = params.get('al_edge_depth')
	al_rounded = params.get('al_rounded')
	ac_edge_depth = params.get('ac_edge_depth')
	frame_code = params.get('frame_code')
	frame_size = params.get('frame_size')
	mat_width = params.get('mat_width')
	min_mat_width = params.get('min_mat_width')
	mat_color = params.get('mat_color')
	curl = params.get('curl')
	curl_width = params.get('curl_width')
	wood_style = params.get('wood_style')
	ald_style = params.get('ald_style')
	forex_edge_color = params.get('forex_edge_color')
	background = params.get('background')
	background_texture = params.get('background_texture')
	background_texture_color = params.get('background_texture_color')
	shadow = params.get('shadow')
	disable_frame_shadow = params.get('disable_frame_shadow')
	pan = params.get('pan')
	tilt = params.get('tilt')
	roll = params.get('roll')
	actual_size = params.get('actual_size')
	size = params.get('size')
	format = params.get('format', 'jpg')

	file_options = [];

	file_options.append(style)

	if style is None:
		return ''

	if style == 'CG' or style == 'CG2':
		if cg_style == 'IW':
			file_options.append(cg_style)
	
		if cg_edge_color is not None:
			file_options.append('EC' + cg_edge_color.upper())

		if cg_depth is not None:
			file_options.append('D' + str(cg_depth))

	elif style == 'AL':
		if al_edge_depth is not None:
			file_options.append('ED' + str(al_edge_depth))

		if al_rounded is not None:
			file_options.append('RD' + str(al_rounded))

	elif style == 'FP':
		if frame_code is not None:
			file_options.append('F' + str(frame_code))

		if frame_size is not None:
			file_options.append('S' + str(frame_size))

		if disable_frame_shadow == True:
			file_options.append('NSHD')

		if min_mat_width is not None:
			file_options.append('MMW' + str(min_mat_width))

		if mat_color is not None:
			file_options.append('MC' + mat_color.upper())

	elif style == 'P':
		if curl is not None:
			file_options.append('C' + str(curl))

			if curl_width is not None:
				file_options.append('CW' + str(curl_width))

		else:
			if frame_code is not None:
				file_options.append('F' + str(frame_code))

			if mat_width is not None:
				file_options.append('MW' + str(mat_width))

			if mat_color is not None:
				file_options.append('MC' + mat_color.upper())

	elif style == 'AC':
		if ac_edge_depth is not None:
			file_options.append('ED' + str(ac_edge_depth))

	elif style == 'WD':
		pass

	elif style == 'ALD':
		if ald_style == 1:
			file_options.append('S' + str(ald_style))

	elif style == 'ACD':
		pass

	elif style == 'F':
		if forex_edge_color is not None:
			file_options.append('EC' + str(forex_edge_color))

	else:
	  return ''

	if wood_style is not None:
		file_options.append('W' + wood_style)

	if background is not None:
		file_options.append('BG' + str(background))

	if background_texture is not None:
		file_options.append('BT' + str(background_texture))

	if background_texture_color is not None:
		file_options.append('BTC' + background_texture_color.upper())

	if shadow == True:
		file_options.append('SHD')

	if pan is not None:
		file_options.append('P' + str(pan))

	if tilt is not None:
		file_options.append('T' + str(tilt))
		
	if roll is not None:
		file_options.append('R' + str(roll))

	if actual_size is not None:
		file_options.append('A' + actual_size.upper())

	if size is not None:
		file_options.append(size.upper())
	else:
		return ''

	file_options = '_'.join(file_options) + '.' + format

	if image_code is not None:
		return __render_location__ + '/v1/' + image_code + '_' + file_options
	elif url is not None:
		clean_url = urllib.quote_plus(url);
		return __render_location__ + '/v1/url/' + file_options + '?url=' + clean_url;

	return file_options
