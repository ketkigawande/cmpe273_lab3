package main 

import (
"fmt"
"net/http"
"encoding/json"
//"io/ioutil"
"github.com/julienschmidt/httprouter"
"strings"
)

type Response struct{
       
       Key string `json:"key"`
       Value string `json:"value"`
}

type Response_simple struct{
       Key string
       Value string 

} 

type Response1 struct{
  Keyvaluearray []Response `json:"response"`
}

var record0[20] Response_simple
var record0_counter int
var record1[20] Response_simple
var record1_counter int 
var record2[20] Response_simple
var record2_counter int  


func Get0(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
      
      var res Response
      key_to_search :=p.ByName("id")
      var id_to_insert int 

          for i:=0;i< len(record0);i++ {

                          if(strings.Compare(record0[i].Key,key_to_search)==0){
                                         id_to_insert=i;

                          }


          }         

      res.Key=record0[id_to_insert].Key
      res.Value=record0[id_to_insert].Value
      
     //fmt.Println("Id:", id1)
      uj, _ := json.Marshal(res)
      rw.Header().Set("Content-Type", "application/json")
      rw.WriteHeader(200)
      fmt.Fprintf(rw, "\n\n\n\n%s", uj)
     //fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
   }


   func Get1(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
      
      var res Response
      key_to_search :=p.ByName("id")
      var id_to_insert int 

      for i:=0;i< len(record1);i++ {

                          if(strings.Compare(record1[i].Key,key_to_search)==0){
                                         id_to_insert=i;

                          }


          }         

      res.Key=record1[id_to_insert].Key
      res.Value=record1[id_to_insert].Value
      
     //fmt.Println("Id:", id1)
      uj, _ := json.Marshal(res)
      rw.Header().Set("Content-Type", "application/json")
      rw.WriteHeader(200)
      fmt.Fprintf(rw, "\n\n\n\n%s", uj)
     //fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
   }


   func Get2(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
      
      var res Response
      key_to_search :=p.ByName("id")
      var id_to_insert int 


      for i:=0;i< len(record2);i++ {

                          if(strings.Compare(record2[i].Key,key_to_search)==0){
                                         id_to_insert=i;

                          }


          }         

      res.Key=record2[id_to_insert].Key
      res.Value=record2[id_to_insert].Value



      
     //fmt.Println("Id:", id1)
      uj, _ := json.Marshal(res)
      rw.Header().Set("Content-Type", "application/json")
      rw.WriteHeader(200)
      fmt.Fprintf(rw, "\n\n\n\n%s", uj)
     //fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
   }


