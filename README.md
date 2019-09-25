# Go WebAssembly 图片缩放

## 开发

设置IDEA开发和编译环境, 参考 https://github.com/golang/go/wiki/Configuring-GoLand-for-WebAssembly .

## 测试

1. 将生成的wasm命名为**image-scale.wasm**并复制到**web**目录中；
2. 复制go安装包中**misc/wasm/wasm_exec.js**到**web**目录中；
3. **web**目录中的内容放入web服务器中访问页面即可测试.

> 注意: web服务器一般需要增加mime配置, 否则会出现异常: `ncaught (in promise) TypeError: Failed to execute 'compile' on 'WebAssembly': Incorrect response MIME type. Expected 'application/wasm'.` .
> 以nginx为例, 修改nginx.conf同文件夹中的**mime.types**, 在types块中添加`application/wasm wasm;` 即可解决问题.
> 推荐直接写一个go的http文件服务器来测试.
