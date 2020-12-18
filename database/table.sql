DROP TABLE IF EXISTS employee;
CREATE TABLE employee(
	id INT NOT NULL AUTO_INCREMENT,
	real_name VARCHAR(11) NOT NULL,
	nick_name VARCHAR(33) NOT NULL,
	english_name VARCHAR(33) NOT NULL,
	sex VARCHAR(33) NOT NULL,
	age int NOT NULL,
	address VARCHAR(33) NOT NULL,
	mobile_phone VARCHAR(33) NOT NULL,
	id_card VARCHAR(33) NOT NULL,
	PRIMARY KEY(id)
);
insert into employee values (666,'lyp','Peter','Peterliang','male',18,'nanchang','18379841098','3607812001');

DROP TABLE IF EXISTS user;
CREATE TABLE user(
    id INT NOT NULL primary key AUTO_INCREMENT,
	username VARCHAR(33) NOT NULL,
	password_hash VARCHAR(1000) NOT NULL,
	employee_id int NOT NULL,
	CONSTRAINT fk_user_employee_id FOREIGN KEY(employee_id) REFERENCES employee(id)
);
insert into user values (1,'Peter','$2a$04$fG1Tq4MQ9KndXX4EB6k8W.6hUIP.q9nmtYLTvty3dvlPLz5fG0Amq',666);

DROP TABLE IF EXISTS department;
CREATE TABLE department(
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
	department_name VARCHAR(33) NOT NULL
);
insert into department values (1,'adminer group');
insert into department values (2,'developer group');
insert into department values (3,'management group');
insert into department values (4,'product group');

DROP TABLE IF EXISTS role;
CREATE TABLE role(
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
	role_name VARCHAR(33) NOT NULL
);
insert into role values (1,'admin');
insert into role values (2,'boss');
insert into role values (3,'manage');
insert into role values (4,'employee');

DROP TABLE IF EXISTS user_role;
CREATE TABLE user_role(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	user_id int NOT NULL,
	role_id int not null,
	constraint fk_user_role_id foreign key(user_id) references user(id),
	constraint fk_role_user_id foreign key(role_id) references role(id)
);
insert into user_role values (1,1,1);

DROP TABLE IF EXISTS user_department;
CREATE TABLE user_department(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	user_id int NOT NULL,
	department_id int not null,
	constraint fk_user_department_id foreign key(user_id) references user(id),
	constraint fk_department_user_id foreign key(department_id) references department(id)
);
insert into user_department values(1,1,1);