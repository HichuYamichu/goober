use std::env;
use std::io;
use std::path::Path;

use rocket::http::ContentType;
use rocket::{post, Data};

use rand::{self, Rng};

use crate::guards::api_key::ApiKey;

const ID_LENGTH: usize = 8;
const BASE62: &[u8] = b"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";

#[post("/upload", data = "<paste>")]
pub fn index(cont_type: &ContentType, paste: Data, _key: ApiKey) -> Result<String, io::Error> {
  let peek = paste.peek();
  let res = peek.iter().map(|&c| c as char).collect::<String>();
  println!("{:?}", res);
  let mut id = String::with_capacity(ID_LENGTH);
  let mut rng = rand::thread_rng();
  for _ in 0..ID_LENGTH {
    id.push(BASE62[rng.gen::<usize>() % 62] as char);
  }
  
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
