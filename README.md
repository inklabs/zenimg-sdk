zenimg-sdk
==========


## Test Coverage

Test cases are stored in tests.json. A bash script can be executed to run all language tests like so:

```bash
$ ./test.sh 
PHP:    [OK]
Ruby:   [OK]
Python: [OK]
JS:     [OK]
```

Errors will show up like the following. This example error signals the missing upper-case X in the actual size.

```bash
$ ./test.sh 
PHP:    [OK]
Ruby:   [OK]
Python: [OK]
JS:
	[FAIL]	http://i.zenimg.com/v1/6b0_CG_A18x18_200X200.jpg
	EXP-->	http://i.zenimg.com/v1/6b0_CG_A18X18_200X200.jpg
```

Individual language tests can be run like so:

```bash
$ php php/test.php 
  [OK]  http://i.zenimg.com/v1/6b0_CG_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG2_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_AC_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_AL_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_WD_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_IW_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_EC0F0_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_D2_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_IW_EC0F0_D2_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG2_IW_EC0F0_D2_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_AC_ED0_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_AC_ED1.25_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_AL_ED0_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_AL_ED1.5_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_AL_ED0_RD1_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_BG0_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_BG1_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_BT1_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_BT1_BTCFF0000_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_SHD_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_P0_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_P-45_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_P45_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_T0_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_T-45_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_T45_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_R0_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_R-45_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_R45_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_P45_T45_R45_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_A18X18_200X200.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_A18X18_300X300.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_300X300.jpg
  [OK]  http://i.zenimg.com/v1/6b0_CG_IW_EC0F0_D0.75_SHD_P45_T45_R45_200X200.jpg
  [OK]  http://i.zenimg.com/v1/url/CG_200X200.jpg?url=http%3A%2F%2Fimgur.com%2Fgallery%2Fm6fak
  [OK]  
  [OK]  
  [OK]  
  [OK]  
  [OK]  
  [OK]  
```
