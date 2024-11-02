#[derive(Debug)]
struct Adult {
	age: i32,
	name: String,
}

impl Adult {
	fn new(age: i32, name: &str) -> Result<Self, String> {
		if age >= 21 {
			return Ok(Adult{age: age, name: name.to_string()});
		}
		return Err("Age must be more or eq 21".to_owned());
	}
}

fn main() {
	let a = Adult::new(20, "X1");
	let b = Adult::new(22, "X2");

	match a {
		Ok(ad1) => println!("{:?}", ad1),
		Err(msg1) => println!("{:?}", msg1),
	}
	match b {
		Ok(ad2) => println!("{:?}", ad2),
		Err(msg2) => println!("{:?}", msg2),
	}
}

