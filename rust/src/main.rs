// src/main.rs

// Definición de constantes para los códigos de color.
const RED: &str = "\x1b[31m";
const GREEN: &str = "\x1b[32m";
const YELLOW: &str = "\x1b[33m";
const BLUE: &str = "\x1b[34m";
const MAGENTA: &str = "\x1b[95m";
const WHITE: &str = "\x1b[97m";

use std::collections::HashMap;
use std::clone::Clone;

struct Rubik {
    values: HashMap<String, String>,
    colors: HashMap<String, String>,
}

impl Rubik {
    // Crea un cubo Rubik ordenado
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

    // Realiza un movimiento del cubo Rubik
    fn rubik_move(&mut self, pos_a: &[&str], pos_n: &[&str]) {
        let temp_values = self.values.clone();
        let temp_colors = self.colors.clone();
        
        for (from, to) in pos_a.iter().zip(pos_n.iter()) {
            self.values.insert((*from).to_string(), temp_values[to].clone());
            self.colors.insert((*from).to_string(), temp_colors[to].clone());
        }
    }

    // Movimiento "U" en sentido horario
    fn move_u(&mut self) {
        let pos_to_change = ["U1", "U2", "U3", "U6", "U9", "U8", "U7", "U4", "R1", "R4", "R7", "F1", "F2", "F3", "L3", "L6", "L9", "B7", "B8", "B9"];
        let pos_new_val = ["U7", "U4", "U1", "U2", "U3", "U6", "U9", "U8", "B7", "B8", "B9", "R7", "R4", "R1", "F1", "F2", "F3", "L9", "L6", "L3"];
        self.rubik_move(&pos_to_change, &pos_new_val);
    }

    // Movimiento "U'" (U en sentido antihorario)
    fn move_up(&mut self) {
        let pos_to_change = ["U1", "U2", "U3", "U6", "U9", "U8", "U7", "U4", "R1", "R4", "R7", "F1", "F2", "F3", "L3", "L6", "L9", "B7", "B8", "B9"];
        let pos_new_val = ["U3", "U6", "U9", "U8", "U7", "U4", "U1", "U2", "F3", "F2", "F1", "L3", "L6", "L9", "B9", "B8", "B7", "R1", "R4", "R7"];
        self.rubik_move(&pos_to_change, &pos_new_val);
    }

    // Movimiento "R" en sentido horario
    fn move_r(&mut self) {
        let pos_to_change = ["R1", "R2", "R3", "R6", "R9", "R7", "R8", "R4", "B3", "B6", "B9", "D7", "D4", "D1", "F3", "F6", "F9", "U3", "U6", "U9"];
        let pos_new_val = ["R7", "R4", "R1", "R2", "R3", "R9", "R6", "R8", "U3", "U6", "U9", "B3", "B6", "B9", "D7", "D4", "D1", "F3", "F6", "F9"];
        self.rubik_move(&pos_to_change, &pos_new_val);
    }

    // Movimiento "R'" (R en sentido antihorario)
    fn move_rp(&mut self) {
        let pos_to_change = ["R1", "R2", "R3", "R6", "R9", "R8", "R7", "R4", "B3", "B6", "B9", "D7", "D4", "D1", "F3", "F6", "F9", "U3", "U6", "U9"];
        let pos_new_val = ["R3", "R6", "R9", "R8", "R7", "R4", "R1", "R2", "D7", "D4", "D1", "F3", "F6", "F9", "U3", "U6", "U9", "B3", "B6", "B9"];
        self.rubik_move(&pos_to_change, &pos_new_val);
    }

    // Movimiento "L" en sentido horario
    fn move_l(&mut self) {
        let pos_to_change = ["L1", "L2", "L3", "L6", "L9", "L7", "L8", "L4", "B1", "B4", "B7", "U1", "U4", "U7", "F1", "F4", "F7", "D3", "D6", "D9"];
        let pos_new_val = ["L7", "L4", "L1", "L2", "L3", "L9", "L6", "L8", "D9", "D6", "D3", "B1", "B4", "B7", "U1", "U4", "U7", "F7", "F4", "F1"];
        self.rubik_move(&pos_to_change, &pos_new_val);
    }

