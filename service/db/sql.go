package db

// SQL Statements

const (
	CreateFileTable = `
		CREATE TABLE Files (
			FileID SERIAL PRIMARY KEY,
			FileName TEXT,
			FileURL TEXT,
			PreviewURL TEXT
		);
	`

	CreateTagTable = `
		CREATE TABLE Tags (
			TagID SERIAL PRIMARY KEY,
			Name TEXT,
			Description TEXT
		);
	`

	CreateLibraryTable = `
		CREATE TABLE Library (
			FileID INT NOT NULL,
			TagID INT NOT NULL,
			PRIMARY KEY (FileID, TagID),
			FOREIGN KEY (TagID) REFERENCES Tags(TagID) ON DELETE CASCADE,
			FOREIGN KEY (FileID) REFERENCES Files(FileID) ON DELETE CASCADE
		);
	`

	QueryLibraryExists = `
		SELECT EXISTS (
			SELECT 1
			FROM information_schema.tables 
			WHERE table_name = 'library'
		);
	`
)
