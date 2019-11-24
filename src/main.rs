#![feature(proc_macro_hygiene)]

use rocket::{catchers, routes};
use rocket_contrib::serve::StaticFiles;
use rocket_contrib::templates::Template;
use uploader::{router, Config};

fn main() {
  std::env::set_var("ROCKET_CLI_COLORS", "off");
  let config = Config::from_file("Config.toml");

  rocket::ignite()
    .manage(config)
    .register(catchers![router::not_found::index])
    .attach(Template::fairing())
    .mount("/static", StaticFiles::from("static"))
    .mount("/files", StaticFiles::from("upload"))
    .mount(
      "/",
      routes![
        router::index::authenticated,
        router::index::redirect,
        router::list::index,
        router::login::login,
        router::login::login_page,
        router::logout::index,
        router::remove::index,
        router::upload::user_auth,
        router::upload::token_auth,
      ],
    )
    .launch();
}
