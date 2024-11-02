use std::fs::read_to_string;
use std::path::Path;
use std::fs::File;
use std::io::prelude::*;
use std::process::exit;
use std::fmt;


static DATABASE: &str = "./db.txt";

fn read_lines(filename: &str) -> Vec<String> {
    read_to_string(filename) 
        .unwrap()  // panic on possible file-reading errors
        .lines()  // split the string into an iterator of string slices
        .map(String::from)  // make each slice into a string
        .collect()  // gather them together into a vector
}

fn save_db(f: &mut File, lines: &Vec<String>) {
	let s = fmt::format(format_args!("{}\n", lines.join("\n")));
	match f.write_all(s.as_bytes()) {
		Err(why) => panic!("couldn't write to {}", why),
		Ok(_) => (),
	}
}


fn main() {
	let mut args: Vec<String> = std::env::args().collect(); 
	args.remove(0);
	if args.len() < 1 {
		eprintln!("Usage: bills <command> [arguments]");
		exit(1);
	}	
	
	let mut lines = vec![];
	let db = Path::new(DATABASE); 
	let display = db.display();
	match db.exists() {
		true => {
			lines = read_lines(DATABASE);								
		},
		false => ()
	}

	match args[0].as_str() {
		"add" => {
			if args.len() != 3 {
				eprintln!("Usage: bills add <item> <price>");
				exit(1);
			}
			let (item, price) = (args[1].clone(), args[2].clone());
			let new_item: String = fmt::format(format_args!("{}::{}", item, price));
			lines.push(new_item);
			let mut file = match File::create(&db) {
				Err(why) => panic!("couldn't create {}: {}", display, why),
				Ok(file) => file,
			};
			save_db(&mut file, &lines);
		},
		"show" => {
			for l in &lines {
				println!("{}", l);
			}
		},
		"remove" => {
			if args.len() != 2 {
				eprintln!("Usage: bills remove <item>");
				exit(1);
			}
			match lines.iter().position(|x| x.split("::").collect::<Vec<&str>>()[0] == args[1]) {
				Some(idx) =>  { lines.remove(idx); },
				None => { println!("item {:? } is not found", args[1]); }, 
			}
			let mut file = match File::create(&db) {
				Err(why) => panic!("couldn't create {}: {}", display, why),
				Ok(file) => file,
			};
			save_db(&mut file, &lines);
		},
		"edit" => {
			if args.len() != 3 {
				eprintln!("Usage: bills edit <item> <new_price>");
				exit(1);
			}
			match lines.iter().position(|x| x.split("::").collect::<Vec<&str>>()[0] == args[1]) {
				Some(idx) =>  { 
					lines.remove(idx); 
					let new_item: String = format!("{}::{}", args[1], args[2]);
					lines.insert(idx, new_item);
				  },
				None => { println!("item {:? } is not found", args[1]); }, 
			}
			let mut file = match File::create(&db) {
				Err(why) => panic!("couldn't create {}: {}", display, why),
				Ok(file) => file,
			};
			save_db(&mut file, &lines);

			
		},
		cmd => { eprintln!("unknown command {:?}", cmd); exit(1) },
	}
	
}
