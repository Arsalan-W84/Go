//BUILD A LOGGER .
//WHAT IS A LOGGER ?
//
//A logger is a tool in software development that records what your application is doing while it runs.

//Types of log levels
// Level :	Purpose
// DEBUG :	Detailed information for debugging
// INFO	 :	Normal application events
// WARN	 :	Something unusual happened
// ERROR :	Something failed
// FATAL :	Critical failure; typically terminates the application
//
// TASK :
// Create an enum of every loglevel and print them in the appropriate format
// Log level : 0 Debug
// Log level : 1 Info , etc.
//
// Purpose : read the code and understand each lines purpose


package main

import "fmt"

//creating a custom type 
type loglevel int

//NOTE 1 : CREATING AN ENUM OF THIS CUSTOM TYPE.
const(
	Logdebug loglevel = iota
	Loginfo 
	Logwarn 
	Logerror
	Logfatal
)

var levelnames = []string{"Debug" , "Info" , "Warn" , "Error" , "Fatal"}


//NOTE 2: adding a method to a custom type , we can call this method UPON a variable
//of this type.
func (l loglevel) String() string {
	if(l< Logdebug || l>Logfatal){
		return "Invalid log"
	}
	return levelnames[l];
}

func printloglevel(l loglevel) {
	//complete the function
	fmt.Printf("Log level: %d %s \n" , l , l.String())
} 

//--------DONT TOUCH BELOW THIS LINE
func main() {
	printloglevel(Logdebug)
	printloglevel(Loginfo)
	printloglevel(Logwarn)
	printloglevel(Logerror)
	printloglevel(Logfatal)
	printloglevel(Logfatal + Logerror)
}
