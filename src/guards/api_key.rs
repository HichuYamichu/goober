use rocket::http::Status;
use rocket::request::{self, FromRequest, Request};
use rocket::{Outcome, State};

use crate::Config;

pub struct ApiKey(String);

#[derive(Debug)]
pub enum ApiKeyError {
  BadCount,
  Missing,
  Invalid,
}

impl<'a, 'r> FromRequest<'a, 'r> for ApiKey {
  type Error = ApiKeyError;

  fn from_request(request: &'a Request<'r>) -> request::Outcome<Self, Self::Error> {
    let config = request.guard::<State<Config>>().unwrap();
    let api_key = &config.api_key;
    let keys: Vec<_> = request.headers().get("x-api-key").collect();
    match keys.len() {
      0 => Outcome::Failure((Status::BadRequest, ApiKeyError::Missing)),
      1 if keys[0] == api_key => Outcome::Success(ApiKey(keys[0].to_string())),
      1 => Outcome::Failure((Status::BadRequest, ApiKeyError::Invalid)),
      _ => Outcome::Failure((Status::BadRequest, ApiKeyError::BadCount)),
    }
  }
}
