package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mysession"))

// Verify whether one player has won or the game has ended in a draw and if the game is still in progress switch the player's turn
func ResultsValidation(c *gin.Context) {

	session, _ := store.Get(c.Request, "mysession")

	if (session.Values["gamestate0"] == session.Values["gamestate1"] && session.Values["gamestate1"] == session.Values["gamestate2"] && session.Values["gamestate2"] == "X") ||
		(session.Values["gamestate0"] == session.Values["gamestate4"] && session.Values["gamestate4"] == session.Values["gamestate8"] && session.Values["gamestate8"] == "X") ||
		(session.Values["gamestate0"] == session.Values["gamestate3"] && session.Values["gamestate3"] == session.Values["gamestate6"] && session.Values["gamestate6"] == "X") ||
		(session.Values["gamestate1"] == session.Values["gamestate4"] && session.Values["gamestate4"] == session.Values["gamestate7"] && session.Values["gamestate7"] == "X") ||
		(session.Values["gamestate2"] == session.Values["gamestate4"] && session.Values["gamestate4"] == session.Values["gamestate6"] && session.Values["gamestate6"] == "X") ||
		(session.Values["gamestate2"] == session.Values["gamestate5"] && session.Values["gamestate5"] == session.Values["gamestate8"] && session.Values["gamestate8"] == "X") ||
		(session.Values["gamestate3"] == session.Values["gamestate4"] && session.Values["gamestate4"] == session.Values["gamestate5"] && session.Values["gamestate5"] == "X") ||
		(session.Values["gamestate6"] == session.Values["gamestate7"] && session.Values["gamestate7"] == session.Values["gamestate8"] && session.Values["gamestate8"] == "X") {
		session.Values["gamestatus"] = ""
		session.Values["gamestatus"] = "Player X has won!"
		session.Values["disable"] = "disabled"
		session.Save(c.Request, c.Writer) // Save the created sessions
		c.Redirect(http.StatusSeeOther, "/display")
		return
	} else if (session.Values["gamestate0"] == session.Values["gamestate1"] && session.Values["gamestate1"] == session.Values["gamestate2"] && session.Values["gamestate2"] == "O") ||
		(session.Values["gamestate0"] == session.Values["gamestate4"] && session.Values["gamestate4"] == session.Values["gamestate8"] && session.Values["gamestate8"] == "O") ||
		(session.Values["gamestate0"] == session.Values["gamestate3"] && session.Values["gamestate3"] == session.Values["gamestate6"] && session.Values["gamestate6"] == "O") ||
		(session.Values["gamestate1"] == session.Values["gamestate4"] && session.Values["gamestate4"] == session.Values["gamestate7"] && session.Values["gamestate7"] == "O") ||
		(session.Values["gamestate2"] == session.Values["gamestate4"] && session.Values["gamestate4"] == session.Values["gamestate6"] && session.Values["gamestate6"] == "O") ||
		(session.Values["gamestate2"] == session.Values["gamestate5"] && session.Values["gamestate5"] == session.Values["gamestate8"] && session.Values["gamestate8"] == "O") ||
		(session.Values["gamestate3"] == session.Values["gamestate4"] && session.Values["gamestate4"] == session.Values["gamestate5"] && session.Values["gamestate5"] == "O") ||
		(session.Values["gamestate6"] == session.Values["gamestate7"] && session.Values["gamestate7"] == session.Values["gamestate8"] && session.Values["gamestate8"] == "O") {
		session.Values["gamestatus"] = ""
		session.Values["gamestatus"] = "Player O has won!"
		session.Values["disable"] = "disabled"
		session.Save(c.Request, c.Writer)
		c.Redirect(http.StatusSeeOther, "/display")
		return
	} else if (session.Values["gamestate0"] != nil && session.Values["gamestate1"] != nil && session.Values["gamestate2"] != nil && session.Values["gamestate3"] != nil && session.Values["gamestate4"] != nil && session.Values["gamestate5"] != nil && session.Values["gamestate6"] != nil && session.Values["gamestate7"] != nil && session.Values["gamestate8"] != nil) {
		session.Values["gamestatus"] = ""
		session.Values["gamestatus"] = "Game ended in a draw!"
		session.Save(c.Request, c.Writer)
		c.Redirect(http.StatusSeeOther, "/display")
		return
	} else {
		if session.Values["currentplayer"] == "X" {
			session.Values["currentplayer"] = "O"
		} else {
			session.Values["currentplayer"] = "X"
		}
		subCurrPlayerDuringGame := fmt.Sprintf("It's %v's turn", session.Values["currentplayer"])
		session.Values["gamestatus"] = subCurrPlayerDuringGame
		session.Save(c.Request, c.Writer)
		c.Redirect(http.StatusSeeOther, "/display")
		return
	}
}

