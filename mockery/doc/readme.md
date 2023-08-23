
mockery，针对于接口生成对应的实现。对于这样一个工具，它应该具有这三个功能：
1. 找到你需要实现的接口；
2. 将对应的实现生成到固定的路径；
3. 根据参数，返回对应的参数。
   所以，我们只需要知道，如何操作mockery，实现上述三种功能即可；

   
```azure
mockery --dir ./article --output ./article/mock --all

mockery --dir ./article --output ./article/mock --name "ArticleRepository"
```
- all：会实现文件下所有接口，每个接口单独一个文件，以接口给文件命名； 
- dir：接口所在的目录；
- output：表示接口实现对象所在的目录；



