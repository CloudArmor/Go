// pronicnumber_test.go
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see pronicnumber.go

package math_test

import (
	"testing"

	"github.com/CloudArmor/Go/math"
)

func TestPronicNumber(t *testing.T) {
	var tests = []struct {
		name          string
		n             int
		expectedValue bool
	}{
		{"-12 is not pronic", -12, false},
		{"0 is pronic", 0, true},
		{"1 is not pronic", 1, false},
		{"2 is pronic", 2, true},
		{"3 is not pronic", 3, false},
		{"4 is not pronic", 4, false},
		{"5 is not pronic", 5, false},
		{"6 is pronic", 6, true},
		{"7 is not pronic", 7, false},
		{"8 is not pronic", 8, false},
		{"9 is not pronic", 9, false},
		{"10 is not pronic", 10, false},
		{"11 is not pronic", 11, false},
		{"12 is pronic", 12, true},
		{"13 is not pronic", 13, false},
		{"14 is not pronic", 14, false},
		{"15 is not pronic", 15, false},
		{"16 is not pronic", 16, false},
		{"17 is not pronic", 17, false},
		{"18 is not pronic", 18, false},
		{"19 is not pronic", 19, false},
		{"20 is pronic", 20, true},
		{"21 is not pronic", 21, false},
		{"22 is not pronic", 22, false},
		{"23 is not pronic", 23, false},
		{"24 is not pronic", 24, false},
		{"25 is not pronic", 25, false},
		{"26 is not pronic", 26, false},
		{"27 is not pronic", 27, false},
		{"28 is not pronic", 28, false},
		{"29 is not pronic", 29, false},
		{"30 is pronic", 30, true},
		{"31 is not pronic", 31, false},
		{"32 is not pronic", 32, false},
		{"33 is not pronic", 33, false},
		{"34 is not pronic", 34, false},
		{"35 is not pronic", 35, false},
		{"36 is not pronic", 36, false},
		{"37 is not pronic", 37, false},
		{"38 is not pronic", 38, false},
		{"39 is not pronic", 39, false},
		{"40 is not pronic", 40, false},
		{"41 is not pronic", 41, false},
		{"42 is pronic", 42, true},
		{"43 is not pronic", 43, false},
		{"44 is not pronic", 44, false},
		{"45 is not pronic", 45, false},
		{"46 is not pronic", 46, false},
		{"47 is not pronic", 47, false},
		{"48 is not pronic", 48, false},
		{"49 is not pronic", 49, false},
		{"50 is not pronic", 50, false},
		{"51 is not pronic", 51, false},
		{"52 is not pronic", 52, false},
		{"53 is not pronic", 53, false},
		{"54 is not pronic", 54, false},
		{"55 is not pronic", 55, false},
		{"56 is pronic", 56, true},
		{"57 is not pronic", 57, false},
		{"58 is not pronic", 58, false},
		{"59 is not pronic", 59, false},
		{"60 is not pronic", 60, false},
		{"61 is not pronic", 61, false},
		{"62 is not pronic", 62, false},
		{"63 is not pronic", 63, false},
		{"64 is not pronic", 64, false},
		{"65 is not pronic", 65, false},
		{"66 is not pronic", 66, false},
		{"67 is not pronic", 67, false},
		{"68 is not pronic", 68, false},
		{"69 is not pronic", 69, false},
		{"70 is not pronic", 70, false},
		{"71 is not pronic", 71, false},
		{"72 is pronic", 72, true},
		{"73 is not pronic", 73, false},
		{"74 is not pronic", 74, false},
		{"75 is not pronic", 75, false},
		{"76 is not pronic", 76, false},
		{"77 is not pronic", 77, false},
		{"78 is not pronic", 78, false},
		{"79 is not pronic", 79, false},
		{"80 is not pronic", 80, false},
		{"81 is not pronic", 81, false},
		{"82 is not pronic", 82, false},
		{"83 is not pronic", 83, false},
		{"84 is not pronic", 84, false},
		{"85 is not pronic", 85, false},
		{"86 is not pronic", 86, false},
		{"87 is not pronic", 87, false},
		{"88 is not pronic", 88, false},
		{"89 is not pronic", 89, false},
		{"90 is pronic", 90, true},
		{"91 is not pronic", 91, false},
		{"92 is not pronic", 92, false},
		{"93 is not pronic", 93, false},
		{"94 is not pronic", 94, false},
		{"95 is not pronic", 95, false},
		{"96 is not pronic", 96, false},
		{"97 is not pronic", 97, false},
		{"98 is not pronic", 98, false},
		{"99 is not pronic", 99, false},
		{"100 is not pronic", 100, false},
		{"101 is not pronic", 101, false},
		{"102 is not pronic", 102, false},
		{"103 is not pronic", 103, false},
		{"104 is not pronic", 104, false},
		{"105 is not pronic", 105, false},
		{"106 is not pronic", 106, false},
		{"107 is not pronic", 107, false},
		{"108 is not pronic", 108, false},
		{"109 is not pronic", 109, false},
		{"110 is pronic", 110, true},
		{"111 is not pronic", 111, false},
		{"112 is not pronic", 112, false},
		{"113 is not pronic", 113, false},
		{"114 is not pronic", 114, false},
		{"115 is not pronic", 115, false},
		{"116 is not pronic", 116, false},
		{"117 is not pronic", 117, false},
		{"118 is not pronic", 118, false},
		{"119 is not pronic", 119, false},
		{"120 is not pronic", 120, false},
		{"121 is not pronic", 121, false},
		{"122 is not pronic", 122, false},
		{"123 is not pronic", 123, false},
		{"124 is not pronic", 124, false},
		{"125 is not pronic", 125, false},
		{"126 is not pronic", 126, false},
		{"127 is not pronic", 127, false},
		{"128 is not pronic", 128, false},
		{"129 is not pronic", 129, false},
		{"130 is not pronic", 130, false},
		{"131 is not pronic", 131, false},
		{"132 is pronic", 132, true},
		{"133 is not pronic", 133, false},
		{"134 is not pronic", 134, false},
		{"135 is not pronic", 135, false},
		{"136 is not pronic", 136, false},
		{"137 is not pronic", 137, false},
		{"138 is not pronic", 138, false},
		{"139 is not pronic", 139, false},
		{"140 is not pronic", 140, false},
		{"141 is not pronic", 141, false},
		{"142 is not pronic", 142, false},
		{"143 is not pronic", 143, false},
		{"144 is not pronic", 144, false},
		{"145 is not pronic", 145, false},
		{"146 is not pronic", 146, false},
		{"147 is not pronic", 147, false},
		{"148 is not pronic", 148, false},
		{"149 is not pronic", 149, false},
		{"150 is not pronic", 150, false},
		{"151 is not pronic", 151, false},
		{"152 is not pronic", 152, false},
		{"153 is not pronic", 153, false},
		{"154 is not pronic", 154, false},
		{"155 is not pronic", 155, false},
		{"156 is pronic", 156, true},
		{"157 is not pronic", 157, false},
		{"158 is not pronic", 158, false},
		{"159 is not pronic", 159, false},
		{"160 is not pronic", 160, false},
		{"161 is not pronic", 161, false},
		{"162 is not pronic", 162, false},
		{"163 is not pronic", 163, false},
		{"164 is not pronic", 164, false},
		{"165 is not pronic", 165, false},
		{"166 is not pronic", 166, false},
		{"167 is not pronic", 167, false},
		{"168 is not pronic", 168, false},
		{"169 is not pronic", 169, false},
		{"170 is not pronic", 170, false},
		{"171 is not pronic", 171, false},
		{"172 is not pronic", 172, false},
		{"173 is not pronic", 173, false},
		{"174 is not pronic", 174, false},
		{"175 is not pronic", 175, false},
		{"176 is not pronic", 176, false},
		{"177 is not pronic", 177, false},
		{"178 is not pronic", 178, false},
		{"179 is not pronic", 179, false},
		{"180 is not pronic", 180, false},
		{"181 is not pronic", 181, false},
		{"182 is pronic", 182, true},
		{"183 is not pronic", 183, false},
		{"184 is not pronic", 184, false},
		{"185 is not pronic", 185, false},
		{"186 is not pronic", 186, false},
		{"187 is not pronic", 187, false},
		{"188 is not pronic", 188, false},
		{"189 is not pronic", 189, false},
		{"190 is not pronic", 190, false},
		{"191 is not pronic", 191, false},
		{"192 is not pronic", 192, false},
		{"193 is not pronic", 193, false},
		{"194 is not pronic", 194, false},
		{"195 is not pronic", 195, false},
		{"196 is not pronic", 196, false},
		{"197 is not pronic", 197, false},
		{"198 is not pronic", 198, false},
		{"199 is not pronic", 199, false},
		{"200 is not pronic", 200, false},
		{"201 is not pronic", 201, false},
		{"202 is not pronic", 202, false},
		{"203 is not pronic", 203, false},
		{"204 is not pronic", 204, false},
		{"205 is not pronic", 205, false},
		{"206 is not pronic", 206, false},
		{"207 is not pronic", 207, false},
		{"208 is not pronic", 208, false},
		{"209 is not pronic", 209, false},
		{"210 is pronic", 210, true},
		{"211 is not pronic", 211, false},
		{"212 is not pronic", 212, false},
		{"213 is not pronic", 213, false},
		{"214 is not pronic", 214, false},
		{"215 is not pronic", 215, false},
		{"216 is not pronic", 216, false},
		{"217 is not pronic", 217, false},
		{"218 is not pronic", 218, false},
		{"219 is not pronic", 219, false},
		{"220 is not pronic", 220, false},
		{"221 is not pronic", 221, false},
		{"222 is not pronic", 222, false},
		{"223 is not pronic", 223, false},
		{"224 is not pronic", 224, false},
		{"225 is not pronic", 225, false},
		{"226 is not pronic", 226, false},
		{"227 is not pronic", 227, false},
		{"228 is not pronic", 228, false},
		{"229 is not pronic", 229, false},
		{"230 is not pronic", 230, false},
		{"231 is not pronic", 231, false},
		{"232 is not pronic", 232, false},
		{"233 is not pronic", 233, false},
		{"234 is not pronic", 234, false},
		{"235 is not pronic", 235, false},
		{"236 is not pronic", 236, false},
		{"237 is not pronic", 237, false},
		{"238 is not pronic", 238, false},
		{"239 is not pronic", 239, false},
		{"240 is pronic", 240, true},
		{"241 is not pronic", 241, false},
		{"242 is not pronic", 242, false},
		{"243 is not pronic", 243, false},
		{"244 is not pronic", 244, false},
		{"245 is not pronic", 245, false},
		{"246 is not pronic", 246, false},
		{"247 is not pronic", 247, false},
		{"248 is not pronic", 248, false},
		{"249 is not pronic", 249, false},
		{"250 is not pronic", 250, false},
		{"251 is not pronic", 251, false},
		{"252 is not pronic", 252, false},
		{"253 is not pronic", 253, false},
		{"254 is not pronic", 254, false},
		{"255 is not pronic", 255, false},
		{"256 is not pronic", 256, false},
		{"257 is not pronic", 257, false},
		{"258 is not pronic", 258, false},
		{"259 is not pronic", 259, false},
		{"260 is not pronic", 260, false},
		{"261 is not pronic", 261, false},
		{"262 is not pronic", 262, false},
		{"263 is not pronic", 263, false},
		{"264 is not pronic", 264, false},
		{"265 is not pronic", 265, false},
		{"266 is not pronic", 266, false},
		{"267 is not pronic", 267, false},
		{"268 is not pronic", 268, false},
		{"269 is not pronic", 269, false},
		{"270 is not pronic", 270, false},
		{"271 is not pronic", 271, false},
		{"272 is pronic", 272, true},
		{"273 is not pronic", 273, false},
		{"274 is not pronic", 274, false},
		{"275 is not pronic", 275, false},
		{"276 is not pronic", 276, false},
		{"277 is not pronic", 277, false},
		{"278 is not pronic", 278, false},
		{"279 is not pronic", 279, false},
		{"280 is not pronic", 280, false},
		{"281 is not pronic", 281, false},
		{"282 is not pronic", 282, false},
		{"283 is not pronic", 283, false},
		{"284 is not pronic", 284, false},
		{"285 is not pronic", 285, false},
		{"286 is not pronic", 286, false},
		{"287 is not pronic", 287, false},
		{"288 is not pronic", 288, false},
		{"289 is not pronic", 289, false},
		{"290 is not pronic", 290, false},
		{"291 is not pronic", 291, false},
		{"292 is not pronic", 292, false},
		{"293 is not pronic", 293, false},
		{"294 is not pronic", 294, false},
		{"295 is not pronic", 295, false},
		{"296 is not pronic", 296, false},
		{"297 is not pronic", 297, false},
		{"298 is not pronic", 298, false},
		{"299 is not pronic", 299, false},
		{"300 is not pronic", 300, false},
		{"301 is not pronic", 301, false},
		{"302 is not pronic", 302, false},
		{"303 is not pronic", 303, false},
		{"304 is not pronic", 304, false},
		{"305 is not pronic", 305, false},
		{"306 is pronic", 306, true},
		{"307 is not pronic", 307, false},
		{"308 is not pronic", 308, false},
		{"309 is not pronic", 309, false},
		{"310 is not pronic", 310, false},
		{"311 is not pronic", 311, false},
		{"312 is not pronic", 312, false},
		{"313 is not pronic", 313, false},
		{"314 is not pronic", 314, false},
		{"315 is not pronic", 315, false},
		{"316 is not pronic", 316, false},
		{"317 is not pronic", 317, false},
		{"318 is not pronic", 318, false},
		{"319 is not pronic", 319, false},
		{"320 is not pronic", 320, false},
		{"321 is not pronic", 321, false},
		{"322 is not pronic", 322, false},
		{"323 is not pronic", 323, false},
		{"324 is not pronic", 324, false},
		{"325 is not pronic", 325, false},
		{"326 is not pronic", 326, false},
		{"327 is not pronic", 327, false},
		{"328 is not pronic", 328, false},
		{"329 is not pronic", 329, false},
		{"330 is not pronic", 330, false},
		{"331 is not pronic", 331, false},
		{"332 is not pronic", 332, false},
		{"333 is not pronic", 333, false},
		{"334 is not pronic", 334, false},
		{"335 is not pronic", 335, false},
		{"336 is not pronic", 336, false},
		{"337 is not pronic", 337, false},
		{"338 is not pronic", 338, false},
		{"339 is not pronic", 339, false},
		{"340 is not pronic", 340, false},
		{"341 is not pronic", 341, false},
		{"342 is pronic", 342, true},
		{"343 is not pronic", 343, false},
		{"344 is not pronic", 344, false},
		{"345 is not pronic", 345, false},
		{"346 is not pronic", 346, false},
		{"347 is not pronic", 347, false},
		{"348 is not pronic", 348, false},
		{"349 is not pronic", 349, false},
		{"350 is not pronic", 350, false},
		{"351 is not pronic", 351, false},
		{"352 is not pronic", 352, false},
		{"353 is not pronic", 353, false},
		{"354 is not pronic", 354, false},
		{"355 is not pronic", 355, false},
		{"356 is not pronic", 356, false},
		{"357 is not pronic", 357, false},
		{"358 is not pronic", 358, false},
		{"359 is not pronic", 359, false},
		{"360 is not pronic", 360, false},
		{"361 is not pronic", 361, false},
		{"362 is not pronic", 362, false},
		{"363 is not pronic", 363, false},
		{"364 is not pronic", 364, false},
		{"365 is not pronic", 365, false},
		{"366 is not pronic", 366, false},
		{"367 is not pronic", 367, false},
		{"368 is not pronic", 368, false},
		{"369 is not pronic", 369, false},
		{"370 is not pronic", 370, false},
		{"371 is not pronic", 371, false},
		{"372 is not pronic", 372, false},
		{"373 is not pronic", 373, false},
		{"374 is not pronic", 374, false},
		{"375 is not pronic", 375, false},
		{"376 is not pronic", 376, false},
		{"377 is not pronic", 377, false},
		{"378 is not pronic", 378, false},
		{"379 is not pronic", 379, false},
		{"380 is pronic", 380, true},
		{"381 is not pronic", 381, false},
		{"382 is not pronic", 382, false},
		{"383 is not pronic", 383, false},
		{"384 is not pronic", 384, false},
		{"385 is not pronic", 385, false},
		{"386 is not pronic", 386, false},
		{"387 is not pronic", 387, false},
		{"388 is not pronic", 388, false},
		{"389 is not pronic", 389, false},
		{"390 is not pronic", 390, false},
		{"391 is not pronic", 391, false},
		{"392 is not pronic", 392, false},
		{"393 is not pronic", 393, false},
		{"394 is not pronic", 394, false},
		{"395 is not pronic", 395, false},
		{"396 is not pronic", 396, false},
		{"397 is not pronic", 397, false},
		{"398 is not pronic", 398, false},
		{"399 is not pronic", 399, false},
		{"400 is not pronic", 400, false},
		{"2147441940 is pronic", 2147441940, true},
		{"9223372033963249500 is pronic", 9223372033963249500, true},
		{"9223372033963249664 is not pronic", 9223372033963249664, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := math.PronicNumber(test.n)
			if result != test.expectedValue {
				t.Errorf("expected value: %v, got: %v", test.expectedValue, result)
			}
		})
	}
}
func BenchmarkPronicNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.PronicNumber(65536)
	}
}
