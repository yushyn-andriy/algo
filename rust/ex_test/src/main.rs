use rs_color::*;

mod greet {
	pub fn hello() {
		println!("Hello");
	}
}

mod math {
	pub fn add(a: i32, b:i32) -> i32 {
		a + b
	}
	pub fn substract(a: i32, b:i32) -> i32 {
		a - b
	}
}



fn main() {
	use greet::hello;
	use math::add;
	hello();
	add(1, 2);
	println!("{}", Color::Green("Hello world").make());
	println!("{}", Color::Cyan("Hello world").make());
	println!("{}", Color::Red("Hello world").make());
	println!("{}", Color::Yellow("Hello world").make());
	println!("{}", Color::Blue("Hello world").make());
	println!("{}", Color::Magenta("Hello world").make());
}

#[cfg(test)]
mod test {
	use crate::math::*;
	#[test]
	fn test_add() {
		assert_eq!(add(1, 2), 4, "1+2 should be eq 3");
	}
	#[test]
	fn test_substruct() {
		assert_eq!(substract(1, 2), 4, "1-2 should be eq 4");
	}
}
