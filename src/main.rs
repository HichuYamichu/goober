#![feature(proc_macro_hygiene)]

use rocket::{catchers, routes};
use rocket_contrib::serve::StaticFiles;
use uploader::router;
use rocket_contrib::templates::Template;

fn main() {
  std::env::set_var("ROCKET_CLI_COLORS", "off");

  rocket::ignite()
    .register(catchers![router::not_found::index])
    .attach(Template::fairing())
    .mount("/static", StaticFiles::from("static"))
    .mount("/files", StaticFiles::from("upload"))
    .mount(
      "/",
      routes![
        router::index::index,
        router::index::user_index,
        router::upload::index,
        router::remove::index,
        router::list::index,
        router::login::login,
        router::login::login_page,
        router::logout::index,
      ],
    )
    .launch();
}