//Store the player's value for the corresponding clicked box, each in a separate session variable
func GameSave(c *gin.Context) {

	session, _ := store.Get(c.Request, "mysession")

	session.Values["playButton"] = c.Request.PostFormValue("playButton")
	cell0 := c.Request.PostFormValue("cell0")
	cell1 := c.Request.PostFormValue("cell1")
	cell2 := c.Request.PostFormValue("cell2")
	cell3 := c.Request.PostFormValue("cell3")
	cell4 := c.Request.PostFormValue("cell4")
	cell5 := c.Request.PostFormValue("cell5")
	cell6 := c.Request.PostFormValue("cell6")
	cell7 := c.Request.PostFormValue("cell7")
	cell8 := c.Request.PostFormValue("cell8")
	session.Values["restartButton"] = c.Request.PostFormValue("restartButton")
	session.Save(c.Request, c.Writer)

	if session.Values["restartButton"] != "" {
		c.Redirect(http.StatusSeeOther, "/restart")
		return
	}
	if cell0 != "" {
		session.Values["gamestate0"] = session.Values["currentplayer"]
		session.Save(c.Request, c.Writer)
		c.Redirect(http.StatusSeeOther, "/userindex")
		return
	}
	if cell1 != "" {
		session.Values["gamestate1"] = session.Values["currentplayer"]
		session.Save(c.Request, c.Writer)
		c.Redirect(http.StatusSeeOther, "/userindex")
		return
	}
	if cell2 != "" {
		session.Values["gamestate2"] = session.Values["currentplayer"]
		session.Save(c.Request, c.Writer)
		c.Redirect(http.StatusSeeOther, "/userindex")
		return
	}
	if cell3 != "" {
		session.Values["gamestate3"] = session.Values["currentplayer"]
		session.Save(c.Request, c.Writer)
		c.Redirect(http.StatusSeeOther, "/userindex")
		return
	}
	if cell4 != "" {
		session.Values["gamestate4"] = session.Values["currentplayer"]
		session.Save(c.Request, c.Writer)
		c.Redirect(http.StatusSeeOther, "/userindex")
		return
	}
	if cell5 != "" {
		session.Values["gamestate5"] = session.Values["currentplayer"]
		session.Save(c.Request, c.Writer)
		c.Redirect(http.StatusSeeOther, "/userindex")
		return
	}
	if cell6 != "" {
		session.Values["gamestate6"] = session.Values["currentplayer"]
		session.Save(c.Request, c.Writer)
		c.Redirect(http.StatusSeeOther, "/userindex")
		return
	}
	if cell7 != "" {
		session.Values["gamestate7"] = session.Values["currentplayer"]
		session.Save(c.Request, c.Writer)
		c.Redirect(http.StatusSeeOther, "/userindex")
		return
	}
	if cell8 != "" {
		session.Values["gamestate8"] = session.Values["currentplayer"]
		session.Save(c.Request, c.Writer)
		c.Redirect(http.StatusSeeOther, "/userindex")
		return
	}
}

