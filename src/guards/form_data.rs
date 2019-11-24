use rocket::http::Status;
use rocket::outcome::Outcome::*;
use rocket::request::{self, FromRequest, Request};
use rocket::response::status;

#[derive(Debug)]
pub struct FormData(pub String);

#[derive(Debug)]
enum FormDataError {
  ContentType,
  Boundary,
}

impl<'a, 'r> FromRequest<'a, 'r> for FormData {
  type Error = FormDataError;

  fn from_request(request: &'a Request<'r>) -> request::Outcome<FormData, Self::Error> {
    let content_type = request.content_type().unwrap();

    if !content_type.is_form_data() {
      return Failure((Status::BadRequest, FormDataError::ContentType));
    }

    let boundary = content_type
      .params()
      .find(|&(k, _)| k == "boundary")
      .ok_or_else(|| {
        Failure((Status::BadRequest, FormDataError::ContentType))
      });

    Ok(FormData(boundary))
  }
}
