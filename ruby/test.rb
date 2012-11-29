require_relative 'zenimg.rb'
require 'json'
require 'awesome_print'

test_path = File.expand_path File.dirname(File.dirname(__FILE__)) + '/tests.json'
tests =  JSON.parse(IO.read(test_path))

status = 0
tests.each{ |test|
  params = test[0]
  expected = test[1]
  
  result = Zenimg::get_img_url(params)
  
  	if expected == result
  		puts "\t[OK]\t" + result
  	else
  		puts "\t[FAIL]\t" + result
  		puts "\tEXP-->\t" + expected
  		status = 1
  	end
}

exit status
