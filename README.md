# API 文档
- [API 文档](#api---)
	+ [basic Response Body]
* [1.获取员工信息](#1---)
* [2.添加员工信息](#2----)
* [3.更新员工信息](#3-----)
* [4.删除员工信息](#4------)

## 1.获取员工信息
- URL: /api/employee

- Method: GET

- Response Body

```json
{
	http.StatusOK,gin.H{
		"Number" :    "6109119121",
		"Name" :       "Peterliang",
		"Profession" : "computer",
		"Task" :       "move trick",
	}
}
```

| 序号 |   参数    |     类型      |         规则          |    简介    |
| :--: | :-------: | :-----------: | :-------------------: | :--------: |
|  1   | Number  |     string      |       数字构成       |     -      |

## 添加员工信息

- URL: /api/employee

- Method: POST

- Response Body

```json
{
	http.StatusOK, gin.H{
		"Number" :    "6109119121",
		"Name" :       "Peterliang",
		"Profession" : "computer",
		"Task" :       "move trick",
	}
}
```

 | 序号 |  参数   |  类型  |  简介  | 必须 |
  | :--: | :-----: | :----: | :----: | :--: |
  |  1   |  Number  | String |  工号  |  Y   |
  |  2   |  Name  | String | 姓名 |  Y   |
  |  3   |   Profession    | String | 专业 |  Y   |
  |  4   | Task | String | 任务 |  Y   |

  ## 修改员工信息

- URL: /api/employee

- Method: PUT

- Response Body

```json
{
	http.StatusOK, gin.H{
		"Number" :    "6109119121",
		"Name" :       "Peterliang",
		"Profession" : "computer",
		"Task" :       "move trick",
	}
}
```

 | 序号 |  参数   |  类型  |  简介  | 必须 |
  | :--: | :-----: | :----: | :----: | :--: |
  |  1   |  Number  | String |  工号  |  Y   |
  |  2   |  Name  | String | 新的姓名 |  N   |
  |  3   |   Profession    | String | 新的专业 |  N   |
  |  4   | Task | String | 新的任务 |  N   |


  ## 删除员工信息

- URL: /api/employee

- Method: DELETE

- Response Body

```json
{
	http.StatusOK, gin.H{
		"Number" :    "6109119121",
		"Name" :       "Peterliang",
		"Profession" : "computer",
		"Task" :       "move trick",
	}
}
```

 | 序号 |  参数   |  类型  |  简介  | 必须 |
  | :--: | :-----: | :----: | :----: | :--: |
  |  1   |  Number  | String |  工号  |  Y   |
