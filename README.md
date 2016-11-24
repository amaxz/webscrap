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

    Query keyword, ex. `./webscrap "iPhone 5s" "SONY Z3"`, output:
    
    ```
    2016/05/13 16:16:09 {"src":"天猫商城","stamp":"1463127369","q":"huawei p9","url":"http://detail.tmall.com/item.htm?id=529420098972","vendor":"华为官方旗舰店","price":"3188.00","title":"Huawei/华为 P9"}
    2016/05/13 16:16:09 {"src":"天猫商城","stamp":"1463127369","q":"huawei p9","url":"http://detail.tmall.com/item.htm?id=529460804269","vendor":"华为官方旗舰店","price":"3688.00","title":"Huawei/华为 P9全网通"}
    2016/05/13 16:16:09 {"src":"天猫商城","stamp":"1463127369","q":"huawei p9","url":"http://detail.tmall.com/item.htm?id=530876708666","vendor":"华为官方旗舰店","price":"3988.00","title":"Huawei/华为 P9 Plus红外线"}
    2016/05/13 16:16:10 {"src":"京东商城","stamp":"1463127370","q":"huawei p9","url":"http://item.jd.com/1959364405.html","vendor":"FUSTAR数码专营店","price":"3148.00","title":"全球购 华为 HUAWEI P9 智能手机 全网通 双卡双4G 钛银灰 3G/32GB 标配"}
    2016/05/13 16:16:10 {"src":"京东商城","stamp":"1463127370","q":"huawei p9","url":"http://item.jd.com/10313117484.html","vendor":"能良数码旗舰店","price":"3688.00","title":"华为(HUAWEI) P9 4G手机 双卡双待 金色 全网通(4G RAM+64G ROM)标配"}
    2016/05/13 16:16:10 {"src":"京东商城","stamp":"1463127370","q":"huawei p9","url":"http://item.jd.com/10320363255.html","vendor":"谷高数码手机专营店","price":"3688.00","title":"华为HUAWEI华为P9全网通双卡双待4G手机 琥珀金 全网通(4G RAM+64G ROM)"}
    2016/05/13 16:16:10 {"src":"京东商城","stamp":"1463127370","q":"huawei p9","url":"http://item.jd.com/10313313379.html","vendor":"能良数码旗舰店","price":"4199.00","title":"华为(HUAWEI) P9 Plus 4G手机 双卡双待 琥珀金 全网通(4G RAM+64G ROM)标配"}
    2016/05/13 16:16:10 {"src":"京东商城","stamp":"1463127370","q":"huawei p9","url":"http://item.jd.com/10313600913.html","vendor":"爱优创杰手机专营店","price":"3988.00","title":"华为 HUAWEI P9 Plus 移动联通电信4G 琥珀金 全网通(4GB+64GB)"}
    2016/05/13 16:16:12 {"src":"苏宁易购","stamp":"1463127372","q":"huawei p9","url":"http://product.suning.com/149363540.html","vendor":"3C瑞业数码手机专营店","price":"4099.00","title":"HUAWEI P9 4GB+64GB 全网通版（陶瓷白）"}
    2016/05/13 16:16:12 {"src":"苏宁易购","stamp":"1463127372","q":"huawei p9","url":"http://product.suning.com/149363538.html","vendor":"蚂蚁客手机数码官方旗舰店","price":"3688.00","title":"HUAWEI P9 4GB+64GB 全网通版（金色）"}
    2016/05/13 16:16:12 {"src":"苏宁易购","stamp":"1463127372","q":"huawei p9","url":"http://product.suning.com/149366439.html","vendor":"锐奇网信数码专营店","price":"2988.00","title":"HUAWEI P9 3GB+32GB 移动定制版（钛银灰）"}
    2016/05/13 16:16:12 {"src":"苏宁易购","stamp":"1463127372","q":"huawei p9","url":"http://product.suning.com/148849482.html","vendor":"锐奇网信数码专营店","price":"3688.00","title":"华为 HUAWEI P9 4GB+64GB版 全网通（金色）"}
    2016/05/13 16:16:12 {"src":"苏宁易购","stamp":"1463127372","q":"huawei p9","url":"http://product.suning.com/148849480.html","vendor":"3C瑞业数码手机专营店","price":"4199.00","title":"华为 HUAWEI P9 4GB+64GB版 全网通（陶瓷白）"}
    2016/05/13 16:16:12 {"src":"苏宁易购","stamp":"1463127372","q":"huawei p9","url":"http://product.suning.com/149534958.html","vendor":"蚂蚁客手机数码官方旗舰店","price":"2988.00","title":"HUAWEI P9 3GB+32GB 移动定制版（皓月银）"}
    2016/05/13 16:16:12 {"src":"苏宁易购","stamp":"1463127372","q":"huawei p9","url":"http://product.suning.com/149366453.html","vendor":"新世纪硅谷数码专营店","price":"2988.00","title":"HUAWEI P9 3GB+32GB 联通定制版（皓月银）"}
    2016/05/13 16:16:12 {"src":"苏宁易购","stamp":"1463127372","q":"huawei p9","url":"http://product.suning.com/149366478.html","vendor":"3C瑞业数码手机专营店","price":"3188.00","title":"HUAWEI P9 3GB+32GB 联通定制版（钛银灰）"}
    2016/05/13 16:16:12 {"src":"苏宁易购","stamp":"1463127372","q":"huawei p9","url":"http://product.suning.com/149363537.html","vendor":"新世纪硅谷数码专营店","price":"3688.00","title":"HUAWEI P9 4GB+64GB 全网通版（玫瑰金）"}
    2016/05/13 16:16:12 {"src":"苏宁易购","stamp":"1463127372","q":"huawei p9","url":"http://product.suning.com/148849479.html","vendor":"锐奇网信数码专营店","price":"3188.00","title":"华为 HUAWEI P9 3GB+32GB版 全网通（钛银灰）"}
    2016/05/13 16:16:12 {"src":"苏宁易购","stamp":"1463127372","q":"huawei p9","url":"http://product.suning.com/148849481.html","vendor":"五洲手机专营店","price":"3699.00","title":"华为 HUAWEI P9 4GB+64GB版 全网通（玫瑰金）"}
    2016/05/13 16:16:12 {"src":"苏宁易购","stamp":"1463127372","q":"huawei p9","url":"http://product.suning.com/148849478.html","vendor":"锐奇网信数码专营店","price":"3188.00","title":"华为 HUAWEI P9 3GB+32GB版 全网通（皓月银）"}
    ```