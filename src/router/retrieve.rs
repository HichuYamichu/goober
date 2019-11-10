use std::fs::File;

use crate::paste_id::PasteID;
use rocket::get;
use rocket::response::content;

#[get("/<id>")]
pub fn index(id: PasteID<'_>) -> Option<content::Plain<File>> {
  let filename = format!("upload/{id}", id = id);
  File::open(&filename).map(|f| content::Plain(f)).ok()
}
