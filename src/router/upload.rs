use std::fs::File;
use std::io;
use std::io::prelude::*;
use std::path::Path;

use rocket::http::ContentType;
use rocket::response::Redirect;
use rocket::{post, Data, State};

use rocket_multipart_form_data::{
  MultipartFormData, MultipartFormDataError, MultipartFormDataField, MultipartFormDataOptions,
  RawField,
};

use rand::{self, Rng};

use crate::guards::api_key::ApiKey;
use crate::guards::user::User;
use crate::Config;

const ID_LENGTH: usize = 8;
const BASE62: &[u8] = b"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";

#[post("/upload", data = "<data>")]
pub fn user_auth(
  cont_type: &ContentType,
  data: Data,
  user_lvl: User,
) -> Result<Redirect, &'static str> {
  let mut size_limit: u64 = 4 * 1000 * 1000 * 100;
  if user_lvl.0 == 1 {
    size_limit = size_limit * 10;
  }

  let mut options = MultipartFormDataOptions::new();
  options
    .allowed_fields
    .push(MultipartFormDataField::raw("file").size_limit(size_limit));

  let mut multipart_form_data = match MultipartFormData::parse(cont_type, data, options) {
    Ok(multipart_form_data) => multipart_form_data,
    Err(err) => match err {
      MultipartFormDataError::DataTooLargeError(_) => return Err("The file is too large."),
      _ => panic!("{:?}", err),
    },
  };

  let image = multipart_form_data.raw.remove("file");

  match image {
    Some(image) => match image {
      RawField::Single(raw) => {
        let file_name = raw.file_name.unwrap_or(generte_id());
        let file_path = format!("upload/{file_name}", file_name = file_name);
        let mut file = File::create(file_path).unwrap();

        match file.write_all(&raw.raw) {
          Ok(()) => {
            let url = format!("/files/{file_name}", file_name = file_name);
            Ok(Redirect::to(url))
          }
          _ => Err("Cannot save the file."),
        }
      }
      RawField::Multiple(_) => unreachable!(),
    },
    None => Err("Please input a file."),
  }
}

#[post("/upload", data = "<data>", rank = 2)]
pub fn token_auth(
  cont_type: &ContentType,
  data: Data,
  _key: ApiKey,
  config: State<Config>,
) -> Result<String, io::Error> {
  let id = generte_id();
  let ext = cont_type.media_type().sub();
  let filename = format!("upload/{id}.{ext}", id = id, ext = ext);
  let url = format!(
    "{host}/files/{id}.{ext}\n",
    host = &config.host,
    id = id,
    ext = ext
  );

  data.stream_to_file(Path::new(&filename))?;
  Ok(url)
}

fn generte_id() -> String {
  let mut id = String::with_capacity(ID_LENGTH);
  let mut rng = rand::thread_rng();
  for _ in 0..ID_LENGTH {
    id.push(BASE62[rng.gen::<usize>() % 62] as char);
  }
  id
}
