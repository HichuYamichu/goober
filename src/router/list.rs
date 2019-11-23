use std::fs;

use rocket::get;

use crate::guards::api_key::ApiKey;
use rocket::response::content;

#[get("/list")]
pub fn index() -> content::Html<std::string::String> {
  let paths = fs::read_dir("upload").unwrap();
  let names = paths.map(|entry| {
    let entry = entry.unwrap();
    let entry_path = entry.path();
    let file_name = entry_path.file_name().unwrap();
    let file_name_as_str = file_name.to_str().unwrap();
    let file_name_as_string = String::from(file_name_as_str);
    format!("<a href=\"files/{path}\">{path}</a><br>", path = file_name_as_string)
  }).collect::<String>();

  return content::Html(names);
}