    // Movimiento "L'" (L en sentido antihorario)
    fn move_lp(&mut self) {
        let pos_to_change = ["L1", "L2", "L3", "L6", "L9", "L8", "L7", "L4", "B1", "B4", "B7", "U1", "U4", "U7", "F1", "F4", "F7", "D3", "D6", "D9"];
        let pos_new_val = ["L3", "L6", "L9", "L8", "L7", "L4", "L1", "L2", "U1", "U4", "U7", "F1", "F4", "F7", "D9", "D6", "D3", "B7", "B4", "B1"];
        self.rubik_move(&pos_to_change, &pos_new_val);
    }

    // Movimiento "F" en sentido horario
    fn move_f(&mut self) {
        let pos_to_change = ["F1", "F2", "F3", "F6", "F9", "F7", "F8", "F4", "L7", "L8", "L9", "U7", "U8", "U9", "R7", "R8", "R9", "D7", "D8", "D9"];
        let pos_new_val = ["F7", "F4", "F1", "F2", "F3", "F9", "F6", "F8", "D7", "D8", "D9", "L7", "L8", "L9", "U7", "U8", "U9", "R7", "R8", "R9"];
        self.rubik_move(&pos_to_change, &pos_new_val);
    }

    // Movimiento "F'" (F en sentido antihorario)
    fn move_fp(&mut self) {
        let pos_to_change = ["F1", "F2", "F3", "F6", "F9", "F8", "F7", "F4", "L7", "L8", "L9", "U7", "U8", "U9", "R7", "R8", "R9", "D7", "D8", "D9"];
        let pos_new_val = ["F3", "F6", "F9", "F8", "F7", "F4", "F1", "F2", "U7", "U8", "U9", "R7", "R8", "R9", "D7", "D8", "D9", "L7", "L8", "L9"];
        self.rubik_move(&pos_to_change, &pos_new_val);
    }

    // Movimiento "D" en sentido horario
    fn move_d(&mut self) {
        let pos_to_change = ["D1", "D2", "D3", "D6", "D9", "D7", "D8", "D4", "R3", "R6", "R9", "F7", "F8", "F9", "L1", "L2", "L3", "B1", "B2", "B3"];
        let pos_new_val = ["D7", "D4", "D1", "D2", "D3", "D9", "D6", "D8", "F7", "F8", "F9", "L3", "L2", "L1", "B3", "B2", "B1", "R9", "R6", "R3"];
        self.rubik_move(&pos_to_change, &pos_new_val);
    }

    // Movimiento "D'" (D en sentido antihorario)
    fn move_dp(&mut self) {
        let pos_to_change = ["D1", "D2", "D3", "D6", "D9", "D8", "D7", "D4", "R3", "R6", "R9", "F7", "F8", "F9", "L1", "L2", "L3", "B1", "B2", "B3"];
        let pos_new_val = ["D3", "D6", "D9", "D8", "D7", "D4", "D1", "D2", "B3", "B2", "B1", "R9", "R6", "R3", "F9", "F8", "F7", "L1", "L2", "L3"];
        self.rubik_move(&pos_to_change, &pos_new_val);
    }

    // Movimiento "B" en sentido horario
    fn move_b(&mut self) {
        let pos_to_change = ["B1", "B2", "B3", "B6", "B9", "B7", "B8", "B4", "R1", "R2", "R3", "U1", "U2", "U3", "L1", "L2", "L3", "D1", "D2", "D3"];
        let pos_new_val = ["B7", "B4", "B1", "B2", "B3", "B9", "B6", "B8", "D3", "D2", "D1", "R3", "R2", "R1", "U3", "U2", "U1", "L3", "L2", "L1"];
        self.rubik_move(&pos_to_change, &pos_new_val);
    }

    // Movimiento "B'" (B en sentido antihorario)
    fn move_bp(&mut self) {
        let pos_to_change = ["B1", "B2", "B3", "B6", "B9", "B8", "B7", "B4", "R1", "R2", "R3", "U1", "U2", "U3", "L1", "L2", "L3", "D1", "D2", "D3"];
        let pos_new_val = ["B3", "B6", "B9", "B8", "B7", "B4", "B1", "B2", "U1", "U2", "U3", "L3", "L2", "L1", "D3", "D2", "D1", "R1", "R2", "R3"];
        self.rubik_move(&pos_to_change, &pos_new_val);
    }

    // Movimiento "sexy" (R, U, R', U')
    fn sexy_move(&mut self) {
        self.move_r();
        self.move_u();
        self.move_rp();
        self.move_up();
    }
}

fn main() {
    // Crear el cubo Rubik
    let mut rubik = Rubik::new();
    
    // Ejemplo de ejecución del movimiento "sexy"
    rubik.sexy_move();

    println!("Sexy move ejecutado correctamente.");
}


