import os
import sys
import json
import zenimg

test_path = open(os.path.join(os.path.dirname(os.path.dirname(__file__)), 'tests.json'))
tests = json.load(test_path)

status = 0;
for test in tests:
	params = test[0]
	expected = test[1]
	
	result = zenimg.get_img_url(params)
	
	if expected == result:
		print "\t[OK]\t" + result
	else:
		print "\t[FAIL]\t" + result
		print "\tEXP-->\t" + expected
		status = 1

sys.exit(status)
