/*
	All right reserved https://github.com/hunterhug/marmot at 2016-2021
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
*/
package miner

import (
	"github.com/hunterhug/marmot/util"
	"math/rand"
	"strings"
)

// Ua Global User-Agent provide
var Ua = map[int]string{}

// UaInit User-Agent init
func UaInit() {
	Ua = map[int]string{
		0:   "Mozilla/5.0 (SymbianOS/9.1; U; [en]; SymbianOS/91 Series60/3.0) AppleWebkit/413 (KHTML, like Gecko) Safari/413",
		1:   "Mozilla/5.0 (SymbianOS/9.1; U; [en]; Series60/3.0 NokiaE60/4.06.0) AppleWebKit/413 (KHTML, like Gecko) Safari/413",
		2:   "Mozilla/5.0 (SymbianOS/9.2; U; Series60/3.1 NokiaN95/10.0.018; Profile/MIDP-2.0 Configuration/CLDC-1.1) AppleWebKit/413 (KHTML, like Gecko) Safari/413",
		3:   "Mozilla/5.0 (SymbianOS/9.3; U; Series60/3.2 NokiaE75-1/110.48.125 Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413",
		4:   "Mozilla/5.0 (SymbianOS/9.4; U; Series60/5.0 Nokia5800d-1/21.0.025; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413",
		5:   "Mozilla/5.0 (SymbianOS/9.4; Series60/5.0 NokiaN97-1/12.0.024; Profile/MIDP-2.1 Configuration/CLDC-1.1; en-us) AppleWebKit/525 (KHTML, like Gecko) BrowserNG/7.1.12344",
		6:   "Mozilla/4.0 (compatible; MSIE 4.0; MSN 2.5; Windows 95)",
		7:   "Mozilla/4.0 (compatible; MSIE 4.0; Windows 95)",
		8:   "Mozilla/4.0 (compatible; MSIE 4.01; MSN 2.5; MSN 2.5; Windows 98)",
		9:   "Mozilla/4.0 (compatible; MSIE 4.01; MSN 2.5; Windows 95)",
		10:  "Mozilla/4.0 (compatible; MSIE 4.01; MSN 2.5; Windows 98)",
		11:  "Mozilla/4.0 (compatible; MSIE 4.01; Windows 95)",
		12:  "Mozilla/4.0 (compatible; MSIE 4.01; Windows 95; Yahoo! enPAN Version Windows 95/NT CD-ROM Edition 1.0.)",
		13:  "Mozilla/4.0 (compatible; MSIE 4.01; Windows 98)",
		14:  "Mozilla/4.0 (compatible; MSIE 4.01; Windows 98; BIGLOBE)",
		15:  "Mozilla/4.0 (compatible; MSIE 4.01; Windows 98; canoncopyer)",
		16:  "Mozilla/4.0 (compatible; MSIE 4.01; Windows 98; Compaq)",
		17:  "Mozilla/4.0 (compatible; MSIE 5.0; Windows 95)",
		18:  "Mozilla/4.0 (compatible; MSIE 5.0; Windows 95; DigExt)",
		19:  "Mozilla/4.0 (compatible; MSIE 5.0; Windows 95; DigExt; i-CABLE)",
		20:  "Mozilla/4.0 (compatible; MSIE 5.0; Windows 95; DigExt; ocnie5-1)",
		21:  "Mozilla/4.0 (compatible; MSIE 5.0; Windows 95; Yahoo! enPAN Version Windows 95/NT CD-ROM Edition 1.0.)",
		22:  "Mozilla/4.0 (compatible; MSIE 5.0; Windows 98)",
		23:  "Mozilla/4.0 (compatible; MSIE 5.0; Windows 98; CNETHomeBuild051099)",
		24:  "Mozilla/4.0 (compatible; MSIE 5.0; Windows 98; DigExt)",
		25:  "Mozilla/4.0 (compatible; MSIE 5.0; Windows 98; DigExt; ocnie5-1)",
		26:  "Mozilla/4.0 (compatible; MSIE 5.0; Windows 98; wn_ie5_en_v1)",
		27:  "Mozilla/4.0 (compatible; MSIE 5.0; Windows 98; Yahoo! enPAN Version Windows 95/NT CD-ROM Edition 1.0.; DigExt)",
		28:  "Mozilla/4.0 (compatible; MSIE 5.0; Windows NT)",
		29:  "Mozilla/4.0 (compatible; MSIE 5.0; Windows NT; DigExt)",
		30:  "Mozilla/4.0 (compatible; MSIE 5.0; Windows 98; DigExt)",
		31:  "Mozilla/4.0 (compatible; MSIE 5.01; MSN 2.5; Windows 98)",
		32:  "Mozilla/4.0 (compatible; MSIE 5.01; Windows 95)",
		33:  "Mozilla/4.0 (compatible; MSIE 5.01; Windows 98)",
		34:  "Mozilla/4.0 (compatible; MSIE 5.01; Windows NT 5.0)",
		35:  "Mozilla/4.0 (compatible; MSIE 5.01; Windows NT 5.0; DigExt)",
		36:  "Mozilla/4.0 (compatible; MSIE 5.01; Windows NT 5.0; TUCOWS)",
		37:  "Mozilla/4.0 (compatible; MSIE 5.01; Windows NT)",
		38:  "Mozilla/4.0 (compatible; MSIE 5.01; Windows NT 5.0; .NET CLR 1.1.4322)",
		39:  "Mozilla/4.0 (compatible; MSIE 5.0; Mac_PowerPC)",
		40:  "Mozilla/4.0 (compatible; MSIE 5.16; Mac_PowerPC)",
		41:  "Mozilla/4.0 (compatible; MSIE 5.17; Mac_PowerPC)",
		42:  "Mozilla/4.0 (compatible; MSIE 5.22; Mac_PowerPC)",
		43:  "Mozilla/4.0 (compatible; MSIE 5.23; Mac_PowerPC)",
		44:  "Mozilla/4.0 (compatible; MSIE 5.5; Windows 95)",
		45:  "Mozilla/4.0 (compatible; MSIE 5.5; Windows 98)",
		46:  "Mozilla/4.0 (compatible; MSIE 5.5; Windows NT 5.0)",
		47:  "Mozilla/4.0 (compatible; MSIE 5.5; Windows NT 5.0; .NET CLR 1.0.3705)",
		48:  "Mozilla/4.0 (compatible; MSIE 5.5; Windows NT 5.0; .NET CLR 1.1.4322)",
		49:  "Mozilla/4.0 (compatible; MSIE 5.5; Windows NT 5.0; by TSG)",
		50:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1)",
		51:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.0.3705; .NET CLR 1.1.4322)",
		52:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322)",
		53:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322; .NET CLR 1.0.3705)",
		54:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322; .NET CLR 2.0.40607)",
		55:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322; .NET CLR 2.0.40607; .NET CLR 1.0.3705)",
		56:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322; NOKTURNAL KICKS ASS)",
		57:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; FDM;",
		58:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; Maxthon; .NET CLR 1.1.4322; .NET CLR 2.0.41115)",
		59:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; .NET CLR 1.1.4322)",
		60:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; .NET CLR 1",
		61:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)",
		62:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0)",
		63:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; .NET CLR 1",
		64:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows 98)",
		65:  "Mozilla/4.0 (compatible; MSIE 6.0; AOL 9.0; Windows NT 5.1; iebar; .NET CLR 1.0.3705)",
		66:  "Mozilla/4.0 (compatible; MSIE 6.0; Win32);",
		67:  "Mozilla/4.0 (compatible; MSIE 6.0; Win32); .NET CLR 1.0.3705)",
		68:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; Hotbar 4.4.6.0)",
		69:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; Win 9x 4.90)",
		70:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 4.0)",
		71:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 4.0; BVG",
		72:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0)",
		73:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; .NET CLR 1.0.3705; .NET CLR 1.1.4322)",
		74:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; .NET CLR 1.1.4322)",
		75:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; .NET CLR 1.1.4322; FDM)",
		76:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; {FF0C8E09-3C86-44CB-834A-B8CEEC80A1D7}; iOpus-I-M)",
		77:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; i-Nav 3.0.1.0F; .NET CLR 1.0.3705; .NET CLR 1.1.4322)",
		78:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; MathPlayer 2.0; .NET CLR 1.1.4322)",
		79:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; Maxthon; .NET CLR 1.1.4322)",
		80:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; T312461)",
		81:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0;)",
		82:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)",
		83:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1);",
		84:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; .NET CLR 1.0.3705)",
		85:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; .NET CLR 1.0.3705; .NET CLR 1.1.4322)",
		86:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; .NET CLR 1.1.4322)",
		87:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; .NET CLR 1.1.4322; .NET CLR 1.0.3705)",
		88:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; .NET CLR 2.0.40607)",
		89:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; Alexa Toolbar)",
		90:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; BrowserBob)",
		91:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; DFO-MPO Internet Explorer 6.0)",
		92:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; ENGINE; Mozilla/4.0 (compatible; MSIE 6.0; Windows NT))",
		93:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; ESB{F40811EE-DF17-4BC9-8785-B362ABF34098}; .NET CLR 1.1.4322)",
		94:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; FDM)",
		95:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; FTDv3 Browser; .NET CLR 1.0.3705; .NET CLR 1.1.4322)",
		96:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; FunWebProducts; .NET CLR 1.1.4322; .NET CLR 2.0.40607)",
		97:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; FunWebProducts; AtHome033)",
		98:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; HCI0449; .NET CLR 1.0.3705)",
		99:  "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; i-NavFourF; .NET CLR 1.1.4322)",
		100: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; Maxthon;",
		101: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; Maxthon; .NET CLR 1.1.4322)",
		102: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; MyIE2; .NET CLR 1.1.4322)",
		103: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; MyIE2; Maxthon; .NET CLR 1.1.4322)",
		104: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; Q312461; FunWebProducts; .NET CLR 1.1.4322)",
		105: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; Woningstichting Den Helder; .NET CLR 1.0.3705)",
		106: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.2; .NET CLR 1.1.4322)",
		107: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.2; .NET CLR 1.1.4322; .NET CLR 2.0.41115)",
		108: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.2; MyIE2; .NET CLR 1.1.4322)",
		109: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.2; MyIE2; Maxthon; .NET CLR 1.1.4322)",
		110: "Mozilla/4.0 (compatible; MSIE 6.0; Windows XP)",
		111: "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1)",
		112: "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; SV1; .NET CLR 1.0.3705; .NET CLR 2.0.50727; .NET CLR 1.1.4322)",
		113: "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322)",
		114: "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; SV1)",
		115: "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1)",
		116: "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; .NET CLR 1.1.4322; .NET CLR 2.0.50727; .NET CLR 3.0.04506.30)",
		117: "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; SLCC1; .NET CLR 2.0.50727; Media Center PC 5.0; .NET CLR 3.0.04506)",
		118: "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; SLCC1; .NET CLR 2.0.50727; Media Center PC 5.0; .NET CLR 3.0.04506; InfoPath.1)",
		119: "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
		120: "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; .NET CLR 2.0.50727; .NET CLR 3.0.04506.30; .NET CLR 3.0.04506.648)",
		121: "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; .NET CLR 2.0.50727; InfoPath.1",
		122: "Mozilla/4.0 (compatible; GoogleToolbar 5.0.2124.2070; Windows 6.0; MSIE 8.0.6001.18241)",
		123: "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; EasyBits Go v1.0; InfoPath.1; .NET CLR 2.0.50727; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729)",
		124: "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0)",
		125: "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; Sleipnir/2.9.8)",
		126: "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.0; Trident/5.0)",
		127: "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0)",
		128: "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)",
		129: "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)",
		130: "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; WOW64; Trident/6.0)",
		131: "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)",
		132: "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0)",
		133: "Mozilla/5.0 (Windows; U; Windows NT 6.0; en-US) AppleWebKit/525.13 (KHTML, like Gecko) Chrome/0.2.149.27 Safari/525.13",
		134: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/525.19 (KHTML, like Gecko) Chrome/1.0.154.48 Safari/525.19",
		135: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/530.5 (KHTML, like Gecko) Chrome/2.0.172.33 Safari/530.5",
		136: "Mozilla/5.0 (Windows; U; Windows NT 6.0; en-US) AppleWebKit/532.0 (KHTML, like Gecko) Chrome/3.0.195.38 Safari/532.0",
		137: "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_3; en-US) AppleWebKit/533.4 (KHTML, like Gecko) Chrome/5.0.375.55 Safari/533.4",
		138: "Mozilla/5.0 (X11; U; Linux i686; en-US) AppleWebKit/534.3 (KHTML, like Gecko) Chrome/6.0.472.63 Safari/534.3",
		139: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/534.3 (KHTML, like Gecko) Chrome/6.0.472.55 Safari/534.3",
		140: "Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/534.7 (KHTML, like Gecko) Chrome/7.0.517.43 Safari/534.7",
		141: "Mozilla/5.0 (Windows; U; Windows NT 6.0; en-US) AppleWebKit/534.10 (KHTML, like Gecko) Chrome/8.0.552.224 Safari/534.10 ChromePlus/1.5.2.0",
		142: "Mozilla/5.0 (en-us) AppleWebKit/534.14 (KHTML, like Gecko; Google Wireless Transcoder) Chrome/9.0.597 Safari/534.14",
		143: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/534.16 (KHTML, like Gecko) Chrome/10.0.648.151 Safari/534.16",
		144: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/534.24 (KHTML, like Gecko) Chrome/11.0.696.71 Safari/534.24",
		145: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/534.24 (KHTML, like Gecko) Iron/11.0.700.2 Chrome/11.0.700.2 Safari/534.24",
		146: "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/534.24 (KHTML, like Gecko) Chrome/11.0.696.65 Safari/534.24",
		147: "Mozilla/5.0 (Windows NT 5.1) AppleWebKit/534.30 (KHTML, like Gecko) Chrome/12.0.742.122 Safari/534.30 ChromePlus/1.6.3.1",
		148: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/534.30 (KHTML, like Gecko) Chrome/12.0.742.112 Safari/534.30",
		149: "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/13.0.782.107 Safari/535.1",
		150: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/535.1 (KHTML, like Gecko) RockMelt/0.9.64.361 Chrome/13.0.782.218 Safari/535.1",
		151: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/13.0.782.112 Safari/535.1",
		152: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/13.0.782.220 Safari/535.1",
		153: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.202 Safari/535.1",
		154: "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.202 Safari/535.1",
		155: "Mozilla/5.0 (Windows NT 5.1) AppleWebKit/535.2 (KHTML, like Gecko) Chrome/15.0.874.120 Safari/535.2",
		156: "Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.2 (KHTML, like Gecko) Ubuntu/10.04 Chromium/15.0.874.106 Chrome/15.0.874.106 Safari/535.2",
		157: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_2) AppleWebKit/535.2 (KHTML, like Gecko) Chrome/15.0.874.106 Safari/535.2",
		158: "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.7 (KHTML, like Gecko) Chrome/16.0.912.75 Safari/535.7",
		159: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/535.7 (KHTML, like Gecko) Chrome/16.0.912.75 Safari/535.7",
		160: "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; SLCC1; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30618; Lunascape 4.7.3)",
		161: "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; .NET CLR 1.1.4322; .NET CLR 2.0.50727; Lunascape 5.0 alpha2)",
		162: "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.1; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; Lunascape 5.0 alpha2)",
		163: "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1) ; SLCC1; .NET CLR 2.0.50727; Media Center PC 5.0; .NET CLR 3.0.04506; Tablet PC 2.0; Lunascape 5.0 alpha2)",
		164: "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; MAGW; .NET4.0C; Lunascape 6.5.8.24780)",
		165: "Mozilla/2.02 (Macintosh; I; PPC)",
		166: "Mozilla/3.01 (Macintosh; I; PPC)",
		167: "Mozilla/4.01 (Macintosh; I; PPC)",
		168: "Mozilla/4.03 [en]C-IMS (Win95; I)",
		169: "Mozilla/4.04 [en] (Win95; I ;Nav)",
		170: "Mozilla/4.04 [en] (Win95; I)",
		171: "Mozilla/4.04 [en] (WinNT; I ;Nav)",
		172: "Mozilla/4.04 [en] (WinNT; I)",
		173: "Mozilla/4.04 [en] (Macintosh; I; PPC Nav)",
		174: "Mozilla/4.04 [en] (X11; I; SunOS 5.5 sun4u)",
		175: "Mozilla/4.05 [en] (Win95; I)",
		176: "Mozilla/4.05 (Macintosh; I; PPC)",
		177: "Mozilla/4.06 [en] (Win98; I)",
		178: "Mozilla/4.06 [en] (Macintosh; I; PPC)",
		179: "Mozilla/4.08 (Macintosh; I; PPC)",
		180: "Mozilla/4.5 [en] (Win95; I)",
		181: "Mozilla/4.5 [en] (Win98; I)",
		182: "Mozilla/4.5 [en] (WinNT; I)",
		183: "Mozilla/4.5 (Macintosh; I; PPC)",
		184: "Mozilla/4.51 [en] (Win95; I)",
		185: "Mozilla/4.51 [en] (Win98; I)",
		186: "Mozilla/4.51 [en] (WinNT; I)",
		187: "Mozilla/4.51 [en] (X11; I; SunOS 5.8 sun4u)",
		188: "Mozilla/4.6 [en] (Win95; I)",
		189: "Mozilla/4.6 [en] (Win98; I)",
		190: "Mozilla/4.6 [en] (WinNT; I)",
		191: "Mozilla/4.6 [en] (WinNT; I)",
		192: "Mozilla/4.7 [en] (WinNT; I)",
		193: "Mozilla/4.7 [en] (Win95; I)",
		194: "Mozilla/4.7 [en] (Win98; I)",
		195: "Mozilla/4.7 [en] (WinNT; I)",
		196: "Mozilla/4.7 [en] (WinNT; I)",
		197: "Mozilla/4.7 [en] (WinNT; U)",
		198: "Mozilla/4.7 [en] (Macintosh; I; PPC)",
		199: "Mozilla/4.76 [en_jp] (X11; U; SunOS 5.8 sun4u)",
		200: "Mozilla/4.76 [en] (X11; U; SunOS 5.8 sun4u)",
		201: "Mozilla/4.78 [en] (X11; U; SunOS 5.9 sun4u)",
		202: "Mozilla/4.8 [en] (X11; U; SunOS 5.7 sun4u)",
		203: "Mozilla/5.0 (Windows; U; Win98; en-JP; m18) Gecko/20001108 Netscape6/6.0",
		204: "Mozilla/5.0 (Windows; U; Windows NT 5.0; en-JP; m18) Gecko/20010131 Netscape6/6.01",
		205: "Mozilla/5.0 (Windows; U; Win 9x 4.90; en-JP; rv:0.9.4) Gecko/20011128 Netscape6/6.2.1",
		206: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-JP; rv:0.9.4.1) Gecko/20020508 Netscape6/6.2.3",
		207: "Mozilla/5.0 (Macintosh; N; PPC; en-JP; macen-pub12) Gecko/20001108 Netscape6/6.0",
		208: "Mozilla/5.0 (Macintosh; U; PPC; en-JP; rv:0.9.2) Gecko/20010726 Netscape6/6.1",
		209: "Mozilla/5.0 (Macintosh; U; PPC; en-JP; rv:0.9.4) Gecko/20011022 Netscape6/6.2",
		210: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-US; rv:0.9.4.1) Gecko/20020315 Netscape6/6.2.2",
		211: "Mozilla/5.0 (Windows; U; Win98; en-US; rv:1.0.2) Gecko/20030208 Netscape/7.02",
		212: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-JP; rv:1.4) Gecko/20030624 Netscape/7.1 (ax)",
		213: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.4) Gecko/20030624 Netscape/7.1 (ax)",
		214: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.7.2) Gecko/20040804 Netscape/7.2 (ax)",
		215: "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.7.2) Gecko/20040805 Netscape/7.2",
		216: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-JP; rv:1.0.2) Gecko/20021120 Netscape/7.01",
		217: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X Mach-O; en-JP; rv:1.4) Gecko/20030624 Netscape/7.1",
		218: "Mozilla/5.0 (X11; U; SunOS sun4u; en-JP; rv:1.0.1) Gecko/20020921 Netscape/7.0",
		219: "Mozilla/5.0 (Windows; U; Windows NT 5.0; en-JP; rv:1.5) Gecko/20031007",
		220: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.6) Gecko/20040113",
		221: "Mozilla/5.0 (Windows; U; Windows NT 5.0; en-US; rv:1.7) Gecko/20040616",
		222: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X Mach-O; en-US; rv:1.6) Gecko/20040113",
		223: "Mozilla/5.0 (X11; U; Linux i686; en-JP; rv:1.2.1) Gecko/20030225",
		224: "Mozilla/5.0 (X11; U; Linux i686; en-JP; rv:1.4.1) Gecko/20031030",
		225: "Mozilla/5.0 (X11; U; FreeBSD i386; en-US; rv:1.7.1) Gecko/20040805",
		226: "Mozilla/5.0 (X11; U; SunOS sun4u; en-JP; rv:1.4) Gecko/20040414",
		227: "Mozilla/5.0 (X11; U; Linux i686; rv:1.7.3) Gecko/20040913 Firefox/0.10",
		228: "Mozilla/5.0 (Windows; U; Windows NT 5.0; rv:1.7.3) Gecko/20040913 Firefox/0.10",
		229: "Mozilla/5.0 (Windows; U; Windows NT 5.1; rv:1.7.3) Gecko/20040913 Firefox/0.10",
		230: "Mozilla/5.0 (X11; U; Linux i686; rv:1.7.3) Gecko/20041001 Firefox/0.10.1",
		231: "Mozilla/5.0 (Windows; U; Windows NT 5.1; rv:1.7.3) Gecko/20041001 Firefox/0.10.1",
		232: "Mozilla/5.0 (Windows; U; Windows NT 5.0; rv:1.7.3) Gecko/20041001 Firefox/0.10.1",
		233: "Mozilla/5.0 (Windows; U; Windows NT 5.2; rv:1.7.3) Gecko/20041001 Firefox/0.10.1",
		234: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-JP; rv:1.6) Gecko/20040206 Firefox/0.8",
		235: "Mozilla/5.0 (Windows; U; Win98; en-US; rv:1.6) Gecko/20040206 Firefox/0.8",
		236: "Mozilla/5.0 (X11; U; Linux i686; en-JP; rv:1.6) Gecko/20040207 Firefox/0.8",
		237: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X Mach-O; en-US; rv:1.6) Gecko/20040206 Firefox/0.8",
		238: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-JP; rv:1.7) Gecko/20040614 Firefox/0.9",
		239: "Mozilla/5.0 (X11; U; Linux i686; es-ES; rv:1.7) Gecko/20040708 Firefox/0.9",
		240: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-JP; rv:1.7) Gecko/20040707 Firefox/0.9.2",
		241: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.7.5) Gecko/20041107 Firefox/0.9.2 StumbleUpon/1.994",
		242: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-JP; rv:1.7) Gecko/20040803 Firefox/0.9.3",
		243: "Mozilla/5.0 (Windows; U; Windows NT 5.0; en-US; rv:1.7) Gecko/20040803 Firefox/0.9.3",
		244: "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.7) Gecko/20040803 Firefox/0.9.3",
		245: "Mozilla/5.0 (Windows; U; Win98; en-US; rv:1.7) Gecko/20040803 Firefox/0.9.3",
		246: "Mozilla/5.0 (Windows; U; Windows NT 5.0; en-US; rv:1.7) Gecko/20040803 Firefox/0.9.3",
		247: "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.4) Gecko/20040803 Firefox/0.9.3",
		248: "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.7) Gecko/20041013 Firefox/0.9.3 (Ubuntu)",
		249: "Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.7) Gecko/20041013 Firefox/0.9.3 (Ubuntu)",
		250: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X Mach-O; en-US; rv:1.7.5) Gecko/20041107 Firefox/1.0",
		251: "Mozilla/5.0 (Windows; U; Win 9x 4.90; nl-NL; rv:1.7.5) Gecko/20041202 Firefox/1.0",
		252: "Mozilla/5.0 (Windows; U; Win 9x 4.90; nl-NL; rv:1.7.5) Gecko/20041202 Firefox/1.0",
		253: "Mozilla/5.0 (Windows; U; Win98; nl-NL; rv:1.7.5) Gecko/20041202 Firefox/1.0",
		254: "Mozilla/5.0 (Windows; U; Windows NT 5.0; de-DE; rv:1.7.5) Gecko/20041108 Firefox/1.0",
		255: "Mozilla/5.0 (Windows; U; Windows NT 5.0; de-DE; rv:1.7.5) Gecko/20041122 Firefox/1.0",
		256: "Mozilla/5.0 (Windows; U; Windows NT 5.0; en-GB; rv:1.7.5) Gecko/20041107 Firefox/1.0",
		257: "Mozilla/5.0 (Windows; U; Windows NT 5.0; en-GB; rv:1.7.5) Gecko/20041110 Firefox/1.0",
		258: "Mozilla/5.0 (Windows; U; Windows NT 5.0; en-US; rv:1.7.5) Gecko/20041107 Firefox/1.0",
		259: "Mozilla/5.0 (Windows; U; Windows NT 5.0; fr-FR; rv:1.7.5) Gecko/20041108 Firefox/1.0",
		260: "Mozilla/5.0 (Windows; U; Windows NT 5.0; it-IT; rv:1.7.5) Gecko/20041110 Firefox/1.0",
		261: "Mozilla/5.0 (Windows; U; Windows NT 5.1; de-DE; rv:1.7.5) Gecko/20041107 Firefox/1.0",
		262: "Mozilla/5.0 (Windows; U; Windows NT 5.1; de-DE; rv:1.7.5) Gecko/20041108 Firefox/1.0",
		263: "Mozilla/5.0 (Windows; U; Windows NT 5.1; de-DE; rv:1.7.5) Gecko/20041122 Firefox/1.0",
		264: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-GB; rv:1.7.5) Gecko/20041110 Firefox/1.0",
		265: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.7.5) Gecko/20041107 Firefox/1.0",
		266: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.7.5) Gecko/20041107 Firefox/1.0 StumbleUpon/1.999",
		267: "Mozilla/5.0 (Windows; U; Windows NT 5.1; es-ES; rv:1.7.5) Gecko/20041210 Firefox/1.0",
		268: "Mozilla/5.0 (Windows; U; Windows NT 5.1; fr-FR; rv:1.7.5) Gecko/20041108 Firefox/1.0",
		269: "Mozilla/5.0 (Windows; U; Windows NT 5.1; nl-NL; rv:1.7.5) Gecko/20041202 Firefox/1.0",
		270: "Mozilla/5.0 (Windows; U; Windows NT 5.1; sv-SE; rv:1.7.5) Gecko/20041108 Firefox/1.0",
		271: "Mozilla/5.0 (Windows; U; Windows NT 5.2; en-US; rv:1.7.5) Gecko/20041107 Firefox/1.0",
		272: "Mozilla/5.0 (Windows; U; Windows NT 5.2; en-US; rv:1.8b) Gecko/20050212 Firefox/1.0+ (MOOX M3)",
		273: "Mozilla/5.0 (Windows; U; WinNT4.0; en-US; rv:1.7.5) Gecko/20041107 Firefox/1.0",
		274: "Mozilla/5.0 (Windows; U; Windows NT 5.0; en-US; rv:1.8b) Gecko/20050118 Firefox/1.0+",
		275: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8b) Gecko/20050118 Firefox/1.0+",
		276: "Mozilla/5.0 (X11; U; FreeBSD i386; en-US; rv:1.7.5) Gecko/20050103 Firefox/1.0",
		277: "Mozilla/5.0 (X11; U; Linux i386; en-US; rv:1.7.5) Gecko/20041109 Firefox/1.0",
		278: "Mozilla/5.0 (X11; U; Linux i686; de-DE; rv:1.7.5) Gecko/20041128 Firefox/1.0 (Debian package 1.0-4)",
		279: "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.7.5) Gecko/20041107 Firefox/1.0",
		280: "Mozilla/5.0 (X11; U; Linux i686; en-JP; rv:1.7.5) Gecko/20041108 Firefox/1.0",
		281: "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.7.5) Gecko/20041111 Firefox/1.0",
		282: "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.7.5) Gecko/20041111 Firefox/1.0 (Debian package 1.0-2)",
		283: "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.7.5) Gecko/20041116 Firefox/1.0 (Ubuntu) (Ubuntu package 1.0-2ubuntu3)",
		284: "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.7.5) Gecko/20041119 Firefox/1.0",
		285: "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.7.5) Gecko/20041123 Firefox/1.0",
		286: "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.7.5) Gecko/20041128 Firefox/1.0 (Debian package 1.0-4)",
		287: "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.7.5) Gecko/20041130 Firefox/1.0",
		288: "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.7.5) Gecko/20041214 Firefox/1.0 StumbleUpon/1.999",
		289: "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.7.5) Gecko/20041219 Firefox/1.0 (Debian package 1.0+dfsg.1-1)",
		290: "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.7.5) Gecko/20050110 Firefox/1.0 (Debian package 1.0+dfsg.1-2)",
		291: "Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.7.5) Gecko/20041107 Firefox/1.0",
		292: "Mozilla/5.0 (X11; U; Linux i686; nl-NL; rv:1.7.5) Gecko/20050221 Firefox/1.0 (Ubuntu) (Ubuntu package 1.0+dfsg.1-6ubuntu1)",
		293: "Mozilla/5.0 (X11; U; Linux i686; rv:1.8b) Gecko/20050124 Firefox/1.0+",
		294: "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.7.6) Gecko/20050306 Firefox/1.0.1 (Debian package 1.0.1-2)",
		295: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.7.6) Gecko/20050223 Firefox/1.0.1",
		296: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.7.6) Gecko/20050225 Firefox/1.0.1",
		297: "Mozilla/5.0 (Windows; U; Windows NT 5.1; de-DE; rv:1.7.6) Gecko/20050226 Firefox/1.0.1",
		298: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X Mach-O; en-JP-mac; rv:1.8) Gecko/20051111 Firefox/1.5",
		299: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en; rv:1.8) Gecko/20051111 Firefox/1.5",
		300: "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.8) Gecko/20051111 Firefox/1.5",
		301: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en; rv:1.8.0.1) Gecko/20060111 Firefox/1.5.0.1",
		302: "Mozilla/5.0 (X11; U; Linux i686; en; rv:1.8.0.2) Gecko/20060308 Firefox/1.5.0.2",
		303: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en; rv:1.8.0.3) Gecko/20060426 Firefox/1.5.0.3",
		304: "Mozilla/5.0 (X11; U; Linux i686; en; rv:1.8.0.4) Gecko/20060508 Firefox/1.5.0.4",
		305: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en; rv:1.8.1.20) Gecko/20081217 Firefox/2.0.0.20",
		306: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.12) Gecko/20080201 Firefox/2.0.0.12",
		307: "Mozilla/5.0 (Macintosh; U; Intel Mac OS X; en-JP-mac; rv:1.8.1.20) Gecko/20081217 Firefox/2.0.0.20",
		308: "Mozilla/5.0 (Windows; U; Windows NT 5.1; en; rv:1.9.0.6) Gecko/2009011913 Firefox/3.0.6",
		309: "Mozilla/5.0 (Windows; U; Windows NT 6.0; en; rv:1.9.0.6) Gecko/2009011913 Firefox/3.0.6 (.NET CLR 3.5.30729)",
		310: "Mozilla/5.0 (Windows; U; Windows NT 6.0; en; rv:1.9.0.17) Gecko/2009122116 Firefox/3.0.17 GTB6 (.NET CLR 3.5.30729)",
		311: "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-JP-mac; rv:1.9.0.6) Gecko/2009011912 Firefox/3.0.6",
		312: "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-JP-mac; rv:1.9.0.6) Gecko/2009011912 Firefox/3.0.6 GTB5",
		313: "Mozilla/5.0 (Windows NT 6.1; rv:2.0) Gecko/20100101 Firefox/4.0",
		314: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.5; rv:2.0) Gecko/20100101 Firefox/4.0",
		315: "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:5.0) Gecko/20100101 Firefox/5.0",
		316: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.5; rv:5.0.1) Gecko/20100101 Firefox/5.0.1",
		317: "Mozilla/5.0 (Windows NT 5.1; rv:6.0) Gecko/20100101 Firefox/6.0",
		318: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.6; rv:6.0.2) Gecko/20100101 Firefox/6.0.2",
		319: "Mozilla/5.0 (Windows NT 6.0; rv:7.0.1) Gecko/20100101 Firefox/7.0.1",
		320: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.7; rv:7.0.1) Gecko/20100101 Firefox/7.0.1",
		321: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.5; rv:8.0.1) Gecko/20100101 Firefox/8.0.1",
		322: "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:9.0.1) Gecko/20100101 Firefox/9.0.1",
		323: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.7; rv:9.0.1) Gecko/20100101 Firefox/9.0.1",
		324: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-jp) AppleWebKit/85.7 (KHTML, like Gecko) Safari/85.6",
		325: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-jp) AppleWebKit/85.7 (KHTML, like Gecko) Safari/85.7",
		326: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; de-de) AppleWebKit/85.8.2 (KHTML, like Gecko) Safari/85.8",
		327: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-us) AppleWebKit/85.8.2 (KHTML, like Gecko) Safari/85.8.1",
		328: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-us) AppleWebKit/85.8.5 (KHTML, like Gecko) Safari/85.8",
		329: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-jp) AppleWebKit/103u (KHTML, like Gecko) Safari/100.1",
		330: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-jp) AppleWebKit/124 (KHTML, like Gecko) Safari/125.1",
		331: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-jp) AppleWebKit/125.2 (KHTML, like Gecko) Safari/125.8",
		332: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-jp) AppleWebKit/125.4 (KHTML, like Gecko) Safari/125.9",
		333: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-us) AppleWebKit/125.5.5 (KHTML, like Gecko) Safari/125.11",
		334: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; de-de) AppleWebKit/125.5.5 (KHTML, like Gecko) Safari/125.12",
		335: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-us) AppleWebKit/125.5.5 (KHTML, like Gecko) Safari/125.12",
		336: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; nl-nl) AppleWebKit/125.5.5 (KHTML, like Gecko) Safari/125.12",
		337: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en) AppleWebKit/125.5.6 (KHTML, like Gecko) Safari/125.12",
		338: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-jp) AppleWebKit/312.1 (KHTML, like Gecko) Safari/312",
		339: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-jp) AppleWebKit/312.1.1 (KHTML, like Gecko) Safari/312",
		340: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-jp) AppleWebKit/312.5 (KHTML, like Gecko) Safari/312.3",
		341: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-jp) AppleWebKit/312.8 (KHTML, like Gecko) Safari/312.5",
		342: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-jp) AppleWebKit/312.8 (KHTML, like Gecko) Safari/312.6",
		343: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-jp) AppleWebKit/412 (KHTML, like Gecko) Safari/412",
		344: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-jp) AppleWebKit/412.6.2 (KHTML, like Gecko) Safari/125.11",
		345: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-jp) AppleWebKit/412.6.2 (KHTML, like Gecko) Safari/412.2.2",
		346: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-jp) AppleWebKit/412.7 (KHTML, like Gecko) Safari/412.5",
		347: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-jp) AppleWebKit/416.11 (KHTML, like Gecko) Safari/416.12",
		348: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-jp) AppleWebKit/418 (KHTML, like Gecko) Safari/417.9.2",
		349: "Mozilla/5.0 (Macintosh; U; Intel Mac OS X; en-jp) AppleWebKit/418 (KHTML, like Gecko) Safari/417.9.3",
		350: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-jp) AppleWebKit/522.11.1 (KHTML, like Gecko) Version/3.0.3 Safari/522.12.1",
		351: "Mozilla/5.0 (Macintosh; U; Intel Mac OS X; en-jp) AppleWebKit/523.10.3 (KHTML, like Gecko) Version/3.0.4 Safari/523.10",
		352: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-jp) AppleWebKit/523.10.6 (KHTML, like Gecko) Version/3.0.4 Safari/523.10.6",
		353: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-jp) AppleWebKit/523.12 (KHTML, like Gecko) Version/3.0.4 Safari/523.12",
		354: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-jp) AppleWebKit/523.12.2 (KHTML, like Gecko) Version/3.0.4 Safari/523.12.2",
		355: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_4_11; en-jp) AppleWebKit/525.13 (KHTML, like Gecko) Version/3.1 Safari/525.13",
		356: "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_3; en-jp) AppleWebKit/525.18 (KHTML, like Gecko) Version/3.1.1 Safari/525.20",
		357: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_4_11; en-jp) AppleWebKit/525.18 (KHTML, like Gecko) Version/3.1.2 Safari/525.22",
		358: "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_5; en-jp) AppleWebKit/525.26.2 (KHTML, like Gecko) Version/3.2 Safari/525.26.12",
		359: "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_6; en-jp) AppleWebKit/525.27.1 (KHTML, like Gecko) Version/3.2.1 Safari/525.27.1",
		360: "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_6; en-jp) AppleWebKit/528.16 (KHTML, like Gecko) Version/4.0 Safari/528.16",
		361: "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_2; en-jp) AppleWebKit/531.21.8 (KHTML, like Gecko) Version/4.0.4 Safari/531.21.10",
		362: "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_3; en-jp) AppleWebKit/531.21.11 (KHTML, like Gecko) Version/4.0.4 Safari/531.21.10",
		363: "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_3; en-jp) AppleWebKit/533.16 (KHTML, like Gecko) Version/5.0 Safari/533.16",
		364: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/534.52.7 (KHTML, like Gecko) Version/5.1.2 Safari/534.52.7",
		365: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/534.57.2 (KHTML, like Gecko) Version/5.1.7 Safari/534.57.2",
		366: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8) AppleWebKit/536.25 (KHTML, like Gecko) Version/6.0 Safari/536.25",
		367: "Mozilla/4.0 (Windows 95;US) Opera 3.62 [en]",
		368: "Mozilla/4.0 (compatible; MSIE 5.0; Mac_PowerPC) Opera 6.0 [en]",
		369: "Mozilla/4.0 (compatible; MSIE 5.0; Windows 2000) Opera 6.01 [en]",
		370: "Mozilla/4.0 (compatible; MSIE 5.0; Windows ME) Opera 6.03 [en]",
		371: "Mozilla/4.0 (compatible; MSIE 5.0; Windows 2000) Opera 6.05 [en]",
		372: "Mozilla/4.0 (compatible; MSIE 5.0; Windows XP) Opera 6.06 [en]",
		373: "Mozilla/4.0 (compatible; MSIE 5.0; Windows 2000) Opera 6.06 [en]",
		374: "Opera 7.11",
		375: "Opera/7.23 (Windows NT 5.0; U) [en]",
		376: "Opera/7.52 (Windows NT 5.1; U) [en]",
		377: "Opera/7.53 (Windows NT 5.0; U) [en]",
		378: "Opera/7.54 (Windows NT 5.0; U) [en]",
		379: "Opera/7.54 (Windows NT 5.1; U) [en]",
		380: "Opera/7.54 (Windows NT 5.1; U)",
		381: "Opera/7.54 (X11; Linux i686; U) [en]",
		382: "Opera/7.54 (X11; Linux i686; U) [sv]",
		383: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1) Opera 7.11 [en]",
		384: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1) Opera 7.21 [pt-BR]",
		385: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1) Opera 7.22 [en]",
		386: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0) Opera 7.23 [en]",
		387: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0) Opera 7.23 [en]",
		388: "Mozilla/4.0 (compatible; MSIE 6.0; X11; Linux i686) Opera 7.23 [en]",
		389: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1) Opera 7.23 [fr]",
		390: "Mozilla/4.0 (compatible; MSIE 6.0; Mac_PowerPC) Opera 7.50 [en]",
		391: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0) Opera 7.53 [en]",
		392: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1) Opera 7.53 [en]",
		393: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1) Opera 7.54 [en]",
		394: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0) Opera 7.54 [en]",
		395: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1) Opera 7.54 [en]",
		396: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0) Opera 7.54u1 [en]",
		397: "Opera/7.54 (Windows 98; U) [en]",
		398: "Mozilla/4.78 (Windows NT 5.1; U) Opera 7.23 [en]",
		399: "Mozilla/5.0 (Windows NT 5.0; U) Opera 7.54 [en]",
		400: "Mozilla/4.0 (compatible; MSIE 6.0; X11; OpenBSD i386) Opera 7.54 [en]",
		401: "Opera/8.0 (X11; Linux i686; U; en)",
		402: "Opera/8.01 (Windows ME; U; en)",
		403: "Opera/8.01 (Windows NT 5.1; U; en)",
		404: "Mozilla/4.0 (compatible; MSIE 6.0; Mac_PowerPC Mac OS X; en) Opera 8.01",
		405: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; en) Opera 8.01",
		406: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; tr) Opera 8.02",
		407: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; en) Opera 8.02",
		408: "Opera/8.5 (Windows NT 5.0; U; en)",
		409: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; en) Opera 8.5",
		410: "Opera/8.51 (Windows NT 5.1; U; en)",
		411: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; en) Opera 8.51",
		412: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; de) Opera 8.52",
		413: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; pl) Opera 8.52",
		414: "Mozilla/5.0 (Windows NT 5.1; U; en) Opera 8.52",
		415: "Opera/8.53 (Windows NT 5.1; U; en)",
		416: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; en) Opera 8.53",
		417: "Opera/8.54 (Windows NT 5.1; U; en)",
		418: "Opera/8.54 (Windows NT 5.0; U; en)",
		419: "Mozilla/5.0 (X11; Linux i686; U; cs) Opera 8.54",
		420: "Mozilla/4.0 (compatible; MSIE 6.0; KDDI-SA39) Opera 8.60 [en]",
		421: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; en) Opera 9.00",
		422: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; en) Opera 9.00",
		423: "Opera/9.00 (Windows NT 5.1; U; en)",
		424: "Opera/9.0 (Windows NT 5.1; U; en)",
		425: "Opera/9.00 (Macintosh; PPC Mac OS X; U; en)",
		426: "Opera/9.02 (Macintosh; PPC Mac OS X; U; en)",
		427: "Opera/9.02 (Windows NT 5.1; U; zh-tw)",
		428: "Opera/9.10 (Windows NT 6.0; U; en)",
		429: "Opera/9.21 (Windows NT 6.0; U; en)",
		430: "Opera/9.22 (Windows NT 5.1; U; en)",
		431: "Opera/9.23 (Windows NT 5.1; U; en)",
		432: "Opera/9.23 (Windows ME; U; en)",
		433: "Opera/9.26 (Windows NT 5.1; U; en)",
		434: "Opera/9.51 (Windows NT 5.1; U; en)",
		435: "Opera/9.52 (Macintosh; Intel Mac OS X; U; en)",
		436: "Opera/9.52 (Windows NT 5.1; U; en)",
		437: "Opera/9.60 (Macintosh; Intel Mac OS X; U; en) Presto/2.1.1",
		438: "Opera/9.60 (Windows NT 5.1; U; en) Presto/2.1.1",
		439: "Opera/9.61 (Windows NT 5.1; U; en) Presto/2.1.1",
		440: "Opera/9.62 (Windows NT 5.1; U; en) Presto/2.1.1",
		441: "Mozilla/4.0 (compatible; MSIE 6.0; X11; Linux i686; en) Opera 10.10",
		442: "Opera/9.80 (Windows NT 6.1; U; en) Presto/2.9.168 Version/11.50",
		443: "Opera/9.80 (Windows NT 6.1; U; en) Presto/2.10.229 Version/11.60",
		444: "Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.10.289 Version/12.00",
	}

	// this *.txt maybe not found if you exec binary, so we just fill several ua
	temp, err := util.ReadFromFile(util.CurDir() + "/config/ua.txt")

	if err == nil {
		uas := strings.Split(string(temp), "\n")
		for i, ua := range uas {
			Ua[i] = strings.TrimSpace(strings.Replace(ua, "\r", "", -1))
		}
	}

}

// RandomUa back random User-Agent
func RandomUa() string {
	length := len(Ua)
	if length == 0 {
		return ""
	}
	return Ua[rand.Intn(length-1)]
}
