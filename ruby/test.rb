require_relative 'zenimg.rb'

puts Zenimg::get_img_url({
	:image_code => '6b0',
	:style => 'CG',
	:cg_style => 'IW',
	:cg_edge_color => '0f0',
	:cg_depth => 2,
	:size => '200X200',
	:format => 'jpg',
})

puts Zenimg::get_img_url({
	:image_code => '6b0',
	:style => 'AL',
	:al_edge_depth => 0,
	:al_rounded => 1,
	:size => '200X200',
	:format => 'jpg',
})

puts Zenimg::get_img_url({
	:image_code => '6b0',
	:style => 'AC',
	:ac_edge_depth => 0,
	:size => '200X200',
	:format => 'jpg',
})

puts Zenimg::get_img_url({
	:image_code => '6b0',
	:style => 'WD',
	:size => '200X200',
	:format => 'jpg',
})


puts Zenimg::get_img_url({
	:url => 'http://imgur.com/gallery/m6fak',
	:style => 'CG',
	:shadow => true,
	:pan => 25,
	:tilt => 25,
	:roll => 25,
	:actual_size => "24x24",
	:size => '200X200',
	:format => 'jpg',
})
