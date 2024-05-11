use std::collections::HashMap;

struct Rubik {
    values: HashMap<String, String>,
    colors: HashMap<String, String>,
}

impl Rubik {
    fn new() -> Self {
        let sides = ["B", "L", "U", "R", "D", "F"];
        let mut values = HashMap::new();
        let mut colors = HashMap::new();

        for side in sides.iter() {
            for i in 1..=9 {
                let index = format!("{}{}", side, i);
                values.insert(index.clone(), index.clone());
                colors.insert(index.clone(), match *side {
                    "U" => "yellow".to_string(),
                    "B" => "magenta".to_string(),
                    "R" => "green".to_string(),
                    "D" => "white".to_string(),
                    "L" => "blue".to_string(),
                    "F" => "red".to_string(),
                    _ => unreachable!(),
                });
            }
        }

        Rubik { values, colors }
    }

    // Función que realiza un movimiento en el cubo Rubik
    fn rubik_move(&mut self, pos_a: &[&str], pos_n: &[&str]) {
        let temp_values = self.values.clone();
        let temp_colors = self.colors.clone();

        for (from, to) in pos_a.iter().zip(pos_n.iter()) {
            // Corrección aquí: desreferenciar `to` que es `&&str` para obtener `&str`
            let value = temp_values.get(*to).unwrap_or(&"Unknown".to_string()).clone();
            let color = temp_colors.get(*to).unwrap_or(&"Unknown".to_string()).clone();
            self.values.insert((*from).to_string(), value);
            self.colors.insert((*from).to_string(), color);
        }
    }

    fn print_pos(&self, pos: &str) -> String {
        match self.values.get(pos) {
            Some(val) => val.clone(),
            None => "?".to_string(),
        }
    }

    fn format_cli(&self) {
        println!("         |--------|");
        println!(
            "         |{}-{}-{}|",
            self.print_pos("B1"),
            self.print_pos("B2"),
            self.print_pos("B3")
        );
        println!(
            "         |{}-{}-{}|",
            self.print_pos("B4"),
            self.print_pos("B5"),
            self.print_pos("B6")
        );
        println!(
            "         |{}-{}-{}|",
            self.print_pos("B7"),
            self.print_pos("B8"),
            self.print_pos("B9")
        );
        println!("|--------|--------|--------|--------|");
        println!(
            "|{}-{}-{}|{}-{}-{}|{}-{}-{}|{}-{}-{}|",
            self.print_pos("L1"),
            self.print_pos("L2"),
            self.print_pos("L3"),
            self.print_pos("U1"),
            self.print_pos("U2"),
            self.print_pos("U3"),
            self.print_pos("R1"),
            self.print_pos("R2"),
            self.print_pos("R3"),
            self.print_pos("D1"),
            self.print_pos("D2"),
            self.print_pos("D3")
        );
        println!(
            "|{}-{}-{}|{}-{}-{}|{}-{}-{}|{}-{}-{}|",
            self.print_pos("L4"),
            self.print_pos("L5"),
            self.print_pos("L6"),
            self.print_pos("U4"),
            self.print_pos("U5"),
            self.print_pos("U6"),
            self.print_pos("R4"),
            self.print_pos("R5"),
            self.print_pos("R6"),
            self.print_pos("D4"),
            self.print_pos("D5"),
            self.print_pos("D6")
        );
        println!(
            "|{}-{}-{}|{}-{}-{}|{}-{}-{}|{}-{}-{}|",
            self.print_pos("L7"),
            self.print_pos("L8"),
            self.print_pos("L9"),
            self.print_pos("U7"),
            self.print_pos("U8"),
            self.print_pos("U9"),
            self.print_pos("R7"),
            self.print_pos("R8"),
            self.print_pos("R9"),
            self.print_pos("D7"),
            self.print_pos("D8"),
            self.print_pos("D9")
        );
        println!("|--------|--------|--------|--------|");
        println!(
            "         |{}-{}-{}|",
            self.print_pos("F1"),
            self.print_pos("F2"),
            self.print_pos("F3")
        );
        println!(
            "         |{}-{}-{}|",
            self.print_pos("F4"),
            self.print_pos("F5"),
            self.print_pos("F6")
        );
        println!(
            "         |{}-{}-{}|",
            self.print_pos("F7"),
            self.print_pos("F8"),
            self.print_pos("F9")
        );
        println!("         |--------|");
    }

