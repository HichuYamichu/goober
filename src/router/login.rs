use rocket::http::{Cookie, Cookies};
use rocket::request::Form;
use rocket::response::{content, Redirect};
use rocket::{get, post, FromForm, State};

use crate::Config;

#[derive(FromForm)]
pub struct Login {
  username: String,
  password: String,
}

#[post("/login", data = "<login>")]
pub fn login(mut cookies: Cookies<'_>, login: Form<Login>, config: State<Config>) -> Redirect {
  let super_user_login = &config.users.super_user_login;
  let super_user_pass = &config.users.super_user_pass;
  let regular_user_login = &config.users.regular_user_login;
  let regular_user_pass = &config.users.regular_user_pass;
  if &login.username == super_user_login && &login.password == super_user_pass {
    cookies.add_private(Cookie::new("user_level", 1.to_string()));
    Redirect::to("/")
  } else if &login.username == regular_user_login && &login.password == regular_user_pass {
    cookies.add_private(Cookie::new("user_level", 0.to_string()));
    Redirect::to("/")
  } else {
    Redirect::to("/login")
  }
}

#[get("/login")]
pub fn login_page() -> content::Html<String> {
  let html = "<!DOCTYPE html>
  <html>
  
  <head>
    <meta charset=\"utf-8\" />
    <meta name=\"viewport\" content=\"width=device-width\" />
    <title>Upload</title>
    <link rel=\"stylesheet\" href=\"static/css/style.css\">
  </head>
  
  <body>
    <form class=\"loginForm\" action=\"/login\" method=\"post\">
      <input type=\"text\" name=\"username\" id=\"username\" />
      <br>
      <input type=\"password\" name=\"password\" id=\"password\" />
      <p><input type=\"submit\" value=\"login\"></p>
    </form>
  
  </body>
  
  </html>"
    .to_string();
  content::Html(html)
}
