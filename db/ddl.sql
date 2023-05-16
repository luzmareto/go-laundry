create table if not exists customer (id varchar(100) primary key, name varchar(100), phone_number varchar(15) unique, is_deleted bool default false);


create table if not exists uom (id varchar(100) primary key, name varchar(30) not null, is_deleted bool default false);


create table if not exists product (id varchar(100) primary key, name varchar(50) not null, price bigint, uom_id varchar(100), is_deleted bool default false,
                                    foreign key(uom_id) references uom(id));


create table if not exists employee (id varchar(100) primary key, name varchar(100), is_deleted bool default false);


create table if not exists bill (id varchar(100) primary key, bill_date date, finish_date date, employee_id varchar(100), customer_id varchar(100),
                                 foreign key(employee_id) references employee(id),
                                 foreign key(customer_id) references customer(id));


create table if not exists bill_detail (id varchar(100) primary key, bill_id varchar(100), product_id varchar(100), product_price bigint, qty int,
                                        foreign key(bill_id) references bill(id),
                                        foreign key(product_id) references product(id));

