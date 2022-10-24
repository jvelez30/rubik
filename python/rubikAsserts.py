# Rubik Asserts 

from assertAux import *

def assertOrderedRubik(rubik) :
    assertion = True
    i = 0
    while ( i < len(positions) ) :
        pos = positions[i];
        if (rubik[0][pos] != orderedValues[i] or rubik[1][pos] != orderedColors[i]) :
            assertion = False;
        
        i+=1;
    
    if (assertion) :
        print(str(str("assertOrderedRubik\n" + str(green)) + "Assert True\n") + str(white),end="");
    else : 
        print(str(str("assertOrderedRubik\n" + str(red)) + "Assert False\n") + str(white),end="");
    

def assertMoveURubik(rubik) :

    global red, green, yellow, blue, magenta, white;
    global positions, moveUValues, moveUColors;
    assertion  = True;
    i = 0;
    while ( i < len(positions) ) :
        pos = positions[i];
        if (rubik[0][pos] != moveUValues[i] or rubik[1][pos] != moveUColors[i]) :
            assertion = False
        
        i+=1;
    
    if (assertion) :
        print(str(str("assertMoveURubik\n" + str(green)) + "Assert True\n") + str(white),end="");
    else : 
        print(str(str("assertMoveURubik\n" + str(red)) + "Assert False\n") + str(white),end="");
    

def assertMoveUPRubik(rubik) :

    global red, green, yellow, blue, magenta, white;
    global positions, moveUPValues, moveUPColors;
    assertion  = True;
    i = 0;
    while ( i < len(positions) ) :
        pos = positions[i];
        if (rubik[0][pos] != moveUPValues[i] or rubik[1][pos] != moveUPColors[i]) :
            assertion = False;
        
        i+=1;
    
    if (assertion) :
        print(str(str("assertMoveUPRubik\n" + str(green)) + "Assert True\n") + str(white),end="");
    else : 
        print(str(str("assertMoveUPRubik\n" + str(red)) + "Assert False\n") + str(white),end="");
    

def assertMoveRRubik(rubik) :

    global red, green, yellow, blue, magenta, white;
    global positions, moveRValues, moveRColors;
    assertion  = True;
    i = 0;
    while ( i < len(positions) ) :
        pos = positions[i];
        if (rubik[0][pos] != moveRValues[i] or rubik[1][pos] != moveRColors[i]) :
            assertion = False;
        
        i+=1;
    
    if (assertion) :
        print(str(str("assertMoveRRubik\n" + str(green)) + "Assert True\n") + str(white),end="");
    else : 
        print(str(str("assertMoveRRubik\n" + str(red)) + "Assert False\n") + str(white),end="");
    

def assertMoveRPRubik(rubik) :

    global red, green, yellow, blue, magenta, white;
    global positions, moveRPValues, moveRPColors;
    assertion  = True;
    i = 0;
    while ( i < len(positions) ) :
        pos = positions[i];
        if (rubik[0][pos] != moveRPValues[i] or rubik[1][pos] != moveRPColors[i]) :
            assertion = False;
        
        i+=1;
    
    if (assertion) :
        print(str(str("assertMoveRPRubik\n" + str(green)) + "Assert True\n") + str(white),end="");
    else : 
        print(str(str("assertMoveRPRubik\n" + str(red)) + "Assert False\n") + str(white),end="");
    

def assertMoveLRubik(rubik) :

    global red, green, yellow, blue, magenta, white;
    global positions, moveLValues, moveLColors;
    assertion  = True;
    i = 0;
    while ( i < len(positions) ) :
        pos = positions[i];
        if (rubik[0][pos] != moveLValues[i] or rubik[1][pos] != moveLColors[i]) :
            assertion = False;
        
        i+=1;
    
    if (assertion) :
        print(str(str("assertMoveLRubik\n" + str(green)) + "Assert True\n") + str(white),end="");
    else : 
        print(str(str("assertMoveLRubik\n" + str(red)) + "Assert False\n") + str(white),end="");
    

def assertMoveLPRubik(rubik) :

    global red, green, yellow, blue, magenta, white;
    global positions, moveLPValues, moveLPColors;
    assertion  = True;
    i = 0;
    while ( i < len(positions) ) :
        pos = positions[i];
        if (rubik[0][pos] != moveLPValues[i] or rubik[1][pos] != moveLPColors[i]) :
            assertion = False;
        
        i+=1;
    
    if (assertion) :
        print(str(str("assertMoveLPRubik\n" + str(green)) + "Assert True\n") + str(white),end="");
    else : 
        print(str(str("assertMoveLPRubik\n" + str(red)) + "Assert False\n") + str(white),end="");
    

