use rocket::http::Status;
use rocket::request::{self, FromRequest, Request};
use rocket::Outcome;
use std::env;

pub struct ApiKey(String);

fn is_valid(provided: &str) -> bool {
  let key = env::var("API_KEY").unwrap();
  provided == key
}

#[derive(Debug)]
pub enum ApiKeyError {
  BadCount,
  Missing,
  Invalid,
}

impl<'a, 'r> FromRequest<'a, 'r> for ApiKey {
  type Error = ApiKeyError;

  fn from_request(request: &'a Request<'r>) -> request::Outcome<Self, Self::Error> {
    let keys: Vec<_> = request.headers().get("x-api-key").collect();
    match keys.len() {
      0 => Outcome::Failure((Status::BadRequest, ApiKeyError::Missing)),
      1 if is_valid(keys[0]) => Outcome::Success(ApiKey(keys[0].to_string())),
      1 => Outcome::Failure((Status::BadRequest, ApiKeyError::Invalid)),
      _ => Outcome::Failure((Status::BadRequest, ApiKeyError::BadCount)),
    }
  }
}
