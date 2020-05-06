package main

import (
	"fmt"
	"html/template"
	"strings"
	"net/http"
)


func ConvText(prefix,suffix,excel string)string{
	strs:=strings.Split(excel,"\r\n")
	ret:=""
	all:=len(strs)
	for l,line:=range strs{
		ret+=prefix+line+suffix
		if l<all-1{
			ret+="\n"

		}
	}
	return ret
}

func quickview(w http.ResponseWriter, r *http.Request) {
	if r.Method=="GET"{
        t, _ := template.ParseFiles("qview.tpl")
        args:= make(map[string]string)
		args["Prefix"]=""
		args["Suffix"]=""
		args["Excel"]=""
		t.Execute(w, args)

//		ret,err:=dbop.ViewTracks()
/*		ret,err:=dbop.SearchMatch(args["From"],args["To"],"","","desc","combined")
		if err==nil && ret!=nil{
			args["Total"]=fmt.Sprintf("%d",len(ret))
			for _,line:=range ret{
				fmt.Fprintf(w,"%s<br>",line)
			}
		}*/
	}else{
		r.ParseForm()
        prefix:= strings.TrimSpace(r.Form["prefix"][0])
        suffix:=strings.TrimSpace(r.Form["suffix"][0])
        excel:=strings.TrimSpace(r.Form["excel"][0])
		value:=ConvText(prefix,suffix,excel)
//		if err==nil {
	        t, _ := template.ParseFiles("qview.tpl")
	        args:= make(map[string]string)
			args["Prefix"]=prefix
			args["Suffix"]=suffix
			args["Excel"]=value
			t.Execute(w, args)
//			}
	}
}

func main() {
/*	http.HandleFunc("/",applnk)
	rootdir:=os.Getenv("PWD")+"/localapp"
	fhandle=http.StripPrefix("/localapp/",http.FileServer(http.Dir(rootdir)))
	http.HandleFunc("/localapp/",localapp)
	http.HandleFunc("/applinks", applnk)
	http.HandleFunc("/appmgr/", appmgr)
	http.HandleFunc("/download",download);*/
	http.HandleFunc("/",quickview)
//	http.HandleFunc("/quickview",quickview);
	err := http.ListenAndServe(":8904", nil)
	if err != nil {
		fmt.Printf("Error:", err)
	}
}