def assertMoveFRubik(rubik) :

    global red, green, yellow, blue, magenta, white;
    global positions, moveFValues, moveFColors;
    assertion  = True;
    i = 0;
    while ( i < len(positions) ) :
        pos = positions[i];
        if (rubik[0][pos] != moveFValues[i] or rubik[1][pos] != moveFColors[i]) :
            assertion = False;
        
        i+=1;
    
    if (assertion) :
        print(str(str("assertMoveFRubik\n" + str(green)) + "Assert True\n") + str(white),end="");
    else : 
        print(str(str("assertMoveFRubik\n" + str(red)) + "Assert False\n") + str(white),end="");
    

def assertMoveFPRubik(rubik) :

    global red, green, yellow, blue, magenta, white;
    global positions, moveFPValues, moveFPColors;
    assertion  = True;
    i = 0;
    while ( i < len(positions) ) :
        pos = positions[i];
        if (rubik[0][pos] != moveFPValues[i] or rubik[1][pos] != moveFPColors[i]) :
            assertion = False;
        
        i+=1;
    
    if (assertion) :
        print(str(str("assertMoveFPRubik\n" + str(green)) + "Assert True\n") + str(white),end="");
    else : 
        print(str(str("assertMoveFPRubik\n" + str(red)) + "Assert False\n") + str(white),end="");
    

def assertMoveBRubik(rubik) :

    global red, green, yellow, blue, magenta, white;
    global positions, moveBValues, moveBColors;
    assertion  = True;
    i = 0;
    while ( i < len(positions) ) :
        pos = positions[i];
        if (rubik[0][pos] != moveBValues[i] or rubik[1][pos] != moveBColors[i]) :
            assertion = False;
        
        i+=1;
    
    if (assertion) :
        print(str(str("assertMoveBRubik\n" + str(green)) + "Assert True\n") + str(white),end="");
    else : 
        print(str(str("assertMoveBRubik\n" + str(red)) + "Assert False\n") + str(white),end="");
    

def assertMoveBPRubik(rubik) :

    global red, green, yellow, blue, magenta, white;
    global positions, moveBPValues, moveBPColors;
    assertion  = True;
    i = 0;
    while ( i < len(positions) ) :
        pos = positions[i];
        if (rubik[0][pos] != moveBPValues[i] or rubik[1][pos] != moveBPColors[i]) :
            assertion = False;
        
        i+=1;
    
    if (assertion) :
        print(str(str("assertMoveBPRubik\n" + str(green)) + "Assert True\n") + str(white),end="");
    else : 
        print(str(str("assertMoveBPRubik\n" + str(red)) + "Assert False\n") + str(white),end="");
    

def assertMoveDRubik(rubik) :

    global red, green, yellow, blue, magenta, white;
    global positions, moveDValues, moveDColors;
    assertion  = True;
    i = 0;
    while ( i < len(positions) ) :
        pos = positions[i];
        if (rubik[0][pos] != moveDValues[i] or rubik[1][pos] != moveDColors[i]) :
            assertion = False;
        
        i+=1;
    
    if (assertion) :
        print(str(str("assertMoveDRubik\n" + str(green)) + "Assert True\n") + str(white),end="");
    else : 
        print(str(str("assertMoveDRubik\n" + str(red)) + "Assert False\n") + str(white),end="");
    

def assertMoveDPRubik(rubik) :

    global red, green, yellow, blue, magenta, white;
    global positions, moveDPValues, moveDPColors;
    assertion  = True;
    i = 0;
    while ( i < len(positions) ) :
        pos = positions[i];
        if (rubik[0][pos] != moveDPValues[i] or rubik[1][pos] != moveDPColors[i]) :
            assertion = False;
        
        i+=1;
    
    if (assertion) :
        print(str(str("assertMoveDPRubik\n" + str(green)) + "Assert True\n") + str(white),end="");
    else : 
        print(str(str("assertMoveDPRubik\n" + str(red)) + "Assert False\n") + str(white),end="");
    

def assertSexyMoveRubik(rubik) :

    global red, green, yellow, blue, magenta, white;
    global positions, sexyMoveValues, sexyMoveColors;
    assertion  = True;
    i = 0;
    while ( i < len(positions) ) :
        pos = positions[i];
        if (rubik[0][pos] != sexyMoveValues[i] or rubik[1][pos] != sexyMoveColors[i]) :
            assertion = False;
        
        i+=1;
    
    if (assertion) :
        print(str(str("assertSexyMoveRubik\n" + str(green)) + "Assert True\n") + str(white),end="");
    else : 
        print(str(str("assertSexyMoveRubik\n" + str(red)) + "Assert False\n") + str(white),end="");