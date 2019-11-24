use rocket::get;
use rocket::http::{Cookie, Cookies};
use rocket::response::Redirect;

#[get("/logout")]
pub fn index(mut cookies: Cookies<'_>) -> Redirect {
  cookies.remove_private(Cookie::named("user_id"));
  Redirect::to("/login")
}
