#![feature(proc_macro_hygiene, decl_macro)]

use std::fs;

#[macro_use]
extern crate serde_derive;

pub mod guards;
pub mod router;

#[derive(Deserialize)]
pub struct Config {
  host: String,
  api_key: String,
  users: Users,
}

#[derive(Deserialize)]
struct Users {
  super_user_login: String,
  super_user_pass: String,
  regular_user_login: String,
  regular_user_pass: String,
}

impl Config {
  pub fn from_file(path: &str) -> Config {
      let toml_str = fs::read_to_string(path).unwrap();
      toml::from_str(&toml_str).unwrap()
  }
}