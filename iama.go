package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/vaxx99/iama/bcd"
)

type block []string

var wd, sp string

func main() {

	switch len(os.Args) {
	case 2:
		wd = os.Args[1]
	case 3:
		wd = os.Args[1]
		sp = os.Args[2]
	default:
		wd = "."
		sp = ""
	}
	os.Chdir(wd)
	f, _ := ioutil.ReadDir(".")
	for _, fn := range f {
		if isTrue(fn.Name()) {
			u := time.Now()
			rec, ft, mT, rc := iama(fn.Name(), sp)
			f, err := os.OpenFile(mT+"."+strings.ToLower(ft), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
			defer f.Close()
			if err != nil {
				panic(err)
			}
			for _, i := range rec {
				_, err = f.WriteString(i + "\n")
			}
			d := float64(time.Now().Sub(u).Seconds())
			fmt.Printf("%s%s%.3f%s%d%s", fn.Name(), "\t-->\t", d, "\t-->\t", rc, "\n")
		}
	}
}

func iama(fn, fs string) ([]string, string, string, int) {
	var rec []string
	_, ft, mT, rc, j, file, e := bcd.Finfo(fn)
	yy := mT[0:4]
	bcd.Err(e)
	defer file.Close()

	i := 0

	if ft == "AMA" {

		for {
			data, e := bcd.Read(file, j)
			j = bcd.Next(data[len(data)-3])
			bcd.Err(e)
			ad := bcd.H2bcd(data)

			if ad[0:6] == "AA9020" {
				b := A9020(ad, yy, fs)
				rec = append(rec, b)
			}
			if ad[0:6] == "AA9021" {
				b := A9021(ad, yy, fs)
				rec = append(rec, b)
			}
			if ad[0:6] == "AA9025" {
				b := A9025(ad, yy, fs)
				rec = append(rec, b)
			}
			if ad[0:6] == "AA9026" {
				b := A9026(ad, yy, fs)
				rec = append(rec, b)
			}
			if j == 0 {
				//exit on "EOF"
				return rec, ft, mT, rc
			}
			i++
		}
	}

	if ft == "IAD" {

		i := 0
		for i < rc {
			data, e := bcd.Read(file, 42)
			bcd.Err(e)
			ad := bcd.H2bcd(data)
			if ad[0:6] == "AA0003" {
				b := A0003(ad, yy, fs)
				rec = append(rec, b)
			} else {
				_, e = bcd.Read(file, 32)
				bcd.Err(e)
			}
			i++
		}
	}
	return rec, ft, mT, rc
}

//A9020 - Structure
func A9020(bcd, yy, sp string) string {
	aa := bcd[0:2] + sp
	aa += bcd[2:6] + sp
	aa += bcd[6:12] + sp
	aa += bcd[12:18] + sp
	aa += bcd[18:20] + sp
	i, _ := strconv.Atoi(bcd[18:20])
	aa += bcd[36-i:36] + sp
	aa += bcd[36:38] + sp
	i, _ = strconv.Atoi(bcd[36:38])
	aa += bcd[70-i:70] + sp
	aa += bcd[70:74] + sp
	aa += bcd[74:78] + sp
	aa += bcd[78:80] + sp
	aa += bcd[80:82] + sp
	aa += bcd[82:84] + sp
	aa += bcd[84:86] + sp
	aa += bcd[86:88] + sp
	aa += yy + bcd[89:99] + sp
	aa += bcd[99:100] + sp
	aa += yy + bcd[101:111] + sp
	aa += bcd[111:112] + sp
	aa += bcd[112:114] + sp
	aa += bcd[114:118] + sp
	aa += bcd[118:122] + sp
	aa += bcd[122:126] + sp
	aa += bcd[126:132] + sp
	aa += bcd[132:138] + sp
	aa += bcd[138:140] + sp
	aa += bcd[140:148] + sp
	aa += bcd[148:150] + sp
	aa += bcd[150:152] + sp
	aa += bcd[152:154] + sp
	aa += bcd[154:156] + sp
	aa += bcd[156:158] + sp
	aa += bcd[158:160] + sp
	return aa
}

//A9021 - Structure
func A9021(bcd, yy, sp string) string {
	aa := bcd[0:2] + sp   //Hexadecimal Identifier
	aa += bcd[2:6] + sp   //Structure Identifier Code
	aa += bcd[6:12] + sp  //Ticket Number
	aa += bcd[12:18] + sp //Sequence Number
	aa += bcd[18:20] + sp
	i, _ := strconv.Atoi(bcd[18:20])
	aa += bcd[36-i:36] + sp //Originating Phone Number
	aa += bcd[36:38] + sp
	i, _ = strconv.Atoi(bcd[36:38])
	aa += bcd[70-i:70] + sp    //Terminating Phone Number
	aa += bcd[70:72] + sp      //Charge Category
	aa += bcd[72:74] + sp      //Nature of Call
	aa += bcd[74:76] + sp      //CDA Indicator
	aa += bcd[76:78] + sp      //LDC Indicator
	aa += bcd[78:80] + sp      //Service Class of Call
	aa += yy + bcd[81:91] + sp //Date and Time of Charging Commencement
	aa += bcd[91:92] + sp      //*ms
	aa += bcd[92:104] + sp     //Date and Time of Call End
	aa += bcd[104:106] + sp    //Cause of Call End
	aa += bcd[106:110] + sp    //Destination
	aa += bcd[110:114] + sp    //Outgoing Trunk Group
	aa += bcd[114:118] + sp    //Incoming Trunk Group
	aa += bcd[118:124] + sp    //Conversation Time
	aa += bcd[124:130] + sp    //Chargeable Duration
	aa += bcd[130:132] + sp    //Class of Rate
	aa += bcd[132:140] + sp    //Fee
	aa += bcd[140:142] + sp    //Trouble Mark
	aa += bcd[142:144] + sp    //Day
	aa += bcd[144:146] + sp    //A-Party Category
	aa += bcd[146:148] + sp    // Type of Call
	aa += bcd[148:150] + sp    //Customer Feature
	aa += bcd[150:152] + sp    //Customer Feature Action
	return aa
}

//A9025 - Structure
func A9025(bcd, yy, sp string) string {
	aa := bcd[0:2] + sp   //Hexadecimal Identifier
	aa += bcd[2:6] + sp   //Structure Identifier Code
	aa += bcd[6:12] + sp  //Ticket Number
	aa += bcd[12:18] + sp //Sequence Number
	aa += bcd[18:20] + sp
	i, _ := strconv.Atoi(bcd[18:20])
	aa += bcd[36-i:36] + sp //Originating Phone Number
	aa += bcd[36:38] + sp
	i, _ = strconv.Atoi(bcd[36:38])
	aa += bcd[70-i:70] + sp    //Terminating Phone Number
	aa += bcd[70:72] + sp      //Charge Category
	aa += bcd[72:74] + sp      //Nature of Call
	aa += bcd[74:76] + sp      //CDA Indicator
	aa += bcd[76:78] + sp      //LDC Indicator
	aa += bcd[78:80] + sp      //Service Class of Call
	aa += yy + bcd[81:91] + sp //Date and Time of Charging Commencement
	aa += bcd[91:92] + sp
	aa += yy + bcd[93:103] + sp //Date and Time of Call End
	aa += bcd[103:104] + sp
	aa += bcd[104:106] + sp //Cause of Call End
	aa += bcd[106:110] + sp //Destination
	aa += bcd[110:114] + sp //Outgoing Trunk Group
	aa += bcd[114:118] + sp //Incoming Trunk Group
	aa += bcd[118:124] + sp //Conversation Time
	aa += bcd[124:130] + sp //Chargeable Duration
	aa += bcd[130:132] + sp //Class of Rate
	aa += bcd[132:140] + sp //Fee
	aa += bcd[140:142] + sp //Trouble Mark
	aa += bcd[142:144] + sp //Day
	aa += bcd[144:146] + sp //A-Party Category
	aa += bcd[146:148] + sp //Type of Call
	aa += bcd[148:150] + sp //Bearer Service
	aa += bcd[150:154] + sp //CUG Interlock Code
	aa += bcd[154:156] + sp //COG OA Indicator
	aa += bcd[156:160] + sp //UUI Messages
	aa += bcd[160:162] + sp //Terminating Access
	aa += bcd[162:164] + sp //Network Indicator
	aa += bcd[164:168] + sp //Release Cause
	aa += bcd[168:170] + sp //Supplementary Service Indicator
	return aa
}

//A9026 - Structure
func A9026(bcd, yy, sp string) string {
	aa := bcd[0:2] + sp   // Hexadecimal Identifier
	aa += bcd[2:6] + sp   // Structure Identifier Code
	aa += bcd[6:12] + sp  // Ticket Number
	aa += bcd[12:18] + sp // Sequence Number
	aa += bcd[18:20] + sp
	i, _ := strconv.Atoi(bcd[18:20])
	aa += bcd[36-i:36] + sp // Originating Phone Number
	aa += bcd[36:38] + sp
	i, _ = strconv.Atoi(bcd[36:38])
	aa += bcd[70-i:70] + sp     // Terminating Phone Number
	aa += bcd[70:72] + sp       // Charge Category
	aa += bcd[72:74] + sp       // Nature of Call
	aa += bcd[74:76] + sp       // CDA Indicator
	aa += bcd[76:78] + sp       // LDC Indicator
	aa += bcd[78:80] + sp       // Service Class of Call
	aa += yy + bcd[81:91] + sp  // Date and Time of Charging Commencement
	aa += bcd[91:92] + sp       //*ms
	aa += yy + bcd[93:103] + sp // Date and Time of Call End
	aa += bcd[103:104] + sp     //*ms
	aa += bcd[104:106] + sp     // Cause of Call End
	aa += bcd[106:110] + sp     // Destination
	aa += bcd[110:114] + sp     // Outgoing Trunk Group
	aa += bcd[114:118] + sp     // Incoming Trunk Group
	aa += bcd[118:124] + sp     // Conversation Time
	aa += bcd[124:130] + sp     // Chargeable Duration
	aa += bcd[130:132] + sp     // Class of Rate
	aa += bcd[132:140] + sp     // Fee
	aa += bcd[140:142] + sp     // Trouble Mark
	aa += bcd[142:144] + sp     // Day
	aa += bcd[144:146] + sp     // A-Party Category
	aa += bcd[146:148] + sp     // Type of Call
	aa += bcd[148:150] + sp     // Bearer Service
	aa += bcd[150:152] + sp     // Supplementary Service Indicator
	aa += bcd[152:154] + sp     // Supplementary Service Action
	return aa
}

//A9050 - Structure
func A9050(bcd, sp string) string {
	aa := bcd[8:10] + sp
	aa += bcd[0:2] + sp
	aa += bcd[2:6] + sp
	aa += bcd[6:10] + sp
	aa += bcd[10:14] + sp
	aa += bcd[14:22] + sp
	aa += bcd[22:28] + sp
	aa += bcd[28:36] + sp
	aa += bcd[36:42] + sp
	aa += bcd[42:46] + sp
	return aa
}

//A9051 - Structure
func A9051(bcd, sp string) string {
	aa := bcd[0:2] + sp
	aa += bcd[2:6] + sp
	aa += bcd[6:10] + sp
	aa += bcd[10:14] + sp
	aa += bcd[14:22] + sp
	aa += bcd[22:28] + sp
	aa += bcd[28:36] + sp
	aa += bcd[36:42] + sp
	aa += bcd[42:46] + sp
	aa += bcd[46:54] + sp
	return aa
}

//A0003 - Structure
func A0003(bcd, yy, sp string) string {
	aa := bcd[0:2] + sp
	aa += bcd[2:6] + sp
	aa += bcd[6:10] + sp
	aa += bcd[10:14] + sp
	i, _ := strconv.Atoi(bcd[15:25]) //DS
	if i == 0 {
		aa += sp
	} else {
		aa += yy + bcd[15:25] + sp
	}
	aa += bcd[25:26] + sp
	aa += bcd[26:28] + sp
	aa += bcd[28:32] + sp
	aa += bcd[32:34] + sp
	i, _ = strconv.Atoi(bcd[32:34]) //DN
	aa += bcd[66-i:66] + sp         // Terminating Phone Number
	aa += bcd[66:70] + sp
	aa += yy + bcd[71:81] + sp //DE
	aa += bcd[81:82] + sp
	aa += bcd[82:84] + sp
	return aa
}

func isTrue(fn string) bool {
	if info, err := os.Stat(fn); err == nil && info.IsDir() {
		return false
	}
	f, _ := os.Open(fn)
	defer f.Close()
	data, _ := bcd.Read(f, 31)
	ad := bcd.H2bcd(data)
	if ad[8:14] == "AA9050" || ad[0:2] == "01" {
		return true
	}
	return false
}
