
trait Body {
}

trait Color {

}


struct Vehicle<T: Body + Color> {
	item: T
}

fn main() {

}

