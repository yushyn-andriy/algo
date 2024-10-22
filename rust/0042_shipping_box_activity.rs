
#[derive(Debug)]
enum Color {
	White,
	Black,
}

struct ShippingBox {
	weight: f64,
	height: f64,
	width: f64,
	length: f64,
	color: Color,
}


impl ShippingBox {
	fn standard_box(color: Color) -> Self {
		ShippingBox{
			weight: 200.0,
			height: 3.0,
			width: 2.0,
			length: 5.0,
			color: color,
		}
	}
	
	fn print_box_info(&self) {
		println!("-------------------------");
		println!("Weight: {:?} Kg", self.weight);
		println!("H: {:?} M", self.height);
		println!("W: {:?} M", self.width);
		println!("L: {:?} M", self.length);
		println!("C: {:?}", self.color);
		println!("-------------------------");
	}

}



fn main() {
	let b = ShippingBox::standard_box(Color::Black);
	b.print_box_info();
}
