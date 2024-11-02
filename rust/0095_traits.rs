

trait Area {
	fn get_area(&self) -> f64;
}


struct Circle {
	r: f64,
}

struct Square {
	side: f64
}

impl Area for Circle {
	fn get_area(&self) -> f64 {
		3.1415*self.r*self.r
	}
}

impl Area for Square {
	fn get_area(&self) -> f64 {
		self.side*self.side
	}

}

fn print_square(obj: &impl Area) {
	println!("The area is {:?}", obj.get_area());
}


fn main() {
	let c = Circle{r: 5.0};
	let s = Square{side: 10.0};

	print_square(&c);
	print_square(&s);

}
