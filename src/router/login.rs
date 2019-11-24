use std::collections::HashMap;

use rocket::http::{Cookie, Cookies};
use rocket::request::Form;
use rocket::response::Redirect;
use rocket::{get, post, FromForm};
use rocket_contrib::templates::Template;

#[derive(FromForm)]
pub struct Login {
  username: String,
  password: String,
}

#[post("/login", data = "<login>")]
pub fn login(mut cookies: Cookies<'_>, login: Form<Login>) -> Redirect {
  if login.username == "1234" && login.password == "1234" {
    cookies.add_private(Cookie::new("user_id", 1.to_string()));
    Redirect::to("/")
  } else {
    Redirect::to("/login")
  }
}

#[get("/login", rank = 2)]
pub fn login_page() -> Template {
  let mut context = HashMap::new();
  context.insert("0", 0);
  Template::render("login", context)
}
