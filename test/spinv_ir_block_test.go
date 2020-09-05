// Lute - 一款对中文语境优化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package test

import (
	"strings"
	"testing"

	"github.com/88250/lute"
)

var spinVditorIRBlockDOMTests = []*parseTest{

	{"31", "<blockquote data-block=\"0\" data-node-id=\"20200905220947-qep8u59\"><p data-block=\"0\" data-node-id=\"\"><br></p><p data-block=\"0\" data-node-id=\"\"><wbr>foo</p></blockquote>", "<blockquote data-block=\"0\" data-node-id=\"20200905220947-qep8u59\"><p data-block=\"0\" data-node-id=\"\"><wbr>foo</p></blockquote>"},
	{"30", "<p data-block=\"0\" data-node-id=\"20200905114748-mh8mcoj\"><span data-type=\"block-ref-embed\" class=\"vditor-ir__node\"><span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200904104256-63oy2pv</span> <span class=\"vditor-ir__blockref\">\"<wbr>\"</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span><span data-block-def-id=\"20200904104256-63oy2pv\" data-render=\"2\" data-type=\"block-render\"></span></span></p>", "<p data-block=\"0\" data-node-id=\"20200905114748-mh8mcoj\"><span data-type=\"block-ref\" class=\"vditor-ir__node vditor-ir__node--expand\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200904104256-63oy2pv</span> <span class=\"vditor-ir__blockref\">&#34;<wbr>&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span></p>"},
	{"29", "<p data-block=\"0\" data-node-id=\"20200815151746-idfp3ma\">foo((20200815093609-63sg22j \"*<wbr>\"))bar</p>", "<p data-block=\"0\" data-node-id=\"20200815151746-idfp3ma\">foo<span data-type=\"block-ref-embed\" class=\"vditor-ir__node vditor-ir__node--expand\"><span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200815093609-63sg22j</span> <span class=\"vditor-ir__blockref\">&#34;*<wbr>&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span><span data-block-def-id=\"20200815093609-63sg22j\" data-render=\"2\" data-type=\"block-render\"></span></span>bar</p>"},
	{"28", "<ul data-marker=\"-\" data-block=\"0\" data-node-id=\"20200831104952-togxq1j\"><li data-marker=\"-\" data-node-id=\"\"><div data-block=\"0\" data-node-id=\"\" data-type=\"math-block\" class=\"vditor-ir__node\"><span data-type=\"math-block-open-marker\">$$</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code data-type=\"math-block\" class=\"language-math\">foo<wbr></code></pre><pre class=\"vditor-ir__preview\" data-render=\"1\"><code class=\"language-math\"><div class=\"vditor-math\" data-math=\"z\\ll1\\Leftrightarrow \\frac{\\lambda^3N}{V}\\ll1\"><mjx-container class=\"MathJax CtxtMenu_Attached_0\" jax=\"SVG\" display=\"true\" tabindex=\"0\" ctxtmenu_counter=\"1331\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"19.109ex\" height=\"5.016ex\" role=\"img\" focusable=\"false\" viewBox=\"0 -1509.2 8446.2 2217.2\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"><defs><path id=\"MJX-1332-TEX-I-7A\" d=\"M347 338Q337 338 294 349T231 360Q211 360 197 356T174 346T162 335T155 324L153 320Q150 317 138 317Q117 317 117 325Q117 330 120 339Q133 378 163 406T229 440Q241 442 246 442Q271 442 291 425T329 392T367 375Q389 375 411 408T434 441Q435 442 449 442H462Q468 436 468 434Q468 430 463 420T449 399T432 377T418 358L411 349Q368 298 275 214T160 106L148 94L163 93Q185 93 227 82T290 71Q328 71 360 90T402 140Q406 149 409 151T424 153Q443 153 443 143Q443 138 442 134Q425 72 376 31T278 -11Q252 -11 232 6T193 40T155 57Q111 57 76 -3Q70 -11 59 -11H54H41Q35 -5 35 -2Q35 13 93 84Q132 129 225 214T340 322Q352 338 347 338Z\"></path><path id=\"MJX-1332-TEX-N-226A\" d=\"M639 -48Q639 -54 634 -60T619 -67H618Q612 -67 536 -26Q430 33 329 88Q61 235 59 239Q56 243 56 250T59 261Q62 266 336 415T615 567L619 568Q622 567 625 567Q639 562 639 548Q639 540 633 534Q632 532 374 391L117 250L374 109Q632 -32 633 -34Q639 -40 639 -48ZM944 -48Q944 -54 939 -60T924 -67H923Q917 -67 841 -26Q735 33 634 88Q366 235 364 239Q361 243 361 250T364 261Q367 266 641 415T920 567L924 568Q927 567 930 567Q944 562 944 548Q944 540 938 534Q937 532 679 391L422 250L679 109Q937 -32 938 -34Q944 -40 944 -48Z\"></path><path id=\"MJX-1332-TEX-N-31\" d=\"M213 578L200 573Q186 568 160 563T102 556H83V602H102Q149 604 189 617T245 641T273 663Q275 666 285 666Q294 666 302 660V361L303 61Q310 54 315 52T339 48T401 46H427V0H416Q395 3 257 3Q121 3 100 0H88V46H114Q136 46 152 46T177 47T193 50T201 52T207 57T213 61V578Z\"></path><path id=\"MJX-1332-TEX-N-21D4\" d=\"M308 524Q318 526 323 526Q340 526 340 514Q340 507 336 499Q326 476 314 454T292 417T274 391T260 374L255 368Q255 367 500 367Q744 367 744 368L739 374Q734 379 726 390T707 416T685 453T663 499Q658 511 658 515Q658 525 680 525Q687 524 690 523T695 519T701 507Q766 359 902 287Q921 276 939 269T961 259T966 250Q966 246 965 244T960 240T949 236T930 228T902 213Q763 137 701 -7Q697 -16 695 -19T690 -23T680 -25Q658 -25 658 -15Q658 -11 663 1Q673 24 685 46T707 83T725 109T739 126L744 132Q744 133 500 133Q255 133 255 132L260 126Q265 121 273 110T292 84T314 47T336 1Q341 -11 341 -15Q341 -25 319 -25Q312 -24 309 -23T304 -19T298 -7Q233 141 97 213Q83 221 70 227T51 235T41 239T35 243T34 250T35 256T40 261T51 265T70 273T97 287Q235 363 299 509Q305 522 308 524ZM792 319L783 327H216Q183 294 120 256L110 250L120 244Q173 212 207 181L216 173H783L792 181Q826 212 879 244L889 250L879 256Q826 288 792 319Z\"></path><path id=\"MJX-1332-TEX-I-3BB\" d=\"M166 673Q166 685 183 694H202Q292 691 316 644Q322 629 373 486T474 207T524 67Q531 47 537 34T546 15T551 6T555 2T556 -2T550 -11H482Q457 3 450 18T399 152L354 277L340 262Q327 246 293 207T236 141Q211 112 174 69Q123 9 111 -1T83 -12Q47 -12 47 20Q47 37 61 52T199 187Q229 216 266 252T321 306L338 322Q338 323 288 462T234 612Q214 657 183 657Q166 657 166 673Z\"></path><path id=\"MJX-1332-TEX-N-33\" d=\"M127 463Q100 463 85 480T69 524Q69 579 117 622T233 665Q268 665 277 664Q351 652 390 611T430 522Q430 470 396 421T302 350L299 348Q299 347 308 345T337 336T375 315Q457 262 457 175Q457 96 395 37T238 -22Q158 -22 100 21T42 130Q42 158 60 175T105 193Q133 193 151 175T169 130Q169 119 166 110T159 94T148 82T136 74T126 70T118 67L114 66Q165 21 238 21Q293 21 321 74Q338 107 338 175V195Q338 290 274 322Q259 328 213 329L171 330L168 332Q166 335 166 348Q166 366 174 366Q202 366 232 371Q266 376 294 413T322 525V533Q322 590 287 612Q265 626 240 626Q208 626 181 615T143 592T132 580H135Q138 579 143 578T153 573T165 566T175 555T183 540T186 520Q186 498 172 481T127 463Z\"></path><path id=\"MJX-1332-TEX-I-4E\" d=\"M234 637Q231 637 226 637Q201 637 196 638T191 649Q191 676 202 682Q204 683 299 683Q376 683 387 683T401 677Q612 181 616 168L670 381Q723 592 723 606Q723 633 659 637Q635 637 635 648Q635 650 637 660Q641 676 643 679T653 683Q656 683 684 682T767 680Q817 680 843 681T873 682Q888 682 888 672Q888 650 880 642Q878 637 858 637Q787 633 769 597L620 7Q618 0 599 0Q585 0 582 2Q579 5 453 305L326 604L261 344Q196 88 196 79Q201 46 268 46H278Q284 41 284 38T282 19Q278 6 272 0H259Q228 2 151 2Q123 2 100 2T63 2T46 1Q31 1 31 10Q31 14 34 26T39 40Q41 46 62 46Q130 49 150 85Q154 91 221 362L289 634Q287 635 234 637Z\"></path><path id=\"MJX-1332-TEX-I-56\" d=\"M52 648Q52 670 65 683H76Q118 680 181 680Q299 680 320 683H330Q336 677 336 674T334 656Q329 641 325 637H304Q282 635 274 635Q245 630 242 620Q242 618 271 369T301 118L374 235Q447 352 520 471T595 594Q599 601 599 609Q599 633 555 637Q537 637 537 648Q537 649 539 661Q542 675 545 679T558 683Q560 683 570 683T604 682T668 681Q737 681 755 683H762Q769 676 769 672Q769 655 760 640Q757 637 743 637Q730 636 719 635T698 630T682 623T670 615T660 608T652 599T645 592L452 282Q272 -9 266 -16Q263 -18 259 -21L241 -22H234Q216 -22 216 -15Q213 -9 177 305Q139 623 138 626Q133 637 76 637H59Q52 642 52 648Z\"></path></defs><g stroke=\"currentColor\" fill=\"currentColor\" stroke-width=\"0\" transform=\"matrix(1 0 0 -1 0 0)\"><g data-mml-node=\"math\"><g data-mml-node=\"mi\"><use xlink:href=\"#MJX-1332-TEX-I-7A\"></use></g><g data-mml-node=\"mo\" transform=\"translate(742.8, 0)\"><use xlink:href=\"#MJX-1332-TEX-N-226A\"></use></g><g data-mml-node=\"mn\" transform=\"translate(2020.6, 0)\"><use xlink:href=\"#MJX-1332-TEX-N-31\"></use></g><g data-mml-node=\"mo\" transform=\"translate(2798.3, 0)\"><use xlink:href=\"#MJX-1332-TEX-N-21D4\"></use></g><g data-mml-node=\"mfrac\" transform=\"translate(4076.1, 0)\"><g data-mml-node=\"mrow\" transform=\"translate(220, 676)\"><g data-mml-node=\"msup\"><g data-mml-node=\"mi\"><use xlink:href=\"#MJX-1332-TEX-I-3BB\"></use></g><g data-mml-node=\"mn\" transform=\"translate(583, 363) scale(0.707)\"><use xlink:href=\"#MJX-1332-TEX-N-33\"></use></g></g><g data-mml-node=\"mi\" transform=\"translate(986.6, 0)\"><use xlink:href=\"#MJX-1332-TEX-I-4E\"></use></g></g><g data-mml-node=\"mi\" transform=\"translate(772.8, -686)\"><use xlink:href=\"#MJX-1332-TEX-I-56\"></use></g><rect width=\"2074.6\" height=\"60\" x=\"120\" y=\"220\"></rect></g><g data-mml-node=\"mo\" transform=\"translate(6668.4, 0)\"><use xlink:href=\"#MJX-1332-TEX-N-226A\"></use></g><g data-mml-node=\"mn\" transform=\"translate(7946.2, 0)\"><use xlink:href=\"#MJX-1332-TEX-N-31\"></use></g></g></g></svg></mjx-container></div></code></pre><span data-type=\"math-block-close-marker\">$$</span></div><p data-block=\"0\" data-node-id=\"\">​bar<wbr></p></li><li data-marker=\"-\" data-node-id=\"\"><div data-block=\"0\" data-node-id=\"\" data-type=\"math-block\" class=\"vditor-ir__node\"><span data-type=\"math-block-open-marker\">$$</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code data-type=\"math-block\" class=\"language-math\">foo<wbr></code></pre><pre class=\"vditor-ir__preview\" data-render=\"1\"><code class=\"language-math\"><div class=\"vditor-math\" data-math=\"z\\sim \\lambda^3 \\sim T^{-3/2}\"><mjx-container class=\"MathJax CtxtMenu_Attached_0\" jax=\"SVG\" display=\"true\" tabindex=\"0\" ctxtmenu_counter=\"1333\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"14.836ex\" height=\"2.161ex\" role=\"img\" focusable=\"false\" viewBox=\"0 -943.3 6557.5 955.3\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"><defs><path id=\"MJX-1334-TEX-I-7A\" d=\"M347 338Q337 338 294 349T231 360Q211 360 197 356T174 346T162 335T155 324L153 320Q150 317 138 317Q117 317 117 325Q117 330 120 339Q133 378 163 406T229 440Q241 442 246 442Q271 442 291 425T329 392T367 375Q389 375 411 408T434 441Q435 442 449 442H462Q468 436 468 434Q468 430 463 420T449 399T432 377T418 358L411 349Q368 298 275 214T160 106L148 94L163 93Q185 93 227 82T290 71Q328 71 360 90T402 140Q406 149 409 151T424 153Q443 153 443 143Q443 138 442 134Q425 72 376 31T278 -11Q252 -11 232 6T193 40T155 57Q111 57 76 -3Q70 -11 59 -11H54H41Q35 -5 35 -2Q35 13 93 84Q132 129 225 214T340 322Q352 338 347 338Z\"></path><path id=\"MJX-1334-TEX-N-223C\" d=\"M55 166Q55 241 101 304T222 367Q260 367 296 349T362 304T421 252T484 208T554 189Q616 189 655 236T694 338Q694 350 698 358T708 367Q722 367 722 334Q722 260 677 197T562 134H554Q517 134 481 152T414 196T355 248T292 293T223 311Q179 311 145 286Q109 257 96 218T80 156T69 133Q55 133 55 166Z\"></path><path id=\"MJX-1334-TEX-I-3BB\" d=\"M166 673Q166 685 183 694H202Q292 691 316 644Q322 629 373 486T474 207T524 67Q531 47 537 34T546 15T551 6T555 2T556 -2T550 -11H482Q457 3 450 18T399 152L354 277L340 262Q327 246 293 207T236 141Q211 112 174 69Q123 9 111 -1T83 -12Q47 -12 47 20Q47 37 61 52T199 187Q229 216 266 252T321 306L338 322Q338 323 288 462T234 612Q214 657 183 657Q166 657 166 673Z\"></path><path id=\"MJX-1334-TEX-N-33\" d=\"M127 463Q100 463 85 480T69 524Q69 579 117 622T233 665Q268 665 277 664Q351 652 390 611T430 522Q430 470 396 421T302 350L299 348Q299 347 308 345T337 336T375 315Q457 262 457 175Q457 96 395 37T238 -22Q158 -22 100 21T42 130Q42 158 60 175T105 193Q133 193 151 175T169 130Q169 119 166 110T159 94T148 82T136 74T126 70T118 67L114 66Q165 21 238 21Q293 21 321 74Q338 107 338 175V195Q338 290 274 322Q259 328 213 329L171 330L168 332Q166 335 166 348Q166 366 174 366Q202 366 232 371Q266 376 294 413T322 525V533Q322 590 287 612Q265 626 240 626Q208 626 181 615T143 592T132 580H135Q138 579 143 578T153 573T165 566T175 555T183 540T186 520Q186 498 172 481T127 463Z\"></path><path id=\"MJX-1334-TEX-I-54\" d=\"M40 437Q21 437 21 445Q21 450 37 501T71 602L88 651Q93 669 101 677H569H659Q691 677 697 676T704 667Q704 661 687 553T668 444Q668 437 649 437Q640 437 637 437T631 442L629 445Q629 451 635 490T641 551Q641 586 628 604T573 629Q568 630 515 631Q469 631 457 630T439 622Q438 621 368 343T298 60Q298 48 386 46Q418 46 427 45T436 36Q436 31 433 22Q429 4 424 1L422 0Q419 0 415 0Q410 0 363 1T228 2Q99 2 64 0H49Q43 6 43 9T45 27Q49 40 55 46H83H94Q174 46 189 55Q190 56 191 56Q196 59 201 76T241 233Q258 301 269 344Q339 619 339 625Q339 630 310 630H279Q212 630 191 624Q146 614 121 583T67 467Q60 445 57 441T43 437H40Z\"></path><path id=\"MJX-1334-TEX-N-2212\" d=\"M84 237T84 250T98 270H679Q694 262 694 250T679 230H98Q84 237 84 250Z\"></path><path id=\"MJX-1334-TEX-N-2F\" d=\"M423 750Q432 750 438 744T444 730Q444 725 271 248T92 -240Q85 -250 75 -250Q68 -250 62 -245T56 -231Q56 -221 230 257T407 740Q411 750 423 750Z\"></path><path id=\"MJX-1334-TEX-N-32\" d=\"M109 429Q82 429 66 447T50 491Q50 562 103 614T235 666Q326 666 387 610T449 465Q449 422 429 383T381 315T301 241Q265 210 201 149L142 93L218 92Q375 92 385 97Q392 99 409 186V189H449V186Q448 183 436 95T421 3V0H50V19V31Q50 38 56 46T86 81Q115 113 136 137Q145 147 170 174T204 211T233 244T261 278T284 308T305 340T320 369T333 401T340 431T343 464Q343 527 309 573T212 619Q179 619 154 602T119 569T109 550Q109 549 114 549Q132 549 151 535T170 489Q170 464 154 447T109 429Z\"></path></defs><g stroke=\"currentColor\" fill=\"currentColor\" stroke-width=\"0\" transform=\"matrix(1 0 0 -1 0 0)\"><g data-mml-node=\"math\"><g data-mml-node=\"mi\"><use xlink:href=\"#MJX-1334-TEX-I-7A\"></use></g><g data-mml-node=\"mo\" transform=\"translate(742.8, 0)\"><use xlink:href=\"#MJX-1334-TEX-N-223C\"></use></g><g data-mml-node=\"msup\" transform=\"translate(1798.6, 0)\"><g data-mml-node=\"mi\"><use xlink:href=\"#MJX-1334-TEX-I-3BB\"></use></g><g data-mml-node=\"mn\" transform=\"translate(583, 413) scale(0.707)\"><use xlink:href=\"#MJX-1334-TEX-N-33\"></use></g></g><g data-mml-node=\"mo\" transform=\"translate(3062.9, 0)\"><use xlink:href=\"#MJX-1334-TEX-N-223C\"></use></g><g data-mml-node=\"msup\" transform=\"translate(4118.7, 0)\"><g data-mml-node=\"mi\"><use xlink:href=\"#MJX-1334-TEX-I-54\"></use></g><g data-mml-node=\"TeXAtom\" transform=\"translate(778, 413) scale(0.707)\"><g data-mml-node=\"mo\"><use xlink:href=\"#MJX-1334-TEX-N-2212\"></use></g><g data-mml-node=\"mn\" transform=\"translate(778, 0)\"><use xlink:href=\"#MJX-1334-TEX-N-33\"></use></g><g data-mml-node=\"TeXAtom\" transform=\"translate(1278, 0)\"><g data-mml-node=\"mo\"><use xlink:href=\"#MJX-1334-TEX-N-2F\"></use></g></g><g data-mml-node=\"mn\" transform=\"translate(1778, 0)\"><use xlink:href=\"#MJX-1334-TEX-N-32\"></use></g></g></g></g></g></svg></mjx-container></div></code></pre><span data-type=\"math-block-close-marker\">$$</span></div><p data-block=\"0\"><wbr></p><p data-block=\"0\" data-node-id=\"\">​bar<wbr></p></li></ul>", "<ul data-marker=\"-\" data-block=\"0\" data-node-id=\"20200831104952-togxq1j\"><li data-marker=\"-\" data-node-id=\"\"><div data-block=\"0\" data-node-id=\"\" data-type=\"math-block\" class=\"vditor-ir__node vditor-ir__node--expand\"><span data-type=\"math-block-open-marker\">$$</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code data-type=\"math-block\" class=\"language-math\">foo<wbr></code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code data-type=\"math-block\" class=\"language-math\">foo</code></pre><span data-type=\"math-block-close-marker\">$$</span></div><p data-block=\"0\" data-node-id=\"\">bar<wbr></p></li><li data-marker=\"-\" data-node-id=\"\"><div data-block=\"0\" data-node-id=\"\" data-type=\"math-block\" class=\"vditor-ir__node vditor-ir__node--expand\"><span data-type=\"math-block-open-marker\">$$</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code data-type=\"math-block\" class=\"language-math\">foo<wbr></code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code data-type=\"math-block\" class=\"language-math\">foo</code></pre><span data-type=\"math-block-close-marker\">$$</span></div><p data-block=\"0\" data-node-id=\"\"><wbr></p><p data-block=\"0\" data-node-id=\"\">bar<wbr></p></li></ul>"},
	{"27", "<div data-block=\"0\" data-node-id=\"20200825235504-ect683e\" data-type=\"code-block\" class=\"vditor-ir__node\"><span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">​</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code>f<wbr></code></pre><pre class=\"vditor-ir__preview\" data-render=\"1\"><div class=\"vditor-copy\"><textarea></textarea><span aria-label=\"复制\" onmouseover=\"this.setAttribute('aria-label', '复制')\" class=\"vditor-tooltipped vditor-tooltipped__w\" onclick=\"this.previousElementSibling.select();document.execCommand('copy');this.setAttribute('aria-label', '已复制')\"><svg><use xlink:href=\"#vditor-icon-copy\"></use></svg></span></div><code class=\"hljs\"></code></pre><span data-type=\"code-block-close-marker\">```</span></div>", "<div data-block=\"0\" data-node-id=\"20200825235504-ect683e\" data-type=\"code-block\" class=\"vditor-ir__node vditor-ir__node--expand\"><span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">\u200b</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code>f<wbr>\n</code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code>f\n</code></pre><span data-type=\"code-block-close-marker\">```</span></div>"},
	{"26", "<div data-block=\"0\" data-node-id=\"20200825234341-mq1uocf\" data-type=\"math-block\" class=\"vditor-ir__node\"><span data-type=\"math-block-open-marker\">$$</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code data-type=\"math-block\" class=\"language-math\">f<wbr></code></pre><pre class=\"vditor-ir__preview\" data-render=\"1\"><code class=\"language-math\"><div class=\"vditor-math\" data-math=\"\"><span class=\"katex-display\"><span class=\"katex\"><span class=\"katex-html\" aria-hidden=\"true\"></span></span></span></div></code></pre><span data-type=\"math-block-close-marker\">$$</span></div>", "<div data-block=\"0\" data-node-id=\"20200825234341-mq1uocf\" data-type=\"math-block\" class=\"vditor-ir__node vditor-ir__node--expand\"><span data-type=\"math-block-open-marker\">$$</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code data-type=\"math-block\" class=\"language-math\">f<wbr></code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code data-type=\"math-block\" class=\"language-math\">f</code></pre><span data-type=\"math-block-close-marker\">$$</span></div>"},
	{"25", "<p data-block=\"0\" data-node-id=\"20200815231623-j492vp3\"><span data-type=\"inline-math\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker\">$</span><code data-newline=\"1\" class=\"vditor-ir__marker vditor-ir__marker--pre\" data-type=\"math-inline\">foo</code><span class=\"vditor-ir__preview\" data-render=\"2\"><code class=\"language-math\">foo</code></span><span class=\"vditor-ir__marker\">$</span></span></p>", "<p data-block=\"0\" data-node-id=\"20200815231623-j492vp3\"><span data-type=\"inline-math\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker\">$</span><code data-newline=\"1\" class=\"vditor-ir__marker vditor-ir__marker--pre\" data-type=\"math-inline\">foo</code><span class=\"vditor-ir__preview\" data-render=\"2\"><code class=\"language-math\">foo</code></span><span class=\"vditor-ir__marker\">$</span></span></p>"},
	{"24", "<p data-block=\"0\" data-node-id=\"20200825174808-w1hxf9h\">​</p>| col1<wbr> | col2 | col3 |\n| --- | --- | --- |\n|  |  |  |\n|  |  |  |", "<table data-block=\"0\" data-type=\"table\" data-node-id=\"20200825174808-w1hxf9h\"><thead><tr><th>col1<wbr></th><th>col2</th><th>col3</th></tr></thead><tbody><tr><td> </td><td> </td><td> </td></tr><tr><td> </td><td> </td><td> </td></tr></tbody></table>"},
	{"23", "<h1 data-block=\"0\" class=\"vditor-ir__node\" data-node-id=\"20200824174550-e8couax\" id=\"ir-f\" data-marker=\"#\"><span class=\"vditor-ir__marker vditor-ir__marker--heading\" data-type=\"heading-marker\"># </span>foo<wbr></h1>", "<h1 data-block=\"0\" class=\"vditor-ir__node vditor-ir__node--expand\" data-node-id=\"20200824174550-e8couax\" id=\"ir-foo\" data-marker=\"#\"><span class=\"vditor-ir__marker vditor-ir__marker--heading\" data-type=\"heading-marker\"># </span>foo<wbr></h1>"},
	{"22", "<p data-block=\"0\" data-node-id=\"20200821122908-xte3idk\">##<wbr><span>foo</span></p>", "<p data-block=\"0\" data-node-id=\"20200821122908-xte3idk\">##<wbr>foo</p>"},
	{"21", "<p data-block=\"0\" data-node-id=\"20200821090811-hrewhtz\">​<span data-type=\"emoji\" class=\"vditor-ir__node\"><span data-render=\"2\">😄</span><span class=\"vditor-ir__marker\">:smile:foo<wbr></span></span></p>", "<p data-block=\"0\" data-node-id=\"20200821090811-hrewhtz\">\u200b<span data-type=\"emoji\" class=\"vditor-ir__node\"><span data-render=\"2\">😄</span><span class=\"vditor-ir__marker\">:smile:</span></span>foo<wbr></p>"},
	{"20", "<p data-block=\"0\" data-node-id=\"20200817164045-vtlwy31\"><span data-type=\"html-entity\" class=\"vditor-ir__node\"><code data-newline=\"1\" class=\"vditor-ir__marker vditor-ir__marker--pre\" data-type=\"html-entity\">&amp;emsp;</code><span class=\"vditor-ir__preview\" data-render=\"2\"></span></span></p>", "<p data-block=\"0\" data-node-id=\"20200817164045-vtlwy31\"><span data-type=\"html-entity\" class=\"vditor-ir__node\"><code data-newline=\"1\" class=\"vditor-ir__marker vditor-ir__marker--pre\" data-type=\"html-entity\">&amp;emsp;</code><span class=\"vditor-ir__preview\" data-render=\"2\"><code>\u2003</code></span></span></p>"},
	{"19", "<p data-block=\"0\" data-node-id=\"20200817131123-tyrwvzd\"><span data-type=\"inline-math\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker\">$</span><code data-newline=\"1\" class=\"vditor-ir__marker vditor-ir__marker--pre\" data-type=\"math-inline\">foo</code><span class=\"vditor-ir__preview\" data-render=\"2\"></span><span class=\"vditor-ir__marker\">$b<wbr></span></span></p>", "<p data-block=\"0\" data-node-id=\"20200817131123-tyrwvzd\"><span data-type=\"inline-math\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker\">$</span><code data-newline=\"1\" class=\"vditor-ir__marker vditor-ir__marker--pre\" data-type=\"math-inline\">foo</code><span class=\"vditor-ir__preview\" data-render=\"2\"><code class=\"language-math\">foo</code></span><span class=\"vditor-ir__marker\">$</span></span>b<wbr></p>"},
	{"16", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20200814215643-77a63qr\"><li data-marker=\"*\" class=\"vditor-task\" data-node-id=\"\"><input checked=\"\" type=\"checkbox\"> foo<wbr></li></ul>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20200814215643-77a63qr\"><li data-marker=\"*\" class=\"vditor-task\" data-node-id=\"\"><input checked=\"\" type=\"checkbox\"/> foo<wbr></li></ul>"},
	{"15", "<p data-block=\"0\" data-node-id=\"20200813185628-tdpjvvr\">foo</p><p data-block=\"0\" data-node-id=\"20200813185636-sp1i1bp\"><span data-type=\"block-ref\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200813185628-tdpjvvr</span> <span class=\"vditor-ir__blockref\">\"bar\"</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span> baz\n```\n[text](addr)\n```<wbr></p>", "<p data-block=\"0\" data-node-id=\"20200813185628-tdpjvvr\">foo</p><p data-block=\"0\" data-node-id=\"20200813185636-sp1i1bp\"><span data-type=\"block-ref\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200813185628-tdpjvvr</span> <span class=\"vditor-ir__blockref\">&#34;bar&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span> baz</p><div data-block=\"0\" data-node-id=\"20200903115853-1ddf908\" data-type=\"code-block\" class=\"vditor-ir__node vditor-ir__node--expand\"><span data-type=\"code-block-open-marker\">```</span><span class=\"vditor-ir__marker vditor-ir__marker--info\" data-type=\"code-block-info\">​</span><pre class=\"vditor-ir__marker--pre vditor-ir__marker\"><code>[text](addr)<wbr>\n</code></pre><pre class=\"vditor-ir__preview\" data-render=\"2\"><code>[text](addr)</code></pre><span data-type=\"code-block-close-marker\">```</span></div>"},
	{"14", "<p data-block=\"0\" data-node-id=\"20200813113846-42h0ba1\">![foo](bar)<wbr></p>", "<p data-block=\"0\" data-node-id=\"20200813113846-42h0ba1\"><span class=\"vditor-ir__node vditor-ir__node--expand\" data-type=\"img\"><span class=\"vditor-ir__marker\">!</span><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">[</span><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">foo</span><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">]</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">bar</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><wbr><img src=\"bar\" alt=\"foo\"/></span></p>"},
	{"13", "<p data-block=\"0\" data-node-id=\"20200811153824-grm842y\">foo</p><blockquote data-block=\"0\" data-node-id=\"20200811153825-amjdbjz\"><p data-block=\"0\" data-node-id=\"\"><wbr><br></p></blockquote>", "<p data-block=\"0\" data-node-id=\"20200811153824-grm842y\">foo</p><p data-block=\"0\" data-node-id=\"20200811153825-amjdbjz\">&gt; <wbr></p>"},
	{"12", "<p data-block=\"0\" data-node-id=\"20200811153040-mrqu125\"><span data-type=\"a\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">[</span><span class=\"vditor-ir__link\">foo</span><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">]</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">bar</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)b<wbr></span></span></p>", "<p data-block=\"0\" data-node-id=\"20200811153040-mrqu125\"><span data-type=\"a\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">[</span><span class=\"vditor-ir__link\">foo</span><span class=\"vditor-ir__marker vditor-ir__marker--bracket\">]</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">bar</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span>b<wbr></p>"},
	{"11", "<h1 data-block=\"0\" class=\"vditor-ir__node\" data-node-id=\"20200825151101-01fo919\" id=\"ir-ba<wbr>\" data-marker=\"#\"><span class=\"vditor-ir__marker vditor-ir__marker--heading\" data-type=\"heading-marker\"># </span>foo<span data-type=\"heading-id\" class=\"vditor-ir__marker\"> {bar<wbr>}</span></h1>", "<h1 data-block=\"0\" class=\"vditor-ir__node vditor-ir__node--expand\" data-node-id=\"20200825151101-01fo919\" id=\"ir-bar&lt;wbr&gt;\" data-marker=\"#\"><span class=\"vditor-ir__marker vditor-ir__marker--heading\" data-type=\"heading-marker\"># </span>foo<span data-type=\"heading-id\" class=\"vditor-ir__marker\"> {bar<wbr>}</span></h1>"},
	{"10", "<p data-block=\"0\" data-node-id=\"20200811140724-dxmm3jo\"><span data-type=\"em\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--bi\">*</span><em data-newline=\"1\">foo</em><span class=\"vditor-ir__marker vditor-ir__marker--bi\">*b<wbr></span></span></p>", "<p data-block=\"0\" data-node-id=\"20200811140724-dxmm3jo\"><span data-type=\"em\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--bi\">*</span><em data-newline=\"1\">foo</em><span class=\"vditor-ir__marker vditor-ir__marker--bi\">*</span></span>b<wbr></p>"},
	{"9", "<p data-block=\"0\" data-node-id=\"20200810211034-9d34ae\"><span data-type=\"block-ref\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200810191413-6f5616</span> <span class=\"vditor-ir__blockref\">\"foo\"</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">) <wbr></span></span></p>", "<p data-block=\"0\" data-node-id=\"20200810211034-9d34ae\"><span data-type=\"block-ref\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">20200810191413-6f5616</span> <span class=\"vditor-ir__blockref\">&#34;foo&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span> <wbr></p>"},
	{"8", "<p data-block=\"0\" data-node-id=\"20200809211933-b81ed7\">* foo<wbr></p>", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20200809211933-b81ed7\"><li data-marker=\"*\" data-node-id=\"\">foo<wbr></li></ul>"},
	{"7", "<p data-block=\"0\" data-node-id=\"20200809184752-a537de\">&gt; foo<wbr></p>", "<blockquote data-block=\"0\" data-node-id=\"20200809184752-a537de\"><p data-block=\"0\" data-node-id=\"\">foo<wbr></p></blockquote>"},
	{"6", "<p data-block=\"0\" data-node-id=\"20200809093825-b06abb\"><span data-type=\"block-ref\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">foo</span> <span class=\"vditor-ir__blockref\">\"text<wbr>\"</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span></p>", "<p data-block=\"0\" data-node-id=\"20200809093825-b06abb\"><span data-type=\"block-ref\" class=\"vditor-ir__node vditor-ir__node--expand\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">foo</span> <span class=\"vditor-ir__blockref\">&#34;text<wbr>&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span></p>"},
	{"5", "<p data-block=\"0\" data-node-id=\"1596459249782\">foo ((bar)) <wbr></p>", "<p data-block=\"0\" data-node-id=\"1596459249782\">foo <span data-type=\"block-ref\" class=\"vditor-ir__node\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">bar</span> <span class=\"vditor-ir__blockref\">&#34;bar&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span> <wbr></p>"},
	{"4", "<p data-block=\"0\" data-node-id=\"1596459249782\">((foo \"text\")<wbr></p>\n", "<p data-block=\"0\" data-node-id=\"1596459249782\">((foo &#34;text&#34;)<wbr></p>"},
	{"3", "<p data-block=\"0\" data-node-id=\"1596459249782\">((foo \"text\"))<wbr></p>\n", "<p data-block=\"0\" data-node-id=\"1596459249782\"><span data-type=\"block-ref\" class=\"vditor-ir__node vditor-ir__node--expand\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">foo</span> <span class=\"vditor-ir__blockref\">&#34;text&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span><wbr></p>"},
	{"2", "<p data-block=\"0\" data-node-id=\"1596459249782\">((foo))<wbr></p>\n", "<p data-block=\"0\" data-node-id=\"1596459249782\"><span data-type=\"block-ref\" class=\"vditor-ir__node vditor-ir__node--expand\"><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">(</span><span class=\"vditor-ir__marker vditor-ir__marker--link\">foo</span> <span class=\"vditor-ir__blockref\">&#34;foo&#34;</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span><span class=\"vditor-ir__marker vditor-ir__marker--paren\">)</span></span><wbr></p>"},
	{"1", "<p data-block=\"0\" data-node-id=\"1\">foo</p><p data-block=\"0\" data-node-id=\"2\"><wbr><br></p>", "<p data-block=\"0\" data-node-id=\"1\">foo</p><p data-block=\"0\" data-node-id=\"2\"><wbr></p>"},
	{"0", "<p data-block=\"0\" data-node-id=\"1\">foo</p><p data-block=\"0\" data-node-id=\"20200811112006-322210\"><wbr><br></p>", "<p data-block=\"0\" data-node-id=\"1\">foo</p><p data-block=\"0\" data-node-id=\"20200811112006-322210\"><wbr></p>"},
}

func TestSpinVditorIRBlockDOM(t *testing.T) {
	luteEngine := lute.New()
	luteEngine.BlockRef = true

	for _, test := range spinVditorIRBlockDOMTests {
		html := luteEngine.SpinVditorIRBlockDOM(test.from)

		if "15" == test.name || "18" == test.name {
			// 去掉最后一个新生成的节点 ID，因为这个节点 ID 是随机生成，每次测试用例运行都不一样，比较没有意义，长度一致即可
			lastNodeIDIdx := strings.LastIndex(html, "data-node-id=")
			length := len("data-node-id=\"20200813190226-1234567\" ")
			html = html[:lastNodeIDIdx] + html[lastNodeIDIdx+length:]
			lastNodeIDIdx = strings.LastIndex(test.to, "data-node-id=")
			test.to = test.to[:lastNodeIDIdx] + test.to[lastNodeIDIdx+length:]
		}

		if test.to != html {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal html\n\t%q", test.name, test.to, html, test.from)
		}
	}
}
