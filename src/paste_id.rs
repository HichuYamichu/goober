use std::borrow::Cow;
use std::fmt;

use rand::{self, Rng};
use rocket::http::RawStr;
use rocket::request::FromParam;

const BASE62: &[u8] = b"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";

pub struct PasteID<'a>(Cow<'a, str>);

impl<'a> PasteID<'a> {
  pub fn new(size: usize) -> PasteID<'static> {
    let mut id = String::with_capacity(size);
    let mut rng = rand::thread_rng();
    for _ in 0..size {
      id.push(BASE62[rng.gen::<usize>() % 62] as char);
    }

    PasteID(Cow::Owned(id))
  }
}

impl<'a> fmt::Display for PasteID<'a> {
  fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
    write!(f, "{}", self.0)
  }
}

fn valid_id(id: &str) -> bool {
  id.chars()
    .all(|c| (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9'))
}

impl<'a> FromParam<'a> for PasteID<'a> {
  type Error = &'a RawStr;

  fn from_param(param: &'a RawStr) -> Result<PasteID<'a>, &'a RawStr> {
    match valid_id(param) {
      true => Ok(PasteID(Cow::Borrowed(param))),
      false => Err(param),
    }
  }
}
