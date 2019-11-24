use rocket::catch;
use rocket::response::Redirect;

#[catch(404)]
pub fn index() -> Redirect {
  Redirect::to("/")
}
