use rocket::get;

#[get("/")]
pub fn index() -> &'static str {
  "Nothing to see here."
}
