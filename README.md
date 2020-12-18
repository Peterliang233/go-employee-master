# API 文档
- [API 文档](#api---)
	+ [basic Response Body]
* [0,用户登录]（#0---）
* [1.获取员工信息](#1---)
* [2.添加员工信息](#2----)
* [3.更新员工信息](#3-----)
* [4.删除员工信息](#4------)

##用户登录
- URL: v1/login

- Method: POST

- Request Body

```json
{
  "username" : "Peter",
  "password": "3607812001lyp"
}
```

- Response Body

```json
{
  "code" : 0,
  "msg" : "登录成功",
  "detail" : "welcome",
  "username" : "Peter",
  "roles" : "admin",
  "token" : "token"
}
```
## 获取员工信息
- URL: v1/user/employee

- Method: GET

- Request Body

```json
{
	"number": "6109119121"
}
```

- Response Body

```json
{
	"code" : 999999,
	"employee":{
		"number" :    "6109119121",
		"name" :       "Peterliang",
		"profession" : "computer",
		"task" :       "move trick"
	},
    "message": "ok"
}
```

| 序号 |   参数    |     类型      |         规则          |    简介    |
| :--: | :-------: | :-----------: | :-------------------: | :--------: |
|  1   | number  |     string      |       数字构成       |     -      |

## 添加员工信息

- URL:v1/user/employee

- Method: POST

- Request Body

```json
{
	"number" :    "6109119121",
	"name" :       "Peterliang",
	"profession" : "computer",
	"task" :       "move trick"
}
```

- Response Body

```json
{
	"code" : 999999,
	"employee":{
		"number" :    "6109119121",
		"name" :       "Peterliang",
		"profession" : "computer",
		"task" :       "move trick"
	},
    "message" : "ok"
}
```

 | 序号 |  参数   |  类型  |  简介  | 必须 |
  | :--: | :-----: | :----: | :----: | :--: |
  |  1   |  number  | String |  工号  |  Y   |
  |  2   |  name  | String | 姓名 |  Y   |
  |  3   |   profession    | String | 专业 |  Y   |
  |  4   | task | String | 任务 |  Y   |

## 修改员工信息

- URL: v1/user/employee

- Method: PUT

- Request Body

```json
{
	"number" :    "6109119121",
	"name" :       "Peterliang",
	"profession" : "computer",
	"task" :       "move trick"
}
```

- Response Body

```json
{
	"code" : 999999,
	"employee":{
		"number" :    "6109119121",
		"name" :       "Peterliang",
		"profession" : "computer",
		"task" :      "move trick"
	},
    "message":"ok"
}
```

 | 序号 |  参数   |  类型  |  简介  | 必须 |
  | :--: | :-----: | :----: | :----: | :--: |
  |  1   |  number  | String |  工号  |  Y   |
  |  2   |  name  | String | 新的姓名 |  N   |
  |  3   |   profession    | String | 新的专业 |  N   |
  |  4   | task | String | 新的任务 |  N   |


## 删除员工信息

- URL: v1/user/employee

- Method: DELETE

- Request Body

```json
{
	"number": "6109119121"
}
```

- Response Body

```json
{
	"code" : 999999,
	"employee":{
		"number" :    "6109119121",
		"name" :       "Peterliang",
		"profession" : "computer",
		"task" :       "move trick"
	},
    "message" : "ok"
}
```

 | 序号 |  参数   |  类型  |  简介  | 必须 |
  | :--: | :-----: | :----: | :----: | :--: |
  |  1   |  number  | String |  工号  |  Y   |
