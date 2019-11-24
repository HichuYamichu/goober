use rocket::http::ContentType;
use rocket::{post, Data};
use std::env;
use std::fs::File;
use std::io;
use std::io::prelude::*;
use std::path::Path;

use rand::{self, Rng};

use rocket_multipart_form_data::{
  MultipartFormData, MultipartFormDataError, MultipartFormDataField, MultipartFormDataOptions,
  RawField,
};

use rocket::response::Redirect;

use crate::guards::api_key::ApiKey;
use crate::guards::user::User;

const ID_LENGTH: usize = 8;
const BASE62: &[u8] = b"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";

#[post("/upload", data = "<data>")]
pub fn index(cont_type: &ContentType, data: Data, user: User) -> Result<Redirect, &'static str> {
  let mut options = MultipartFormDataOptions::new();
  options
    .allowed_fields
    .push(MultipartFormDataField::raw("file").size_limit(2048 * 2048 * 64 * 1000));

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
          _ => Err("Cannnot save the file."),
        }
      }
      RawField::Multiple(_) => unreachable!(),
    },
    None => Err("Please input a file."),
  }
}

#[post("/upload", data = "<paste>", rank = 2)]
pub fn sharex(cont_type: &ContentType, paste: Data, _key: ApiKey) -> Result<String, io::Error> {
  let id = generte_id();
  let ext = cont_type.media_type().sub();
  let filename = format!("upload/{id}.{ext}", id = id, ext = ext);
  let url = format!(
    "{host}/files/{id}.{ext}\n",
    host = env::var("HOST").unwrap(),
    id = id,
    ext = ext
  );

  paste.stream_to_file(Path::new(&filename))?;
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
