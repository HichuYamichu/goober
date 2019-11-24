use rocket::outcome::IntoOutcome;
use rocket::request::{self, FromRequest, Request};

#[derive(Debug)]
pub struct User(pub usize);

impl<'a, 'r> FromRequest<'a, 'r> for User {
  type Error = std::convert::Infallible;

  fn from_request(request: &'a Request<'r>) -> request::Outcome<User, Self::Error> {
    request
      .cookies()
      .get_private("user_level")
      .and_then(|cookie| cookie.value().parse().ok())
      .map(|id| User(id))
      .or_forward(())
  }
}
