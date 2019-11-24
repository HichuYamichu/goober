use std::fs;
use std::io;
use std::path::Path;

use rocket::delete;

use crate::guards::user::User;

#[delete("/remove/<resorce>")]
pub fn index(resorce: String, user_lvl: User) -> io::Result<()> {
  if user_lvl.0 == 1 {
    let filename = format!("upload/{file}", file = resorce);
    if Path::new(&filename).exists() {
      fs::remove_file(filename)?;
    }
  }
  Ok(())
}
