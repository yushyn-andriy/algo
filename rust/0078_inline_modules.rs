
mod greet {
	pub fn hello() {
		println!("Hello");
	}
}

mod math {
	pub fn add(a: i32, b:i32) -> i32 {
		a + b
	}
}



fn main() {
	use greet::hello;
	use math::add;
	hello();
	add(1, 2);

}

#[cfg(test)]
mod test {
	use math::add;
	#[test]
	fn test_add() {
		assert_eq!(add(1, 2), 4, "1+2 should be eq 4");
	}
}