//Handle what to display in the html template
func Display(c *gin.Context) {

	session, _ := store.Get(c.Request, "mysession")

	//Display the first player's turn
	if session.Values["playButton"] == nil {
		session.Values["currentplayer"] = "X"
		session.Values["subCurrentPlayer"] = fmt.Sprintf("It's %v's turn", session.Values["currentplayer"])
		session.Values["gamestatus"] = session.Values["subCurrentPlayer"]
		session.Save(c.Request, c.Writer)
	}

	//Choose what to display in box 1 of the grid; whether X, O or a radio input
	if  session.Values["gamestate0"] == "X" { 
		session.Values["cellValue1"] = "X"
	} else if session.Values["gamestate0"] == "O" {
		session.Values["cellValue1"] = "O"
	} else {
		subSessDisable1 := fmt.Sprintf("<input type='radio' name='cell0' value='0' %v onclick='document.getElementById(\"play-btn\").click();'/>", session.Values["disable"])
		if session.Values["disable"] == "disabled" {
			session.Values["cellValue1"] = template.HTML(subSessDisable1)	
		} else {
			session.Values["cellValue1"] = template.HTML("<input type='radio' name='cell0' value='0' onclick='document.getElementById(\"play-btn\").click();'/>")
		}
	}


	//Choose what to display in box 2 of the grid; whether X, O or a radio input
	if  session.Values["gamestate1"] == "X" { 
		session.Values["cellValue2"] = "X"
	} else if session.Values["gamestate1"] == "O" {
		session.Values["cellValue2"] = "O"
	} else {
		subSessDisable2 := fmt.Sprintf("<input type='radio' name='cell1' value='1' %v onclick='document.getElementById(\"play-btn\").click();'/>", session.Values["disable"])
		if session.Values["disable"] == "disabled" {
			session.Values["cellValue2"] = template.HTML(subSessDisable2)	
		} else {
			session.Values["cellValue2"] = template.HTML("<input type='radio' name='cell1' value='1' onclick='document.getElementById(\"play-btn\").click();'/>")
		}
	}


	//Choose what to display in box 3 of the grid; whether X, O or a radio input
	if  session.Values["gamestate2"] == "X" { 
		session.Values["cellValue3"] = "X"
	} else if session.Values["gamestate2"] == "O" {
		session.Values["cellValue3"] = "O"
	} else {
		subSessDisable3 := fmt.Sprintf("<input type='radio' name='cell2' value='2' %v onclick='document.getElementById(\"play-btn\").click();'/>", session.Values["disable"])
		if session.Values["disable"] == "disabled" {
			session.Values["cellValue3"] = template.HTML(subSessDisable3)	
		} else {
			session.Values["cellValue3"] = template.HTML("<input type='radio' name='cell2' value='2' onclick='document.getElementById(\"play-btn\").click();'/>")
		}
	}


	//Choose what to display in box 4 of the grid; whether X, O or a radio input
	if  session.Values["gamestate3"] == "X" { 
		session.Values["cellValue4"] = "X"
	} else if session.Values["gamestate3"] == "O" {
		session.Values["cellValue4"] = "O"
	} else {
		subSessDisable4 := fmt.Sprintf("<input type='radio' name='cell3' value='3' %v onclick='document.getElementById(\"play-btn\").click();'/>", session.Values["disable"])
		if session.Values["disable"] == "disabled" {
			session.Values["cellValue4"] = template.HTML(subSessDisable4)	
		} else {
			session.Values["cellValue4"] = template.HTML("<input type='radio' name='cell3' value='3' onclick='document.getElementById(\"play-btn\").click();'/>")
		}
	}


	//Choose what to display in box 5 of the grid; whether X, O or a radio input
	if  session.Values["gamestate4"] == "X" { 
		session.Values["cellValue5"] = "X"
	} else if session.Values["gamestate4"] == "O" {
		session.Values["cellValue5"] = "O"
	} else {
		subSessDisable5 := fmt.Sprintf("<input type='radio' name='cell4' value='4' %v onclick='document.getElementById(\"play-btn\").click();'/>", session.Values["disable"])
		if session.Values["disable"] == "disabled" {
			session.Values["cellValue5"] = template.HTML(subSessDisable5)	
		} else {
			session.Values["cellValue5"] = template.HTML("<input type='radio' name='cell4' value='4' onclick='document.getElementById(\"play-btn\").click();'/>")
		}
	}


	//Choose what to display in box 6 of the grid; whether X, O or a radio input
	if  session.Values["gamestate5"] == "X" { 
		session.Values["cellValue6"] = "X"
	} else if session.Values["gamestate5"] == "O" {
		session.Values["cellValue6"] = "O"
	} else {
		subSessDisable6 := fmt.Sprintf("<input type='radio' name='cell5' value='5' %v onclick='document.getElementById(\"play-btn\").click();'/>", session.Values["disable"])
		if session.Values["disable"] == "disabled" {
			session.Values["cellValue6"] = template.HTML(subSessDisable6)	
		} else {
			session.Values["cellValue6"] = template.HTML("<input type='radio' name='cell5' value='5' onclick='document.getElementById(\"play-btn\").click();'/>")
		}
	}


	//Choose what to display in box 7 of the grid; whether X, O or a radio input
	if  session.Values["gamestate6"] == "X" { 
		session.Values["cellValue7"] = "X"
	} else if session.Values["gamestate6"] == "O" {
		session.Values["cellValue7"] = "O"
	} else {
		subSessDisable7 := fmt.Sprintf("<input type='radio' name='cell6' value='6' %v onclick='document.getElementById(\"play-btn\").click();'/>", session.Values["disable"])
		if session.Values["disable"] == "disabled" {
			session.Values["cellValue7"] = template.HTML(subSessDisable7)	
		} else {
			session.Values["cellValue7"] = template.HTML("<input type='radio' name='cell6' value='6' onclick='document.getElementById(\"play-btn\").click();'/>")
		}
	}


	//Choose what to display in box 8 of the grid; whether X, O or a radio input
	if  session.Values["gamestate7"] == "X" { 
		session.Values["cellValue8"] = "X"
	} else if session.Values["gamestate7"] == "O" {
		session.Values["cellValue8"] = "O"
	} else {
		subSessDisable8 := fmt.Sprintf("<input type='radio' name='cell7' value='7' %v onclick='document.getElementById(\"play-btn\").click();'/>", session.Values["disable"])
		if session.Values["disable"] == "disabled" {
			session.Values["cellValue8"] = template.HTML(subSessDisable8)	
		} else {
			session.Values["cellValue8"] = template.HTML("<input type='radio' name='cell7' value='7' onclick='document.getElementById(\"play-btn\").click();'/>")
		}
	}


	//Choose what to display in box 9 of the grid; whether X, O or a radio input
	if  session.Values["gamestate8"] == "X" { 
		session.Values["cellValue9"] = "X"
	} else if session.Values["gamestate8"] == "O" {
		session.Values["cellValue9"] = "O"
	} else {
		subSessDisable9 := fmt.Sprintf("<input type='radio' name='cell8' value='8' %v onclick='document.getElementById(\"play-btn\").click();'/>", session.Values["disable"])
		if session.Values["disable"] == "disabled" {
			session.Values["cellValue9"] = template.HTML(subSessDisable9)	
		} else {
			session.Values["cellValue9"] = template.HTML("<input type='radio' name='cell8' value='8' onclick='document.getElementById(\"play-btn\").click();'/>")
		}
	}

	// Save the created sessions
	session.Save(c.Request, c.Writer)
	// Go to the html and send along a map of key - value pairs to be displayed in the html page
	c.HTML(200, "template.html", gin.H{"gamestatus": session.Values["gamestatus"], "disable": session.Values["disable"], "cellvalue1": session.Values["cellValue1"], "cellvalue2": session.Values["cellValue2"], "cellvalue3": session.Values["cellValue3"], "cellvalue4": session.Values["cellValue4"], "cellvalue5": session.Values["cellValue5"], "cellvalue6": session.Values["cellValue6"], "cellvalue7": session.Values["cellValue7"], "cellvalue8": session.Values["cellValue8"], "cellvalue9": session.Values["cellValue9"]})
}

// Reset the games' parameters and values to start the game afresh
func Restart(c *gin.Context) {

	session, _ := store.Get(c.Request, "mysession")

		session.Values["currentplayer"] = ""
		session.Values["gamestate0"] = ""
		session.Values["gamestate1"] = ""
		session.Values["gamestate2"] = ""
		session.Values["gamestate3"] = ""
		session.Values["gamestate4"] = ""
		session.Values["gamestate5"] = ""
		session.Values["gamestate6"] = ""
		session.Values["gamestate7"] = ""
		session.Values["gamestate8"] = ""
		session.Values["gamestatus"] = ""
		session.Options.MaxAge = -1  //Set the maximum age of the sessions store
		session.Values["currentplayer"] = "X"
		session.Values["gamestatus"] = session.Values["subCurrentPlayer"]
		session.Save(c.Request, c.Writer)

		c.Redirect(303, "/display")
	
}