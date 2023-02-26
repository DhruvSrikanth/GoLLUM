# Feedback for Milestone 2
Overall very well done! Structure of the AST code is clear, very well documented readme and helpful comments inside code where neccessary, I think you're very well set up for completing the compiler project and you're making great progress. <br>
You have implemented and tested mostly all the semantic analysis, but I would still recommend that if you have time, to write some extra tests to confirm that you have really considered everything when it comes to semantic analysis. Below are some of the problems I discovered. 

1. There is a small problem with your control-flow checker, to my understanding, you're currently using an OR assignment clause, which basically means that if any statement within the function has a return, then you'd consider it correct. However, it should be implemented so that `Any function with a return type must return a valid value (of its return type) along all control flow paths.` as per language specification (also it'll be difficult to check dyanmically), meaning that the following functions should be considered erroneous. One idea for implementation may be to check the last statement of a function and check if it is a return statement, but how exactly you want to do this is up to you. 
```
func bar (num int) int {
    if (num == 0) {
        return 1;
    } else {
        if (num >= 1) {
            return 1;
        } else {
            // This inner else block is missing a return statement.
        }
    }
}

func foo (num int) int {
    // A for loop is not neccessarily executed, so this return statement is not guarenteed.
    for(false){
        return 0;
    }
}
```


2. For your Delete typechecking, it's a bit problematic, the following code produces a sematic error: dist.cord has not bee declared. I believe this is caused by  you checking for the full String() of the expression (dist.cord), instead of the first part of the id.
```
type coord struct {
    x int;
    y int;
};

type distance struct {
    cord * coord;
    distance int;
};

func calculateDistance(playerCoord *coord, originX int, originY int) int{
    var newCoord *coord;
    var dist *distance;

    dist = new distance;
    dist.cord = new coord;

    delete dist.cord;
    return 0;
}
```
3. You may want to check for the existance of a main function, since it is used as a entry point to the execution of your code. 
4. You're not checking that the condition of an if and for block evaluates to a boolean value, it may makes sense for a integer value to pass the type check (still not recommended), but currently you're also allowing struct types to be the conditional, which will be very problematic.
5. Another thing I'm not seeing here is checks for redeclaration? <br>
`type Point2D struct { ... }` <br>
`var Point2D int;` <br>
This should produce an error because type declarations and global variables have global scope so there cannot be a global variable declaration and type decalaration of the same scope. Check "Lanaguage Overview" page under "Language Semantics".