    fn sexy_move(&mut self) {
        self.move_r();
        self.move_u();
        self.move_rp();
        self.move_up();
    }

    // Function to rotate the R face
    fn move_r(&mut self) {
        let pos_a = ["R1", "R2", "R3", "R6", "R9", "R7", "R8", "R4", "B3", "B6", "B9", "D7", "D4", "D1", "F3", "F6", "F9", "U3", "U6", "U9"];
        let pos_n = ["R7", "R4", "R1", "R2", "R3", "R9", "R6", "R8", "U3", "U6", "U9", "B3", "B6", "B9", "D7", "D4", "D1", "F3", "F6", "F9"];
        self.rubik_move(&pos_a, &pos_n);
    }

    // Function to rotate the R face in reverse
    fn move_rp(&mut self) {
        let pos_a = ["R1", "R2", "R3", "R6", "R9", "R8", "R7", "R4", "B3", "B6", "B9", "D7", "D4", "D1", "F3", "F6", "F9", "U3", "U6", "U9"];
        let pos_n = ["R3", "R6", "R9", "R8", "R7", "R4", "R1", "R2", "D7", "D4", "D1", "F3", "F6", "F9", "U3", "U6", "U9", "B3", "B6", "B9"];
        self.rubik_move(&pos_a, &pos_n);
    }

    // Function to rotate the U face
    fn move_u(&mut self) {
        let pos_a = ["U1", "U2", "U3", "U6", "U9", "U8", "U7", "U4", "R1", "R4", "R7", "F1", "F2", "F3", "L3", "L6", "L9", "B7", "B8", "B9"];
        let pos_n = ["U7", "U4", "U1", "U2", "U3", "U6", "U9", "U8", "B7", "B8", "B9", "R7", "R4", "R1", "F1", "F2", "F3", "L9", "L6", "L3"];
        self.rubik_move(&pos_a, &pos_n);
    }

    // Function to rotate the U face in reverse
    fn move_up(&mut self) {
        let pos_a = ["U1", "U2", "U3", "U6", "U9", "U8", "U7", "U4", "R1", "R4", "R7", "F1", "F2", "F3", "L3", "L6", "L9", "B7", "B8", "B9"];
        let pos_n = ["U3", "U6", "U9", "U8", "U7", "U4", "U1", "U2", "F3", "F2", "F1", "L3", "L6", "L9", "B9", "B8", "B7", "R1", "R4", "R7"];
        self.rubik_move(&pos_a, &pos_n);
    }

    // Function to rotate the L face
    fn move_l(&mut self) {
        let pos_a = ["L1", "L2", "L3", "L6", "L9", "L7", "L8", "L4", "B7", "B4", "B1", "D3", "D6", "D9", "F1", "F4", "F7", "U1", "U4", "U7"];
        let pos_n = ["L7", "L4", "L1", "L2", "L3", "L9", "L6", "L8", "U1", "U4", "U7", "B7", "B4", "B1", "D3", "D6", "D9", "F1", "F4", "F7"];
        self.rubik_move(&pos_a, &pos_n);
    }

    // Function to rotate the L face in reverse
    fn move_lp(&mut self) {
        let pos_a = ["L1", "L2", "L3", "L6", "L9", "L7", "L8", "L4", "B7", "B4", "B1", "D3", "D6", "D9", "F1", "F4", "F7", "U1", "U4", "U7"];
        let pos_n = ["L3", "L6", "L9", "L8", "L7", "L4", "L1", "L2", "D9", "D6", "D3", "F7", "F4", "F1", "U7", "U4", "U1", "B1", "B4", "B7"];
        self.rubik_move(&pos_a, &pos_n);
    }