def printPos(rubik, pos):
    return rubik[1][pos]+rubik[0][pos]+white

def formatCli(rubik):
    print("         |--------|"+"\n" +
                "         |"+printPos(rubik, "B1")+"-"+printPos(rubik, "B2")+"-"+printPos(rubik, "B3")+"|"+"\n" +
                "         |"+printPos(rubik, "B4")+"-"+printPos(rubik, "B5")+"-"+printPos(rubik, "B6")+"|"+"\n" +
                "         |"+printPos(rubik, "B7")+"-"+printPos(rubik, "B8")+"-"+printPos(rubik, "B9")+"|"+"\n" +
                "|--------|--------|--------|--------|"+"\n" +
                "|"+printPos(rubik, "L1")+"-"+printPos(rubik, "L2")+"-"+printPos(rubik, "L3")+"|" +
                printPos(rubik, "U1")+"-"+printPos(rubik, "U2")+"-"+printPos(rubik, "U3") +
                "|"+printPos(rubik, "R1")+"-"+printPos(rubik, "R2")+"-"+printPos(rubik, "R3")+"|" +
                printPos(rubik, "D1")+"-"+printPos(rubik, "D2")+"-"+printPos(rubik, "D3")+"|"+"\n" +
                "|"+printPos(rubik, "L4")+"-"+printPos(rubik, "L5")+"-"+printPos(rubik, "L6")+"|" +
                printPos(rubik, "U4")+"-"+printPos(rubik, "U5")+"-"+printPos(rubik, "U6") +
                "|"+printPos(rubik, "R4")+"-"+printPos(rubik, "R5")+"-"+printPos(rubik, "R6")+"|" +
                printPos(rubik, "D4")+"-"+printPos(rubik, "D5")+"-"+printPos(rubik, "D6")+"|"+"\n" +
                "|"+printPos(rubik, "L7")+"-"+printPos(rubik, "L8")+"-"+printPos(rubik, "L9")+"|" +
                printPos(rubik, "U7")+"-"+printPos(rubik, "U8")+"-"+printPos(rubik, "U9") +
                "|"+printPos(rubik, "R7")+"-"+printPos(rubik, "R8")+"-"+printPos(rubik, "R9")+"|" +
                printPos(rubik, "D7")+"-"+printPos(rubik, "D8")+"-"+printPos(rubik, "D9")+"|"+"\n" +
                "|--------|--------|--------|--------|"+"\n" +
                "         |"+printPos(rubik, "F1")+"-"+printPos(rubik, "F2")+"-"+printPos(rubik, "F3")+"|"+"\n" +
                "         |"+printPos(rubik, "F4")+"-"+printPos(rubik, "F5")+"-"+printPos(rubik, "F6")+"|"+"\n" +
                "         |"+printPos(rubik, "F7")+"-"+printPos(rubik, "F8")+"-"+printPos(rubik, "F9")+"|"+"\n" +
                "         |--------|"+"\n")

def printPosHtml(rubik, pos):
    return "<span style=\"background-color: "+translateColorToHtml(rubik[pos]['c'])+"\">"+rubik[pos]['v']+"</span>"

