package dbop

import (
"database/sql"
_"strings"
"log"
"os"
"errors"
"fmt"
_"github.com/Go-SQL-Driver/MySQL"
_"time"
)

var db *sql.DB

type AppInfo struct{
	ID int64
	Name string
	Version string
	Vender string
	Url string
	Descr string
	Icon string
	Cost float64
	Sell float64
	Online int64
}

type StoreInfo struct{ // one table for each user
	ID int64
	Name string
	Level int64
	Descr string
}

type TrackInfo struct{
	StoreID int64
	AppID int64
	RegTime string
}

func init(){
	var err error
	db,err=sql.Open("mysql","applnk:2huoyige@tcp(123.206.55.31:3306)/applnk")
	if err!=nil{
		log.Println("Open database error:",err)
		os.Exit(1)
	}
}
/*

func (info* UserInfo)SaveInfo() error{
	dbinfo,_:=FindUser(info.Username)
	if dbinfo!=nil{
	}else{
		return errors.New("SaveInfo: user not found")
	}
	query:=fmt.Sprintf("update users set pwsha256='%s,descr='%s',face='%s',phone='%s' where uid=%d",info.Password,info.Descr,info.Face,info.Phone,info.UID)
	if _,err:=db.Exec(query);err!=nil{
		log.Println("Update db error:",err)
		return err
	}
	return nil
}

func FindStoreName(username string) ([]* StoreInfo,error){
	query:=fmt.Sprintf("select * from users where username='%s'",username)
	res,err:=db.Query(query)
	if err!=nil{
		log.Println("find user query error:",err)
		return nil,err
	}
	if res.Next(){
		info:=new(UserInfo)
		if err:=res.Scan(&info.UID,	&info.Username,
				&info.Password,&info.Descr,&info.Face,
				&info.Phone,&info.RegTime);err!=nil{
			log.Println("Query error:",err)
			return nil,err
		}
		return info,nil
	}
	return nil,nil
}


func ListUsers()([]*UserInfo,error){
	ret:=make([]*UserInfo,0,20)
	query:="select * from users"
	res,err:=db.Query(query)
	if err!=nil{
		log.Println("Query all users error:",err)
		return nil,err
	}
	for res.Next(){
		info:=new(UserInfo)
		if err:=res.Scan(&info.UID,	&info.Username,
				&info.Password,&info.Descr,&info.Face,
				&info.Phone,&info.RegTime);err!=nil{
			log.Println("Get object from db result  error:",err)
			return nil,err
		}else{
			ret=append(ret,info)
		}
	}
	return ret,nil
}*/

func FindStoreID(id int64) (* StoreInfo,error){
	query:=fmt.Sprintf("select * from stores where id='%d'",id)
	res,err:=db.Query(query)
	if err!=nil{
		log.Println("find store query error:",err)
		return nil,err
	}
	if res.Next(){
		info:=new(StoreInfo)
		if err:=res.Scan(&info.ID,	&info.Name,
				&info.Level, &info.Descr);err!=nil{
			log.Println("Query error:",err)
			return nil,err
		}
		return info,nil
	}
	return nil,nil
}

func FindApp(id int64) (* AppInfo,error){
	query:=fmt.Sprintf("select * from apps where id='%d'",id)
	res,err:=db.Query(query)
	if err!=nil{
		log.Println("find store query error:",err)
		return nil,err
	}
	if res.Next(){
		info:=new(AppInfo)
		if err:=res.Scan(&info.ID,	&info.Name,
				&info.Version, &info.Vender, &info.Url, &info.Descr,
				&info.Icon,&info.Cost,&info.Sell, &info.Online);err!=nil{
			log.Println("Query error:",err)
			return nil,err
		}
		return info,nil
	}
	return nil,nil
}

