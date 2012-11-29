var fs = require('fs');
eval(fs.readFileSync('./zenimg.js').toString());

var test_path = (require('path').dirname(require('path').dirname(require.main.filename))) + '/tests.json';
var tests = require(test_path);

var status = 0;
tests.forEach(function(item) {
	var params = item[0];
	var expected = item[1];
	
	result = Zenimg.get_img_url(params);
	
	if (expected == result) {
 		console.log("\t[OK]\t" + result)
	} else {
		console.log("\t[FAIL]\t" + result)
		console.log("\tEXP-->\t" + expected)
		status = 1
	}
});

process.exit(status)
