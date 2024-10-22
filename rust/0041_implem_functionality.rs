
struct Employe {
	name: String,
	wage: f64,
}


impl Employe {
	fn show_wage(e: &Employe) {
		println!("{:?}", e.wage);
	}

	fn inc_wage(&mut self) {
		self.wage = self.wage + 100.0;
	}
}


fn main() {
	let mut empl = Employe{name:"Mike".to_string(), wage: 1129.0};
	Employe::show_wage(&empl);
	empl.inc_wage();
	Employe::show_wage(&empl);
	empl.inc_wage();
	Employe::show_wage(&empl);
}
