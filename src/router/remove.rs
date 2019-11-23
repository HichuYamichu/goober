use std::io;
use std::fs;
use std::path::Path;

use rocket::delete;

use crate::guards::api_key::ApiKey;

#[delete("/remove/<resorce>")]
pub fn index(resorce: String, _key: ApiKey) -> io::Result<()> {
  let filename = format!("upload/{file}", file = resorce);
  if Path::new(&filename).exists() {
    fs::remove_file(filename)?;
  }
  Ok(())
}
