# SQL Tuning  
## インデックスを貼る。  
### WHEREに入っているカラムにインデックスを貼る
```
SELECT * FROM `users` WHERE `age` = 29
```
ageにインデックスをつける。
```
ALTER TABLE `users` ADD INDEX `age_idx`(`age`)
```
インデックスつけれているかを確認する。
```
EXPLAIN SELECT * FROM `users` WHERE `age` = 29
```
possible_keysにインデックス名が表示されること、rowsが著しく減少していることを確認する。  
インデックスを張り替えたい場合は、
```
ALTER TABLE `users` DROP INDEX `age_idx`, ADD INDEX `age_name_idx`(`age`,`name`)
```
### ソートの場合
```
SELECT * FROM `users` ORDER BY `register_date`
```
ソートされるカラムにインデックスをつける
```
ALTER TABLE `users` ADD INDEX `register_date_idx`(`register_date`);
```
### ユニークなインデックス
カラムの値がユニークな場合は、こちらを使用する。
```
SELECT * FROM `users` WHERE `email` = 'tanaka@example.com';
```
```
ALTER TABLE `users` ADD UNIQUE INDEX `email_idx`(`email`);
```
### 複合インデックスをつける
例えば、
```
SELECT * FROM `users` WHERE `age` = 29 AND `gender` = 'male'
```
2つのカラムにインデックスをつける。
```
ALTER TABLE `users` ADD INDEX `age_gender_idx`(`age`,`gender`);
```
降順に変える。
```
ALTER TABLE `users` DROP INDEX `age_gender_idx`, ADD INDEX `age_gender_idx`(`age`,`gender`,DESC);
```
### 複合インデックスの順番
```
SELECT * FROM `users` WHERE `age` = 29 AND `gender` = 'male'
```
- age,gender
- gender,age

で迷う。
レコード数少ないほうを先に選ぶ。ageが２００、genderが１００００ならage,gender
```
select age, count(*) from users group by age;
select gender, count(*) from users group by gender;
```

## N+1問題
### 見つけ方
スロークエリログを取得して実行数多いとこが怪しい。
### JOINを使う
テーブルを結合するので、結合に時間かかる。
けど、簡単。
```
type User struct{
    ID          int         `db:"id" json:"id"`
    AccountName string      `db:"account_name" json:"account_name"`
    Passhash    string      `db:"passhash" json:"passhash"`
    Authority   int         `db:"authority" json:"authority"`
    DelFlg      int         `db:"del_flg" json:"del_flg"`
    CreatedAt   time.Time   `db:"created_at" json:"created_at"`
}

type Post struct{
    ID          int         `db:"id" json:"id"`
    UserID
}
```