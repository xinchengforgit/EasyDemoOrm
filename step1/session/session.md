#### 关于session结构体

session结构体包含三个成员变量

~~~
db * sql.DB //使用sql.open()方法连接数据库后返回的指针
第二个和第三个成员变量用来凭借SQL语句和SQL语句中占位符的对应值，用户调用Raw()方法即可改变这两个变量的值
~~~

#### 封装Exec(),Query(),QueryRow()三个原生方法

~~~
从代码易得
~~~

目的：

- 封装有 2 个目的，一是统一打印日志（包括 执行的SQL 语句和错误日志）。
- 二是执行完成后，清空 `(s *Session).sql` 和 `(s *Session).sqlVars` 两个变量。这样 Session 可以复用，开启一次会话，可以执行多次 SQL。