func Put0(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
      
      reply:=Response{}

      json.NewDecoder(r.Body).Decode(&reply) 

      key:=p.ByName("id")
      val:=p.ByName("value")
      fmt.Println("key---->value")
      fmt.Println(key)
      fmt.Println(val)

      //store it in the array
       record0[record0_counter].Key=key
       record0[record0_counter].Value=val
       record0_counter=record0_counter+1

      
        var res Response
        res.Key=key
        res.Value=val
      

     //fmt.Println("Id:", id1)
      uj, _ := json.Marshal(res)
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(200)
      fmt.Fprintf(w, "\n\n\n\n%s", uj)
     //fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
      
  }

  func Put1(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

       fmt.Println("her in put1 function")
      
      reply:=Response{}

      json.NewDecoder(r.Body).Decode(&reply) 

      key:=p.ByName("id")
      val:=p.ByName("value")
      fmt.Println("key---->value")
      fmt.Println(key)
      fmt.Println(val)

      //store it in the array
       record1[record1_counter].Key=key
       record1[record1_counter].Value=val
       record1_counter=record1_counter+1

      
        var res Response
        res.Key=key
        res.Value=val
      

     //fmt.Println("Id:", id1)
      uj, _ := json.Marshal(res)
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(200)
      fmt.Fprintf(w, "\n\n\n\n%s", uj)
     //fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name")))
      
  }

  func Put2(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
      
      reply:=Response{}

      json.NewDecoder(r.Body).Decode(&reply)       
      //id1:=p.ByName("id")
      key:=p.ByName("id")
      val:=p.ByName("value")

      //store it in the array
       record2[record2_counter].Key=key
       record2[record2_counter].Value=val

       record2_counter=record2_counter+1

      
      var res Response
      res.Key=key
      res.Value=val
       
      
      //fmt.Println("Id:", id1)
      uj, _ := json.Marshal(res)
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(200)
      fmt.Fprintf(w, "\n\n\n\n%s", uj)
     //fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
      
  }

  func Gett0(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
      
      var Res Response1
      var k int 
      fmt.Println("I am in get all function")
      var indStruct Response

      fmt.Println("Length");
      fmt.Println(len(record0))

                          for k=0;k<len(record0);k++ {


                                      if(strings.Compare(record0[k].Key,"")==0){
                                             break;
                                      }
                                      
                                  }

                  fmt.Println(k)

                 for j:=0;j<k;j++ {


      indStruct.Key=record0[j].Key
      indStruct.Value=record0[j].Value
      Res.Keyvaluearray = append(Res.Keyvaluearray,indStruct)
  }

     
      uj, _ := json.Marshal(Res)
      rw.Header().Set("Content-Type", "application/json")
      rw.WriteHeader(200)
      fmt.Fprintf(rw, "\n\n\n\n%s", uj)
     
   }


   func Gett1(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
      
      var Res Response1
      var k int 
      fmt.Println("I am in get all function")
      var indStruct Response

      fmt.Println("Length");
      fmt.Println(len(record1))

                          for k=0;k<len(record1);k++ {


                                      if(strings.Compare(record1[k].Key,"")==0){
                                             break;
                                      }
                                      
                                  }

                  fmt.Println(k)

                 for j:=0;j<k;j++ {


      indStruct.Key=record1[j].Key
      indStruct.Value=record1[j].Value
      Res.Keyvaluearray = append(Res.Keyvaluearray,indStruct)
  }

     
      uj, _ := json.Marshal(Res)
      rw.Header().Set("Content-Type", "application/json")
      rw.WriteHeader(200)
      fmt.Fprintf(rw, "\n\n\n\n%s", uj)
     
   }

   func Gett2(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
      
      var Res Response1
      var k int 
      fmt.Println("I am in get all function")
      var indStruct Response

      fmt.Println("Length");
      fmt.Println(len(record2))

                          for k=0;k<len(record2);k++ {


                                      if(strings.Compare(record2[k].Key,"")==0){
                                             break;
                                      }
                                      
                                  }

                  fmt.Println(k)

                 for j:=0;j<k;j++ {


      indStruct.Key=record2[j].Key
      indStruct.Value=record2[j].Value
      Res.Keyvaluearray = append(Res.Keyvaluearray,indStruct)
  }

     
      uj, _ := json.Marshal(Res)
      rw.Header().Set("Content-Type", "application/json")
      rw.WriteHeader(200)
      fmt.Fprintf(rw, "\n\n\n\n%s", uj)
     
   }




  func main() {
      
    r := httprouter.New()
    q := httprouter.New()
    s := httprouter.New()

    


     //r.POST("/keys",postt)
     r.GET("/keys/:id",Get0)
     q.GET("/keys/:id",Get1)
     s.GET("/keys/:id",Get2)
    
    r.GET("/keys",Gett0)
    q.GET("/keys",Gett1)
    s.GET("/keys",Gett2)
     
     r.PUT("/keys/:id/:value",Put0)
     q.PUT("/keys/:id/:value",Put1)
     s.PUT("/keys/:id/:value",Put2)

    go func(){
  server3 := http.Server{
    Addr:    "0.0.0.0:3000",
    Handler: r,
  } 
server3.ListenAndServe()
}()
  
go func(){
  server1 := http.Server{
    Addr:    "0.0.0.0:3001",
    Handler: q,
  }
  server1.ListenAndServe()
}()
    server2 := http.Server{
    Addr:    "0.0.0.0:3002",
    Handler: s,
  }
  server2.ListenAndServe()

     
 }
