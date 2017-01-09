Dsp接口测试工具使用说明
============
## 使用平台 ##
Linux操作系统

##使用用途##
检查Dsp自己的接口是否与芒果TV提供的协议文档相吻合

## 使用方法 ##


1. 在Linux平台上赋予该二进制可执行权限， chmod 777 dsp_check_tool
2. 运行dsp_check_tool  运行命令 ./dsp_check_tool -h ，查看帮助信息
3. 根据不同的测试需求选择相应的参数。带参数运行dsp_check_tool，参数具体含义如下：  
    - -d    `打印详细  `   
    - -t string `检查类型, pc_front: pc 前贴片, pc_pause: pc 暂停, app_front: app 前贴片, app_pause: app 暂停, pd_deal:首选交易, app_banner:App Banner (Default "pc_front")`    
    - -u string  `dsp 使用的通信url (default "http://127.0.0.1/mgtv/bid")`   
4. 当选择-d参数后，工具会打印出检查的详细信息。如果检查通过，则会返回 `Congratulations, test passed` 提示；如果检查不通过则会打印出对应的错误提示。

## 检查项 ##

- creative_id 和 ad_url
	-  有且只能有一个存在
	-  ad_url必须包含"http://"
- click_through_url：必须存在落地页
- iurl：
	-  必须存在曝光检测地址url
	-  上报event为0
	-  url中必须包含宏替换字段`%%SETTLE_PRICE%%`
- curl：必须存在点击检测地址
- ctype：素材类型必须与请求的一致
- duration：素材时长必须介于[min_playtime,max_playtime]之间
- price：
   -  rtb广告中必须大于或者等于请求中的min_cpm_price
   -  首选交易广告必须大于请求中的price
- title：必须返回素创意标题


