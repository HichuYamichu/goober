use std::collections::HashMap;
use rocket_contrib::templates::Template;

use rocket::get;

use crate::guards::user::User;
use rocket::response::Redirect;

#[get("/")]
pub fn user_index(_user: User) -> Template {
   let mut context = HashMap::new();
  context.insert("0", 0);
  Template::render("index", context)
}

#[get("/", rank = 2)]
pub fn index() -> Redirect {
  Redirect::to("/login")
}
