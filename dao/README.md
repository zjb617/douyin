需要设计的表
- User
  - id
  - name
  - password
  - follow_count
  - follower_count

```sql
create table user (
    id bigint unsigned auto_increment,
    username varchar(32) not null,
    password varchar(32) not null,
    follow_count bigint default 0,
    follower_count bigint default 0,
    primary key (id)
)engine=innodb default charset=utf8;
```

- Video
  - id
  - uid
  - play_url
  - cover_url
  - favorite_count
  - comment_count
  - title
  - create_time

```sql
create table video (
    id bigint unsigned auto_increment,
    uid bigint unsigned not null,
    play_url varchar(100) not null,
    cover_url varchar(100) not null,
    favorite_count bigint default 0,
    comment_count bigint default 0,
    title varchar(20) not null,
    create_time bigint not null,
    primary key (id),
    foreign key (uid) references user(id)
)engine=innodb default charset=utf8;
```

- Comment
  - id
  - uid
  - video_id
  - content
  - create_date

```sql
create table comment (
    id bigint unsigned auto_increment,
    uid bigint unsigned not null,
    video_id bigint unsigned not null,
    content varchar(1000) not null,
    create_date bigint not null,
    primary key (id),
    foreign key (uid) references user(id),
    foreign key (video_id) references video(id)
)engine=innodb default charset=utf8;
```

- Message
  - id
  - uid
  - to_uid
  - content
  - create_time

```sql
create table message (
    id bigint unsigned auto_increment,
    uid bigint unsigned not null,
    to_uid bigint unsigned not null,
    content varchar(1000) not null,
    create_time bigint not null,
    primary key (id),
    foreign key (uid) references user(id),
    foreign key (to_uid) references user(id)
)engine=innodb default charset=utf8;
```

关系表
- UserFollow（多对多）
  - id
  - uid
  - follow_uid

```sql
create table userfollow (
    id bigint unsigned auto_increment,
    uid bigint unsigned not null,
    follow_uid bigint unsigned not null,
    primary key (id),
    foreign key (uid) references user(id),
    foreign key (follow_uid) references user(id)
)engine=innodb default charset=utf8;
```

- VideoFavorite（多对多）
  - id
  - uid
  - video_id

```sql
create table videofavorite (
    id bigint unsigned auto_increment,
    uid bigint unsigned not null,
    video_id bigint unsigned not null,
    primary key (id),
    foreign key (uid) references user(id),
    foreign key (video_id) references video(id)
)engine=innodb default charset=utf8;
```
