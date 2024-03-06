package handler

import (
	"encoding/json"
	"net/http"
	"web-server/internal/model"
	"web-server/util"
)

func BookHandler(w http.ResponseWriter, r *http.Request) {
	books := []model.Book{
		{ID: 1, Name: "Pride and Prejudice", ISBN: "978-0679783268"},
		{ID: 2, Name: "To Kill a Mockingbird", ISBN: "978-0061120084"},
		{ID: 3, Name: "Beloved", ISBN: "978-1400033416"},
		{ID: 4, Name: "The Handmaid's Tale", ISBN: "978-0385490818"},
		{ID: 5, Name: "Frankenstein", ISBN: "978-0486282114"},
		{ID: 6, Name: "Jane Eyre", ISBN: "978-0141441146"},
		{ID: 7, Name: "Wuthering Heights", ISBN: "978-0141439556"},
		{ID: 8, Name: "The Bell Jar", ISBN: "978-0060837020"},
		{ID: 9, Name: "Middlemarch", ISBN: "978-0141439549"},
		{ID: 10, Name: "A Room of One's Own", ISBN: "978-0156787338"},
		{ID: 11, Name: "The Color Purple", ISBN: "978-0156028356"},
		{ID: 12, Name: "Little Women", ISBN: "978-0147514011"},
		{ID: 13, Name: "Rebecca", ISBN: "978-0380730407"},
		{ID: 14, Name: "1984", ISBN: "978-0451524935"},
		{ID: 15, Name: "Moby-Dick", ISBN: "978-1503280786"},
		{ID: 16, Name: "The Great Gatsby", ISBN: "978-0743273565"},
		{ID: 17, Name: "Ulysses", ISBN: "978-1514645901"},
		{ID: 18, Name: "The Catcher in the Rye", ISBN: "978-0316769488"},
		{ID: 19, Name: "Don Quixote", ISBN: "978-0060934347"},
		{ID: 20, Name: "The Brothers Karamazov", ISBN: "978-0374528379"},
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(books); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Failed to encode books")
		return
	}
}