func (info* TrackInfo)RegisterVisit() error{
/*    tm:=time.Now().Local()
    info.RegTime=fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())*/
    query:=fmt.Sprintf("insert into tracks (storeid, appid) values (%d,%d)",info.StoreID,info.AppID)
    if _,err:=db.Exec(query);err!=nil{
        log.Println("Insert db error:",err)
        return err
    }
	return nil
}

func ViewTracks() ([]string,error){
	var vtime, name, app string
	ret:=make ([]string, 0, 100)
	query:="select tracks.visit,stores.name,apps.name from tracks,stores,apps where tracks.storeid=stores.id and tracks.appid=apps.id  order by tracks.visit";
	res,err:=db.Query(query)
	if err!=nil{
		log.Println("Query quick view of visit tracks error",err)
		return nil,err
	}
	for res.Next(){
		if err:=res.Scan(&vtime,&name,&app);err!=nil{
			log.Println ("Get object from result error:",err)
			return nil,err
		}else{
			ret=append(ret,fmt.Sprintf("%s   %-20s  %s\n",vtime,name,app))
		}
	}
	return ret,nil
}

func GetAllApps(storeid int64)([]*AppInfo,error){
	if storeid<1000{
		return nil,errors.New("Invalid storeid")
	}
	ret:=make([]*AppInfo,0,50)
	query:="select * from apps"
	res,err:=db.Query(query)
	if err!=nil{
		log.Println("Query all apps error:",err)
		return nil,err
	}
	for res.Next(){
		info:=new(AppInfo)
		if err:=res.Scan(&info.ID,	&info.Name,
				&info.Version, &info.Vender,&info.Url,
				 &info.Descr,&info.Icon,
				&info.Cost,&info.Sell, &info.Online);err!=nil{
			log.Println("Get object from db result  error:",err)
			return nil,err
		}else{
			ret=append(ret,info)
		}
	}
	return ret,nil
}
/*
func AddUser(info *UserInfo) error{
//	return nil
// Add user info in db,add user msg table in db
	if find,_:=FindUser(info.Username);find!=nil{
		return errors.New("User already exists")
	}
	tm:=time.Now().Local()
	info.RegTime=fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
	query:=fmt.Sprintf("insert into users (username,pwsha256,descr,face,phone,regtime) values ('%s','%s','%s','%s','%s','%s')",info.Username,info.Password,info.Descr,info.Face,info.Phone,info.RegTime)
	if result,err:=db.Exec(query);err!=nil{
		log.Println("Insert db error:",err)
		return err
	}else{
		info.UID ,_= result.LastInsertId()
		query=fmt.Sprintf("create table `msg%d` (`msgid` int(11) not null AUTO_INCREMENT, `type` smallint(3) not null, `content` varchar(1024), `fromuid` int(11) not null, `arrived` tinyint(1) not null, `svrstamp` datetime, PRIMARY KEY(`msgid`)) default charset=utf8",info.UID)
		//query=fmt.Sprintf("create table `msg%d` (`msgid` int(11) not null AUTO_INCREMENT, `type` smallint(3) not null, `content` varchar(1024), `fromuid` int(11) not null, `arrived` tinyint(1) not null, `svrstamp` datetime, PRIMARY KEY(`msgid`)) default character set=utf8",info.UID)
		if _,err:=db.Exec(query);err!=nil{
			log.Println("Create msg table error:",err)
			return err
		}
		return nil
	}
}

func DelUser(name string, passwd string)error{
	info,err:=FindUser(name)
	if err!=nil{
		log.Println("Del user error:",err)
		return err
	}
	if info==nil{
		return errors.New("User not found")
	}
	if passwd!=info.Password{
		return errors.New("Username/Password is incorrect")
	}
	query:=fmt.Sprintf("delete from users where username='%s'",info.Username)
	if _,err:=db.Exec(query);err!=nil{
		log.Println("Delete user failed:",err)
		return err
	}
	query=fmt.Sprintf("drop table if exists msg%d",info.UID)
	db.Exec(query)
	return nil
}*/
