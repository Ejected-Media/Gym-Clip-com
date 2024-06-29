# Gym-Clip.com ~

package main


// .
  
import (
  "os"
  "log"
		
  "text/template"
  "net/http"
		
  "cloud.google.com/go"
    
    
)



type htmlPageData struct {
    pageTitle string
    pagePath string
    pageList []pageNav
    
}


type pageNav struct {
    pageTitle string
    pageLink string
}






func app_welcome_center_page() {
    
    
}



// . appHandler
func appHandler(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/app" {
    	http.NotFound(w, r)
    	return
    }
    
// ,

  pageTitle := "www.Gym-Clip.com/ - Website App"
  pagePath := r.URL.Path
  
  pageType := ".."
  
  
  pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
      pageList: []pageNav {
          { pageTitle: "one", pageLink: "one"},
          { pageTitle: "two", pageLink: "two"},
          { pageTitle: "three", pageLink: "three"},
      },
  	
  }  //. .  pageData
  
  
    if pagePath == "/clients" {
      pageTitle = "Clients Page"
      pageList = pageList
  }
  
    if pagePath == "/workouts" {
      pageTitle = "Workouts Page"
      pageList = pageList
  }
  
    if pagePath == "/videos" {
      pageTitle = "Videos Page"
      pageList = pageList
  }
  
    if pagePath == "/gym-page" {
      pageTitle = "Gym Page"
      pageList = pageList
  }
  
  
  
    if pagePath == "/splash" {
      pageTitle = "Splash Page"
      pageList = pageList
  }
  
    if pagePath == "/discord" {
      pageTitle = "Discord App Page"
      pageList = pageList
  }
  
    if pagePath == "/facebook" {
      pageTitle = "Facebook App Page"
      pageList = pageList
  }
  
  
  
    if pagePath == "/flutter" {
      pageTitle = "Flutter Page"
      pageList = pageList
  }
  
    if pagePath == "/desktop" {
      pageTitle = "Desktop App Page"
      pageList = pageList
  }
  
    if pagePath == "/html-app" {
      pageTitle = "HTML-App Page"
      pageList = pageList
  }
  
  
  pageFilePath := template.Must(template.ParseFiles("layout_main_page.html"))
  pageFilePath.Execute(w, pageData)
  
}  //  .  appHandler


//  .  html url routes 
//  .  as well as terminal cli logs

func main() {

  appName := "www.Gym-Clip.com/ - Website App"


  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/account", indexHandler)
  http.HandleFunc("/profile", indexHandler)
  
  http.HandleFunc("/clients", indexHandler)
  http.HandleFunc("/workouts", indexHandler)
  http.HandleFunc("/videos", indexHandler)
  http.HandleFunc("/gym-page", indexHandler)
  
  http.HandleFunc("/splash", indexHandler)
  http.HandleFunc("/discord", indexHandler)
  http.HandleFunc("/facebook", indexHandler)
  
  http.HandleFunc("/flutter", indexHandler)
  http.HandleFunc("/desktop", indexHandler)
  http.HandleFunc("/html-app", indexHandler)



// -- -
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
    log.Printf("Loading _webapp with default port")
  }
  
  
  log.Printf("_webapp is active and Listening on port %s", port)

  log.Printf("// -- - %s", appName)
  log.Printf("_webapp now loaded and running at http://localhost:%s", port)

// -- - 
  if err := http.ListenAndServe(":"+port, nil); err != nil {
    log.Fatal("Error Starting the HTTP Server :", err)
    return
  }


}
