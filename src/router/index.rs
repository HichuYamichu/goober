use rocket::get;
use rocket::response::{Redirect, content};

use crate::guards::user::User;

#[get("/")]
pub fn authenticated(_user: User) -> content::Html<String> {
  let html = "<!DOCTYPE html>
  <html>
  
  <head>
     <meta charset=\"utf-8\" />
     <meta name=\"viewport\" content=\"width=device-width\" />
     <title>Upload</title>
     <link rel=\"stylesheet\" href=\"static/css/style.css\">
  </head>
  
  <body>
     <form class=\"uploadForm\" action=\"/upload\" method=\"POST\" enctype=\"multipart/form-data\">
        <input type=\"file\" name=\"file\">
        <button type=\"submit\">Upload</button>
     </form>
  </body>
  
  </html>".to_string();
  content::Html(html)
}

#[get("/", rank = 2)]
pub fn redirect() -> Redirect {
  Redirect::to("/login")
}
