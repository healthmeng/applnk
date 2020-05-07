package main

import (
	"fmt"
	"html/template"
	"strings"
	"net/http"
	"os"
	"bufio"
	"log"
)


func ConvText(prefix,suffix,excel string)string{
	strs:=strings.Split(excel,"\n")
	ret:=""
	all:=len(strs)
	for l,line:=range strs{
		line=strings.TrimSpace(strings.Trim(line,"\r"))
		if line!=""{
			if strings.HasPrefix(line,prefix) && strings.HasSuffix(line,suffix){
				ret+=line
			}else{
				ret+=prefix+line+suffix
			}
		}
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
        excel:=r.Form["excel"][0]
        //excel:=strings.TrimSpace(r.Form["excel"][0])
		value:=ConvText(prefix,suffix,excel)
//		if err==nil {
	        t, _ := template.ParseFiles("qview.tpl")
	        args:= make(map[string]string)
			args["Prefix"]=prefix
			args["Suffix"]=suffix
			args["Excel"]=value
			t.Execute(w, args)
//			}
		log.Printf("From: %s\n",r.RemoteAddr)
	}
}

func ShowLog(w http.ResponseWriter, r *http.Request) {
	if r.Method=="GET"{
		r.ParseForm()
		if r.Form.Get("clear")=="1"{
			file,_:=os.Create("nohup.out")
			file.Close()
			
		}else{
			if file,err:=os.Open("nohup.out");err==nil{
				defer file.Close()
				rd:=bufio.NewReader(file)
				for{
					line,_,err:=rd.ReadLine()
					if err!=nil{
						break
					}
					fmt.Fprintln(w,string(line))
				}
			}
		}
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
	http.HandleFunc("/quickview",quickview)
	http.HandleFunc("/log",ShowLog)
	err := http.ListenAndServe(":8904", nil)
	if err != nil {
		fmt.Printf("Error:", err)
	}
}
