package handlers

import (
	models "gofiber_restapi/app/models"
	utils "gofiber_restapi/app/utils"
	"gofiber_restapi/database"

	"github.com/gofiber/fiber/v2"
)

func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []models.Note

	// find all notes in the database
	db.Find(&notes)

	// Else return notes
	return c.JSON(utils.Response(notes, 200, true, ""))
}

func CreateNote(c *fiber.Ctx) error {
	db := database.DB
	note := new(models.Note)

	// Store the body in the note and return error if encountered
	err := c.BodyParser(note)

	if err != nil {
		return c.Status(500).JSON(utils.Response(err, 500, false, "Please review the request object")) //fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Add a uuid to the note
	// note.ID = uuid.New()
	// Create the Note and return error if encountered
	err = db.Create(&note).Error
	if err != nil {
		return c.Status(500).JSON(utils.Response(err, 500, false, "Please review the request object")) //fiber.Map{"status": "error", "message": "Could not create note", "data": err})
	}

	// Return the created note
	return c.JSON(utils.Response(note, 200, true, "")) //fiber.Map{"status": "success", "message": "Created Note", "data": note})
}

func GetNote(c *fiber.Ctx) error {
	db := database.DB
	var note models.Note

	// Read the param noteId
	id := c.Params("noteId")

	// Find the note with the given Id
	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	// if note.ID == uuid.Nil {
	//     return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	// }

	// Return the note with the Id
	return c.JSON(utils.Response(note, 200, true, "")) //fiber.Map{"status": "success", "message": "Notes Found", "data": note})
}

func UpdateNote(c *fiber.Ctx) error {
	type updateNote struct {
		Title    string `json:"title"`
		SubTitle string `json:"sub_title"`
		Text     string `json:"Text"`
	}
	db := database.DB
	var note models.Note

	// Read the param noteId
	id := c.Params("noteId")

	// Find the note with the given Id
	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	// if note.ID == uuid.Nil {
	//     return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	// }

	// Store the body containing the updated data and return error if encountered
	var updateNoteData updateNote
	err := c.BodyParser(&updateNoteData)
	if err != nil {
		return c.Status(500).JSON(utils.Response(err, 500, false, "Please review the request object")) //fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Edit the note
	note.Title = updateNoteData.Title
	note.SubTitle = updateNoteData.SubTitle
	note.Text = updateNoteData.Text

	// Save the Changes
	db.Save(&note)

	// Return the updated note
	return c.JSON(utils.Response(note, 200, true, "")) //fiber.Map{"status": "success", "message": "Notes Found", "data": note})
}

func DeleteNote(c *fiber.Ctx) error {
	db := database.DB
	var note models.Note

	// Read the param noteId
	id := c.Params("noteId")

	// Find the note with the given Id
	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	// if note.ID == uuid.Nil {
	//     return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	// }

	// Delete the note and return error if encountered
	err := db.Delete(&note, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(utils.Response(nil, 404, false, "Failed to delete note")) //fiber.Map{"status": "error", "message": "Failed to delete note", "data": nil})
	}

	// Return success message
	return c.JSON(utils.Response(nil, 200, true, "Deleted note")) //fiber.Map{"status": "success", "message": "Deleted Note"})
}
