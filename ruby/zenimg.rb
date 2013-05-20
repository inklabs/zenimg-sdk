#  Copyright 2012 Ink Labs, LLC
#
#  v.1.0.4
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

require 'cgi'

class Zenimg
	
	@@render_location = 'http://i.zenimg.com'
	@@pan_angles = [
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
	]

	def self.get_img_url(params)
		image_code = params['image_code']
		url = params['url']
		style = params['style']
		cg_style = params['cg_style']
		cg_edge_color = params['cg_edge_color']
		cg_depth = params['cg_depth']
		al_edge_depth = params['al_edge_depth']
		al_rounded = params['al_rounded']
		ac_edge_depth = params['ac_edge_depth']
		frame_code = params['frame_code']
		mat_width = params['mat_width']
		mat_color = params['mat_color']
		wood_style = params['wood_style']
		background = params['background']
		background_texture = params['background_texture']
		background_texture_color = params['background_texture_color']
		shadow = params['shadow']
		pan = params['pan']
		tilt = params['tilt']
		roll = params['roll']
		actual_size = params['actual_size']
		size = params['size']
		format = params['format'] || 'jpg'

		file_options = [];

		file_options << style

		case style
		when 'CG', 'CG2'
			if cg_style == 'IW'
				file_options << cg_style
			end

			if cg_edge_color != nil
				file_options << 'EC' + cg_edge_color.upcase
			end

			if cg_depth != nil
				file_options << 'D' + cg_depth.to_s
			end
		when 'P'
			if frame_code != nil
				file_options << 'F' + frame_code.to_s
			end

			if mat_width != nil
				file_options << 'MW' + mat_width.to_s
			end

			if mat_color != nil
				file_options << 'MC' + mat_color.upcase
			end
		when 'AL'
			if al_edge_depth != nil
				file_options << 'ED' + al_edge_depth.to_s
			end

			if al_rounded != nil
				file_options << 'RD' + al_rounded.to_s
			end
		when 'AC'
			if ac_edge_depth != nil
				file_options << 'ED' + ac_edge_depth.to_s
			end
		when 'WD'
		else
		  return ''
		end

		if wood_style != nil
			file_options << 'W' + wood_style
		end

		if background != nil
			file_options << 'BG' + background.to_s
		end

		if background_texture != nil
			file_options << 'BT' + background_texture.to_s
		end

		if background_texture_color != nil
			file_options << 'BTC' + background_texture_color.upcase
		end

		if shadow == true
			file_options << 'SHD'
		end

		if pan != nil
			file_options << 'P' + pan.to_s
		end
	
		if tilt != nil
			file_options << 'T' + tilt.to_s
		end
	
		if roll != nil
			file_options << 'R' + roll.to_s
		end

		if actual_size != nil
			file_options << 'A' + actual_size.upcase
		end

		if size != nil
			file_options << size.upcase
		else
		  return ''
		end

		file_options = file_options.join('_') + '.' + format

		if image_code != nil
			return @@render_location + '/v1/' + image_code + '_' + file_options
		elsif url != nil
			clean_url = CGI.escape(url);
			return @@render_location + '/v1/url/' + file_options + '?url=' + clean_url;
		end
		
		return ''
	end
end