    // Function to rotate the D face
    fn move_d(&mut self) {
        let pos_a = ["D1", "D2", "D3", "D6", "D9", "D8", "D7", "D4", "R3", "R6", "R9", "F7", "F8", "F9", "L1", "L2", "L3", "B1", "B2", "B3"];
        let pos_n = ["D7", "D4", "D1", "D2", "D3", "D6", "D9", "D8", "B1", "B2", "B3", "R3", "R6", "R9", "F7", "F8", "F9", "L1", "L2", "L3"];
        self.rubik_move(&pos_a, &pos_n);
    }

    // Function to rotate the D face in reverse
    fn move_dp(&mut self) {
        let pos_a = ["D1", "D2", "D3", "D6", "D9", "D8", "D7", "D4", "B1", "B2", "B3", "L1", "L4", "L7", "R3", "R6", "R9", "F7", "F8", "F9"];
        let pos_n = ["D3", "D6", "D9", "D8", "D7", "D4", "D1", "D2", "L7", "L4", "L1", "F7", "F8", "F9", "B1", "B2", "B3", "R9", "R6", "R3"];
        self.rubik_move(&pos_a, &pos_n);
    }

    // Function to rotate the F face
    fn move_f(&mut self) {
        let pos_a = ["F1", "F2", "F3", "F6", "F9", "F7", "F8", "F4", "L7", "L8", "L9", "U7", "U8", "U9", "R7", "R8", "R9", "D7", "D8", "D9"];
        let pos_n = ["F7", "F4", "F1", "F2", "F3", "F9", "F6", "F8", "D7", "D8", "D9", "L7", "L8", "L9", "U7", "U8", "U9", "R7", "R8", "R9"];
        self.rubik_move(&pos_a, &pos_n);
    }

    // Function to rotate the F face in reverse
    fn move_fp(&mut self) {
        let pos_a = ["F1", "F2", "F3", "F6", "F9", "F8", "F7", "F4", "L7", "L8", "L9", "U7", "U8", "U9", "R7", "R8", "R9", "D7", "D8", "D9"];
        let pos_n = ["F3", "F6", "F9", "F8", "F7", "F4", "F1", "F2", "U7", "U8", "U9", "R7", "R8", "R9", "D7", "D8", "D9", "L7", "L8", "L9"];
        self.rubik_move(&pos_a, &pos_n);
    }

    // Function to rotate the B face
    fn move_b(&mut self) {
        let pos_a = ["B1", "B2", "B3", "B6", "B9", "B8", "B7", "B4", "L1", "L2", "L3", "U1", "U2", "U3", "R1", "R2", "R3", "D1", "D2", "D3"];
        let pos_n = ["B7", "B4", "B1", "B2", "B3", "B6", "B9", "B8", "U1", "U2", "U3", "R1", "R2", "R3", "D1", "D2", "D3", "L1", "L2", "L3"];
        self.rubik_move(&pos_a, &pos_n);
    }

    // Function to rotate the B face in reverse
    fn move_bp(&mut self) {
        let pos_a = ["B1", "B2", "B3", "B6", "B9", "B8", "B7", "B4", "L1", "L2", "L3", "U1", "U2", "U3", "R1", "R2", "R3", "D1", "D2", "D3"];
        let pos_n = ["B3", "B6", "B9", "B8", "B7", "B4", "B1", "B2", "D1", "D2", "D3", "L1", "L2", "L3", "U1", "U2", "U3", "R1", "R2", "R3"];
        self.rubik_move(&pos_a, &pos_n);
    }
}

fn main() {
    // Crear un nuevo cubo Rubik ordenado
    let mut rubik = Rubik::new();

    // Imprimir el estado inicial del cubo
    println!("Estado inicial del cubo:");
    rubik.format_cli();

    // Realizar algunos movimientos
    rubik.move_r();
    rubik.move_u();
    rubik.move_rp();
    rubik.move_up();

    // Imprimir el estado después de realizar movimientos
    println!("\nEstado del cubo después de algunos movimientos:");
    rubik.format_cli();
}

