## Web Scrap 

Webscrap is a price fetcher of commodity sells on e-commerce website like `JD.com`, `Suning.com`, `Tmall.com`, etc

### How to build

* win32

        env GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -ldflags -s -o webscrap-win32.exe webscrap

* linux

        env GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -ldflags -s -o webscrap-linux webscrap

* OS X

        env GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags -w -o webscrap-darwin webscrap

<!--
env GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -ldflags -s -o webscrap-win32.exe webscrap
env GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags -w -o webscrap-darwin webscrap
env GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -ldflags -s -o webscrap-linux webscrap
-->

### How to run

Usage: `./webscrap [-o <path>] [-f <path> | <keyword>...]`
    
* `-f path`

    Input file name (alternative with environment `$WOEGO_WEBSCRAP_FILE`), the input file is a plain text file of keywords which splited with line seperator '\n'

* `-o path`

    Output file name (default "./output/S`${UnixTime}`.txt"), search result output with json format line text prefixed with timestamp  

* `-s number`

    Minimum sleep duration second (default 10)
    
* `<keyword>...`

    Query keyword, ex. `./webscrap "huawei p9" "SONY Z3"`, output:
    
```
========================== 天猫商城: huawei p9 ==========================
 1  3688.00 华为官方旗舰店  等更多商家 <Huawei/华为 P9全网通> (http://detail.tmall.com/item.htm?id=529460804269)
 4  4388.00 华为官方旗舰店  等更多商家 <Huawei/华为 P9 plus红外线> (http://detail.tmall.com/item.htm?id=530876708666)
 5  3188.00 华为官方旗舰店  等更多商家 <Huawei/华为 P9> (http://detail.tmall.com/item.htm?id=529420098972)
10  5088.00 华为莫问专卖店  等更多商家 <Huawei/华为 P9 Plus全网通HUAWEI P9 Plus> (http://detail.tmall.com/item.htm?id=538015904116)
11  3188.00 爱要美数码专营店  等更多商家 <Huawei/华为 P9 3GB+32GB> (http://detail.tmall.com/item.htm?id=539907130215)
12  3788.00 广东电信亿品汇专卖店  等更多商家 <Huawei/华为 P9 4GB+64GB> (http://detail.tmall.com/item.htm?id=540392109982)

========================== 京东商城: huawei p9 ==========================
 5  2949.00 谷高数码手机专营店 <华为HUAWEI华为P9双卡双待4G手机 灰色 全网通(3G RAM+32G ROM)标配> (http://item.jd.com/10313430790.html)
 6  4688.00 能良数码旗舰店 <华为(HUAWEI) P9 Plus 4G手机 双卡双待 琥珀灰 全网通(4G RAM+128G ROM)标配> (http://item.jd.com/10393084083.html)
 8  4188.00 君问手机旗舰店 <华为（HUAWEI）p9 plus智能手机 双卡双待 琥珀金 全网通(4G+128G)标配> (http://item.jd.com/10967402384.html)
 9  4688.00 爱优创杰手机专营店 <华为 HUAWEI P9 Plus 移动联通电信4G手机 琥珀金 全网通(4G RAM+128G ROM)标配> (http://item.jd.com/10313600912.html)
11  3988.00 话机世界官方旗舰店 <华为(HUAWEI) P9 Plus 4G手机 双卡双待 琥珀金 全网通(4G RAM+64G ROM)> (http://item.jd.com/10601115291.html)
12  2988.00 易道手机专营店 <华为(HUAWEI) P9 4G手机 32G标配 全网通双卡双待 皓月银 移动版(3G RAM+32G ROM)标配> (http://item.jd.com/10658071700.html)
13  4888.00 谷高数码手机专营店 <华为HUAWEI P9 Plus双卡双待4G手机 灰色 全网通(4G RAM+128G ROM)套装> (http://item.jd.com/10380973103.html)
14          达沃手机专营店 <华为(HUAWEI) P9 琥珀灰 全网通(3GB+32GB)标配> (https://ccc-x.jd.com/dsp/nc?ext=)
15  5388.00 CBI国际数码全球购专营店 <全球购 (HUAWEI) 华为 P9 Plus 双卡双待 全网通4G 智能手机 金色 128G> (http://item.jd.com/1966634438.html)
18  3488.00 君问手机旗舰店 <华为(HUAWEI)p9 64G智能手机 双卡双待 金色 全网通(4G+64G)标配> (http://item.jd.com/10678599497.html)
20  4299.00 弘利源全球购数码通讯专营店 <华为HUAWEI 华为P9 Plus 双卡双待 智能手机 琥珀金 64G> (http://item.jd.com/1970942501.html)
21  3249.00 河北启迪手机专营店 <华为(HUAWEI) P9 全网通手机 金色 4GB RAM+64GB ROM版> (http://item.jd.com/10941841658.html)
22  3788.00 能良数码旗舰店 <华为(HUAWEI) P9 4G手机 双卡双待 琥珀灰 全网通(4G RAM+64G ROM)标配> (http://item.jd.com/10441345381.html)
23  2598.00 滨泽数码手机旗舰店 <华为HUAWEI华为P9 双卡双待4G手机 皓月银 移动4G(3G RAM+32G ROM) 标配版> (http://item.jd.com/10743941046.html)
24  2988.00 GIT通讯数码旗舰店 <全球购 (HUAWEI) 华为P9 双卡双待智能手机 皓月银 全网通 3G/32GB 标配> (http://item.jd.com/1962557098.html)
26  1888.00 surepromise手机海外旗舰店 <【全球购】华为（HUAWEI)P9 Lite 16GB 5.2寸高清屏幕Andriod6. 白色> (http://item.jd.com/1964108672.html)
29  5388.00 SOULDIO全球购专营店 <全球购 (HUAWEI) 华为 P9 Plus 双卡双待 通4G 智能手机 金色 128G> (http://item.jd.com/1969767072.html)

========================== 苏宁易购: huawei p9 ==========================
11  3096.00 亨通达数码官方旗舰店 <华为(HUAWEI) P9 双卡双待4G手机 琥珀金 全网通(4G RAM+64G ROM）高配版送P9点纹透明保护壳(标配自带)+Type-C转换头 全新原封 现货急速发 货票同行> (http://product.suning.com/0070147903/172825518.html)
12  3688.00 苏宁自营 <华为(HUAWEI) P9 双卡双待4G手机 琥珀金 全网通(4G RAM+64G ROM）高配版立减500 现货急速发 送P9点纹透明保护壳(标配自带)+Type-C转换头> (http://product.suning.com/0070080854/148179309.html)
28  3086.00 3C瑞业数码手机专营店 <华为 HUAWEI P9 4GB+64GB版 全网通（琥珀金）【送开窗智能保护套+自带保护壳+Type-C转换头+自拍杆！现货速发！顺丰包邮！】> (http://product.suning.com/0070071664/148849482.html)
31  2878.00 亨通达数码官方旗舰店 <华为(HUAWEI) P9 双卡双待4G手机 皓月银 全网通(3G RAM+32G ROM)送P9点纹透明保护壳(标配自带)+Type-C转换头 全新原封 现货急速发 货票同行> (http://product.suning.com/0070147903/173296619.html)
36  3688.00 亨通达数码官方旗舰店 <华为(HUAWEI) P9 Plus 双卡双待4G手机 琥珀金 全网通(4G RAM+64G ROM) 高配送华为原装Type C数据线+P9 plus点纹透明保护壳(标配自带)+Type-C转换头 全新原封 货票同行> (http://product.suning.com/0070147903/173285593.html)
37  3988.00 苏宁自营 <华为(HUAWEI) P9 Plus 双卡双待4G手机 琥珀金 全网通(4G RAM+64G ROM) 高配现货急速发 购买即送点纹透明保护壳(标配自带)+Type-C转换头> (http://product.suning.com/0070080854/151440878.html)
40  3199.00 橙子数码手机专营店 <华为(HUAWEI) P9 双卡双待4G手机 玛瑙红 全网通(4G RAM+64G ROM）高配版现货急速发 送P9点纹透明保护壳(标配自带)+华为原装Type C 数据线+Type-C转换头> (http://product.suning.com/0070147903/176301592.html)
41  3088.00 华科手机专营店 <华为(HUAWEI) P9 双卡双待4G手机 琥珀金 全网通(4G RAM+64G ROM）高配版现货急速发 送P9点纹透明保护壳(标配自带)+Type-C转换头> (http://product.suning.com/0070080854/149363538.html)
42  3086.00 3C瑞业数码手机专营店 <华为(HUAWEI) P9 双卡双待4G手机 琥珀金 全网通(4G RAM+64G ROM）高配版立减500 全新原封 现货急速发 送P9点纹透明保护壳(标配自带)+Type-C转换头> (http://product.suning.com/0070080854/148849482.html)
46  2485.00 安铎数码专营店 <华为(HUAWEI) P9 双卡双待4G手机 皓月银 移动版(3G RAM+32G ROM)送P9点纹透明保护壳(标配自带)+Type-C转换头 全新原封 现货急速发 货票同行> (http://product.suning.com/0070147903/173296623.html)
48  3315.00 橙子数码手机专营店 <华为(HUAWEI) P9 双卡双待4G手机 陶瓷白 全网通(4G RAM+64G ROM) 高配版全新原封 现货急速发 购买即送P9点纹透明保护壳(标配自带)+Type-C转换头> (http://product.suning.com/0070080854/149363540.html)
49  3488.00 橙子数码手机专营店 <华为(HUAWEI) P9 双卡双待4G手机 托帕蓝 全网通(4G RAM+64G ROM) 高配版现货急速发 送P9点纹透明保护壳(标配自带)+华为原装Type C 数据线+Type-C转换头> (http://product.suning.com/0070147903/176301591.html)
54  2498.00 3C瑞业数码手机专营店 <华为(HUAWEI) P9 双卡双待4G手机 钛银灰 移动版(3G RAM+32G ROM) 标配版立减400 现货急速发 送点纹透明保护壳（标配自带）+Type-C转换头> (http://product.suning.com/0070080854/147702581.html)
58  3688.00 锐易达手机专营店 <HUAWEI P9 4GB+64GB 全网通版（琥珀灰）后置1200万徕卡双摄像头！大光圈拍摄，指纹识别！> (http://product.suning.com/0070060611/157255696.html)
62  3688.00 亨通达数码官方旗舰店 <华为(HUAWEI) P9 双卡双待4G手机 陶瓷白 全网通(4G RAM+64G ROM) 高配版送P9点纹透明保护壳(标配自带)+华为原装Type C 数据线+Type-C转换头 全新原封 现货急速发 货票同行> (http://product.suning.com/0070147903/172830438.html)

========================== 天猫商城: SONY Z3 ==========================

========================== 京东商城: SONY Z3 ==========================
 1  3039.00 魔力世界数码生活专营店 <SONY Z3> (http://item.jd.com/1959117309.html)
 2  2999.00 索尼手机旗舰店 <索尼(SONY)Xperia Z3+ E6533 防水防尘 移动联通双4G手机 香槟金> (http://item.jd.com/1673109941.html)
 3  2160.00 XINKE全球购专营店 <索尼(SONY)Xperia Z3+ E6533 移动联通双4G手机 湖水绿 Z3+E6533 双4G版 双卡> (http://item.jd.com/1967531086.html)
 4  2168.00 数字创想科技（香港）专营店 <索尼 SONY E6533 Z3+ 双卡双待 移动联通双4G智能防水手机 涧湖绿 Z3+E6533 双4G版 双卡> (http://item.jd.com/1962284151.html)
 5  2199.00 Kyushu Star专营店 <索尼 SONY Xperia Z3+ E6533 E6553 防水防尘 正品 白色 Z3+E6533 双4G版 双卡> (http://item.jd.com/1956061650.html)
 6  2299.00 FUSTAR数码专营店 <索尼 SONY Xperia Z3+ E6533 双卡双4G手机 防水防尘 涧湖绿> (http://item.jd.com/1957574590.html)
 7  2388.00 CBI国际数码全球购专营店 <SONY（索尼）XperiaZ3+E6533双卡 移动联通智能4G手机 港版 绿色双卡双待 移动联通4G 32G> (http://item.jd.com/1967537094.html)
 8  2388.00 HK数码国际全球购专营店 <索尼 SONY Xperia Z3+ E6533 Dual移动联通双4G 双卡防水防尘手机 黛丽黑 Z3+E6533 双4G版 双卡> (http://item.jd.com/1964741712.html)
 9  2245.00 eBay精选 <【eBay精选】索尼Sony Xperia Z3＋ E6533 S6935 4G智能手机> (http://item.jd.com/1964134510.html)
10  2399.00 Quantum海外专营店 <索尼手机 Z3/Z3+ D6683 港行正品 SONY手机 Z3 黛丽黑> (http://item.jd.com/1959454359.html)
11  2245.00 eBay精选 <【eBay精选】索尼Sony Xperia Z3＋ E6533 S6938 4G智能手机> (http://item.jd.com/1964132930.html)
12  2588.00 SOULDIO全球购专营店 <SONY（索尼）Xperia Z3+ E6533 双卡双待 移动联通4G智能手机 港版 绿色 32G> (http://item.jd.com/1964965445.html)
13  2788.00 港岛海外购专营店 <索尼手机 Z3/Z3+ D6683 港行正品 SONY手机 Z3+ 涧湖色> (http://item.jd.com/1971543547.html)
15  2399.00 量子科技专营店 <全球购 索尼手机 Z3/Z3+ D6683 港行正品 SONY手机 Z3 黛丽黑> (http://item.jd.com/1954180300.html)
16  2388.00 GIT通讯数码旗舰店 <全球购 海外版 SONY(索尼）Z3 双卡 移动 联通 4G手机 黑色 16GB 官方标配> (http://item.jd.com/1952763836.html)
17  2798.00 香港智诚博扬专营店 <全球购 SONY(索尼) D6683 Z3 港行 双卡双待 移动 联通4G 手机 香槟金> (http://item.jd.com/1957323537.html)
18  3446.00 eBay精选 <【eBay精选】索尼/Sony Xperia Z3 Dual 16GB 黑色 智能手机> (http://item.jd.com/1965463415.html)
19  2245.00 eBay精选 <【eBay精选】索尼（SONY）Xperia Z3+E6533 双卡双待 4G 智能手机> (http://item.jd.com/1964207869.html)
20  2828.00 香港智诚博扬专营店 <全球购 SONY（索尼）E6533 Z3+ 双卡双待 移动联通智能手机 涧湖绿> (http://item.jd.com/1957327812.html)
21  2662.00 eBay精选 <【eBay精选】索尼/Sony Xperia Z3+ 5.2吋 32GB 蓝色 智能手机> (http://item.jd.com/1961028133.html)
22  2388.00 易淘国际专营店 <SONY（索尼）Xperia Z3+ E6533 Z3 双卡移动联通智能4G手机 港版 金色双卡双待 移动联通4G 32G> (http://item.jd.com/1971441534.html)
23  2399.00 环球名品海外专营店 <全球购 SONY（索尼）E6533 Z3+ 双卡双待 移动联通智能手机 古铜金> (http://item.jd.com/1961604147.html)
24  2368.00 eBay精选 <【eBay精选】索尼/Sony Xperia Z3+ Dual 32GB 白色 智能手机> (http://item.jd.com/1964232119.html)
25  2388.00 弘利源全球购数码通讯专营店 <SONY（索尼）Xperia Z3+ E6533 金 色 32G> (http://item.jd.com/1966932599.html)
26  2388.00 CBI国际数码全球购专营店 <SONY（索尼）Xperia Z3+ E6533 Z3 双卡双待 4G智能手机 港版 白色 32G> (http://item.jd.com/1963953967.html)
27  2099.00 Citiwide 国际专营店 <索尼(SONY) Xperia Z3 Dual D6633 16GB 双卡版 铜色> (http://item.jd.com/1973652279.html)

========================== 苏宁易购: SONY Z3 ==========================
 2  2488.00 數字創想科技（香港）有限公司海外专营店 <索尼（Sony）Z3+ E6533 防水防尘 双卡双待 涧湖绿色 32G 日本品牌赠送钢化贴膜+转换头！限时限量特价回馈！防水 防尘 轻薄机身，电池续航力特长 2070万像素索尼镜头> (http://product.suning.com/0070097650/137833659.html)
 4  2388.00 數字創想科技（香港）有限公司海外专营店 <索尼（Sony）Z3+ E6553 防水防尘 单卡 白色 32G 日本品牌下单即送钢化膜+转换头！SONY大法好！专业防水！强劲骁龙™ 810 处理器! 2070 万画素相机> (http://product.suning.com/0070097650/145805424.html)
 5  2488.00 數字創想科技（香港）有限公司海外专营店 <索尼（Sony）Z3+ E6533 防水防尘 双卡双待 黑色 32G 日本品牌赠送钢化贴膜+转换头！限时限量特价回馈！防水 防尘 轻薄机身，电池续航力特长 2070万像素索尼镜头> (http://product.suning.com/0070097650/137102678.html)
 6  3946.00 苏宁海外购自营店 <SONY XPERIA Z3TABLET COMPACT LTE BLACK 16GB SGP621轻巧更为至上 骁龙801 四核心 特大电池容量 防水平板电脑> (http://product.suning.com/0070088237/124295874.html)
 7  2488.00 數字創想科技（香港）有限公司海外专营店 <索尼（Sony）Z3+ E6533 防水防尘 双卡双待 白色 32G 日本品牌赠送钢化贴膜+转换头！限时限量特价回馈！防水 防尘 轻薄机身，电池续航力特长 2070万像素索尼镜头> (http://product.suning.com/0070097650/137141498.html)
 ```
