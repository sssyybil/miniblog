
数据库表创建成功后，可直接根据表生成相应的 Go Model 文件。使用到的工具为 [GORM gen tool](https://gorm.io/gen/gen_tool.html)。

* 安装 gen tool

```shell
go install gorm.io/gen/tools/gentool@latest
```
* 命令详解
```markdown
gentool -h

Usage of gentool:
 -c string
       config file path 
 -db string
       input mysql or postgres or sqlite or sqlserver. consult[https://gorm.io/docs/connecting_to_the_database.html] (default "mysql")
 -dsn string
       consult[https://gorm.io/docs/connecting_to_the_database.html]
 -fieldNullable
       generate with pointer when field is nullable
 -fieldWithIndexTag
       generate field with gorm index tag
 -fieldWithTypeTag
       generate field with gorm column type tag
 -modelPkgName string
       generated model code's package name
 -outFile string
       query code file name, default: gen.go
 -outPath string
       specify a directory for output (default "./dao/query")
 -tables string
       enter the required data table or leave it blank
 -onlyModel
       only generate models (without query file)
 -withUnitTest
       generate unit test for query code
 -fieldSignable
       detect integer field's unsigned type, adjust generated data type
```


# 使用
```shell
```


