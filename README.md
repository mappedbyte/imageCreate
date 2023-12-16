# imageCreate
Microsoft Bing Image Designer  微软图像创建器 golang接口实现

Bing Image Creator是微软出品的AI文生图平台，它的用户界面非常友好易于上手。你只需输入简短的短语或关键词即可生成图像，而无需进行复杂的设置或调整。


先上几张生成的照片：
![图1.jpg](image%2F%CD%BC1.jpg)
![图2.jpg](image%2F%CD%BC2.jpg)
![图3.jpg](image%2F%CD%BC3.jpg)
![图4.jpg](image%2F%CD%BC4.jpg)

地址:https://cn.bing.com/images/create?FORM=GENEXP 需要科学上网。

进入之后获取Cookie,并且填充到[imageDesigner.cnf](imageDesigner.cnf)

接口提取自官网网页,同样需要在配置文件中配置代理

接口1: 服务地址:9999/submit?message=xxx

![请求1.jpg](image%2F%C7%EB%C7%F31.jpg)

接口2: 服务地址:9999/result/{id}

![请求2.jpg](image%2F%C7%EB%C7%F32.jpg)



