use std::fs;

use rocket::get;

use rocket_contrib::templates::Template;

#[derive(Serialize)]
struct Context {
  names: Vec<String>,
}

#[get("/list")]
pub fn index() -> Template {
  let paths = fs::read_dir("upload").unwrap();
  let names = paths
    .map(|entry| {
      let entry = entry.unwrap();
      let entry_path = entry.path();
      let file_name = entry_path.file_name().unwrap();
      let file_name_as_str = file_name.to_str().unwrap();
      let file_name_as_string = String::from(file_name_as_str);
      file_name_as_string
    })
    .collect::<Vec<String>>();

  Template::render("list", &Context { names: names })
}
