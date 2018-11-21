package main

import (
"fmt"
"dbop"
"html/template"
"net/http"
"strconv"
)

func AppEdit() interface{} {
	applist, _ := dbop.GetAllApps(0xffff) // Manager
	str := ""
	for _, app := range applist {
		online := "上线"
		if app.Online == 0 {
			online = "下线"
		}
		str += fmt.Sprintf("<input type=\"text\" name=\"app%d\" value=\"%s\" size=6 readonly /> <a target=\"_blank\" class=\"linkto\" href=\"%s\">链接</a>  状态:%s  <input type=\"button\" value=\"编辑\" onclick=\"oneditapp('%s',%d)\" /> <input type=\"button\" value=\"删除\" onclick=\"ondelapp('%s',%d)\" /><br>\n", app.ID, app.Name, app.Url, online, app.Name, app.ID, app.Name, app.ID)
		//str+=fmt.Sprintf("ID:%-3d名称:<input type=\"text\" name=\"app%d\" value=\"%s\" size=6 readonly />  <a target=\"_blank\" class=\"linkto\" href=\"%s\">链接</a> 状态:<select><option value=\"online\" %s>上线</option><option value=\"offline\" %s>下线</option></select><br>\n",app.ID,app.ID,app.Name,app.Url,online,offline)
	}
	return template.HTML(str)
}

func appmgr(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpfunc := make(template.FuncMap)
		tpfunc["AppEdit"] = AppEdit
		t := template.New("appmgr.tpl")
		t = t.Funcs(tpfunc)
		t, err := t.ParseFiles("appmgr.tpl")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		id := r.Form["editid"][0]
		fmt.Fprintf(w, "Edit id %s\n", id)
	}
}

func addapp(w http.ResponseWriter, r *http.Request) {
fmt.Println("addapp: ",r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("addapp.tpl")
		args := make(map[string]string)
		args["APPNAME"] =""
		args["APPURL"] ="" 
		args["APPICON"] ="" 
		args["SELON"] = "selected"
		t.Execute(w, args)
	} else {
		r.ParseForm()
		info:=&dbop.AppInfo {
				Name:r.Form["appname"][0],
				Url:r.Form["appurl"][0],
				Icon:r.Form["appicon"][0],
				Online:1 }
		if err := info.Insert(); err == nil {
			t, _ := template.ParseFiles("addapp.tpl")
			t.Execute(w,nil)
		}

	}
}

func editapp(w http.ResponseWriter, r *http.Request) {
fmt.Println("editapp: ",r.Method)
	if r.Method == "GET" {
		r.ParseForm()
		appstr := r.Form.Get("appid")
		t, _ := template.ParseFiles("editapp.tpl")
		args := make(map[string]string)
		if appstr != "" {
			if appid, err := strconv.ParseInt(appstr, 10, 64); err == nil {
				if appInfo, err := dbop.FindApp(appid); err == nil {
					args["APPNAME"] = appInfo.Name
					args["APPURL"] = appInfo.Url
					args["APPICON"] = appInfo.Icon
					if appInfo.Online == 0 {
						args["SELON"] = ""
						args["SELOFF"] = "selected"
					} else {
						args["SELON"] = "selected"
						args["SELOFF"] = ""
					}
				}
			}
		}
		t.Execute(w, args)
	} else {
		r.ParseForm()
		idstr := r.Form["editid"][0]
		if editid, err := strconv.ParseInt(idstr, 10, 64); err == nil {
			if info, err := dbop.FindApp(editid); err == nil {
				info.Name = r.Form["appname"][0]
				info.Url = r.Form["appurl"][0]
				info.Icon = r.Form["appicon"][0]
				if r.Form.Get("status") == "online" {
					info.Online = 1
				} else {
					info.Online = 0
				}
				//info.SaveInfo()
				if err := info.SaveInfo(); err == nil {
					t, _ := template.ParseFiles("editok.tpl")
					t.Execute(w,nil)
				//	http.Redirect(w, r, "/appmgr", http.StatusFound)
				}

			}
		}
	}
}

func delapp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delapp: ", r.URL.Path) // update database, to be done
	r.ParseForm()
	appstr := r.Form.Get("appid")
	if appstr != "" {
		if appid, err := strconv.ParseInt(appstr, 10, 64); err == nil {
			dbop.DelApp(appid)
		}
	}
	fmt.Println("redirected")
	http.Redirect(w, r, "/appmgr", http.StatusFound)
}