def formatHtml(rubik):
    print("<table>")
    print(" <tr><td></td><td>"+printPosHtml(rubik, "B1")+"-"+printPosHtml(rubik, "B2")+"-"+printPosHtml(rubik, "B3")+"</td></tr>")
    print(" <tr><td></td><td>"+printPosHtml(rubik, "B4")+"-"+printPosHtml(rubik, "B5")+"-"+printPosHtml(rubik, "B6")+"</td></tr>")
    print(" <tr><td></td><td>"+printPosHtml(rubik, "B7")+"-"+printPosHtml(rubik, "B8")+"-"+printPosHtml(rubik, "B9")+"</td></tr>")
    print(" <tr><td>"+printPosHtml(rubik, "L1")+"-"+printPosHtml(rubik, "L2")+"-"+printPosHtml(rubik, "L3")+"</td>")
    print("     <td>"+printPosHtml(rubik, "U1")+"-"+printPosHtml(rubik, "U2")+"-"+printPosHtml(rubik, "U3")+"</td>")
    print("     <td>"+printPosHtml(rubik, "R1")+"-"+printPosHtml(rubik, "R2")+"-"+printPosHtml(rubik, "R3")+"</td>")
    print("     <td>"+printPosHtml(rubik, "D1")+"-"+printPosHtml(rubik, "D2")+"-"+printPosHtml(rubik, "D3")+"</td></tr>")
    print(" <tr><td>"+printPosHtml(rubik, "L4")+"-"+printPosHtml(rubik, "L5")+"-"+printPosHtml(rubik, "L6")+"</td>")
    print("     <td>"+printPosHtml(rubik, "U4")+"-"+printPosHtml(rubik, "U5")+"-"+printPosHtml(rubik, "U6")+"</td>")
    print("     <td>"+printPosHtml(rubik, "R4")+"-"+printPosHtml(rubik, "R5")+"-"+printPosHtml(rubik, "R6")+"</td>")
    print("     <td>"+printPosHtml(rubik, "D4")+"-"+printPosHtml(rubik, "D5")+"-"+printPosHtml(rubik, "D6")+"</td></tr>")
    print(" <tr><td>"+printPosHtml(rubik, "L7")+"-"+printPosHtml(rubik, "L8")+"-"+printPosHtml(rubik, "L9")+"</td>")
    print("     <td>"+printPosHtml(rubik, "U7")+"-"+printPosHtml(rubik, "U8")+"-"+printPosHtml(rubik, "U9")+"</td>")
    print("     <td>"+printPosHtml(rubik, "R7")+"-"+printPosHtml(rubik, "R8")+"-"+printPosHtml(rubik, "R9")+"</td>")
    print("     <td>"+printPosHtml(rubik, "D7")+"-"+printPosHtml(rubik, "D8")+"-"+printPosHtml(rubik, "D9")+"</td></tr>")
    print(" <tr><td></td><td>"+printPosHtml(rubik, "F1")+"-"+printPosHtml(rubik, "F2")+"-"+printPosHtml(rubik, "F3")+"</td></tr>")
    print(" <tr><td></td><td>"+printPosHtml(rubik, "F4")+"-"+printPosHtml(rubik, "F5")+"-"+printPosHtml(rubik, "F6")+"</td></tr>")
    print(" <tr><td></td><td>"+printPosHtml(rubik, "F7")+"-"+printPosHtml(rubik, "F8")+"-"+printPosHtml(rubik, "F9")+"</td></tr>")
    print("</table>")

// # def translateColorToHtml(color):
// #     match color:
// #         case red:
// #             return "red"
// #         case green:
// #             return "lightgreen"
// #         case yellow:
// #             return "yellow"
// #         case blue:
// #             return "lightblue"
// #         case magenta:
// #             return "orange"
// #         case white:
// #             return "white" 
def main():
    rubik = getOrderedRubik()
    // # print(rubik)

    // # for j in range (0,6):
    // #    rubik=sexyMove(rubik)


    // # Filter get var
    // # movesParam = filter_input(1, 'm', 513);
    // # // If moves are setted via string query
    // # if (typeof movesParam !== 'undefined') {
    // #     moves = str_split(movesParam);
    // #     moves.foreach(move => {
    // #         switch (move) {
    // #             case "U":
    // #                 rubik = moveU(rubik);
    // #                 break;
    // #             case 'u':
    // #                 rubik = moveUP(rubik);
    // #                 break;
    // #             case "R":
    // #                 rubik = moveR(rubik);
    // #                 break;
    // #             case 'r':
    // #                 rubik = moveRP(rubik);
    // #                 break;
    // #             case "L":
    // #                 rubik = moveL(rubik);
    // #                 break;
    // #             case 'l':
    // #                 rubik = moveLP(rubik);
    // #                 break;
    // #             case "F":
    // #                 rubik = moveF(rubik);
    // #                 break;
    // #             case 'f':
    // #                 rubik = moveFP(rubik);
    // #                 break;
    // #             case "B":
    // #                 rubik = moveB(rubik);
    // #                 break;
    // #             case 'b':
    // #                 rubik = moveBP(rubik);
    // #                 break;
    // #             case "D":
    // #                 rubik = moveD(rubik);
    // #                 break;
    // #             case 'd':
    // #                 rubik = moveDP(rubik);
    // #                 break;
    // #         
    // #     )
    // # 
    rubik=moveDP(rubik)
    // #if typeof argv[0] is not "undefined") {
    //     # if command line cli format is displayed
    formatCli(rubik)
    // # else {
    //     # else html format is displayed
    // #    formatHtml(rubik);

    // # var_dump($rubik);

if __name__ == "__main__":
    main()