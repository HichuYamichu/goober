use rocket::catch;

#[catch(404)]
pub fn index() -> &'static str {
  "404 Not found."
}
