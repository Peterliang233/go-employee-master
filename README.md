# api文档
- [API 文档](#api文档)
	+ [basic Response Body]
- [1.用户登录](#用户登录)
- [2.获取员工信息](#获取员工信息)
- [3.添加员工信息](#添加员工信息)
- [4.更新员工信息](#更新员工信息)
- [5.删除员工信息](#删除员工信息)

## 用户登录
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
  "code": 0,
  "detail": "welcome",
  "msg": "登录成功",
  "roles": [
    "admin"
  ],
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlBldGVyIiwicm9sZXMiOlsiYWRtaW4iXSwiZXhwIjoxNjA4NTM0NjIyLCJpc3MiOiJnby1zeXMtZW1wbG95ZWUifQ.Eezj0z8o2i0Map_cKygRsRvv1YkHktwQZyP9zsDsFJE",
  "username": "Peter"
}
```

| 序号 |   参数    |     类型      |         规则          |    简介    |
| :--: | :-------: | :-----------: | :-------------------: | :--------: |
| 1 | username | varchar(33) | 用户名 | 登录用户 |
| 2 | password |varchar(33)  | 用户密码 | - |

## 获取员工信息
- URL: v1/user/employee

- Method: GET

- Request Body

```json
{
	"username" : "aaa"
}
```

- Response Body

```json
{
	"code": 2000,
	"data": {
		"departments": [
			"developer group"
		],
		"employee": {
			"id": 666,
			"real_name": "Peterliang",
			"nick_name": "xiaoliang",
			"english_name": "peter",
			"sex": "male",
			"age": 19,
			"address": "ruijing",
			"mobile_phone": "18379841098",
			"id_card": "3607812001"
		},
		"roles": [
			"admin"
		],
		"username": "Peter"
	},
	"msg": "ok"
}
```

| 序号 |   参数    |     类型      |         规则          |    简介    |
| :--: | :-------: | :-----------: | :-------------------: | :--------: |
|1 | username | varchar(33) | 字符构成  | 用户名  |

## 添加员工信息

- URL:v1/user/employee

- Method: POST

- Request Body

```json
{
	"id" : 333,
	"real_name" : "zhanjianpeng",
	"nick_name" : "xiaozhan",
	"english_name" : "zjp",
	"sex" : "male",
	"age" : 19,
	"address" : "jian",
	"mobile_phone" : "188423342323",
	"id_card" : "3607812001",
	"username" :"zjp",
	"password" : "3607812001lyp",
	"role" : "employee",
	"department" : "adminer group"
}
```

- Response Body

```json
{
	"code": 5,
	"data": {
		"NewEmployee": {
			"id": 333,
			"real_name": "zhanjianpeng",
			"nick_name": "xiaozhan",
			"english_name": "zjp",
			"sex": "male",
			"age": 19,
			"address": "jian",
			"mobile_phone": "188423342323",
			"id_card": "3607812001",
			"username": "zjp",
			"password": "$2a$10$Dw92N36Zps.i/pmnJlUkkuY7fNKbSTCykW3ZNTKAM/C6OeeegPl5W",
			"role": "employee",
			"department": "adminer group"
		}
	},
	"msg": "创建用户成功"
}
```
| 序号 |   参数    |     类型      |         规则          |    简介    |
| :--: | :-------: | :-----------: | :-------------------: | :--------: |
|  1   | id  |     unit      |       数字构成       |     -      |
| 2 | real_name | varchar(33) | 字符构成  | 真实姓名  |
| 3 | nick_name | varchar(33) | 字符构成  | 昵称  |
| 4 | english_name | varchar(33) | 字符构成  | 英文名  |
| 5 | sex | varchar(33) | 字符构成  | 性别  |
| 6 | address | varchar(33) | 字符构成  | 地址  |
| 7 | mobile_phone | varchar(33) | 字符构成  | 手机号  |
| 8 | id_card | varchar(33) | 字符构成  | 省份证号  |
|9 | username | varchar(33) | 字符构成  | 用户名  |

## 更新员工信息

- URL: v1/user/employee

- Method: PUT

- Request Body

```json
{
  "id" : 333,
  "real_name" : "zhanjianpeng",
  "nick_name" : "xiaozhan",
  "english_name" : "zjp",
  "sex" : "male",
  "age" : 19,
  "address" : "jian",
  "mobile_phone" : "188423342323",
  "id_card" : "3607812001"
}
"username" : "aaa",
```

- Response Body

```json
{
  "code": 5,
  "msg": "ok",
  "employee": {
    "id" : 333,
    "real_name" : "zhanjianpeng",
    "nick_name" : "xiaozhan",
    "english_name" : "zjp",
    "sex" : "male",
    "age" : 19,
    "address" : "jian",
    "mobile_phone" : "188423342323",
    "id_card" : "3607812001"
  }
}
```
| 序号 |   参数    |     类型      |         规则          |    简介    |
| :--: | :-------: | :-----------: | :-------------------: | :--------: |
|  1   | id  |     unit      |       数字构成       |     -      |
| 2 | real_name | varchar(33) | 字符构成  | 真实姓名  |
| 3 | nick_name | varchar(33) | 字符构成  | 昵称  |
| 4 | english_name | varchar(33) | 字符构成  | 英文名  |
| 5 | sex | varchar(33) | 字符构成  | 性别  |
| 6 | address | varchar(33) | 字符构成  | 地址  |
| 7 | mobile_phone | varchar(33) | 字符构成  | 手机号  |
| 8 | id_card | varchar(33) | 字符构成  | 省份证号  |
|9 | username | varchar(33) | 字符构成  | 用户名  |

## 删除员工信息

- URL: v1/user/employee

- Method: DELETE

- Request Body

```json
{
	"username" : "lyp" 
}
```

- Response Body

```json
{
	"code": 2000,
	"data": {
		"username": "zjp"
	},
	"msg": "ok"
}
```
| 序号 |   参数    |     类型      |         规则          |    简介    |
| :--: | :-------: | :-----------: | :-------------------: | :--------: |
| 1 | username | varchar(33) | 字符构成  | 用户名  |