#![feature(proc_macro_hygiene)]

use rocket::{catchers, routes};
use uploader::router;

fn main() {
  std::env::set_var("ROCKET_CLI_COLORS", "off");

  rocket::ignite()
    .register(catchers![router::not_found::index])
    .mount(
      "/",
      routes![
        router::index::index,
        router::upload::index,
        router::retrieve::index,
      ],
    )
    .launch();
}
