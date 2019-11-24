#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use] extern crate serde_derive;

pub mod router;
pub mod guards;
