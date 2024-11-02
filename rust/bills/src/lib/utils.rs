use std::io;


pub fn get_input() -> Result<String, String> {
	let mut s = String::new();
	match io::stdin().read_line(&mut s) {
		Ok(_) => return Ok(s.trim().to_owned()),
		Err(_error) => return Err(
		"something went wrong when reading input".to_owned())
	}
}

pub fn greet() {
	println!("Hello world");
}

